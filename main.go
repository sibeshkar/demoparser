package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
	"path/filepath"

	//"path/filepath"
	"time"
	//"fmt"

	"github.com/sibeshkar/demoparser/encoders"
	"github.com/sibeshkar/demoparser/logger"
	vnc "github.com/sibeshkar/demoparser/vnc"
)

type VNCDemonstration struct {
	batches []DemoBatch
}

type DemoBatch struct {
	obs       image.Image
	done      bool
	reward    float32
	info      string
	timestamp string
}

func Process(filename string, vnc_demo *VNCDemonstration, framerate *int, speedupFactor *float64) error {

	fastFramerate := int(float64(*framerate) * (*speedupFactor))

	fbsFile := filename
	logLevel := "debug"

	logger.SetLogLevel(logLevel)

	encs := []vnc.Encoding{
		&vnc.ZRLEEncoding{},
		&vnc.CursorPseudoEncoding{},
	}

	fbs, err := vnc.NewFbsConn(
		fbsFile,
		encs,
	)
	if err != nil {
		logger.Error("failed to open fbs reader:", err)
		return err
	}

	//launch video encoding process:
	vcodec := &encoders.X264ImageEncoder{FFMpegBinPath: "/usr/bin/ffmpeg", Framerate: *framerate}
	//vcodec := &encoders.DV8ImageEncoder{}
	//vcodec := &encoders.DV9ImageEncoder{}
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	logger.Tracef("current dir: %s", dir)
	go vcodec.Run("./output.mp4")

	//screenImage := image.NewRGBA(image.Rect(0, 0, int(fbs.Width()), int(fbs.Height())))
	screenImage := vnc.NewVncCanvas(int(fbs.Width()), int(fbs.Height()))
	screenImage.DrawCursor = true // modify for drawing cursor

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

			vcodec.Encode(screenImage.Image)

			timeTarget := timeStart.Add(frameDuration)
			timeLeft := timeTarget.Sub(time.Now())

			//writeImageToFile(screenImage.Image, timeTarget)
			batch := DemoBatch{
				obs:       screenImage.Image,
				done:      false,
				reward:    0.0,
				info:      "{}",
				timestamp: timeTarget.String(),
			}
			vnc_demo.batches = append(vnc_demo.batches, batch) //.Add(1 * time.Millisecond)

			if timeLeft > 0 {
				time.Sleep(timeLeft)
				//logger.Error("sleeping= ", timeLeft)
			}
		}
	}()

	msgReader := vnc.NewFBSPlayHelper(fbs)
	//loop over all messages, feed images to video codec:
	for {
		_, err := msgReader.ReadFbsMessage(true, *speedupFactor)

		//vcodec.Encode(screenImage.Image)
		if err != nil {
			break
			//os.Exit(-1)

		}
		//vcodec.Encode(screenImage)
	}
	return err

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
	err := Process("demo/recording_1565093465/server.rbs", vncDemo, frameRate, speedup)
	if err != nil {
		logger.Infof("Error while processing %v", err)
	}

	fmt.Println(len(vncDemo.batches))

}
