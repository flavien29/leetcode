package main

import fmt "fmt"
import (
	"bytes"
)

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	rows := make([][]rune, numRows)
	row, pos := 0, 0
	for _, c := range s {
		rows[row] = append(rows[row], c)
		if row == 0 {
			pos = 1
		} else if row == numRows-1 {
			pos = -1
		}
		row += pos
	}

	var out bytes.Buffer
	for _, r := range rows {
		out.WriteString(string(r))
	}
	return out.String()
}

func main() {
	var str string = "PAYPALISHIRING"
	var res string = convert(str, 3)

	fmt.Printf("%v", res)
}
