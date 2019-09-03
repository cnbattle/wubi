package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	data, err := ioutil.ReadFile("./resource/86版_全码.txt")
	if err != nil {
		panic(err)
	}
	dict := `package wubi

// 86 版五笔编码，总共 21004 个汉字，其中包括一个特殊的 〇
// 字典，key 表示 unicode, value = 表示五笔编码。如果有多个编码，则用','分隔。编码的排序是按长度。
// Warning: Auto-generated file, don't edit.
var Dict = map[string]string{
`
	for _, line := range strings.Split(string(data), "\n") {
		var wubiKey string

		lineS := strings.Split(line, "\t")
		if len(lineS) <= 2 {
			continue
		}
		if len(lineS) == 3 {
			wubiKey = lineS[2]
		} else {
			wubiKey = strings.Join(lineS[2:], ",")
		}

		dict += fmt.Sprintf("\t\"%v\":\"%v\", // %v \n", lineS[0], wubiKey, lineS[1])

	}
	dict += `}`
	WriteWithIoutil("dict.go", dict)
}

//使用ioutil.WriteFile方式写入文件,是将[]byte内容写入文件,如果content字符串中没有换行符的话，默认就不会有换行符
func WriteWithIoutil(name, content string) {
	data := []byte(content)
	if ioutil.WriteFile(name, data, 0644) != nil {
		fmt.Println("写入文件err:", content)
	}
}
