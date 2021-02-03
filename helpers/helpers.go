package helpers

import "log"

// HandleError is a reusable error checking function
func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
