package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func (p *Product) UploadDocument(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Print(err)
		http.Error(rw, "Input is wrong", http.StatusInternalServerError)
		return
	}

	name := r.Form.Get("name")
	if name == "" {
		http.Error(rw, "name value not specified", http.StatusBadRequest)
		return
	}

	f, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(rw, "name value not specified", http.StatusBadRequest)
		return
	}
	defer f.Close()

	// get the file extension
	fileExtention := strings.ToLower(filepath.Ext(handler.Filename))
	// create Folders
	path := filepath.Join(".", "files")
	_ = os.Mkdir(path, os.ModePerm)
	fullPath := path + "/" + name + fileExtention

	// open and copy files
	file, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		http.Error(rw, "name value not specified", http.StatusBadRequest)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, f)
	if err != nil {
		http.Error(rw, "name value not specified", http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)

}


