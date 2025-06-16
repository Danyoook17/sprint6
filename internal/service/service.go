package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Используем импортированную мапу из пакета morse
var morseCodeMap = morse.DefaultMorse

// Создаем обратную мапу на основе импортированной
var reverseMorseCodeMap = func() map[string]rune {
	m := make(map[string]rune)
	for k, v := range morseCodeMap {
		m[v] = k
	}
	return m
}()

// Конвертация текста в азбуку Морзе
func ToMorse(text string) string {
	text = strings.ToUpper(text)
	var morse []string
	for _, ch := range text {
		if code, ok := morseCodeMap[ch]; ok {
			morse = append(morse, code)
		}
	}
	return strings.Join(morse, " ")
}

// Конвертация азбуки Морзе в текст
func ToText(morse string) string {
	var text []rune
	words := strings.Split(morse, " / ")
	for _, word := range words {
		letters := strings.Split(word, " ")
		for _, l := range letters {
			if ch, ok := reverseMorseCodeMap[l]; ok {
				text = append(text, ch)
			}
		}
		text = append(text, ' ')
	}
	return strings.TrimSpace(string(text))
}

// Основной метод конвертации
func Convert(str string) (string, error) {
	if len(str) == 0 {
		return "", errors.New("empty string")
	}
	res := StringOrMorse(str)
	return res, nil
}

// Определение типа строки и выбор метода конвертации
func StringOrMorse(sm string) string {
	if strings.ContainsFunc(sm, func(r rune) bool {
		return r != '-' && r != '.' && !strings.ContainsRune(" \t\n", r)
	}) {
		return ToMorse(sm)
	}
	return ToText(sm)
}
