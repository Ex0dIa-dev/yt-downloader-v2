package helper

import (
	"fmt"
	"log"
	"os"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckResponseStatusCode(code int) error {
	if code > 299 {
		//return errors.New(fmt.Sprintf("[-]Error in processing request, status code: %v", code))
		return fmt.Errorf("[-]Error in processing request, status code: %v", code)
	}
	return nil
}

func FileExists(filename string) bool {

	stat, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !stat.IsDir()
}
