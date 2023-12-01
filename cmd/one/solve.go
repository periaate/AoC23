// A threaded variant for solving the day one problem.
package main

import (
	"aoc/one"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

var sum int

var wg = sync.WaitGroup{}

func Adder(ch chan int) {
	for {
		sum += <-ch
		wg.Done()
	}
}

func main() {

	if len(os.Args) == 0 {
		if sum := one.Solve(one.GetInput()); sum == 54473 {
			fmt.Println("Success")
			os.Exit(1)
		}
		fmt.Println("Failure")
		os.Exit(0)
	}

	sumCh := make(chan int)

	go Adder(sumCh)

	input := one.GetInput()
	var size int64 = 1_000_000_000
	times := int(size) / len(input)

	start := time.Now()

	for i := 0; i < times; i++ {
		wg.Add(1)
		go func() {
			sumCh <- one.Solve(input)
		}()

	}

	wg.Wait()
	t := time.Now()
	elapsed := t.Sub(start)

	perItem := elapsed.Nanoseconds() / size

	p := message.NewPrinter(language.English)

	if sum != 54473*times {
		log.Fatalln("Unsuccessful")
	}

	p.Printf("Ran over %v items in %v seconds	===	%v ns per item", number.Decimal(size), number.Decimal(elapsed.Seconds()), number.Decimal(perItem))
}
