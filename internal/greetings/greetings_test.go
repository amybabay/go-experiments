package greetings

import "testing"

func TestHello(t *testing.T) {
    name := "Amy"
    desiredOutput := "Hi, Amy. Welcome!"
    ret := Hello(name)
    if ret != desiredOutput {
        t.Errorf("Hello(%v) == %v, want %v", name, ret, desiredOutput)
    }
}
