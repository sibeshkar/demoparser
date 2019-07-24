package vnc

import (
	"bytes"
	"encoding/binary"
	"io"
	"io/ioutil"
	"log"
	"os"

	//"vncproxy/common"
	//"vncproxy/encodings"
	"github.com/golang/protobuf/proto"
	"github.com/sibeshkar/demoparser/logger"
	pb "github.com/sibeshkar/vncproxy/proto"
	//"vncproxy/encodings"
	//"vncproxy/encodings"
)

type ProtoReader struct {
	demonstration    *pb.Demonstration
	currentTimestamp int
	index            int
}

func NewProtoReader(filename string) (*ProtoReader, error) {

	in, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	demonstration := &pb.Demonstration{}
	if err := proto.Unmarshal(in, demonstration); err != nil {
		log.Fatalln("Failed to parse demonstration file:", err)
	}

	return &ProtoReader{demonstration: demonstration, index: 0}, err
}

func (rbs *ProtoReader) ReadFbUpdate() (*pb.FramebufferUpdate, error) {
	fbupdate := rbs.demonstration.Fbupdates[rbs.index]
	rbs.currentTimestamp = int(rbs.demonstration.Initmsg.GetStartTime() + fbupdate.GetTimestamp())
	rbs.index += 1
	return fbupdate, nil
}

func (fbs *ProtoReader) Read(p []byte) (n int, err error) {

	return 0, nil
}

func (fbs *ProtoReader) Write(p []byte) (n int, err error) {

	return 0, nil
}

func (rbs *ProtoReader) CurrentTimestamp() int {
	return rbs.currentTimestamp
}

func (rbs *ProtoReader) ReadStartSession() (*ServerInit, error) {
	initMsg := &ServerInit{}
	demo := rbs.demonstration

	initMsg.FBWidth = uint16(demo.Initmsg.GetFBWidth())
	initMsg.FBHeight = uint16(demo.Initmsg.GetFBHeight())
	initMsg.NameText = []byte(demo.Initmsg.GetDesktopName())
	initMsg.PixelFormat = PixelFormat{
		BPP:        uint8(demo.Initmsg.PixelFormat.GetBPP()),
		Depth:      uint8(demo.Initmsg.PixelFormat.GetDepth()),
		BigEndian:  uint8(demo.Initmsg.PixelFormat.GetBigEndian()),
		TrueColor:  uint8(demo.Initmsg.PixelFormat.GetTrueColor()),
		RedMax:     uint16(demo.Initmsg.PixelFormat.GetRedMax()),
		GreenMax:   uint16(demo.Initmsg.PixelFormat.GetGreenMax()),
		BlueMax:    uint16(demo.Initmsg.PixelFormat.GetBlueMax()),
		RedShift:   uint8(demo.Initmsg.PixelFormat.GetRedShift()),
		BlueShift:  uint8(demo.Initmsg.PixelFormat.GetRedShift()),
		GreenShift: uint8(demo.Initmsg.PixelFormat.GetRedShift()),
	}

	// fmt.Printf("FBHeight: %v \n", demo.Initmsg.GetFBHeight())
	// fmt.Printf("FBWidth: %v \n", demo.Initmsg.GetFBWidth())
	// fmt.Printf("RfbHeader: %v \n", demo.Initmsg.GetRfbHeader())
	// fmt.Printf("RfbVersion: %v \n", demo.Initmsg.GetRfbVersion())
	// fmt.Printf("SecType: %v \n", demo.Initmsg.GetSecType())
	// fmt.Printf("StartTime: %v \n", demo.Initmsg.GetStartTime())
	// fmt.Printf("DesktopName: %v \n", demo.Initmsg.GetDesktopName())
	// fmt.Printf("PixelFormat: %v \n", demo.Initmsg.GetPixelFormat())
	return initMsg, nil
}

type FbsReader struct {
	reader           io.ReadCloser
	buffer           bytes.Buffer
	currentTimestamp int
	//pixelFormat      *PixelFormat
	//encodings        []IEncoding
}

func (fbs *FbsReader) Close() error {
	return fbs.reader.Close()
}

func (fbs *FbsReader) CurrentTimestamp() int {
	return fbs.currentTimestamp
}

func (fbs *FbsReader) Read(p []byte) (n int, err error) {
	if fbs.buffer.Len() < len(p) {
		seg, err := fbs.ReadSegment()

		if err != nil {
			logger.Error("FBSReader.Read: error reading FBSsegment: ", err)
			return 0, err
		}
		fbs.buffer.Write(seg.bytes)
		fbs.currentTimestamp = int(seg.timestamp)
	}
	return fbs.buffer.Read(p)
}

//func (fbs *FbsReader) CurrentPixelFormat() *PixelFormat { return fbs.pixelFormat }
//func (fbs *FbsReader) CurrentColorMap() *common.ColorMap       { return &common.ColorMap{} }
//func (fbs *FbsReader) Encodings() []IEncoding { return fbs.encodings }

