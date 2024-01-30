package main

import (
	"testing"

	"github.com/reddyduggempudi/github-actions-demo-go/hello"
)

func TestGreeting(t *testing.T) {
	result := hello.Greet()
	if result != "Hello GitHub Actions." {
		t.Errorf("Greet() = %s; want Hello GitHub Actions.", result)
	}
}
