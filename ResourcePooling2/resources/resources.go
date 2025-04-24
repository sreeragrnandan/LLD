package resources

import "time"

// Resource contains property of resource
type resource struct {
	ID           int
	CreatedTime  time.Time
	LastUsedTime time.Time
	Data         string
	IsClosed     bool
}

// Resource interface contains the function signature of resource
type Resource interface {
	IsAlive() bool
	Close()
	GetDataID() (id int, data string)
	UpdateLastUsed() time.Time
}

// NewResource returns new resource
func NewResource(id int, data string) Resource {
	return &resource{
		ID:          id,
		CreatedTime: time.Now(),
		Data:        data,
		IsClosed:    false,
	}
}
