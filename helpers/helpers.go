package helpers

import (
	"log"
)

// HandleError is a reusable error checking function
func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//splitSearchResults
func splitSearchResults(resSlice []string) []string {
	size := 5
	var pages []string
	var j int
	for i := 0; i < len(resSlice); i += size{
		j += size
		if j > len(resSlice) {
			j = len(resSlice)
		}
		// do what do you want to with the sub-slice, here just printing the sub-slice
		pages = resSlice[i:j]
	}
}