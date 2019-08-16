package vnc

import (
	"net"
	"time"

	"github.com/sibeshkar/demoparser/logger"
	pb_demo "github.com/sibeshkar/demoparser/proto"
)

type ProtoConn struct {
	ProtoReader
	colorMap    ColorMap
	encodings   []Encoding
	fbHeight    uint16
	fbWidth     uint16
	desktopName string
	pixelFormat PixelFormat
	startTime   uint32
}

func (c *ProtoConn) Conn() net.Conn {
	return nil
}

func (c *ProtoConn) Config() interface{} {
	return nil
}

func (c *ProtoConn) Protocol() string {
	return "RFB 003.008"
}
func (c *ProtoConn) PixelFormat() PixelFormat {
	return c.pixelFormat
}

func (c *ProtoConn) Close() error {
	return nil
}

func (c *ProtoConn) SetPixelFormat(pf PixelFormat) error {
	c.pixelFormat = pf
	return nil
}

func (c *ProtoConn) ColorMap() ColorMap                       { return c.colorMap }
func (c *ProtoConn) SetColorMap(cm ColorMap)                  { c.colorMap = cm }
func (c *ProtoConn) Encodings() []Encoding                    { return c.encodings }
func (c *ProtoConn) SetEncodings([]EncodingType) error        { return nil }
func (c *ProtoConn) Width() uint16                            { return c.fbWidth }
func (c *ProtoConn) Height() uint16                           { return c.fbHeight }
func (c *ProtoConn) SetWidth(w uint16)                        { c.fbWidth = w }
func (c *ProtoConn) SetHeight(h uint16)                       { c.fbHeight = h }
func (c *ProtoConn) DesktopName() []byte                      { return []byte(c.desktopName) }
func (c *ProtoConn) SetDesktopName(d []byte)                  { c.desktopName = string(d) }
func (c *ProtoConn) Flush() error                             { return nil }
func (c *ProtoConn) Wait()                                    {}
func (c *ProtoConn) SetProtoVersion(string)                   {}
func (c *ProtoConn) SetSecurityHandler(SecurityHandler) error { return nil }
func (c *ProtoConn) SecurityHandler() SecurityHandler         { return nil }
func (c *ProtoConn) GetEncInstance(typ EncodingType) Encoding {
	for _, enc := range c.encodings {
		if enc.Type() == typ {
			return enc
		}
	}
	return nil
}

func (c *ProtoConn) GetStartTime() uint32 {
	return c.startTime
}

func NewProtoConn(filename string, encs []Encoding, readStart bool) (*ProtoConn, error) {
	rbs, err := NewProtoReader(filename)
	if err != nil {
		logger.Error("failed to open fbs reader:", err)
		return nil, err
	}

	rbsConn := &ProtoConn{ProtoReader: *rbs}
	rbsConn.encodings = encs

	if err != nil {
		logger.Error("Error reading starTime", err)
	}

	if readStart {
		initMsg, err := rbs.ReadStartSession()
		if err != nil {
			logger.Error("failed to open read fbs start session:", err)
			return nil, err
		}

		rbsConn.SetPixelFormat(initMsg.PixelFormat)
		rbsConn.SetHeight(initMsg.FBHeight)
		rbsConn.SetWidth(initMsg.FBWidth)
		rbsConn.SetDesktopName([]byte(initMsg.NameText))

	}

	return rbsConn, nil
}

type ProtoPlayHelper struct {
	Conn *ProtoConn
	//Fbs              VncStreamFileReader
	//serverMessageMap map[uint8]ServerMessage
	//firstSegDone bool
	startTime int
}

func NewProtoPlayer(r *ProtoConn) *ProtoPlayHelper {
	h := &ProtoPlayHelper{Conn: r}
	h.startTime = int(time.Now().UnixNano() / int64(time.Millisecond))
	return h
}

