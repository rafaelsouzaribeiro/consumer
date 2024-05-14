package kafka

type Reader struct {
	Brokers []string
}

func NewBrokers(brokers *[]string) *Reader {

	return &Reader{
		Brokers: *brokers,
	}
}
