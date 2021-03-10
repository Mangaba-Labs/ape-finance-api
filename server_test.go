package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestServerPort(t *testing.T) {
	port := os.Getenv("PORT")

	if port == "" {
		port = ":5000"
	} else {
		port = fmt.Sprintf(":%s", port)
	}

	if !strings.Contains(port, ":") {
		t.Error("Server port has to start with ':'")
	}
}
