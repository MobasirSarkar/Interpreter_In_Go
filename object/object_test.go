package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	dif1 := &String{Value: "mobasir is goat"}
	dif2 := &String{Value: "mobasir is goat"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}
	if dif1.HashKey() != dif2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}
	if hello1.HashKey() == dif1.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}
}
