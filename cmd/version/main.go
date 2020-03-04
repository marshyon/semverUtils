package main

import (
	"fmt"
	"log"

	architecture "github.com/marshyon/semverUtils"
	"github.com/marshyon/semverUtils/storage/git"
)

func main() {
	dbm := git.Db{}

	vs := architecture.NewVersionService(dbm)

	currentVersion, _, err := vs.GetCurrentVersion()
	if err != nil {
		log.Fatalf("failed to get current version: %s\n", err)
	}
	fmt.Printf("%s\n", currentVersion)

}
