package streamer

import (
	"fmt"

	"github.com/xfrr/goffmpeg/transcoder"
)

type Encoder interface {
	EncodeToMP4(v *Video, baseFileName string) error
	EncodeToHLS(v *Video, baseFileName string) error
}

type VideoEncoder struct{}

func (ve *VideoEncoder) EncodeToMP4(v *Video, baseFileName string) error {
	trans := new(transcoder.Transcoder)
	outputPath := fmt.Sprintf("%s/%s.mp4", v.OutputDir, baseFileName)
	err := trans.Initialize(v.InputFile, outputPath)
	if err != nil {
		return err
	}

	trans.MediaFile().SetVideoCodec("libx264")
	done := trans.Run(false)

	err = <-done
	if err != nil {
		return err
	}

	return nil
}

func (ve *VideoEncoder) EncodeToHLS(v *Video, baseFileName string) error {
	return nil
}
