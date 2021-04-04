package errorhandling

import (
	"errors"
	"log"
	"os"
)

func CheckError(err, errorType error) {
	if errors.Is(err, errorType) {
		log.Printf("Error: %s\n", errorType)
		os.Exit(1)
	}
}
