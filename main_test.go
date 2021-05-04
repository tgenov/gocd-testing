package main

import "testing"
 
func TestConvert(t *testing.T) {
    if Run() != "Hello World!" {
		t.Fail()
	}
}