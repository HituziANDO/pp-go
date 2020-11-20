package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrettyPrint(buf []byte, indent int) (*string, error) {
	var dec interface{}
	if err := json.Unmarshal(buf, &dec); err != nil {
		return nil, err
	}

	space := ""
	for i := 0; i < indent; i++ {
		space += " "
	}

	enc, err := json.MarshalIndent(dec, "", space)
	if err != nil {
		return nil, err
	}

	str := string(enc)
	return &str, nil
}

func PrettyPrintOutput(buf []byte, indent int) {
	str, err := PrettyPrint(buf, indent)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*str)
}
