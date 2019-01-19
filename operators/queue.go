package operators

import (
	"errors"
	"sync"
)

type Queue struct {
	msgList []Message
	mutex   *sync.Mutex
}

//AddMessage to the queue
func (q *Queue) AddMessage(m Message) error {

	q.mutex.Lock()

	q.msgList = append(q.msgList, m)

	q.mutex.Unlock()

	return nil

}

//GetMessage gets next message from the queue
func (q *Queue) GetMessage() (*Message, error) {

	q.mutex.Lock()
	defer q.mutex.Unlock()

	if len(q.msgList) <= 0 {

		return nil, errors.New("List is empty")
	}
	msg := &q.msgList[0]
	q.msgList = q.msgList[1:]

	return msg, nil

}

//GetNewQueue
func GetNewQueue() *Queue {
	queue := Queue{}
	queue.mutex = &sync.Mutex{}
	return &queue
}
