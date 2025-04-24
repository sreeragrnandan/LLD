package main

// Store ...
type Store interface {
	Get(key string) (value string)
	Set(key string, value string) (status bool)
	Delete(key string) (status bool)
	Begin()
	Commit()
	RoolBack()
}

// Transactions ...
type Transactions struct {
	updates map[string]string // pending updates
	deletes map[string]bool   //pending deletes
}

// KeyValStore ...
type KeyValStore struct {
	globalData   map[string]string
	transactions []*Transactions
}

// NewStore ...
func NewStore() Store {
	return &KeyValStore{
		globalData:   make(map[string]string),
		transactions: make([]*Transactions, 0),
	}
}

// NewTxn ...
func (kvs *KeyValStore) NewTxn() *Transactions {
	return &Transactions{
		updates: make(map[string]string),
		deletes: make(map[string]bool),
	}
}

// Begin ...
func (kvs *KeyValStore) Begin() {
	kvs.transactions = append(kvs.transactions, kvs.NewTxn())
}

// Commit ...
func (kvs *KeyValStore) Commit() {
	if kvs.transactions == nil {
		return
	}
	currTxn := kvs.transactions[len(kvs.transactions)-1]
	for key := range currTxn.deletes {
		delete(kvs.globalData, key)
	}
	for key, value := range currTxn.updates {
		kvs.globalData[key] = value
	}
	if len(kvs.transactions) > 0 {
		if len(kvs.transactions) == 1 {
			kvs.transactions = nil
		} else {
			kvs.transactions = kvs.transactions[:len(kvs.transactions)-1]
		}
	}
}

// RoolBack ...
func (kvs *KeyValStore) RoolBack() {
	kvs.transactions = nil
}

// Get ...
func (kvs *KeyValStore) Get(key string) (val string) {
	if kvs.transactions != nil {
		currTxn := kvs.transactions[len(kvs.transactions)-1]
		if _, deleted := currTxn.deletes[val]; deleted {
			return ""
		}
		if val, exists := currTxn.updates[val]; exists {
			return val
		}
	} else {
		val = kvs.globalData[key]
	}
	return
}

// Set ...
func (kvs *KeyValStore) Set(key string, value string) (status bool) {
	if kvs.transactions != nil {
		currTxn := kvs.transactions[len(kvs.transactions)-1]
		currTxn.updates[key] = value
		delete(currTxn.deletes, key)
	} else {
		kvs.globalData[key] = value
	}
	return true
}

// Delete ...
func (kvs *KeyValStore) Delete(key string) (status bool) {
	if kvs.transactions != nil {
		currTxn := kvs.transactions[len(kvs.transactions)-1]
		currTxn.deletes[key] = true
		delete(currTxn.updates, key)
	} else {
		delete(kvs.globalData, key)
	}
	return true
}
func main() {
	kvs := NewStore()
	kvs.Begin()
}
