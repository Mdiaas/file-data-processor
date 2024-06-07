package workerchannel

type job struct {
	Execute func()
}
type workerChannel struct {
	Channel chan job
}

func (wc *workerChannel) Add(function func()) {
	job := job{
		Execute: function,
	}
	wc.Channel <- job
}

func (wc *workerChannel) worker() {
	for {
		select {
		case job := <-wc.Channel:
			job.Execute()
		}
	}
}
