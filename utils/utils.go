package utils

import (
	"fmt"
)

// HandleError prints out an error that occurred
func HandleError(err error) {
	fmt.Printf("%s", err)
}
