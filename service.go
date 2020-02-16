package architecture

// Accessor interface is used to access and abstract storage back-ends
type Accessor interface {
	Save(n int, p Version)
	Retrieve() map[int]Version
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
func (vs VersionService) Get() (map[int]Version, error) {
	v := vs.a.Retrieve()
	// if v.Tag == "" {
	// 	return Version{}, fmt.Errorf("no version or versions with id of %d", n)
	// }
	return v, nil
}

// Save method used to access data through
// Version service and the Save method
// the Save method is implemented by the storage backend
func (vs VersionService) Save(n int, p Version) {
	vs.a.Save(n, p)
}

// NewVersionService creates a new service to action
// save and retrieve operations
func NewVersionService(a Accessor) VersionService {
	return VersionService{
		a: a,
	}
}
