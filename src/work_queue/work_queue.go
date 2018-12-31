package work_queue

type Worker interface {
	Run() interface{}
}

type WorkQueue struct {
	Jobs         chan Worker
	Results      chan interface{}
	StopRequests chan int
	NumWorkers   uint
}

// Create a new work queue capable of doing nWorkers simultaneous tasks, expecting to queue maxJobs tasks.
func Create(nWorkers uint, maxJobs uint) *WorkQueue {
	q := new(WorkQueue)
	q.Jobs = make(chan Worker, maxJobs)
	q.Results = make(chan interface{})
	q.StopRequests = make(chan int, maxJobs)
	q.NumWorkers = nWorkers
	for i := 0; i < int(nWorkers); i++ {
		go q.worker()
	}
	return q
}

// A worker goroutine that processes tasks from .Jobs unless .StopRequests has a message saying to halt now.
func (queue WorkQueue) worker() {
	running := true
	length := len(queue.StopRequests)
	// Run tasks from the queue, unless we have been asked to stop.
	for running {

		if queue.Jobs != nil && length == 0 {
			for incoming := range queue.Jobs {
				queue.Results <- incoming.Run()
			}
		} else {
			return
		}
	}

	// TODO: listen on the .Jobs channel for incoming tasks
	// TODO: run tasks by calling .Run()
	// TODO: send the return value back on Results channel
	// TODO: exit (return) when a signal is sent on StopRequests
}

func (queue WorkQueue) Enqueue(work Worker) {
	queue.Jobs <- work
	// TODO: put the work into the Jobs channel so a worker can find it and start the task.
}

func (queue WorkQueue) Shutdown() {
	queue.StopRequests <- int(queue.NumWorkers)
	close(queue.StopRequests)
	// TODO: tell workers to stop processing tasks.
}
