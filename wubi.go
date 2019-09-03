package wubi

import (
	"errors"
	"strconv"
	"strings"
)

func Get(sText string) (unicode []string, err error) {
	for _, value := range []rune(sText) {
		char, err := getChar(string(value))
		if err != nil {
			return
		}
		unicode = append(unicode, char)
	}
	return
}

func getChar(char string) (string, error) {
	if len(char) != 1 {
		return "", errors.New("char len is error:" + strconv.Itoa(len(char)))
	}
	key := strings.Replace(strconv.QuoteToASCII(char), `\u`, ``, -1)
	key = strings.Replace(strings.ToUpper(key), `"`, ``, -1)
	if _, ok := dict[key]; ok {
		return strings.Split(dict[key], ",")[0], nil
	}
	return char, nil
}
