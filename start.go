package grpool

func (w *worker) start() {
	go func() {
		var job Job
		for {
			// worker free, add it to pool
			w.workerPool <- w

			select {
			case job = <-w.jobChannel:
				job()
			case stop := <-w.stop:
				if stop {
					w.stop <- true
					return
				}
			}
		}
	}()
}
