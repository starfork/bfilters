package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func main() {
	//str := "08234С我是罗德"
	re := regexp.MustCompile("[\u4e00-\u9fa50-9a-zA-Z]+")
	//ms := re.FindAllString(str, -1)
	//fmt.Println(ms)
	lines, _ := ReadLines("./txt.txt")

	//rs := ""
	tf, _ := os.OpenFile("rs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	for _, v := range lines {
		tmp := re.FindAllString(v, -1)
		//fmt.Println(tmp)
		if len(tmp) > 0 {
			tf.WriteString(strings.Join(tmp, "") + "\n")
		}
	}

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
