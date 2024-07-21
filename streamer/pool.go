package streamer

import "fmt"

type VideoDispatcher struct {
	WorkerPool chan chan VideoProcessingJob
	maxWorkers int
	jobQueue   chan VideoProcessingJob
	Processor  Processor
}

type videoWorker struct {
	id         int
	jobQueue   chan VideoProcessingJob
	workerPool chan chan VideoProcessingJob
}

func newVideoWorker(id int, workerPool chan chan VideoProcessingJob) videoWorker {
	fmt.Println("newVideoWorker: creating video worker id", id)
	return videoWorker{
		id:         id,
		jobQueue:   make(chan VideoProcessingJob),
		workerPool: workerPool,
	}
}

func (vd *VideoDispatcher) Run() {
	fmt.Println("vd.Run: starting worker pool by running workers")
	for i := 0; i < vd.maxWorkers; i++ {
		fmt.Println("vd.Run: starting worker id", i+1)
		worker := newVideoWorker(i+1, vd.WorkerPool)
		worker.start()
	}

	go vd.dispatch()
}

func (w videoWorker) start() {
	fmt.Println("w.start(): starting worker id", w.id)
	go func() {
		for {
			w.workerPool <- w.jobQueue
			job := w.jobQueue
			w.processVideoJob((<-job).Video)
		}
	}()
}

func (vd *VideoDispatcher) dispatch() {
	for {
		job := <-vd.jobQueue
		fmt.Println("vd.dispatch: sending job", job.Video.ID, "to worker job queue")

		go func() {
			workerJobQueue := <-vd.WorkerPool
			workerJobQueue <- job
		}()
	}
}

func (w videoWorker) processVideoJob(video Video) {
	fmt.Println("w.processVideoJob: starting encode on video", video.ID)
	video.encode()
}
