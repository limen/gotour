package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type errorLevel uint

func TestIota(t *testing.T) {
	const (
		info errorLevel = 1 << iota
		warn
		fatal
	)
	assert.Equal(t, info, errorLevel(1<<0))
	assert.Equal(t, warn, errorLevel(1<<1))
	assert.Equal(t, fatal, errorLevel(1<<2))
	const (
		book = iota
		drink
		shoes
		clothes
	)
	assert.Equal(t, book, 0)
	assert.Equal(t, drink, 1)
	assert.Equal(t, shoes, 2)
	assert.Equal(t, clothes, 3)
	const (
		toys = iota
		_
		_
		foods
	)
	assert.Equal(t, toys, 0)
	assert.Equal(t, foods, 3)
}
