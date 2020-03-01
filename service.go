package architecture

// Accessor interface is used to access and abstract storage back-ends
type Accessor interface {
	Save(n int, p Version)
	Retrieve() (map[int]Version, int)
}

// Version struct is the main
// data componenent
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
	// if v.Tag == "" {
	// 	return Version{}, fmt.Errorf("no version or versions with id of %d", n)
	// }
	return v, l, nil
}

// Save method used to access data through
// Version service and the Save method
// the Save method is implemented by the storage backend
func (vs VersionService) Save(n int, p Version) {
	vs.a.Save(n, p)
}

// func (vs VersionService) Level() int {
// 	// fmt.Printf("HERE>>>>> %#v\n", vs.a.Level())
// 	return vs.a.Level()
// }

// NewVersionService creates a new service to action
// save and retrieve operations
func NewVersionService(a Accessor) VersionService {
	return VersionService{
		a: a,
	}
}
