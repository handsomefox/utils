// mut is a package that provides a generic mutex struct, which allows to store data inside of it
// to better identify which data is mutex protected.
package mut

import "sync"

// Mutex is a struct that wraps the data and the sync.Mutex inside of it.
// This should better indicate which data is mutex protected and make
// coding less error prone.
type Mutex[T any] struct {
	mutexData *MutexData[T]
}

// MutexData is the actual data.
// It is a separate struct in order for us to avoid copying mutexes.
type MutexData[T any] struct {
	value T
	mutex sync.Mutex
}

// New returns a new mutex ptr.
func New[T any](data T) *Mutex[T] {
	return &Mutex[T]{
		mutexData: &MutexData[T]{
			mutex: sync.Mutex{},
			value: data,
		},
	}
}

// Lock is the same as sync.Mutex.Lock().
func (mutex *Mutex[T]) Lock() {
	mutex.mutexData.mutex.Lock()
}

// Unlock is the same as sync.Mutex.Unlock().
func (mutex *Mutex[T]) Unlock() {
	mutex.mutexData.mutex.Unlock()
}

// TryLock is the same as sync.Mutex.TryLock().
func (mutex *Mutex[T]) TryLock() bool {
	return mutex.mutexData.mutex.TryLock()
}

// Ptr returns a pointer to internal data without locking.
func (mutex *Mutex[T]) Ptr() *T {
	return &mutex.mutexData.value
}

// Data returns a copy of internal data with locking.
func (mutex *Mutex[T]) DataLocked() T {
	mutex.Lock()
	data := mutex.mutexData.value
	mutex.Unlock()
	return data
}

// Data returns a copy the internal data without locking.
func (mutex *Mutex[T]) Data() T {
	return mutex.mutexData.value
}

// SetLocked changes internal value with locking.
func (mutex *Mutex[T]) SetLocked(other T) {
	mutex.Lock()
	mutex.mutexData.value = other
	mutex.Unlock()
}

// Set changes internal value without locking.
func (mutex *Mutex[T]) Set(other T) {
	mutex.mutexData.value = other
}
