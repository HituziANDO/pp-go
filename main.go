package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Command options
	f := flag.String("file", "", "Input JSON file.")
	i := flag.Int("indent", 2, "The number of space to indent.")
	flag.Parse()

	buf, err := readJson(*f)
	if err != nil {
		log.Fatal(err)
	}

	PrettyPrintOutput(*buf, *i)
}

func readJson(file string) (*[]byte, error) {
	if file == "" {
		if len(flag.Args()) > 0 {
			// コマンドライン引数のとき
			//   pp '{"key1":"val1","key2":"val2"}'
			str := flag.Arg(0)
			buf := []byte(str)
			return &buf, nil
		} else {
			// パイプライン処理
			//   cat test.json | pp
			reader := os.Stdin
			buf, err := ioutil.ReadAll(reader)
			if err != nil {
				return nil, err
			}
			return &buf, nil
		}
	} else {
		// オプション指定
		//   pp -file=test.json
		buf, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}
		return &buf, nil
	}
}
