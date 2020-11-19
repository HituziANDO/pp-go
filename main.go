package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Command options
	f := flag.String("file", "", "Input JSON file")
	i := flag.Int("indent", 2, "The number of space")
	flag.Parse()

	buf, err := readJson(*f)
	if err != nil {
		log.Fatal(err)
	}

	str, err := prettyPrint(*buf, *i)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*str)
}

func readJson(file string) (*[]byte, error) {
	if file == "" {
		if len(flag.Args()) > 0 {
			// コマンドライン引数のとき
			//   ./main '{"key1": "val1", "key2": "val2"}'
			str := flag.Arg(0)
			buf := []byte(str)
			return &buf, nil
		} else {
			// パイプライン渡し
			//   cat hoge.json | ./main
			reader := os.Stdin
			buf, err := ioutil.ReadAll(reader)
			if err != nil {
				return nil, err
			}
			return &buf, nil
		}
	} else {
		// オプション指定
		//   ./main -file=hoge.json
		buf, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}
		return &buf, nil
	}
}

func prettyPrint(buf []byte, indent int) (*string, error) {
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
