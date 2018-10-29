package handlers

import (
	"log"
)

// CheckErrors throws a panic if error is different from nil
func CheckErrors(e error) {
	if e != nil {
		log.Panicf("Something went wrong: %s", e)
	}
}
