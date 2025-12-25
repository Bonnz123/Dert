package dert

import "strings"

type TypeData struct {
	Slice, Pointer, Byte, String, Reader string
}


type resultError struct {
	Message string
}

type error interface {
	Error() string
}

func (e *resultError) Error() string {
	return e.Message
}

var TypeDataList = TypeData{
	Pointer: "Pointer",
	Byte: "Byte",
	String: "String",
	Reader: "Reader",
}

func SayHello() string {
	return "Hello World"
}

func SayMyName(name string, result TypeData) (any, error) {

	if name == "" {
		return nil, &resultError{
			Message: "Name is empty",
		}
	}

	if result.Pointer == TypeDataList.Pointer {
		return &name, nil
	} else if result.String == TypeDataList.String {
		return name, nil
	} else if result.Byte == TypeDataList.Byte {
		return []byte(name), nil
	} else if result.Reader == TypeDataList.Reader {
		return strings.NewReader(name), nil
	}

	return nil, &resultError{
		Message: "Error User Input",
	}

}


func SayAuthor() string{
	return "Hello Bonzz"
}