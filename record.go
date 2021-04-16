package drip

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
)

// A Recorder represents the file reconstruction structure.
type Recorder struct {
	File      string
	data      []byte
	name      string
	length    int
	payload   []byte
	positions map[byte][]int
}

// NewRecorder represents the creation of the instance responsible for reading and reconstructing the original file from the json file.
func NewRecorder(s string) *Recorder {
	return &Recorder{File: s}
}

// open file with information about the location of the bytes.
func (r *Recorder) openFile() error {
	file, err := ioutil.ReadFile(r.File)
	if err != nil {
		return err
	}
	r.data = file
	return nil

}

// Record is responsible for the reconstruction of the original file by reading the json file with location information of the bytes.
func (r *Recorder) Record() error {
	err := r.convertDigest()
	if err != nil {
		return err
	}
	r.setPayload()
	err = r.createFile()
	if err != nil {
		return err
	}
	return nil
}

// Create file.
func (r *Recorder) createFile() error {
	fileName := fmt.Sprintf("%v", r.name)
	err := ioutil.WriteFile(fileName, r.payload, fs.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// Set payload for file.
func (r *Recorder) setPayload() {
	data := make([]byte, 0)

	for i := 0; i < r.length; i++ {
		for k, v := range r.positions {
			for _, b := range v {
				if i == b {
					data = append(data, k)
				}
			}
		}
	}
	r.payload = data
}

// Convert data from json for struct Digest
func (r *Recorder) convertDigest() error {
	err := r.openFile()
	if err != nil {
		return err
	}
	var d *Digester
	err = json.Unmarshal(r.data, &d)
	if err != nil {
		return err
	}
	m := make(map[byte][]int)
	for _, v := range d.Positions {
		m[v.Key] = v.Values
	}

	r.name = d.Name
	r.length = d.Length
	r.positions = m
	return nil
}
