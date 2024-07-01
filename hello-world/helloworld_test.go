package helloworld

import "testing"

func TestHelloWorld(t *testing.T) {
	got := HelloWorld()
	want := "Hello, World!"

	if got != want {
		t.Errorf("got: %v, want: %v ", got, want)
	}

}
