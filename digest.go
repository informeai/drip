package drip

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

// A Digester represents a structure for digesting files.
type Digester struct {
	Name      string `json:"name"`
	data      []byte
	Length    int    `json:"length"`
	Positions []Keys `json:"positions"`
}

// Keys represents the key structure of byte slices.
type Keys struct {
	Key    byte  `json:"key"`
	Values []int `json:"values"`
}

// NewDigest creates a new instance of Digest with the name of the file.
func NewDigester(s string) *Digester {
	return &Digester{Name: s}
}

// Data will return the contents of the file in byte format.
func (d *Digester) Data() ([]byte, error) {
	if len(d.data) == 0 {
		return nil, errors.New("not content of file")
	}
	return d.data, nil
}

// Digest creates the .json file with all the information of bytes and locations.
func (d *Digester) Digest() error {
	d.setData()
	d.Length = len(d.data)
	d.setPositions()
	err := d.createFile()
	if err != nil {
		return err
	}

	return nil
}

// Positions return all positions of bytes
func (d *Digester) GetPositions() []Keys {
	return d.Positions
}

// Open the file.
func (d *Digester) openFile() ([]byte, error) {
	file, err := ioutil.ReadFile(d.Name)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (d *Digester) createFile() error {
	nameFile := strings.TrimSuffix(d.Name, filepath.Ext(d.Name))
	content, err := json.Marshal(d)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("%v.json", nameFile), []byte(content), fs.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
func (d *Digester) setData() {
	file, err := d.openFile()
	if err != nil {
		log.Fatal(err)
	}
	d.data = file

}

// Take the position of each byte in the file.
func (d *Digester) setPositions() {
	m := make(map[byte][]int)
	for k, v := range d.data {
		if s, ok := m[v]; ok {

			m[v] = append(s, k)
		} else {
			s := make([]int, 0)
			m[v] = append(s, k)
		}
	}
	keys := make([]Keys, 0)
	for jk, jv := range m {
		key := Keys{Key: jk, Values: jv}
		keys = append(keys, key)

	}
	d.Positions = keys
}
