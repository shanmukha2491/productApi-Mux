package handlers

import (
	"net/http"
	"os"
	"path/filepath"
)

func (p *Product) DownloadFile(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("content", "application/octet-stream")

	filename := r.URL.Query().Get("name")
	directory := filepath.Join("files", filename)

	ff, err := os.Open(directory)
	if err != nil {
		http.Error(rw, "Cannot open the file", http.StatusInternalServerError)
		return
	}
	defer ff.Close()

	rw.Header().Set("Content-Disposition", "attachment;filename="+filepath.Base(directory))

	http.ServeFile(rw, r, directory)

}
