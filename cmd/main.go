package main

import (
	"flag"
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

func main() {
	framerate := 10
	speedupFactor := 3.0
	fastFramerate := int(float64(framerate) * speedupFactor)

	var logLevel = flag.String("logLevel", "info", "change logging level")
	var protoFile = flag.String("protoFile", "demo/proto.rbs", "file name of demonstration")

	flag.Parse()
	logger.SetLogLevel(*logLevel)

	if len(*protoFile) <= 1 {
		logger.Errorf("please provide a fbs file name")
		return
	}
	if _, err := os.Stat(*protoFile); os.IsNotExist(err) {
		logger.Errorf("File doesn't exist", err)
		return
	}
	encs := []vnc.Encoding{
		&vnc.ZRLEEncoding{},
		&vnc.CursorPseudoEncoding{},
	}

	fbs, err := vnc.NewProtoConn(
		*protoFile,
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
	screenImage := vnc.NewVncCanvas(int(fbs.Width()), int(fbs.Height()))
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
			f, err := os.Create("imgs/img_" + timeTarget.String() + ".jpg")
			if err != nil {
				panic(err)
			}
			defer f.Close()
			jpeg.Encode(f, screenImage.Image, nil)
			//.Add(1 * time.Millisecond)
			if timeLeft > 0 {
				time.Sleep(timeLeft)
				//logger.Error("sleeping= ", timeLeft)
			}
		}
	}()

	msgReader := vnc.NewProtoPlayer(fbs)
	//loop over all messages, feed images to video codec:
	for {
		_, err := msgReader.ReadMessage(true, speedupFactor)

		//vcodec.Encode(screenImage.Image)
		if err != nil {
			os.Exit(-1)
		}
		//vcodec.Encode(screenImage)
	}
}
