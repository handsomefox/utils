// mut is a package that provides a generic mutex struct, which allows to store data inside of it
// to better identify which data is mutex protected.
package mut

import "sync"

// Mutex is a struct that wraps the data and the sync.Mutex inside of it.
// This should better indicate which data is mutex protected and make
// coding less error prone.
type Mutex[T any] struct {
	mdata *MutexData[T]
}

// MutexData is the actual data.
// It is a separate struct in order for us to avoid copying mutexes.
type MutexData[T any] struct {
	mu  sync.Mutex
	val T
}

// New returns a new mutex ptr
func New[T any](data T) *Mutex[T] {
	return &Mutex[T]{
		mdata: &MutexData[T]{
			mu:  sync.Mutex{},
			val: data,
		},
	}
}

// Lock is the same as sync.Mutex.Lock()
func (m *Mutex[T]) Lock() {
	m.mdata.mu.Lock()
}

// Unlock is the same as sync.Mutex.Unlock()
func (m *Mutex[T]) Unlock() {
	m.mdata.mu.Unlock()
}

// TryLock is the same as sync.Mutex.TryLock()
func (m *Mutex[T]) TryLock() bool {
	return m.mdata.mu.TryLock()
}

// Ptr returns a pointer to internal data without locking
func (m *Mutex[T]) Ptr() *T {
	return &m.mdata.val
}

// Data returns a copy of internal data with locking
func (m *Mutex[T]) DataLocked() T {
	m.Lock()
	data := m.mdata.val
	m.Unlock()
	return data
}

// Data returns a copy the internal data without locking.
func (m *Mutex[T]) Data() T {
	return m.mdata.val
}

// SetLocked changes internal value with locking.
func (m *Mutex[T]) SetLocked(other T) {
	m.Lock()
	m.mdata.val = other
	m.Unlock()
}

// Set changes internal value without locking
func (m *Mutex[T]) Set(other T) {
	m.mdata.val = other
}
