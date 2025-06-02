package service

import (
	"errors"
	"strings"
)

var morseCodeMap = map[rune]string{
	'A': ".-", 'B': "-...", 'C': "-.-.",
	'D': "-..", 'E': ".", 'F': "..-.",
	'G': "--.", 'H': "....", 'I': "..",
	'J': ".---", 'K': "-.-", 'L': ".-..",
	'M': "--", 'N': "-.", 'O': "---",
	'P': ".--.", 'Q': "--.-", 'R': ".-.",
	'S': "...", 'T': "-", 'U': "..-",
	'V': "...-", 'W': ".--", 'X': "-..-",
	'Y': "-.--", 'Z': "--..",
	'0': "-----", '1': ".----", '2': "..---",
	'3': "...--", '4': "....-", '5': ".....",
	'6': "-....", '7': "--...", '8': "---..",
	'9': "----.",
	' ': "/",
	'А': ".-", 'Б': "-...", 'В': ".--", 'Г': "--.", 'Д': "-..",
	'Е': ".", 'Ё': ".", 'Ж': "...-", 'З': "--..", 'И': "..",
	'Й': ".---", 'К': "-.-", 'Л': ".-..", 'М': "--", 'Н': "-.",
	'О': "---", 'П': ".--.", 'Р': ".-.", 'С': "...", 'Т': "-",
	'У': "..-", 'Ф': "..-.", 'Х': "....", 'Ц': "-.-.", 'Ч': "---.",
	'Ш': "----", 'Щ': "--.-", 'Ъ': "--.--", 'Ы': "-.--", 'Ь': "-..-",
	'Э': "..-..", 'Ю': "..--", 'Я': ".-.-",
}
var reverseMorseCodeMap = func() map[string]rune {
	m := make(map[string]rune)
	for k, v := range morseCodeMap {
		m[v] = k
	}
	return m
}()

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

func Convert(sm string) (string, error) {
	if len(sm) == 0 {
		return "", errors.New("empty string")
	}
	var res string
	for _, ch := range sm {
		if ch != '.' && ch != '-' && ch != ' ' {
			res = ToMorse(sm)
		} else {
			res = ToText(sm)
		}
	}
	return res, nil
}
