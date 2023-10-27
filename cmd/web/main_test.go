package main

import "testing"

// go test -coverprofile=coverage.out && go tool cover -html=coverage.out

func TestRun(t *testing.T) {
	_, err := run()

	if err != nil {
		t.Error("failed run()")
	}
}
