package handlers

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {

	page := `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <title>Document</title>
  </head>
  <body>
    <form
      enctype="multipart/form-data"
      action="http://localhost:8080/upload"
      method="post"
    >
      <input type="file" name="myFile" />
      <input type="submit" value="upload" />
    </form>
  </body>
</html>`

	// записываем содержимое файла в тело ответа
	w.Write([]byte(page))
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	// парсим форму
	r.ParseMultipartForm(10 << 20) // 10 MB

	// получаем файл из формы
	file, _, err := r.FormFile("myFile")
	if err != nil {
		log.Fatal(err)
		http.Error(w, "ошибка при получении файла", http.StatusInternalServerError)
		return
	}

	// закрываем файл
	defer file.Close()

	var b []byte

	// создаем сканер для чтения файла
	scanner := bufio.NewScanner(file)

	// читаем файл построчно
	for scanner.Scan() {
		// вызываем функцию обработки для текущей строки
		b = append(b, []byte(scanner.Text())...)
	}
	// проверяем ошибки во время чтения
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
		http.Error(w, "ошибка при чтении файла", http.StatusInternalServerError)
		return
	}

	fmt.Println(string(b))
	// конвертируем строку
	res, err := service.Convert(string(b))
	if err != nil {
		log.Fatal(err)
		http.Error(w, "пустая строка", http.StatusInternalServerError)
		return
	}
	fmt.Println(string(res))
	// создаем локальный файл
	fileName := time.Now().UTC().String() + ".txt"
	rightFileName := strings.ReplaceAll(fileName, ":", "_")
	rightFileName, err = filepath.Abs(rightFileName)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// записываем результат конвертации в локальный файл
	if err := os.WriteFile(rightFileName, []byte(res), 0755); err != nil {
		log.Fatal(err)
		http.Error(w, "ошибка записи файла", http.StatusInternalServerError)
		return
	}

	// записываем результат в тело ответа
	w.Write([]byte(res))
}
