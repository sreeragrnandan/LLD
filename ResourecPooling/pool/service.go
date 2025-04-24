package pool

import "resource-pooling/resource"

func (p *resourcePool) Acquire() *resource.Resource {
	p.lock.Lock()
	defer p.lock.Unlock()

	for len(p.resources) == 0 {
		p.available.Wait()
	}
	return <-p.resources
}

func (p *resourcePool) Release(resource *resource.Resource) {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.resources <- resource

	p.available.Signal()

	return
}
