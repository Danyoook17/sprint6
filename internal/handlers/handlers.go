package handlers

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
// Парсинг 
	err := r.ParseMultipartForm(0)
	if err != nil {
		http.Error(w, "Ошибка загрузки", http.StatusInternalServerError)
		return
	}
// получение файла 
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Файл не найден", http.StatusInternalServerError)
		return
	}
	defer file.Close()
// чтение файла 
	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	// Конвертируем данные
	convertedString, err := service.Convert(string(content))
	if err != nil {
		http.Error(w, "Ошибка конвертации", http.StatusInternalServerError)
		return
	}

	// Создаем имя для нового файла
	timestamp := time.Now().UTC().String()
	extension := filepath.Ext(handler.Filename)
	newFilename := fmt.Sprintf("%s_converted%s", timestamp, extension)

	// Записываем результат в файл
	err = ioutil.WriteFile(newFilename, []byte(convertedString), 0644)
	if err != nil {
		http.Error(w, "Ошибка записи файла", http.StatusInternalServerError)
		return
	}

	// Отправляем результат
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, convertedString)
}
