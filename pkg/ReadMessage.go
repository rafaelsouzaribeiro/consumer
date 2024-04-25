package pkg

type ReadMessage struct {
	Topic     string
	GroupID   string
	Value     string
	Partition int
	Key       []byte
	Headers   []Header
}

type Header struct {
	Key   string
	Value string
}
