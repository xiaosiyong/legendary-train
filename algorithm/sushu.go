package algorithm

import "fmt"

func FindAllSuShu(max int) {
	origin, wait := make(chan int), make(chan struct{})
	process(origin, wait)
	for i := 2; i <= max; i++ {
		origin <- i
	}
	close(origin)
	<-wait
}

func process(forHandle chan int, wait chan struct{}) {
	go func() {
		cur, ok := <-forHandle
		if !ok {
			close(wait)
			return
		}
		fmt.Println(cur)
		out := make(chan int)
		process(out, wait)
		for v := range forHandle {
			if v%cur != 0 {
				out <- v
			}
		}
		close(out)
	}()
}
