package workerchannel

type WorkerChannel interface {
	Add(function func())
}

func New(numberOfWorkers int) WorkerChannel {
	channel := make(chan job)
	w := &workerChannel{
		Channel: channel,
	}
	for i := 0; i < numberOfWorkers; i++ {
		go w.worker()
	}
	return w
}
