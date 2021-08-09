package helper

import (
	"errors"
	"fmt"
	"log"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckResponseStatusCode(code int) error {
	if code > 299 {
		return errors.New(fmt.Sprintf("[-]Error in processing request, status code: %v", code))
	}
	return nil
}
