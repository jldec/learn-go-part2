package counter_test

import (
	"testing"

	"github.com/jldec/learn-go-part2/counter"
)

func TestCounter(t *testing.T) {
	cnt := counter.New()
	val := cnt.Get()

	if val != 0 {
		t.Errorf("Inital value: %d\nexpected: 0", val)
	}

	cnt.Inc()
	cnt.Inc()
	cnt.Inc()

	val = cnt.Get()

	if val != 3 {
		t.Errorf("After 3x Inc() value: %d\nexpected: 3", val)
	}

	val = cnt.Get()

	if val != 3 {
		t.Errorf("After 3x Inc() + 1 Get() value: %d\nexpected: 3", val)
	}
}

func BenchmarkCounter(b *testing.B) {
	i := 0
	cnt := counter.New()

	for i < b.N {
		cnt.Inc()
		i++
	}

	b.Logf("%d iterations, counter = %d", i, cnt.Get())
}

// TODO: concurrency test
