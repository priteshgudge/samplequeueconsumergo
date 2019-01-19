package main

import (
	"fmt"
	"sync"

	"github.com/priteshgudge/samplequeueconsumergo/operators"
)

func main() {
	queue := operators.GetNewQueue()
	consumer := operators.GetNewConsumer(queue)
	go consumer.RunConsumer()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		//defer wg.Done()
		fmt.Printf("Added action one\n")
		queue.AddMessage(operators.Message{"one"})
	}()
	go func() {
		//defer wg.Done()
		fmt.Printf("Added action two\n")
		queue.AddMessage(operators.Message{"two"})
	}()
	go func() {
		//defer wg.Done()
		fmt.Printf("Added action three\n")
		queue.AddMessage(operators.Message{"three"})
	}()
	wg.Wait()

}
