package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "../index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// парсим форму
	r.ParseMultipartForm(10) // 10 MB
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// получаем файл из формы
	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "ошибка при получении файла", http.StatusInternalServerError)
		return
	}

	// закрываем файл
	defer file.Close()

	// читаем тело ответа
	b, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "ошибка при чтении файла", http.StatusInternalServerError)
		return
	}

	// конвертируем строку
	res, err := service.Convert(string(b))
	if err != nil {
		http.Error(w, "пустая строка", http.StatusInternalServerError)
		return
	}

	// создаем локальный файл
	fileName := time.Now().UTC().String() + filepath.Ext(header.Filename)
	rightFileName := strings.ReplaceAll(fileName, ":", "_")
	rightFileName, err = filepath.Abs(rightFileName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// записываем результат конвертации в локальный файл
	if err := os.WriteFile(rightFileName, []byte(res), 0755); err != nil {
		http.Error(w, "ошибка записи файла", http.StatusInternalServerError)
		return
	}

	// устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// записываем результат в тело ответа
	_, err = w.Write([]byte(res))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
