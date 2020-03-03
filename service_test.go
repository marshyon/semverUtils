package architecture

import (
	"fmt"

	"github.com/cucumber/godog"
)

// DbMock map of Versions by integer key
type DbMock struct {
	Dbm         map[int]Version
	CommitLevel int
}

var (
	dbmMock DbMock
	vs      VersionService
	currV   string
	currL   int
)

// Save method for git backend
func (m DbMock) Save(n int, p Version) {
	m.Dbm[n] = p
}

// Retrieve method for git backend
func (m DbMock) Retrieve() (map[int]Version, int) {
	return m.Dbm, m.CommitLevel
}

func weHaveNoCurrentVersion() error {

	dbmMock = DbMock{}
	dbmMock.Dbm = make(map[int]Version)
	dbmMock.CommitLevel = 2
	vs = NewVersionService(dbmMock)

	return nil
}

func weShouldHaveANewVersionOf(arg1 string) error {

	currentVersion, _, _ := vs.GetCurrentVersion()
	nextVersion := vs.GetNextVersion(currentVersion, 0)

	if nextVersion != arg1 {
		return fmt.Errorf("something went wrong, expected version [%s] actual was [%s]", arg1, nextVersion)
	}
	return nil
}

func weHaveACurrentVersionOf(arg1 string) error {
	currV = arg1
	return nil
}

func weRelease(arg1 int) error {
	currL = arg1
	return nil
}

func weShouldHaveReturnedANewVersionOf(arg1 string) error {
	nextVersion := vs.GetNextVersion(currV, currL)
	if nextVersion != arg1 {
		return fmt.Errorf("something went wrong, expected version [%s] actual was [%s]", arg1, nextVersion)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^we have no current version$`, weHaveNoCurrentVersion)
	s.Step(`^we should have a new version of "([^"]*)"$`, weShouldHaveANewVersionOf)
	s.Step(`^we have a current version of "([^"]*)"\$$`, weHaveACurrentVersionOf)
	s.Step(`^we release (\d+)\$$`, weRelease)
	s.Step(`^we should have returned a new version of "([^"]*)"\$$`, weShouldHaveReturnedANewVersionOf)
}
