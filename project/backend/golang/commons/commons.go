package commons

import "log"

func ErrorTester(successfulMessage, errorMessage string, err error) {
	if err != nil {
		log.Fatal(errorMessage, err.Error())
	} else {
		log.Println(successfulMessage)
	}
}
