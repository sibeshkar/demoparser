package vnc

import (
	"github.com/sibeshkar/demoparser/logger"
	pb "github.com/sibeshkar/vncproxy/proto"
)

func FrameBufferUpdateRead(fbupdate *pb.FramebufferUpdate, rbs *ProtoConn) (*FramebufferUpdate, error) {
	msg := &FramebufferUpdate{}
	msg.NumRect = uint16(len(fbupdate.GetRectangles()))

	logger.Debugf("-------Reading FrameBuffer update with %d rects-------", msg.NumRect)
	var err error
	for i, fbrect := range fbupdate.GetRectangles() {
		logger.Debugf("----------RECT %d----------", i)
		rect := NewRectangle()
		if err := rect.ReadProto(fbrect, rbs); err != nil {
			logger.Debug(err)
			return nil, err
		}
		logger.Debugf("----End RECT #%d Info (%dx%d) encType:%s", i, rect.Width, rect.Height, rect.EncType)

		msg.Rects = append(msg.Rects, rect)

	}

	return msg, err
}