func (h *ProtoPlayHelper) ReadRecordMessage(batch *VncRecordBuffer, SyncWithTimestamps bool, SpeedFactor float64) (pb_demo.Message, error) {
	rbs := h.Conn
	event, _, timestamp, err := rbs.ProtoReader.ReadRecordUpdate()

	//logger.Errorf("Record event readed is %v", event)

	if err != nil {
		logger.Errorf("Error occurred while reading ProtoReader Record file %v", err)
		return event, err
	}

	startTimeMsgHandling := time.Now()

	fbupdateTimestamp := timestamp

	millisSinceStart := int(startTimeMsgHandling.UnixNano()/int64(time.Millisecond)) - h.startTime

	adjustedTimeStamp := float64(fbupdateTimestamp) / SpeedFactor

	millisToSleep := adjustedTimeStamp - float64(millisSinceStart)

	if millisToSleep > 0 {
		time.Sleep(time.Duration(millisToSleep) * time.Millisecond)
	}
	//time.Sleep(50 * time.Millisecond)
	batch.Push(event)

	return event, err

}

func (h *ProtoPlayHelper) ReadEventMessage(batch *VncEventBuffer, SyncWithTimestamps bool, SpeedFactor float64) (ClientMessage, error) {
	rbs := h.Conn
	event, _, timestamp, err := rbs.ProtoReader.ReadEventUpdate()

	if err != nil {
		logger.Errorf("Error occurred while reading ProtoReader client file %v", err)
		return nil, err
	}

	startTimeMsgHandling := time.Now()

	fbupdateTimestamp := timestamp

	millisSinceStart := int(startTimeMsgHandling.UnixNano()/int64(time.Millisecond)) - h.startTime

	adjustedTimeStamp := float64(fbupdateTimestamp) / SpeedFactor

	millisToSleep := adjustedTimeStamp - float64(millisSinceStart)

	if millisToSleep > 0 {
		time.Sleep(time.Duration(millisToSleep) * time.Millisecond)
	}

	batch.Push(event)

	return event, err

}

//This will be looped multiple times in the main
func (h *ProtoPlayHelper) ReadFBMessage(SyncWithTimestamps bool, SpeedFactor float64) (*FramebufferUpdate, error) {
	rbs := h.Conn
	fbUpdate, _, err := rbs.ProtoReader.ReadFbUpdate()
	if err != nil {
		logger.Errorf("Error occurred while reading ProtoReader %v", err)
		return nil, err
	}

	startTimeMsgHandling := time.Now()

	fbupdateTimestamp := fbUpdate.GetTimestamp()

	millisSinceStart := int(startTimeMsgHandling.UnixNano()/int64(time.Millisecond)) - h.startTime

	adjustedTimeStamp := float64(fbupdateTimestamp) / SpeedFactor

	millisToSleep := adjustedTimeStamp - float64(millisSinceStart)

	logger.Debugf("Time: startTimeMsg: %v, fbupdateTimeStamp: %v, millisSinceStart: %v, adjustedTimestamp: %v, millisToSleep: %v", startTimeMsgHandling, fbupdateTimestamp, millisSinceStart, adjustedTimeStamp, millisToSleep)

	if millisToSleep > 0 {
		time.Sleep(time.Duration(millisToSleep) * time.Millisecond)
	}

	parsedFbupdate, err := FrameBufferUpdateRead(fbUpdate, rbs)

	// logger.Debugf("millisSinceStart: %v, adjestedTimeStamp: %v, millisToSleep: %v", millisSinceStart, adjestedTimeStamp, millisToSleep)

	// if millisToSleep > 0 && SyncWithTimestamps {
	// 	logger.Debug("STEP 1")

	// 	time.Sleep(time.Duration(millisToSleep) * time.Millisecond)
	// } else if millisToSleep < -400 {
	// 	logger.Debug("STEP 2")
	// 	logger.Errorf("rendering time is noticeably off, change speedup factor: videoTimeLine: %f, currentTime:%d, offset: %f", adjestedTimeStamp, millisSinceStart, millisToSleep)
	// }

	//logger.Debugf("Error occurred while reading FrameBufferUpdateRead %v", err)

	return parsedFbupdate, err

}
