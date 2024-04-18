package types

type BirdName string

type Wings struct {
	Length      int
	Width       int
	Fly         bool
	FlyDistance int
}

type WingsInterface interface {
	Fly() bool
}