func NewFbsReader(fbsFile string) (*FbsReader, error) {

	reader, err := os.OpenFile(fbsFile, os.O_RDONLY, 0644)
	if err != nil {
		logger.Error("NewFbsReader: can't open fbs file: ", fbsFile)
		return nil, err
	}
	return &FbsReader{reader: reader}, //encodings: []IEncoding{
		//	//&encodings.CopyRectEncoding{},
		//	//&encodings.ZLibEncoding{},
		//	//&encodings.ZRLEEncoding{},
		//	//&encodings.CoRREEncoding{},
		//	//&encodings.HextileEncoding{},
		//	&TightEncoding{},
		//	//&TightPngEncoding{},
		//	//&EncCursorPseudo{},
		//	&RawEncoding{},
		//	//&encodings.RREEncoding{},
		//},
		nil

}

func (fbs *FbsReader) ReadStartSession() (*ServerInit, error) {

	initMsg := ServerInit{}
	reader := fbs.reader

	var framebufferWidth uint16
	var framebufferHeight uint16
	var SecTypeNone uint32
	//read rfb header information (the only part done without the [size|data|timestamp] block wrapper)
	//.("FBS 001.000\n")
	bytes := make([]byte, 12)
	_, err := reader.Read(bytes)
	if err != nil {
		logger.Error("FbsReader.ReadStartSession: error reading rbs init message - FBS file Version:", err)
		return nil, err
	}

	//read the version message into the buffer, it is written in the first fbs block
	//RFB 003.008\n
	bytes = make([]byte, 12)
	_, err = fbs.Read(bytes)
	if err != nil {
		logger.Error("FbsReader.ReadStartSession: error reading rbs init - RFB Version: ", err)
		return nil, err
	}

	//push sec type and fb dimensions
	binary.Read(fbs, binary.BigEndian, &SecTypeNone)
	if err != nil {
		logger.Error("FbsReader.ReadStartSession: error reading rbs init - SecType: ", err)
	}

	//read frame buffer width, height
	binary.Read(fbs, binary.BigEndian, &framebufferWidth)
	if err != nil {
		logger.Error("FbsReader.ReadStartSession: error reading rbs init - FBWidth: ", err)
		return nil, err
	}
	initMsg.FBWidth = framebufferWidth

	binary.Read(fbs, binary.BigEndian, &framebufferHeight)
	if err != nil {
		logger.Error("FbsReader.ReadStartSession: error reading rbs init - FBHeight: ", err)
		return nil, err
	}
	initMsg.FBHeight = framebufferHeight

	//read pixel format
	pixelFormat := &PixelFormat{}
	binary.Read(fbs, binary.BigEndian, pixelFormat)
	if err != nil {
		logger.Error("FbsReader.ReadStartSession: error reading rbs init - Pixelformat: ", err)
		return nil, err
	}

	initMsg.PixelFormat = *pixelFormat

	//read desktop name
	var desknameLen uint32
	binary.Read(fbs, binary.BigEndian, &desknameLen)
	if err != nil {
		logger.Error("FbsReader.ReadStartSession: error reading rbs init - deskname Len: ", err)
		return nil, err
	}
	initMsg.NameLength = desknameLen

	bytes = make([]byte, desknameLen)
	fbs.Read(bytes)
	if err != nil {
		logger.Error("FbsReader.ReadStartSession: error reading rbs init - desktopName: ", err)
		return nil, err
	}

	initMsg.NameText = bytes

	return &initMsg, nil
}

func (fbs *FbsReader) ReadSegment() (*FbsSegment, error) {
	reader := fbs.reader
	var bytesLen uint32

	//read length
	err := binary.Read(reader, binary.BigEndian, &bytesLen)
	if err != nil {
		logger.Error("FbsReader.ReadSegment: reading len, error reading rbs file: ", err)
		return nil, err
	}

	paddedSize := (bytesLen + 3) & 0x7FFFFFFC

	//read bytes
	bytes := make([]byte, paddedSize)
	_, err = reader.Read(bytes)
	if err != nil {
		logger.Error("FbsReader.ReadSegment: reading bytes, error reading rbs file: ", err)
		return nil, err
	}

	//remove padding
	actualBytes := bytes[:bytesLen]

	//read timestamp
	var timeSinceStart uint32
	binary.Read(reader, binary.BigEndian, &timeSinceStart)
	if err != nil {
		logger.Error("FbsReader.ReadSegment: read timestamp, error reading rbs file: ", err)
		return nil, err
	}

	//timeStamp := time.Unix(timeSinceStart, 0)
	seg := &FbsSegment{bytes: actualBytes, timestamp: timeSinceStart}
	return seg, nil
}

type FbsSegment struct {
	bytes     []byte
	timestamp uint32
}
