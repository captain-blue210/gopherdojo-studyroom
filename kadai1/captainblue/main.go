package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"gopherdojo-studyroom/kadai1/captainblue/imgconv"
)

var (
	beforeExt, afterExt string
)

func init() {
	flag.StringVar(&beforeExt, "b", ".jpg", "変換前の画像の拡張子を指定してください。")
	flag.StringVar(&afterExt, "a", ".png", "変換後の画像の拡張子を指定してください。")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) <= 0 {
		log.Fatal("引数を指定してください。")
	}

	names, err := imgconv.GetImgFiles(args[0])
	if err != nil {
		log.Fatal("ファイルの取得に失敗しました。")
	}

	for _, name := range names {
		if filepath.Ext(name) == ".jpg" {
			err := imgconv.ConvertImgFIle(name)
			if err != nil {
				log.Fatal("画像の変換に失敗しました。")
			} else {
				fmt.Println("画像の変換が完了しました。")
			}
		}
	}
}
