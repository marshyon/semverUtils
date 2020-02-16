package main

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
)

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
