package resource

import "time"

// Resource contains information about a resource
type Resource struct {
	ID          int
	Data        string
	CreatedTime time.Time
	IsClosed    bool
}
