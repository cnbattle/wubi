package wubi

import (
	"strconv"
	"strings"
)

// Get 获取字符串的笔记码
func Get(sText string) (unicode []string, err error) {
	for _, value := range []rune(sText) {
		char, err2 := GetChar(string(value))
		if err2 != nil {
			err = err2
			return
		}
		unicode = append(unicode, char)
	}
	return
}

// getChar 获取单字的笔记码
func GetChar(char string) (string, error) {

	key := strings.Replace(strconv.QuoteToASCII(char), `\u`, ``, -1)
	key = strings.Replace(strings.ToUpper(key), `"`, ``, -1)
	if _, ok := dict[key]; ok {
		return strings.Split(dict[key], ",")[0], nil
	}
	return char, nil
}
