package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type hint struct {
	Hint string
}

func appendHello(chi chan *hint, cho chan *hint) {
	h := <-chi
	h.Hint += "Hello"
	cho <- h
}

func appendWorld(chi chan *hint, cho chan *hint) {
	h := <-chi
	h.Hint += "World"
	cho <- h
}

// Test a pipeline which have two pipes.
// Each pipe appends a string to the previous pipe's output
func TestPipeline(t *testing.T) {
	h := &hint{
		Hint: "",
	}
	ch1 := make(chan *hint, 1)
	ch2 := make(chan *hint, 1)
	ch3 := make(chan *hint, 1)
	ch1 <- h
	go func() {
		appendHello(ch1, ch2)
	}()
	go func() {
		appendWorld(ch2, ch3)
	}()

	hint := <-ch3
	assert.Equal(t, "HelloWorld", hint.Hint)

	ch1 <- hint
	go func() {
		appendHello(ch1, ch2)
	}()
	go func() {
		appendWorld(ch2, ch3)
	}()
	assert.Equal(t, "HelloWorldHelloWorld", (<-ch3).Hint)
}
