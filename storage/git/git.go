package git

import "github.com/marshyon/codeStructure"

// Db map of Versions by integer key
type Db map[int]architecture.Version

// Save method for git backend
func (m Db) Save(n int, p architecture.Version) {
	m[n] = p
}

// Retrieve method for git backend
func (m Db) Retrieve() map[int]architecture.Version {
	return m
}
