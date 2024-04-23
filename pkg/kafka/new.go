package kafka

type Reader struct {
	Brokers []string
}

func NewReader(brokers []string) *Reader {

	return &Reader{
		Brokers: brokers,
	}
}
