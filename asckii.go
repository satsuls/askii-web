package main

import (
	"io/ioutil"
	"strings"
)

func askii(s string, font string) string {
	data, _ := ioutil.ReadFile("./font/" + font + ".txt")

	str := string(data)
	newLineCounter := 0
	arrToRet := ""

	arrStrArgs := strings.Split(s, "\\n")

	for _, arrStrArgsValue := range arrStrArgs {

		for i := 1; i <= 8; i++ {

			for _, valueArgs := range arrStrArgsValue {

				for _, value := range str {

					if newLineCounter == (int(valueArgs)-32)*9+i {

						if value == '\n' {
							break
						}
						arrToRet += string(value)

					} else if value == '\n' {
						newLineCounter++
					}
				}
				newLineCounter = 0
			}
			arrToRet += "\n"
		}
	}

	return arrToRet
}
