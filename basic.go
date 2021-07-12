package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	result := multiplyme(3, 6)
	fmt.Println(result)

	r2, mylen, myname := adder(3, 5, 4, 8, 9)
	fmt.Println(r2, mylen, myname)

}

func prt() {
	var batman string = "I'm batman"
	fmt.Println(batman)

	check := 3.651631511115
	fmt.Printf("%v, %T, \n", check, check)

	var (
		name   = "Shivesh"
		age    = 21
		skills = "linux, cybersecurity, coding"
	)

	fmt.Println(name, ",", age, "years old,", "skilled in", skills)
}

func Simple_input() {
	var str string
	fmt.Scanln(&str) // only takes single word input
	fmt.Println(str)

}

func R_input() {

	reader := bufio.NewReader(os.Stdin)                                   // like scanner sc declaration
	fmt.Println("Enter your rating: ")                                    // msg before inputing
	myrating, _ := reader.ReadString('\n')                                // input variable
	mynumrating, _ := strconv.ParseFloat(strings.TrimSpace(myrating), 64) // converting input from String to float
	fmt.Println(mynumrating + 2)                                          // printing

}

func aray() {
	var num = [3]string{"one", "two", "three"}
	fmt.Println(num)

	var nm [3]int
	nm[0] = 1
	nm[1] = 2
	nm[2] = 3

	fmt.Println(nm)
}

func point() {
	var p float64 = 99.2
	p_ref := &p // & is used to store memory address of p in p_ref

	fmt.Println(p)
	fmt.Println(p_ref)
	fmt.Println(*p_ref) // * is used to get the value inside the address stored in p_ref
}

func slice() {
	var things = []string{"one", "two", "three"}
	fmt.Println(things)

	things = append(things, "four")
	fmt.Println(things)

	things = append(things[1:])
	fmt.Println(things)

	things = append(things[1 : len(things)-1])
	fmt.Println(things)

	heros := make([]string, 3, 3) // intitialize heros to either slice or map
	heros[0] = "thor"
	heros[1] = "spiderman"
	heros[2] = "ironman"
	fmt.Println(heros)

	heros = append(heros, "krrish")
	fmt.Println(heros)
}

func maps() {

	// new - only allocates -> no init of memory
	// make - allocates and init -> non zeroes storage

	score := make(map[string]int)
	score["shivesh"] = 89
	fmt.Println(score)

	score["hitesh"] = 45
	score["ron"] = 57
	score["jhon"] = 65
	score["don"] = 75

	fmt.Println(score)

	delete(score, "jhon")

	for key, value := range score {
		fmt.Printf("Score of %v is %v\n", key, value)
	}

}

type User struct {
	Name  string
	Email string
	age   int
}

func struc() {
	rob := User{"rob", "rob@lco.dev", 34}
	fmt.Printf("%+v\n", rob)
	fmt.Printf("%+v\n", rob.Email)

	var sam = new(User)
	sam.Name = "sam"
	sam.Email = "sam@lco.dev"
	sam.age = 22
	fmt.Printf("%+v\n", sam)

	var tobby = &User{"tobby", "tobby@lco.dev", 22}
	fmt.Printf("%+v\n", tobby)
}

func loop() {
	start := 0
	things := []string{"alpha", "beta", "gamma", "kappa", "delta"}
	fmt.Println(things)

	for i := 0; i < len(things); i++ {
		fmt.Println(things[i])
	}

	for i := range things {
		fmt.Println(things[i])
	}

	fmt.Println("\nNew\n")

	for start < len(things) {
		fmt.Println(things[start])
		start += 1
	}

	for i, j := range "XabCd" {
		fmt.Printf("The index number of %v is %v\n", j, i)
	}

}

func multiplyme(v1 int, v2 int) int {
	return v1 * v2
}

func adder(values ...int) (int, int, string) {
	sum := 0
	for i := range values {
		sum = sum + values[i]
	}
	length := len(values)
	name := "just for fun"
	return sum, length, name
}
