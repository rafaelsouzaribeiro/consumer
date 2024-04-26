package pkg

import "time"

type ReadMessage struct {
	Topic     []string
	GroupID   string
	Value     string
	Partition int
	Key       []byte
	Headers   []Header
	Time      time.Time
}

type Header struct {
	Key   string
	Value string
}
