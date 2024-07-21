package main

import (
	"fmt"
	"streamer"
)

func main() {
	const numJobs = 1
	const numWorkers = 2

	notifyChan := make(chan streamer.ProcessingMessage, numJobs)
	defer close(notifyChan)

	videoQueue := make(chan streamer.VideoProcessingJob, numJobs)
	defer close(videoQueue)

	wp := streamer.New(videoQueue, numWorkers)

	wp.Run()

	// video := wp.NewVideo(1, "./input/puppy1.mp4", "./output", "mp4", notifyChan, nil)
	ops := &streamer.VideoOptions{
		RenameOutput: true,
		SegmentDuration: 10,
		MaxRate1080:     "1200k",
		MaxRate720p:     "600k",
		MaxRate480p:     "400k",
	}
	video := wp.NewVideo(1, "./input/puppy1.mp4", "./output", "hls", notifyChan, ops)

	videoQueue <- streamer.VideoProcessingJob{Video: video}

	for i := 1; i <= numJobs; i++ {
		msg := <-notifyChan
		fmt.Println("i:", i, "msg:", msg)
	}

	fmt.Println("Done!")
}
