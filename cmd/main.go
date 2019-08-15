package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image/draw"
	"image/jpeg"
	"os"

	//"path/filepath"
	"time"
	//"fmt"

	"github.com/sibeshkar/demoparser/logger"
	pb_demo "github.com/sibeshkar/demoparser/proto"
	vnc "github.com/sibeshkar/demoparser/vnc"
	//"github.com/sibeshkar/vncproxy/logger"
	//"github.com/amitbet/vnc2video/encoders"
	//"github.com/amitbet/vnc2video/logger"
)

//
//

func ProcessFile(serverfile string, clientfile string, recordfile string, vnc_demo *pb_demo.Demonstration, framerate *int, speedupFactor *float64, logLevel string) error {

	fastFramerate := int(float64(*framerate) * (*speedupFactor))

	logger.SetLogLevel(logLevel)

	encs := []vnc.Encoding{
		&vnc.ZRLEEncoding{},
		&vnc.CursorPseudoEncoding{},
	}

	obs_fbs, err := vnc.NewProtoConn(
		serverfile,
		encs,
		true,
	)

	event_fbs, err := vnc.NewProtoConn(
		clientfile,
		encs,
		true,
	)

	record_fbs, err := vnc.NewProtoConn(
		recordfile,
		encs,
		false,
	)
	if err != nil {
		logger.Error("failed to open fbs reader:", err)
		//return nil, err
	}

	fmt.Printf("The times are: %v, %v, %v", record_fbs.GetStartTime(), event_fbs.GetStartTime(), obs_fbs.GetStartTime())

	//launch video encoding process:
	// vcodec := &encoders.X264ImageEncoder{FFMpegBinPath: "./ffmpeg", Framerate: framerate}
	// //vcodec := &encoders.DV8ImageEncoder{}
	// //vcodec := &encoders.DV9ImageEncoder{}
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// logger.Tracef("current dir: %s", dir)
	// go vcodec.Run("./output.mp4")

	//screenImage := image.NewRGBA(image.Rect(0, 0, int(fbs.Width()), int(fbs.Height())))
	recordBuffer := vnc.NewVncRecordBuffer()
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
			batch := &pb_demo.Batch{
				Timestamp: timeTarget.String(),
			}
			timeLeft := timeTarget.Sub(time.Now())
			//writeImageToFile(screenImage.Image, timeTarget)
			_, events := eventBuffer.Flush()
			_, records := recordBuffer.Flush()

			obs_event, _ := returnEvent("obs", screenImage)
			action_event, _ := returnEvent("actions", events)
			record_event, _ := returnEvent("records", records)

			batch.Iterators = append(batch.Iterators, obs_event)
			batch.Iterators = append(batch.Iterators, action_event)
			batch.Iterators = append(batch.Iterators, record_event)

			vnc_demo.Batches = append(vnc_demo.Batches, batch)
			//.Add(1 * time.Millisecond)
			if timeLeft > 0 {
				time.Sleep(timeLeft)
				//logger.Error("sleeping= ", timeLeft)
			}
		}
	}()

	obsMsgReader := vnc.NewProtoPlayer(obs_fbs)
	eventMsgReader := vnc.NewProtoPlayer(event_fbs)
	recordMsgReader := vnc.NewProtoPlayer(record_fbs)
	//loop over all messages, feed images to video codec:

	go func() {
		for {
			recordMsgReader.ReadRecordMessage(recordBuffer, true, *speedupFactor)

			//vcodec.Encode(screenImage.Image)

			//vcodec.Encode(screenImage)
		}

	}()

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

	var frameRate = flag.Int("fps", 20, "the fps")
	var speedup = flag.Float64("speedup", 1.0, "speedupfactor")
	var logLevel = flag.String("logLevel", "info", "speedupfactor")
	flag.Parse()
	vncDemo := &pb_demo.Demonstration{}
	err := ProcessFile("demo/recording_1565842338/server.rbs", "demo/recording_1565842338/client.rbs", "demo/recording_1565842338/record.rbs", vncDemo, frameRate, speedup, *logLevel)

	if err != nil {
		logger.Infof("Error while processing %v", err)
	}
	i := 0
	for _, batch := range vncDemo.Batches {
		fmt.Println("---------BATCH NUMBER------", i)
		fmt.Print(batch.GetIterators()[1].GetEvents())
		i += 1
		if i == 10 {
			break
		}
		//time.Sleep(1 * time.Second)
	}
	//
	// for _, batch := range vncDemo.Batches {
	// 	fmt.Println("---------BATCH NUMBER------", i)
	// 	iterators := batch.GetIterators()
	// 	for _, iter := range iterators {
	// 		fmt.Print(iter.GetType(), len(iter.GetEvents()))

	// 	}
	// 	i += 1
	// }

}

func returnEvent(event_type string, events ...interface{}) (*pb_demo.Events, error) {
	pb_events := &pb_demo.Events{
		Type: event_type,
	}

	var err error

	for _, event := range events {
		b, err := json.Marshal(event)
		if err != nil {
			return nil, err
		}
		pb_event := &pb_demo.Event{
			Event: b,
		}

		pb_events.Events = append(pb_events.Events, pb_event)

	}

	return pb_events, err
}
