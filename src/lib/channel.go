package src

import (
	"encoding/json"
	"fmt"
	"sync"
)

var (
	Ch       chan map[TaskCode][]byte
	initOnce sync.Once
)

// TaskCode a code that address task
type TaskCode int

// Task is entity of task for channel
type Task struct {
	Code TaskCode
	Job  func(msg []byte)
}

const (
	// T1 is task 1
	T1 TaskCode = 1

	// T2 is task 2
	T2 TaskCode = 2

	// T3 is task 3
	T3 TaskCode = 3

	t4 TaskCode = 4
)

// NewTask create new task
func NewTask(code TaskCode, t func(msg []byte)) Task {
	return Task{
		Code: code,
		Job:  t,
	}
}

// ListenChannel will listen channel connection
func ListenChannel(task ...Task) {
	// To trigger channel at first time
	SendToChannel("", t4)

	for {
		select {
		case msgChan := <-Ch:
			if len(task) > 0 {
				for _, t := range task {
					if msg, ok := msgChan[t.Code]; ok {
						t.Job(msg)
					}
				}
			} else {
				var msg interface{}
				for _, msgByte := range msgChan {
					err := json.Unmarshal(msgByte, &msg)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	}
}

// SendToChannel send message to channel
func SendToChannel(msg interface{}, code TaskCode) error {
	initOnce.Do(func() {
		Ch = make(chan map[TaskCode][]byte, 10)
	})

	bytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	mapData := make(map[TaskCode][]byte)
	mapData[code] = bytes
	Ch <- mapData
	return nil
}
