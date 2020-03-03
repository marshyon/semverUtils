package architecture

import (
	"fmt"
	"log"

	"github.com/blang/semver"
)

// Accessor interface is used to access and abstract the back-end
type Accessor interface {
	Save(n int, p Version)
	Retrieve() (map[int]Version, int)
}

// Version struct used to store a version
type Version struct {
	Tag string
}

// VersionService uses accessor interface
type VersionService struct {
	a Accessor
}

// Get method used to access data through
// Version service and the Retrieve method
// the Retrieve method is implemented by the storage backend
func (vs VersionService) Get() (map[int]Version, int, error) {
	v, l := vs.a.Retrieve()
	return v, l, nil
}

// Save method used to access data through
// Version service and the Save method
// the Save method is implemented by the storage backend
func (vs VersionService) Save(n int, p Version) {
	vs.a.Save(n, p)
}

func (vs VersionService) GetCurrentVersion() (string, int, error) {
	res, level, err := vs.Get()
	return res[1].Tag, level, err
}

func (vs VersionService) GetNextVersion(currentVersion string, level int) (nextVersion string) {

	if currentVersion == "" {
		nextVersion = "0.0.1"
	} else {

		v, err := semver.Make(currentVersion)
		if err != nil {
			log.Fatalf("failed to create semver : %s\n", err)
		}

		if level == 0 {
			v.Major = v.Major + 1
			v.Minor = 0
			v.Patch = 0
		} else if level == 1 {
			v.Minor = v.Minor + 1
			v.Patch = 0
		} else if level == 2 {
			v.Patch = v.Patch + 1
		}

		nextVersion = fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	}

	return nextVersion
}

// NewVersionService creates a new service to action
// save and retrieve operations
func NewVersionService(a Accessor) VersionService {
	return VersionService{
		a: a,
	}
}
