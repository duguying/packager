package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func Upload(filepath string, md5 string, secretKey string, url string) (err error) {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)

	addFileField(w, "file", filepath)
	addField(w, "md5", md5)
	addField(w, "key", secretKey)

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return err
	}

	// set header
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("Referer", "packager - https://github.com/duguying/packager")

	var client http.Client
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	io.Copy(os.Stderr, res.Body) // Replace this with Status.Code check
	return err
}

func addField(w *multipart.Writer, field string, value string) error {
	fw, err := w.CreateFormField(field)
	if err != nil {
		return err
	}

	_, err = fw.Write([]byte(value))
	if err != nil {
		return err
	}

	return nil
}

func addFileField(w *multipart.Writer, field string, filepath string) error {
	fd, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer fd.Close()

	filename := fd.Name()

	fw, err := w.CreateFormFile(field, filename) //这里的file很重要，必须和服务器端的FormFile一致
	if err != nil {
		return err
	}

	// Write file field from file to upload
	_, err = io.Copy(fw, fd)
	if err != nil {
		return err
	}
	w.Close()
	return nil
}
