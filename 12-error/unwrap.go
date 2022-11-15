package main

import (
	"errors"
	"fmt"
)

var ErrFeesNotSubmitted error = errors.New("fees not submitted")
var ErrAdmissionCancelled error = errors.New("admission cancelled")
var ErrDocumentNotSubmitted error = errors.New("document not submitted")

func main() {
	err := admission()
	fmt.Println(err)
	err = errors.Unwrap(err)
	fmt.Println(err)

	err = errors.Unwrap(err)
	fmt.Println(err)
}
func admission() error {
	err := fees()
	if err != nil {
		return fmt.Errorf("%w %v", err, ErrAdmissionCancelled)
	}
	return nil
}

func fees() error {
	err := documents()
	if err != nil {
		return fmt.Errorf("%w %v", err, ErrFeesNotSubmitted)
	}
	return nil
}

func documents() error {
	return fmt.Errorf("%w", ErrDocumentNotSubmitted)
}
