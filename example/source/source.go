package source

// 構造体の定義
type Person struct {
	Name  string
	Age   int
	Shark *struct {
		Teeth int
	}
	I8      int8
	I16     int16
	I32     int32
	B       bool
	Bytet   byte
	Float32 float32
	Com     complex64
}

type Person2 struct {
	Name  string
	Age   int
	Shark *struct {
		Teeth int
	}
}
