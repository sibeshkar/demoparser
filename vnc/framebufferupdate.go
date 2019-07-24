package vnc

import pb "github.com/sibeshkar/vncproxy/proto"

func FrameBufferUpdateRead(fbupdate *pb.FramebufferUpdate, rbs *ProtoConn) (*FramebufferUpdate, error) {
	msg := &FramebufferUpdate{}
	msg.NumRect = uint16(len(fbupdate.GetRectangles()))
	var err error
	for _, fbrect := range fbupdate.GetRectangles() {
		rect := NewRectangle()
		if err := rect.ReadProto(fbrect, rbs); err != nil {
			return nil, err
		}

		msg.Rects = append(msg.Rects, rect)

	}

	return msg, err
}
