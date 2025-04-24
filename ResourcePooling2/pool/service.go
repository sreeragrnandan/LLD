package pool

import (
	"fmt"
	"resource-pooling/resources"
	"time"
)

// Acquire resource from pool
func (p *resourcePool) Acquire() (resources.Resource, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	// Set a 10-second timeout
	timeout := time.After(10 * time.Second)

	for len(p.resources) == 0 {
		select {
		case <-timeout: // Timeout reached
			return nil, fmt.Errorf("error while acquiring resource: timeout after 10 seconds")
		default:
			// p.available.Wait() // Wait until a resource is released
		}
	}

	select {
	case resource := <-p.resources:
		id, _ := resource.GetDataID()
		fmt.Println("Acquired resource with ID", id)
		return resource, nil
	default:
		err := fmt.Errorf("Error on acquiring")
		return nil, err
	}
	// resource := <-p.resources

}

// Release resource from pool
func (p *resourcePool) Release(resource resources.Resource) (err error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	select {
	case p.resources <- resource:
		// id, data := resource.GetDataID()
		// fmt.Printf("Released Resource with Id: %d data: %s \n", id, data)
	default:
		return fmt.Errorf("error while releasing a resource ")
	}
	return
}

// Close the resource pool
func (p *resourcePool) Close() {
	p.lock.Lock()
	defer p.lock.Unlock()
	close(p.resources)
}
