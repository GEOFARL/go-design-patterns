package streamer

type ProcessingMessage struct {
	ID         int
	Successful bool
	Message    string
	OutputFile string
}

type VideoProcessingJob struct {
	Video Video
}

type Processor struct {
	Engine Encoder
}

type Video struct {
	ID           int
	InputFile    string
	OutputDir    string
	EncodingType string
	NotifyChan   chan ProcessingMessage
	Options *VideoOptions
	Encoder Processor
}

type VideoOptions struct {
	RenameOutput    bool
	SegmentDuration int
	MaxRate1080     string
	MaxRate720p     string
	MaxRate480p     string
}

func (vd *VideoDispatcher) NewVideo(id int, input, output, encType string, notifyChan chan ProcessingMessage, ops *VideoOptions) Video {
	if ops == nil {
		ops = &VideoOptions{}
	}

	return Video{
		ID: id,
		InputFile: input,
		OutputDir: output,
		EncodingType: encType,
		NotifyChan: notifyChan,
		Encoder: vd.Processor,
		Options: ops,
	}
}

func (v *Video) encode() {}

func New(jobQueue chan VideoProcessingJob, maxWorkers int) *VideoDispatcher {
	workerPool := make(chan chan VideoProcessingJob, maxWorkers)

	// TODO: implement processor logic
	var e VideoEncoder
	p := Processor{
		Engine: &e,
	}

	return &VideoDispatcher{
		jobQueue:   jobQueue,
		maxWorkers: maxWorkers,
		WorkerPool: workerPool,
		Processor:  p,
	}
}
