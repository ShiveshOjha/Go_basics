package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter ID: ")
	id, _ := reader.ReadString('\n')

	// regex
	reg1, _ := regexp.Compile("^[1-9]{2}[A-Z]{3}[0-9]{4}")
	fmt.Println(reg1.MatchString(id))

	fmt.Println("Enter PAN number: ")
	pan, _ := reader.ReadString('\n')

	reg2, _ := regexp.Compile("^[A-Z]{5}[0-9]{4}[A-Z]{1}")
	fmt.Println(reg2.MatchString(pan))

	// fmt.Println("Enter sentence: ")
	// s, _ := reader.ReadString('\n')

	// re, _ := regexp.Compile("^[A-Z]{1}[\w+\s+\w+]+")
	// fmt.Println(re.ReplaceAllString(s, " "))
	// fmt.Println(re.MatchString(s))

}
