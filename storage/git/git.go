package git

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	architecture "github.com/marshyon/codeStructure"
)

// Db map of Versions by integer key
// type dbmap map[int]architecture.Version

// Db struct with db map
type Db struct {
	Dbm         map[int]architecture.Version
	CommitLevel int
}

// Save method for git backend
func (m Db) Save(n int, p architecture.Version) {
	m.Dbm[n] = p
}

// Retrieve method for git backend
func (m Db) Retrieve() (map[int]architecture.Version, int) {

	m.Dbm = make(map[int]architecture.Version)
	highestCommit := 2
	exitStatus, output, err := runSystemCmd("git --no-pager log --decorate=short --no-color")
	if err != nil {
		log.Fatalf("ERROR running command [%s] [%d]", err, exitStatus)
	}
	versions, commitTypes := parseGitLogDecoratedOutput(output)

	fmt.Printf("commit types : %v\n", commitTypes)
	vs := architecture.NewVersionService(m)
	if len(versions) == 0 {

		tag := architecture.Version{
			Tag: "0.0.1",
		}
		vs.Save(1, tag)
	} else {
		for _, v := range versions {
			tag := architecture.Version{
				Tag: v,
			}
			vs.Save(len(m.Dbm)+1, tag)

		}
		// calculate highest level of commit message
		highestCommit = calculateHighestCommit(commitTypes)
	}

	return m.Dbm, highestCommit
}

func calculateHighestCommit(cts []string) int {
	overallLevel := 2
	if len(cts) == 0 {
		return overallLevel
	}
	currentLevel := 2
	for _, ct := range cts {
		fmt.Printf(">>[%s]\n", strings.ToLower(ct))

		switch commit := strings.ToLower(ct); commit {
		case "breaking change":
			currentLevel = 0
		case "feature":
			currentLevel = 1
		case "chore":
			currentLevel = 2
		case "documentation":
			currentLevel = 2
		case "style":
			currentLevel = 2
		case "refactor":
			currentLevel = 2
		case "test":
			currentLevel = 2
		case "fix":
			currentLevel = 2
		}
		if currentLevel < overallLevel {
			overallLevel = currentLevel
		}
	}
	return overallLevel
}
func parseGitLogDecoratedOutput(output string) (versions []string, commitTypes []string) {
	lines := strings.Split(output, "\n")
	versionPresent := false
	for _, line := range lines {

		vstr, ok := extractSemVerTag(line)
		if ok {
			versions = append(versions, vstr)
			versionPresent = true
		}
		cstr, ok := extractSemCommit(line)
		if ok && versionPresent == false {
			commitTypes = append(commitTypes, cstr)
		}
	}
	return versions, commitTypes

}

func extractSemVerTag(s string) (versionString string, ok bool) {
	var rgx = regexp.MustCompile(`tag:.+?([0-9\.]+).*?\)`)

	rs := rgx.FindStringSubmatch(s)

	if len(rs) > 0 {
		versionString = rs[1]

		return versionString, true
	}
	return "", false
}

func extractSemCommit(s string) (commitString string, ok bool) {

	var rgx = regexp.MustCompile(`(?i)(BREAKING CHANGE|build|chore|ci|docs|style|refactor|perf|test|feature|fix)\s*:\s*\S+`)

	rs := rgx.FindStringSubmatch(s)

	if len(rs) > 0 {
		commitString = rs[1]
		fmt.Printf("commit string [%s]\n", commitString)
		return commitString, true
	}
	return "", false
}

func runSystemCmd(cmdStr string) (cmdStatus int, output string, err error) {

	cmdArgs := strings.Split(cmdStr, " ")
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmdOutput := &bytes.Buffer{}
	cmdError := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	cmd.Stderr = cmdError
	var returnString string
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-time.After(30 * time.Second):
		if err := cmd.Process.Kill(); err != nil {
			returnString := fmt.Sprintf("failed to kill [%s]", cmd)
			return 2, returnString, errors.New("failed to kill process")
		}
		returnString := fmt.Sprintf("Process timed out [%s]", cmd)
		return 2, returnString, errors.New("process killed as timeout reached")
	case err := <-done:
		if err != nil {
			combinedOutput := fmt.Sprintf("%s %s %s", cmdError.Bytes(), err, cmdOutput.Bytes())
			re := regexp.MustCompile("([0-9])")
			errStr := fmt.Sprintf("%s", err)
			strMatch := re.FindAllString(errStr, -1)
			i, err := strconv.Atoi(strMatch[0])
			if err != nil {
				i = 3
			}
			c := strings.TrimSpace(combinedOutput)

			return i, c, errors.New("command completed with errors")
		} else {
			returnString = string(cmdOutput.Bytes())
		}
	}

	return 0, returnString, nil
}
