package vnc

import (
	pb_demo "github.com/sibeshkar/demoparser/proto"
	//"github.com/sibeshkar/vncproxy/logger"
	//"github.com/amitbet/vnc2video/encoders"
	//"github.com/amitbet/vnc2video/logger"
)

type VncEventBuffer struct {
	Events []ClientMessage
	index  int
}

func NewVncEventBuffer() *VncEventBuffer {
	new_buffer := &VncEventBuffer{
		Events: []ClientMessage{},
		index:  0,
	}
	return new_buffer
}

func (batch *VncEventBuffer) Reset() {
	batch.Events = []ClientMessage{}
	batch.index += 1
}

func (batch *VncEventBuffer) Flush() (int, []ClientMessage) {
	events := batch.Events
	idx := batch.index
	batch.Reset()
	return idx, events
}

func (batch *VncEventBuffer) Push(c ClientMessage) error {
	var err error
	batch.Events = append(batch.Events, c)
	return err
}

type VncRecordBuffer struct {
	Events []pb_demo.Message
	index  int
}

func NewVncRecordBuffer() *VncRecordBuffer {
	new_buffer := &VncRecordBuffer{
		Events: []pb_demo.Message{},
		index:  0,
	}
	return new_buffer
}

func (batch *VncRecordBuffer) Reset() {
	batch.Events = []pb_demo.Message{}
	batch.index += 1
}

func (batch *VncRecordBuffer) Flush() (int, []pb_demo.Message) {
	events := batch.Events
	idx := batch.index
	batch.Reset()
	return idx, events
}

func (batch *VncRecordBuffer) Push(p pb_demo.Message) error {
	var err error
	batch.Events = append(batch.Events, p)
	return err
}
