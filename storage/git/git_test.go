package git

import (
	"testing"

	architecture "github.com/marshyon/semverUtils"
)

// DbMock map of Versions by integer key
type DbMock struct {
	Dbm         map[int]architecture.Version
	CommitLevel int
}

// Save method for git backend
func (m DbMock) Save(n int, p architecture.Version) {
	m.Dbm[n] = p
}

// Retrieve method for git backend
func (m DbMock) Retrieve() (map[int]architecture.Version, int) {
	return m.Dbm, 2
}

var committests = []struct {
	in  string
	out int
}{
	{"breaking change", 0},
	{"feature", 1},
	{"chore", 2},
	{"documentation", 2},
	{"style", 2},
	{"refactor", 2},
	{"test", 1},
	{"fix", 2},
}

func TestCommitLevel(t *testing.T) {

	for _, cl := range committests {
		testLevel := commitLevel(cl.in)
		if testLevel != cl.out {
			t.Errorf("expected [%s][%d] but got [%d]", cl.in, cl.out, testLevel)
		}

	}
}

func TestGetVersions(t *testing.T) {
	dbmMock := DbMock{}
	dbmMock.Dbm = make(map[int]architecture.Version)
	dbmMock.CommitLevel = 2

	v1 := architecture.Version{
		Tag: "v0.0.2",
	}

	v2 := architecture.Version{
		Tag: "v0.0.1",
	}

	vs := architecture.NewVersionService(dbmMock)
	vs.Save(1, v1)
	vs.Save(2, v2)

	res, _, err := vs.Get()
	if err != nil {
		t.Errorf("error returned getting %s : %s\n", res[1].Tag, err)
	}

	if res[1].Tag != "v0.0.2" {
		t.Errorf("expected v0.0.2, got [%s]\n", res[1].Tag)
	}

	res, _, err = vs.Get()
	if err != nil {
		t.Errorf("error returned getting %s : %s\n", res[2].Tag, err)
	}
}
