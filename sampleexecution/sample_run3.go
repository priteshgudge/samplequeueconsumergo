package sample

import (
	"fmt"

	"github.com/priteshgudge/samplequeueconsumergo/operators"
)

func sample3() {
	queue := operators.GetNewQueue()
	consumer := operators.GetNewConsumer(queue)
	go consumer.RunConsumer()
	ch := make(chan int, 3)
	go func() {
		fmt.Printf("Added action one\n")
		queue.AddMessage(operators.Message{"one"})
		ch <- 1
	}()
	go func() {
		fmt.Printf("Added action two\n")
		queue.AddMessage(operators.Message{"two"})
		ch <- 2
	}()
	go func() {
		fmt.Printf("Added action three\n")
		queue.AddMessage(operators.Message{"three"})
		ch <- 3
	}()
	for i := 0; i < 3; i++ {
		<-ch
	}
	fmt.Println("All Done")

}
