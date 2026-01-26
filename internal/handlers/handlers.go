package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {

	// получаем текущую директорию
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// меняем текущую директорию на родительскую
	err = os.Chdir("..")

	// получаем новую текущую директорию
	curDir, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// формируем абсолютный путь к файлу
	name := filepath.Join(curDir, "index.html")

	// читаем файл
	s, err := os.ReadFile(name)
	if err != nil {
		http.Error(w, "ошибка при чтении файла", http.StatusInternalServerError)
		return
	}

	// записываем содержимое файла в тело ответа
	w.Write(s)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	// 1 +
	r.ParseMultipartForm(10 << 20) // 10 MB

	// 2.1 получаем файл из формы +
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "ошибка при получении файла", http.StatusInternalServerError)
		return
	}
	// 2.2 закрываем файл +
	defer file.Close()

	// получаем текущую директорию
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// формируем абсолютный путь к файлу
	name := filepath.Join(curDir, "Sprint6", handler.Filename)

	// читаем файл
	s, err := os.ReadFile(name)
	if err != nil {
		http.Error(w, "ошибка при чтении файла", http.StatusInternalServerError)
		return
	}

	res, err := service.Convert(string(s))
	if err != nil {
		log.Fatal(err)
	}

	//5, 6
	fileName := time.Now().UTC().String() + ".txt"
	rightFileName := strings.ReplaceAll(fileName, ":", "_")
	absWayFile := filepath.Join(curDir, "Sprint6", rightFileName)

	if err := os.WriteFile(absWayFile, []byte(res), 0755); err != nil {
		http.Error(w, "ошибка записи файла", http.StatusInternalServerError)
		return
	}

	// 7 +
	w.Write([]byte(res))
}
