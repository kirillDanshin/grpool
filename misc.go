package grpool

// JobDone tells that one of your jobs is done.
// In case you are using WaitAll fn, you should call this method
// every time your job is done.
//
// If you are not using WaitAll then we assume you have your own way of synchronizing.
func (p *Pool) JobDone() {
	p.wg.Done()
}

// WaitCount is how many jobs we should wait when calling WaitAll.
// It is using WaitGroup Add/Done/Wait
func (p *Pool) WaitCount(count int) {
	p.wg.Add(count)
}

// WaitAll waits for all jobs to finish.
func (p *Pool) WaitAll() {
	p.wg.Wait()
}

// Release will release resources used by pool
func (p *Pool) Release() {
	p.dispatcher.stop <- true
	<-p.dispatcher.stop
}
