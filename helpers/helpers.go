package helpers

import (
	"log"
	"math"
)


// HandleError is a reusable error checking function
func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// SplitIntoPages is a function which splits slice of libraries into multiple slices
func SplitIntoPages(data []string, pageSize int) [][]string {
	pageCount := int(math.Ceil(float64(len(data)) / float64(pageSize)))
	pages := make([][]string, pageCount)

	for i := 0; i < pageCount; i++ {
		end := int(math.Min(float64(len(data)), float64((i+1)*pageSize)))
		pages[i] = data[i*pageSize : end]
	}

	return pages
}
