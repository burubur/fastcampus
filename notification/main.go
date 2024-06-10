package main

import (
	"fmt"
	"time"
)

func main() {
	// tasks channel - done
	// dispatcher - in progress
	// worker - done
	// live worker

	taskQueueCH := make(chan Notification, 10)
	dispatcher := Dispatcher{
		MaxWorker: 3,
		TaskQueue: taskQueueCH,
	}

	dispatcher.StartWorkers()

	// simulate to send notifications using dispatcher
	go func() {
		for i := 1; ; i++ {
			notificationContent := Notification{
				ID:      i,
				Message: fmt.Sprintf("pesan ke: %d", i),
			}

			dispatcher.AssignTask(notificationContent)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// wait forever
	select {}
}

type Notification struct {
	ID      int
	Message string
}

func (n Notification) Send() {
	fmt.Printf("mengirimkan pesan dengan content: %s\n", n.Message)
	time.Sleep(200 * time.Millisecond)
	fmt.Println("pengiriman pesan selesai")
}

type Worker struct {
	ID        int
	TaskQueue chan Notification
}

func (w Worker) Start() {
	go func() {
		for notification := range w.TaskQueue {
			fmt.Printf("worker dengan id: %d sedang mengirim notif\n", notification.ID)
			notification.Send()
			fmt.Printf("worker dengan id: %d telah mengirim notif\n", notification.ID)
		}
	}()
}

type Dispatcher struct {
	TaskQueue chan Notification
	MaxWorker int
}

func (d *Dispatcher) StartWorkers() {
	// start worker
	for i := 1; i <= d.MaxWorker; i++ {
		worker := Worker{
			ID:        i,
			TaskQueue: d.TaskQueue,
		}
		worker.Start()
	}
}

func (d *Dispatcher) AssignTask(notification Notification) {
	d.TaskQueue <- notification
}
