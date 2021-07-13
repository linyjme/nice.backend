package error

import "testing"
import "fmt"

func TestError(t *testing.T) {
	err := Open("/Users/5lmh/Desktop/go/src/test.txt")
	switch v := err.(type) {
	case *PathError:
		fmt.Println("get path error,", v)
	default:
	}

}
