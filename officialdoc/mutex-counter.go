package officialdoc

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]

}

func TestInc() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go c.inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
