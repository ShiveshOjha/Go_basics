package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup // for better synchronization
var mut sync.Mutex    // for better resource allocation to avoid any deadlock

func count(n int) {
	for i := 1; i <= n; i++ {
		time.Sleep(time.Second) // to delay execution for better visualization
		fmt.Println(i)
	}
	fmt.Println("")

	defer wg.Done()

}

func webUD(website string) {

	if res, err := http.Get(website); err != nil {
		fmt.Println(website, "is down")
	} else {
		mut.Lock()
		fmt.Printf("[%v] %v is up\n", res.StatusCode, website)
		mut.Unlock()
	}

	defer wg.Done() // decrement 1 to waitGroup for every completion of goroutine(thread) work

}

func main() {

	// executes method without go prefix, simultaneously executes methods with go prefix that have common output with without go prefix method
	// calls count function 2 times to change value of n, that is, 5 and 8
	// execution leading to printing of outputs common to count(8) only
	// if we don't use WaitGroup
	{
		// go count(11)
		// go count(4)
		// go count(5)
		// go count(8)
		// wg.Add(4)

	}

	// concurrency with WaitGroup
	{

		websites := []string{
			"https://stackoverflow.com/",
			"https://github.com/",
			"https://www.linkedi15n.com/",
			"http://medium.com/",
			"https://golang.org/",
			"https://www.udemy.com/",
			"https://www.coursera.org/",
			"https://wesionary.team/",
		}

		for _, webapp := range websites {
			go webUD(webapp)
			wg.Add(1) // for every creation of goroutine(thread) increment 1 to waitGroup
		}

	}

	wg.Wait() //to make goroutine(thread) wait until all goroutine tasks are complete

}
