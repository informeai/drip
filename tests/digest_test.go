package tests

import (
	"testing"

	. "github.com/informeai/drip"
)

var d = NewDigester("file.txt")

func TestNewDigest(t *testing.T) {
	// t.Log(d)
}
func TestDigest(t *testing.T) {
	err := d.Digest()
	if err != nil {
		t.Fatal(err)
	}
}
func TestData(t *testing.T) {
	_, err := d.Data()
	if err != nil {
		t.Fatal(err)
	}
	// fmt.Println(data)
}
