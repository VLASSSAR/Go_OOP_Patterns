package pkg

type Consumer interface {
	GetName() string
	Update(pubName string)
}
