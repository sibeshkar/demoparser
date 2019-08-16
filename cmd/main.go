package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	//"path/filepath"
	"time"
	//"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/sibeshkar/demoparser/logger"
	pb_demo "github.com/sibeshkar/demoparser/proto"
	"github.com/sibeshkar/demoparser/utils"
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

	//fmt.Printf("The times are: %v, %v, %v", record_fbs.GetStartTime(), event_fbs.GetStartTime(), obs_fbs.GetStartTime())

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

			obs_event, _ := returnImage("obs", &screenImage.Image)
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
	var directory = flag.String("directory", "demo/recording_1565930", "record directory")
	var logLevel = flag.String("logLevel", "info", "log level")

	flag.Parse()
	fileDir := utils.AbsPathify(*directory)
	vncDemo := &pb_demo.Demonstration{}
	err := ProcessFile(fileDir+"/server.rbs", fileDir+"/client.rbs", fileDir+"/record.rbs", vncDemo, frameRate, speedup, *logLevel)

	if err != nil {
		logger.Infof("Error while processing %v", err)
	}
	// i := 0
	// for _, batch := range vncDemo.Batches {
	// 	fmt.Println("---------BATCH NUMBER------", i)
	// 	fmt.Print(batch.GetIterators()[1].GetEvents())
	// 	i += 1
	// 	//time.Sleep(1 * time.Second)
	// }

	writeDemoFile(vncDemo)

	// in, err := ioutil.ReadFile(fname)
	// if err != nil {
	// 	log.Fatalln("Error reading file:", err)
	// }
	// book := &pb_demo.Demonstration{}
	// if err := proto.Unmarshal(in, book); err != nil {
	// 	log.Fatalln("Failed to parse address book:", err)
	// }

	// fmt.Print("Length of Demonstrations is :", len(book.GetBatches()))
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

func writeDemoFile(vncDemo *pb_demo.Demonstration) {

	out, err := proto.Marshal(vncDemo)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

	fname := "demo_" + strconv.FormatInt(time.Now().Unix(), 10) + ".rbs"
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	fmt.Println(fname)

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

func returnImage(event_type string, img *draw.Image) (*pb_demo.Events, error) {
	pb_events := &pb_demo.Events{
		Type: event_type,
	}

	var err error

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, *img, nil)

	pb_event := &pb_demo.Event{
		Event: buf.Bytes(),
	}

	pb_events.Events = append(pb_events.Events, pb_event)

	return pb_events, err
}
