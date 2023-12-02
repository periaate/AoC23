package main

import (
	"aoc/lib/util"
	"aoc/two"
	"fmt"
)

func main() {
	res := two.SolvePartTwo(util.ReadLines(two.Fp))
	fmt.Println(res)
	res = two.SolvePartOne(util.ReadLines(two.Fp))
	fmt.Println(res)
}

// var sum int

// var wg = sync.WaitGroup{}

// func Adder(ch chan int) {
// 	for {
// 		sum += <-ch
// 		wg.Done()
// 	}
// }

// const workerCount = 12

// func Worker(sumCh chan int, ch chan []string) {
// 	for {
// 		input := <-ch
// 		sumCh <- two.SolvePartTwo(input)
// 	}
// }

// func main() {
// 	input := util.ReadLines(two.Fp)

// 	sumCh := make(chan int)
// 	workCh := make(chan []string)

// 	for i := workerCount; i > 0; i-- {
// 		go Worker(sumCh, workCh)
// 	}

// 	go Adder(sumCh)

// 	var size int64 = 20_000_000
// 	times := int(size) / len(input)

// 	start := time.Now()

// 	for i := 0; i < times; i++ {
// 		wg.Add(1)
// 		go func() {
// 			workCh <- input
// 		}()

// 	}

// 	wg.Wait()
// 	t := time.Now()
// 	elapsed := t.Sub(start)

// 	perItem := elapsed.Nanoseconds() / size

// 	p := message.NewPrinter(language.English)

// 	if sum != two.CorrectPartTwo*times {
// 		log.Fatalln("Unsuccessful")
// 	}

// 	p.Printf("Ran over %v items in %v seconds	===	%v ns per item", number.Decimal(size), number.Decimal(elapsed.Seconds()), number.Decimal(perItem))
// }
