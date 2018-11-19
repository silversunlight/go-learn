package officialdoc

import "fmt"

/**
select 会等待，直到有一个case能运行，如果有多个同时准备好了，就随机运行一个
*/
func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("this is default")
		}
	}
}

func TestFibonacci2() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}
