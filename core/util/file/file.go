package file

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func GetFileReader(path string) (*bytes.Buffer, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("media", filepath.Base(path))
	if err != nil {
		return nil, err
	}
	fh, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	_, _ = io.Copy(fileWriter, fh)

	bodyWriter.Close()
	return bodyBuf, err
}
