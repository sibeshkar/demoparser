package vnc

//Player -> Parser
type Parser interface {
	ReadMessage(bool, float64)
}

//ProtoReader -> Reader
type Reader interface {
	Read() interface{}
}
