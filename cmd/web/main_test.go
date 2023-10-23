package main

import "testing"

// go test -coverprofile=coverage.out && go tool cover -html=coverage.out

func TestRun(t *testing.T) {
	err := run()

	if err != nil {
		t.Error("failed run()")
	}
}
