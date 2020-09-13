package main

import (
	"flag"
	"log"
	"os"

	"github.com/captain-blue210/gopherdojo-studyroom/kadai1/captainblue/imgconv"
)

var (
	dirPath, beforeExt, afterExt string
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&dirPath, "p", wd, "変換したい画像のディレクトリパスを指定してください。")
	flag.StringVar(&beforeExt, "b", ".jpg", "変換前の画像の拡張子を指定してください。")
	flag.StringVar(&afterExt, "a", ".png", "変換後の画像の拡張子を指定してください。")
	flag.Parse()
	flag.Args()

	names, err := imgconv.GetImgSrcFiles(dirPath, beforeExt)
	if err != nil {
		log.Fatalf("ファイルの取得に失敗しました。 %v", err)
	}

	if err := imgconv.ConvertImgFIles(names, afterExt); err != nil {
		log.Fatalf("ファイルの変換に失敗しました。 %v", err)
	}
}
