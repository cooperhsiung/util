package util

import "fmt"

func TypeOf(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
