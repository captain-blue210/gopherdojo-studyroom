package imgconv

import (
	"os"
	"path/filepath"
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
