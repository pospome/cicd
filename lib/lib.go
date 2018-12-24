package lib

type MyStruct struct {
	Name string
}

func NewMyStruct(name string) *MyStruct {
	return &MyStruct{
		Name: name,
	}
}
