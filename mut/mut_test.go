package mut_test

import (
	"sync"
	"testing"

	"mut"
)

func TestNew(t *testing.T) {
	protected := mut.New(make([]int, 0))

	data := protected.DataLocked()
	data = append(data, 1, 2, 3, 4, 5)
	protected.SetLocked(data)

	protected.Lock()
	defer protected.Unlock()

	data2 := protected.Data()

	if len(data2) != len(data) {
		t.Fatal()
	}
}

func TestLockUnlock(t *testing.T) {
	protected := mut.New("test")
	protected.Lock()

	if protected.TryLock() == true {
		t.Fatal("TryLock did not fail")
	}

	protected.Unlock()

	if protected.TryLock() == false {
		t.Fatal("TryLock did not succeed")
	}

	protected.Unlock()
}

func TestPtr(t *testing.T) {
	protected := mut.New("test")

	protected.Lock()
	ptr1 := protected.Ptr()
	protected.Unlock()

	protected.Lock()
	ptr2 := protected.Ptr()
	protected.Unlock()

	if ptr1 != ptr2 {
		t.Fatal("this should never happen")
	}
}

func TestConcurrent(t *testing.T) {
	protected := mut.New(make([]int, 0))

	appender := func(i int, slice []int) []int {
		return append(slice, i)
	}

	var wg sync.WaitGroup
	n := 100
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int, mu *mut.Mutex[[]int]) {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()

			slice := mu.Data()
			mu.Set(appender(i, slice))
		}(i, protected)
	}

	wg.Wait()

	if len(protected.Data()) != n {
		t.Fatal("Something went wrong")
	}
}

type exampleStruct struct {
	prot *mut.Mutex[string]
}

func newExampleStruct() *exampleStruct {
	return &exampleStruct{
		prot: mut.New(""),
	}
}

func TestStruct(t *testing.T) {
	s := newExampleStruct()

	s.prot.SetLocked(s.prot.DataLocked() + "Hello")
	if s.prot.DataLocked() != "Hello" {
		t.Fatal(s)
	}
}
