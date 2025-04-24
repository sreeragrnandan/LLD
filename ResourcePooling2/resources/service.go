package resources

import "time"

// IsAlive check the resource is alive,closed or not
func (r *resource) IsAlive() bool {
	return !r.IsClosed && time.Since(r.LastUsedTime) < 5*time.Minute
}

// Close To close a perticular resorce
func (r *resource) Close() {
	r.IsClosed = true
}

// GetDataID: Retruns Data and ID of thr resource
func (r *resource) GetDataID() (id int, data string) {
	return r.ID, r.Data
}

func (r *resource) UpdateLastUsed() time.Time {
	r.LastUsedTime = time.Now()
	return r.LastUsedTime
}
