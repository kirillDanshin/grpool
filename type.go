package grpool

import "sync"

// Job represents user request, function which should be executed in some worker.
type Job func()

// Pool represents g-pool
type Pool struct {
	JobQueue   chan Job
	dispatcher *dispatcher
	wg         sync.WaitGroup
}

// dispatcher accepts jobs from clients, and waits for first free worker to deliver job
type dispatcher struct {
	workerPool chan *worker
	jobQueue   chan Job
	stop       chan bool
}

// worker is a goroutine instance which can accept client jobs
type worker struct {
	workerPool chan *worker
	jobChannel chan Job
	stop       chan bool
}
