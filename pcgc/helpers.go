package pcgc

import "log"

func panicOnUnrecoverableError(err error) {
	if err != nil {
		log.Panicf("Did not expect a failure, but got: %v", err)
	}
}

func logError(action func() error) {
	err := action()
	if err != nil {
		log.Printf("Unexpected error: %v", err)
	}
}
