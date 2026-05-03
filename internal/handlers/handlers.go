package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	service "sprint-6/internal/service"
	"time"
)

func MainHandle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if req.Method != http.MethodGet {
		http.Error(w, fmt.Sprintf("эта ручка не поддерживает %s запросы", req.Method),
			http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, req, "./index.html")
}

func UploadHandle(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("эта ручка не поддерживает %s запросы", req.Method),
			http.StatusMethodNotAllowed)
		return
	}

	req.Body = http.MaxBytesReader(w, req.Body, 10<<20)
	err := req.ParseMultipartForm(1 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, fileHandler, err := req.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	converted, err := service.UniversalConverter(string(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filename := time.Now().UTC().Format("20060102_150405") + filepath.Ext(fileHandler.Filename)
	localFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer localFile.Close()

	_, err = localFile.WriteString(converted)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(converted))

}
