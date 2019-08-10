package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"

	//"path/filepath"
	"time"
	//"fmt"

	"github.com/sibeshkar/demoparser/logger"
	vnc "github.com/sibeshkar/demoparser/vnc"
	//"github.com/sibeshkar/vncproxy/logger"
	//"github.com/amitbet/vnc2video/encoders"
	//"github.com/amitbet/vnc2video/logger"
)

//
//

type VNCDemonstration struct {
	batches []DemoBatch
}

type DemoBatch struct {
	obs       image.Image
	actions   []vnc.ClientMessage
	done      bool
	reward    float32
	info      string
	timestamp string
}

func ProcessFile(serverfile string, clientfile string, vnc_demo *VNCDemonstration, framerate *int, speedupFactor *float64, logLevel string) error {

	fastFramerate := int(float64(*framerate) * (*speedupFactor))

	logger.SetLogLevel(logLevel)

	encs := []vnc.Encoding{
		&vnc.ZRLEEncoding{},
		&vnc.CursorPseudoEncoding{},
	}

	obs_fbs, err := vnc.NewProtoConn(
		serverfile,
		encs,
	)

	event_fbs, err := vnc.NewProtoConn(
		clientfile,
		encs,
	)
	if err != nil {
		logger.Error("failed to open fbs reader:", err)
		//return nil, err
	}

	//launch video encoding process:
	// vcodec := &encoders.X264ImageEncoder{FFMpegBinPath: "./ffmpeg", Framerate: framerate}
	// //vcodec := &encoders.DV8ImageEncoder{}
	// //vcodec := &encoders.DV9ImageEncoder{}
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// logger.Tracef("current dir: %s", dir)
	// go vcodec.Run("./output.mp4")

	//screenImage := image.NewRGBA(image.Rect(0, 0, int(fbs.Width()), int(fbs.Height())))

	eventBuffer := vnc.NewVncEventBuffer()
	screenImage := vnc.NewVncCanvas(int(obs_fbs.Width()), int(obs_fbs.Height()))
	screenImage.DrawCursor = false

	for _, enc := range encs {
		myRenderer, ok := enc.(vnc.Renderer)

		if ok {
			myRenderer.SetTargetImage(screenImage)
		}
	}

	go func() {
		frameMillis := (1000.0 / float64(fastFramerate)) - 1 //a couple of millis, adjusting for time lost in software commands
		frameDuration := time.Duration(frameMillis * float64(time.Millisecond))
		//logger.Error("milis= ", frameMillis)

		for {
			timeStart := time.Now()

			//vcodec.Encode(screenImage.Image)

			timeTarget := timeStart.Add(frameDuration)
			timeLeft := timeTarget.Sub(time.Now())
			//writeImageToFile(screenImage.Image, timeTarget)
			_, events := eventBuffer.Flush()
			batch := DemoBatch{
				obs:       screenImage.Image,
				actions:   events,
				done:      false,
				reward:    0.0,
				info:      "{}",
				timestamp: timeTarget.String(),
			}
			vnc_demo.batches = append(vnc_demo.batches, batch) //.Add(1 * time.Millisecond)

			//.Add(1 * time.Millisecond)
			if timeLeft > 0 {
				time.Sleep(timeLeft)
				//logger.Error("sleeping= ", timeLeft)
			}
		}
	}()

	obsMsgReader := vnc.NewProtoPlayer(obs_fbs)
	eventMsgReader := vnc.NewProtoPlayer(event_fbs)
	//loop over all messages, feed images to video codec:

	go func() {
		for {
			eventMsgReader.ReadEventMessage(eventBuffer, true, *speedupFactor)

			//vcodec.Encode(screenImage.Image)

			//vcodec.Encode(screenImage)
		}

	}()
	for {
		_, err := obsMsgReader.ReadFBMessage(true, *speedupFactor)

		//vcodec.Encode(screenImage.Image)
		if err != nil {
			return nil
		}
		//vcodec.Encode(screenImage)
	}
}

func writeImageToFile(image draw.Image, timeTarget time.Time) {
	f, err := os.Create("imgs/img_" + timeTarget.String() + ".jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, image, nil)
}

func main() {

	var frameRate = flag.Int("fps", 20, "change logging level")
	var speedup = flag.Float64("speedup", 1.0, "speedupfactor")
	flag.Parse()
	vncDemo := &VNCDemonstration{}
	err := ProcessFile("demo/recording_1565211076/server.rbs", "demo/recording_1565211076/client.rbs", vncDemo, frameRate, speedup, "debug")
	if err != nil {
		logger.Infof("Error while processing %v", err)
	}

	fmt.Println(len(vncDemo.batches))

	for _, batch := range vncDemo.batches {
		fmt.Println(batch.actions, len(batch.actions))

	}

}
