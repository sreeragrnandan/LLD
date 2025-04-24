package pool

import (
	"resource-pooling/resource"
	"sync"
)

// resourcePool contains property of a resource pool
type resourcePool struct {
	poolSize  int
	resources chan *resource.Resource
	lock      sync.Mutex
	available *sync.Cond
}

// ResourcePool interface contains function signature to manage resource pool
type ResourcePool interface {
	Acquire() *resource.Resource
	Release(reaource *resource.Resource)
}

// NewResourcePool creates a new resource pool constructor for resource pool
func NewResourcePool(poolSize int) ResourcePool {
	pool := &resourcePool{
		poolSize:  poolSize,
		resources: make(chan *resource.Resource, poolSize),
		lock:      sync.Mutex{},
	}
	pool.available = sync.NewCond(&pool.lock)

	// inititlize pool with resources
	for i := 0; i < poolSize; i++ {
		pool.resources <- &resource.Resource{ID: i}
	}
	return pool
}
