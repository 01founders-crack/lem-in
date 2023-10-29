package errors

import "fmt"

func ErrInvalidDataFormat(err error) error {
	return fmt.Errorf("Invalid Data Format, %s", err)
}

func ErrPaths(err error) error {
	return fmt.Errorf("Path Error, %s", err)
}
