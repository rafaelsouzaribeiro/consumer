package entity

type Iconsumer interface {
	Receive(consumer *Consumer)
}
