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

	res, err := ps.Get()
	if err != nil {
		log.Fatalf("failed to get : %s\n", err)
	}
	// for i, val := range res {
	// 	fmt.Printf("==>[%d][%s]\n", i, val)
	// }
	fmt.Printf("CURRENT VERSION length of slice : [%d] first element : [%s]\n", len(res), res[1].Tag)
	fmt.Printf("highest commit level == [[%#v]]\n", ps.Level())
	v, err := semver.Make(res[1].Tag)
	if err != nil {
		log.Fatalf("failed to create semver : %s\n", err)
	}
	fmt.Printf("Major: %d\n", v.Major)
	fmt.Printf("Minor: %d\n", v.Minor)
	fmt.Printf("Patch: %d\n", v.Patch)
	level := ps.Level()
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

	// fmt.Printf("Major: %d\n", v.Major)
	// fmt.Printf("Minor: %d\n", v.Minor)
	// fmt.Printf("Patch: %d\n", v.Patch)
	nextVersion := fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	fmt.Printf("NEXT VERSION : [%s]\n", nextVersion)

}
