package tests

import (
	"testing"

	. "github.com/informeai/drip"
)

var r = NewRecorder("file.json")

func TestNewRecord(t *testing.T) {
	t.Log(r)
}
func TestRecord(t *testing.T) {
	err := r.Record()
	if err != nil {
		t.Fatal(err)
	}
}
