package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	now := time.Now()

	file, err := os.Open("./txt.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	tf, err := os.OpenFile("rs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	defer tf.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(tf)

	re := regexp.MustCompile("[\u4e00-\u9fa50-9a-zA-Z]+")
	//var wg sync.WaitGroup
	for scanner.Scan() {
		tmp := re.FindAllString(scanner.Text(), -1)
		if len(tmp) > 0 {
			writer.WriteString(strings.Join(tmp, "") + "\n")
		}
	}
	writer.Flush()
	fmt.Println(time.Now().Sub(now))
}
