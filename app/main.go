package main

import (
	"fmt"
	"streamer"
)

func main() {
	const numJobs = 4
	const numWorkers = 1

	notifyChan := make(chan streamer.ProcessingMessage, numJobs)
	defer close(notifyChan)

	videoQueue := make(chan streamer.VideoProcessingJob, numJobs)
	defer close(videoQueue)

	wp := streamer.New(videoQueue, numWorkers)
	fmt.Println("wp: ", wp)

	wp.Run()

	video := wp.NewVideo(1, "./input/puppy1.mp4", "./output", "mp4", notifyChan, nil)

	videoQueue <- streamer.VideoProcessingJob{Video: video}
}
