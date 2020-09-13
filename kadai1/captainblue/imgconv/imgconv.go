package imgconv

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// Extension List
const (
	PNG  string = ".png"
	JPG  string = ".jpg"
	JPEG string = ".jpeg"
)

// GetImgSrcFiles : get image files from dir
func GetImgSrcFiles(dirname string, beforeExt string) ([]string, error) {
	var filenames []string
	err := filepath.Walk(dirname,
		func(path string, info os.FileInfo, err error) error {
			fmt.Println(filepath.Ext(path), beforeExt)
			if filepath.Ext(path) == beforeExt {
				filenames = append(filenames, path)
			}
			return nil
		})

	if err != nil {
		return nil, err
	}
	return filenames, nil
}

func convert(name string, afterExt string, img image.Image) error {
	// 出力先ファイルを用意する
	path := strings.Replace(name, filepath.Ext(name), afterExt, 1)
	output, err := os.Create(path)
	if err != nil {
		return err
	}
	defer output.Close()

	switch afterExt {
	case PNG:
		if err := png.Encode(output, img); err != nil {
			return fmt.Errorf("pngへの変換に失敗しました。%v", err)
		}
	case JPG, JPEG:
		if err := jpeg.Encode(output, img, nil); err != nil {
			return fmt.Errorf("jpegへの変換に失敗しました。%v", err)
		}
	}
	return nil
}

func decodeFileToImg(name string) (image.Image, error) {
	// 変換する画像ファイルを読み込む
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 画像オブジェクトに変換する
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// ConvertImgFIles : convert Image files
func ConvertImgFIles(names []string, afterExt string) error {
	for _, name := range names {
		img, err := decodeFileToImg(name)
		if err != nil {
			return err
		}

		if err := convert(name, afterExt, img); err != nil {
			return err
		}
	}
	return nil
}
