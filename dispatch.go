package grpool

func (d *dispatcher) dispatch() {
	for {
		select {
		case job := <-d.jobQueue:
			worker := <-d.workerPool
			worker.jobChannel <- job
		case stop := <-d.stop:
			if stop {
				for i := 0; i < cap(d.workerPool); i++ {
					worker := <-d.workerPool

					worker.stop <- true
					<-worker.stop
				}

				d.stop <- true
				return
			}
		}
	}
}
