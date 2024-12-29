package utils

import (
	"log"
)

// HandleError logs a generic message if an error is not nil.
func HandleError(err error) error {
	if err != nil {
		log.Println("Um erro ocorreu:", err)
	}
	return err
}
