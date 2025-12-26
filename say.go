package dert

import "strings"

type TypeData string


type resultError struct {
	Message string
}

type error interface {
	Error() string
}

func (e *resultError) Error() string {
	return e.Message
}

const (
	TypeDataString = TypeData("string")
	TypeDataByte = TypeData("byte")
	TypeDataReader = TypeData("reader")
	TypeDataPointer = TypeData("pointer")

)

func SayHello() string {
	return "Hello World"
}

func SayMyName(name string, result TypeData) (any, error) {

	if name == "" {
		return nil, &resultError{
			Message: "Name is empty",
		}
	}

	if result == TypeDataPointer {
		return &name, nil
	} else if result == TypeDataString {
		return name, nil
	} else if result == TypeDataByte {
		return []byte(name), nil
	} else if result == TypeDataReader {
		return strings.NewReader(name), nil
	}

	return nil, &resultError{
		Message: "Error User Input",
	}

}


func SayAuthor() string{
	return "Hello Bonzz"
}