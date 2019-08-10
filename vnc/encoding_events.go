package vnc

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
