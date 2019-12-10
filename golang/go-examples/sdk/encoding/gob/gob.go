package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

var (
	err error
)

func main() {
	var buf bytes.Buffer
	s := []byte("hello demon")

	defer debugPrintError(err)

	if err = gob.NewEncoder(&buf).Encode(s); err == nil {
		dec := gob.NewDecoder(bytes.NewBuffer(buf.Bytes()))

		var res []byte

		if err = dec.Decode(&res); err == nil {
			fmt.Println(string(res))
		}
	}
}

func debugPrintError(err error) {
	if err != nil {
		log.Println(err)
	}
}
