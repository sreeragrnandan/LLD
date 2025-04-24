package pool

import (
	"resource-pooling/resources"
	"sync"
)

// resourcePool contains have property of a resourcePool
type resourcePool struct {
	poolSize  int
	resources chan resources.Resource
	lock      sync.Mutex
	available *sync.Cond
}

// ResourcePool interface contains function signature of resourcePool
type ResourcePool interface {
	// Acquire resource from pool
	Acquire() (res resources.Resource, err error)
	// Release resource from pool
	Release(resource resources.Resource) (err error)
	// Close the resource pool
	Close()
}

// NewResourcePool constructs a new resource pool
func NewResourcePool(size int) ResourcePool {
	pool := &resourcePool{
		poolSize:  size,
		resources: make(chan resources.Resource, size),
		lock:      sync.Mutex{},
	}
	pool.available = sync.NewCond(&pool.lock)

	for i := 0; i < size; i++ {
		pool.resources <- resources.NewResource(i, "data of resource "+string(rune(i)))
	}
	return pool
}
