package sample

import (
	"fmt"

	"github.com/priteshgudge/samplequeueconsumergo/operators"
)

func sample1() {
	queue := operators.GetNewQueue()
	consumer := operators.GetNewConsumer(queue)
	go consumer.RunConsumer()
	go func() {
		fmt.Printf("Added action one\n")
		queue.AddMessage(operators.Message{"one"})
	}()
	go func() {
		fmt.Printf("Added action two\n")
		queue.AddMessage(operators.Message{"two"})
	}()
	go func() {
		fmt.Printf("Added action three\n")
		queue.AddMessage(operators.Message{"three"})
	}()

	fmt.Println("All Done")

}
