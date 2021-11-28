package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 6 {
		fmt.Println(os.Args)
		fmt.Println("\n-->hexer by Marc Teufel 2021<--\n\nParameters are missing!\n\nStart hexer:\nhexer [columns] [rows] [start] [offset] [digitsToFormat]\n\nExample:\nhexer 40 25 0x0400 1 4")
		os.Exit(1)
	}

	columns, err := strconv.ParseInt(os.Args[1], 10, 0)

	if err != nil {
		fmt.Println("First input parameter (columns) must be integer")
		os.Exit(1)
	}

	rows, err := strconv.ParseInt(os.Args[2], 10, 0)

	if err != nil {
		fmt.Println("Second input parameter (rows) must be integer")
		os.Exit(1)
	}

	start, err := strconv.ParseInt(os.Args[3], 16, 0)

	if err != nil {
		fmt.Println("Third input parameter (start) must be integer in hex (eg. 0a00)")
		os.Exit(1)
	}

	offset, err := strconv.ParseInt(os.Args[4], 10, 0)

	if err != nil {
		fmt.Println("4th input parameter (offset) must be integer dec")
		os.Exit(1)
	}

	digitsToPrintArg, err := strconv.ParseInt(os.Args[5], 10, 0)
	var digitsToPrint int

	if err != nil {
		fmt.Println("5th input parameter (format) must be integer dec and defines the length of the formatted string (filled with 0)")
		os.Exit(1)
	} else {
		digitsToPrint = int(digitsToPrintArg)
	}

	chars := make([]string, columns+1, 100)
	var c int64 = 1
	var r int64 = 0
	var idx int64 = 0

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)

	for r <= rows {

		if r == 0 {
			chars[0] = ""
		} else {
			chars[0] = strconv.FormatInt(r, 10)
		}

		idx++
		c = 1
		for c <= columns {
			if r == 0 {
				chars[idx] = strconv.FormatInt(c, 10)
			} else {
				chars[idx] = PadLeft(strconv.FormatInt(start, 16), "0", digitsToPrint-len(strconv.FormatInt(start, 16)))
				start = start + offset
			}

			idx++
			c++
		}

		t.AppendRow(CreateRow(chars))
		if r == 0 {
			t.AppendSeparator()
		}
		r++
		idx = 0

	}

	t.Render()

}

func CreateRow(chars []string) []interface{} {
	var new []interface{} = make([]interface{}, len(chars))
	for i, v := range chars {
		new[i] = v
	}
	return new
}

func PadLeft(s, p string, count int) string {
	ret := make([]byte, len(p)*count+len(s))

	b := ret[:len(p)*count]
	bp := copy(b, p)
	for bp < len(b) {
		copy(b[bp:], b[:bp])
		bp *= 2
	}
	copy(ret[len(b):], s)
	return string(ret)
}
