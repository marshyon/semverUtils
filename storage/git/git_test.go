package git

import (
	"fmt"
	"testing"

	architecture "github.com/marshyon/codeStructure"
)

// DbMock map of Versions by integer key
type DbMock map[int]architecture.Version

// Save method for git backend
func (m DbMock) Save(n int, p architecture.Version) {
	m[n] = p
}

// Retrieve method for git backend
func (m DbMock) Retrieve() map[int]architecture.Version {
	return m
}

func TestGetVersions(t *testing.T) {
	dbmMock := DbMock{}

	v1 := architecture.Version{
		Tag: "v0.0.2",
	}

	v2 := architecture.Version{
		Tag: "v0.0.1",
	}

	vs := architecture.NewVersionService(dbmMock)
	vs.Save(1, v1)
	vs.Save(2, v2)

	res, err := vs.Get()
	fmt.Printf("[%#v][%s]\n", res[1].Tag, err)
	if err != nil {
		t.Errorf("error returned storing %s : %s\n", res[1].Tag, err)
	}

	if res[1].Tag != "v0.0.2" {
		t.Errorf("expected v0.0.2, got [%s]\n", res[1].Tag)
	}

	res, err = vs.Get()
	fmt.Printf("[%#v][%s]\n", res[2].Tag, err)
	if err != nil {
		t.Errorf("error returned storing %s : %s\n", res[2].Tag, err)
	}

	// res, err = vs.Get(3)
	// if err == nil {
	// 	t.Errorf("something went wrong retrieving a non existen record %s : %s\n", res[3].Tag, err)
	// }
}
