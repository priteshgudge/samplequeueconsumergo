package operators

import (
	"time"
)

//Consumer  performs actions provided by the messages
type Consumer struct {
	queueReference *Queue
}

//GetNewConsumer constructor for the Consumer
func GetNewConsumer(q *Queue) *Consumer {
	return &Consumer{q}
}

//RunConsumer keeps waiting for messages to arrive into the queue
func (c *Consumer) RunConsumer() {
	for {

		msg, err := c.queueReference.GetMessage()

		if err != nil {
			continue
		}
		msg.performAction()

		time.Sleep(time.Second)
	}

}
