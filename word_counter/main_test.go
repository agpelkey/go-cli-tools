package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	exp := 4

	res := count(b, false, false) //Note - will need to adjust exp to account for bytes if testing.
	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}