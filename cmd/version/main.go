package main

import (
	"fmt"

	architecture "github.com/marshyon/codeStructure"
	"github.com/marshyon/codeStructure/storage/git"
)

func main() {
	dbm := git.Db{}

	p1 := architecture.Version{
		Tag: "go2.3.4",
	}

	p2 := architecture.Version{
		Tag: "go2.1.0",
	}

	ps := architecture.NewVersionService(dbm)
	ps.Save(1, p1)
	ps.Save(2, p2)
	res, err := ps.Get()

	fmt.Printf(">>>>\n%#v\n%s\n", res, err)
	// if err != nil {
	// 	fmt.Printf("Error getting result for key 1 : %s\n", err)
	// } else {
	// 	fmt.Printf("got %s from key 1\n", res.Tag)
	// }
	// res, err = ps.Get(2)
	// if err != nil {
	// 	fmt.Printf("Error getting result for key 2 : %s\n", err)
	// } else {
	// 	fmt.Printf("got %s from key 2\n", res.Tag)
	// }
	// res, err = ps.Get(3)
	// if err != nil {
	// 	fmt.Printf("Error getting result for key 3 : %s\n", err)
	// }
}
