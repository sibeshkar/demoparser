package vnc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"

	//"vncproxy/common"
	//"vncproxy/encodings"

	"github.com/matttproud/golang_protobuf_extensions/pbutil"
	"github.com/sibeshkar/demoparser/logger"
	pb "github.com/sibeshkar/vncproxy/proto"
	//"vncproxy/encodings"
	//"vncproxy/encodings"
)

type ProtoReader struct {
	reader           *os.File
	currentTimestamp int
}

func NewProtoReader(filename string) (*ProtoReader, error) {

	reader, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		logger.Errorf("unable to open file: %s, error: %v", filename, err)

	}

	return &ProtoReader{reader: reader}, err
}

func (rbs *ProtoReader) ReadFbUpdate() (*pb.FramebufferUpdate, int, error) {
	fbupdate := &pb.FramebufferUpdate{}
	len, err := pbutil.ReadDelimited(rbs.reader, fbupdate)
	{

	}
	return fbupdate, len, err

}

func (rbs *ProtoReader) ReadEventUpdate() (ClientMessage, int, uint32, error) {
	msgType := &pb.MessageType{}
	var eventBlank ClientMessage
	var timestamp uint32
	var len int
	var err error
	pbutil.ReadDelimited(rbs.reader, msgType)
	if msgType.GetType() == uint32(4) {
		event := &pb.KeyEvent{}
		len, err = pbutil.ReadDelimited(rbs.reader, event)
		key_event := KeyEvent{
			Down: uint8(event.GetDown()),
			Key:  Key(event.GetKey()),
		}
		timestamp = event.GetTimestamp()
		fmt.Printf("Key event is %v", key_event)
		return &key_event, len, timestamp, err
	} else if msgType.GetType() == uint32(5) {
		event := &pb.PointerEvent{}
		len, err = pbutil.ReadDelimited(rbs.reader, event)
		pointer_event := PointerEvent{
			Mask: uint8(event.GetMask()),
			X:    uint16(event.GetX()),
			Y:    uint16(event.GetY()),
		}
		timestamp = event.GetTimestamp()
		fmt.Println("Pointer event is ", pointer_event)
		return &pointer_event, len, timestamp, err

	}

	return eventBlank, len, timestamp, err

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
	initMsgRead := &pb.InitMsg{}
	pbutil.ReadDelimited(rbs.reader, initMsgRead)

	initMsg.FBWidth = uint16(initMsgRead.GetFBWidth())
	initMsg.FBHeight = uint16(initMsgRead.GetFBHeight())
	initMsg.NameText = []byte(initMsgRead.GetDesktopName())
	initMsg.PixelFormat = PixelFormat{
		BPP:        uint8(initMsgRead.PixelFormat.GetBPP()),
		Depth:      uint8(initMsgRead.PixelFormat.GetDepth()),
		BigEndian:  uint8(initMsgRead.PixelFormat.GetBigEndian()),
		TrueColor:  uint8(initMsgRead.PixelFormat.GetTrueColor()),
		RedMax:     uint16(initMsgRead.PixelFormat.GetRedMax()),
		GreenMax:   uint16(initMsgRead.PixelFormat.GetGreenMax()),
		BlueMax:    uint16(initMsgRead.PixelFormat.GetBlueMax()),
		RedShift:   uint8(initMsgRead.PixelFormat.GetRedShift()),
		BlueShift:  uint8(initMsgRead.PixelFormat.GetRedShift()),
		GreenShift: uint8(initMsgRead.PixelFormat.GetRedShift()),
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
