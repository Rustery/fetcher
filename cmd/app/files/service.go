package files

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func LoadUrlContent(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	content, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	return content, nil
}

func SaveToFile(url *url.URL, content []byte) error {

	path, fileName, err := GetFileNameDataFromUrl(url)
	if err != nil {
		return err
	}
	fullPath := fmt.Sprintf("./web/%s/%s", url.Hostname(), path)
	err = os.MkdirAll(fullPath, os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.Create(fullPath + "/" + fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(content)
	if err != nil {
		return err
	}

	return nil
}

func GetFileNameDataFromUrl(url *url.URL) (string, string, error) {
	path, file := filepath.Split(url.Path)
	ext := filepath.Ext(file)
	if len(ext) == 0 {
		return url.Path, "index.html", nil
	}
	return path, file, nil
}
