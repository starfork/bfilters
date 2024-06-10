package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	//str := "08234С我是罗德"
	re := regexp.MustCompile("[\u4e00-\u9fa50-9a-zA-Z]+")
	//ms := re.FindAllString(str, -1)
	//fmt.Println(ms)
	lines, err := ReadLines("./txt.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	//rs := ""
	tf, err := os.OpenFile("rs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	defer tf.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	writer := bufio.NewWriter(tf)
	for _, v := range lines {
		tmp := re.FindAllString(v, -1)
		//fmt.Println(tmp)
		if len(tmp) > 0 {
			writer.WriteString(strings.Join(tmp, "") + "\n")
		}
	}
	writer.Flush()

}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt != "" {
			lines = append(lines, scanner.Text())
		}
	}
	return lines, scanner.Err()
}
