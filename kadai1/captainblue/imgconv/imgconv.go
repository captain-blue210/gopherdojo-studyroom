package imgconv

import (
	//"image"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func GetImgFiles(dirname string) ([]string, error) {
	var filenames []string
	err := filepath.Walk(dirname,
		func(path string, info os.FileInfo, err error) error {
			filenames = append(filenames, path)
			return nil
		})

	if err != nil {
		return nil, err
	}
	return filenames, nil
}

func ConvertImgFIle(name string) error {
	fmt.Println("画像の変換を開始します。")
	// 変換する画像ファイルを読み込む
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	// 画像オブジェクトに変換する
	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}

	// 出力先ファイルを用意する
	path := strings.Replace(name, filepath.Ext(name), ".png", 1)
	output, err := os.Create(path)
	if err != nil {
		return err
	}
	defer output.Close()

	// 画像ファイルを出力する
	err = png.Encode(output, img)
	if err != nil {
		return err
	}
	return nil
}
