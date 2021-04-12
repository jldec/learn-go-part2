package counter

import (
	"encoding/json"
	"net/http"
)

// Thread-safe counter
// Uses 2 Channels to coordinate reads and writes.
// Must be initialized with New().
type Counter struct {
	readCh  chan uint64
	writeCh chan int
}

// New() is required to initialize a Counter.
func New() *Counter {
	c := &Counter{
		readCh:  make(chan uint64),
		writeCh: make(chan int),
	}

	// The actual counter value lives inside this goroutine.
	// It can only be accessed for R/W via one of the channels.
	go func() {
		var count uint64 = 0
		for {
			select {
			// Reading from readCh is equivalent to reading count.
			case c.readCh <- count:
			// Writing to the writeCh increments count.
			case <-c.writeCh:
				count++
			}
		}
	}()

	return c
}

// Increment counter by pushing an arbitrary int to the write channel.
func (c *Counter) Inc() {
	c.check()
	c.writeCh <- 1
}

// Get current counter value from the read channel.
func (c *Counter) Get() uint64 {
	c.check()
	return <-c.readCh
}

func (c *Counter) check() {
	if c.readCh == nil {
		panic("Uninitialized Counter, requires New()")
	}
}

// return struct for JSON response to Handler
type results struct {
	Counter uint64
}

func (c *Counter) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		enc.SetIndent("", "  ")
		err := enc.Encode(results{
			c.Get(),
		})
		if err != nil {
			http.Error(w, "Error encoding result to JSON: "+err.Error(), 500)
		}
	})
}
