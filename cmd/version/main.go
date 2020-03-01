package main

import (
	"fmt"
	"log"

	"github.com/blang/semver"
	architecture "github.com/marshyon/codeStructure"
	"github.com/marshyon/codeStructure/storage/git"
)

func main() {
	dbm := git.Db{}

	ps := architecture.NewVersionService(dbm)

	res, level, err := ps.Get()
	if err != nil {
		log.Fatalf("failed to get : %s\n", err)
	}

	fmt.Printf("CURRENT VERSION length of slice : [%d] first element : [%s]\n", len(res), res[1].Tag)
	fmt.Printf("highest commit level == [[%#v]]\n", level)
	v, err := semver.Make(res[1].Tag)
	if err != nil {
		log.Fatalf("failed to create semver : %s\n", err)
	}
	fmt.Printf("Major: %d\n", v.Major)
	fmt.Printf("Minor: %d\n", v.Minor)
	fmt.Printf("Patch: %d\n", v.Patch)

	fmt.Printf("CommitLevel=>[%d][%T]\n", level, level)

	if level == 0 {
		fmt.Printf("MAJOR\n")
		v.Major = v.Major + 1
		v.Minor = 0
		v.Patch = 0
	} else if level == 1 {
		fmt.Printf("MINOR\n")
		v.Minor = v.Minor + 1
	} else if level == 2 {
		v.Patch = v.Patch + 1
		fmt.Printf("MINOR\n")
	}

	nextVersion := fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	fmt.Printf("NEXT VERSION : [%s]\n", nextVersion)

}
