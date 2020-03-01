package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/blang/semver"
)

var (
	versionList semver.Versions
)

func main() {

	// os.Chdir("../go")

	mydir, err := os.Getwd()
	if err == nil {
		fmt.Println(mydir)
	}

	versionStrings := make([]string, 0)
	// versionList := semver.Versions
	versionList := []semver.Version{}
	// versionStrings := []string

	// versionStrings = append(versionStrings, "0.0.1")
	// versionStrings = append(versionStrings, "1.0.1")
	// versionStrings = append(versionStrings, "5.0.1")
	// versionStrings = append(versionStrings, "1.100.1")
	// versionStrings = append(versionStrings, "3.0.1")
	// versionStrings = append(versionStrings, "0.0.2")
	// versionStrings = append(versionStrings, "1.9.0")

	for _, v := range versionStrings {
		fmt.Printf("[%s]\n", v)
		v, ve := semver.Make(v)
		if ve != nil {
			log.Fatalf("failed parting [%s] : %s\n", v, ve)
		}
		fmt.Println(v)
		// fmt.Printf("[%T]\n", v)
		versionList = append(versionList, v)
	}

	// sorted := semver.Sort(versionList)
	semver.Sort(versionList)
	fmt.Println(versionList)
	// v, err := semver.Make("0.0.1-alpha.preview+123.github")
	// if err != nil {
	// 	log.Fatalf("failed to create semver : %s\n", err)
	// }
	// fmt.Printf("v is a [%T]\n", v)
	// fmt.Printf("Major: %d\n", v.Major)
	// fmt.Printf("Minor: %d\n", v.Minor)
	// fmt.Printf("Patch: %d\n", v.Patch)
	// fmt.Printf("Pre: %s\n", v.Pre)
	// fmt.Printf("Build: %s\n", v.Build)

	// fmt.Println(semver.Sort(svers))
	exitStatus, output, err := runSystemCmd("git --no-pager log --decorate=short --no-color")
	if err != nil {
		log.Fatalf("ERROR running command [%s] [%d]", err, exitStatus)
	}
	parseGitLogDecoratedOutput(output)
}

func parseGitLogDecoratedOutput(output string) {
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		fmt.Printf(":: %s\n", line)
		extractSemVerTag(line)
	}

}

func extractSemVerTag(s string) {
	var rgx = regexp.MustCompile(`tag:.+?([0-9\.]+).*?\)`)

	rs := rgx.FindStringSubmatch(s)

	if len(rs) > 0 {
		tagStr := rs[1]
		fmt.Printf("tag : [%s]\n", tagStr)
	}
}
