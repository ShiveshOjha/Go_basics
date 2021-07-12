package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	file, _ := os.Open("text.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var str string = ""
	fin := make(map[string]int)
	c := 0

	// removing special characters and numbers

	reg, _ := regexp.Compile("[^a-zA-Z]+")

	for scanner.Scan() {
		str = str + reg.ReplaceAllString(scanner.Text(), "") + " "
	}

	st := strings.Fields(str)

	// counting

	for i := 0; i < len(st); i++ {

		for j := 0; j < len(st); j++ {

			if strings.ToLower(st[i]) == strings.ToLower(st[j]) {
				c++
			}
		}

		if i < len(st)-1 {
			if st[i] != st[i+1] {
				fin[st[i]] = int(c)
				// fmt.Printf("%v-%v\n", st[i], c)
			}
		}
		if i == len(st)-1 {
			fin[st[i]] = int(c)
			// fmt.Printf("%v-%v\n", st[i], c)
		}
		c = 0

	}

	fmt.Println(fin)

}
