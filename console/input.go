package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	answer := generateNumber(100)
	hint, try := getHint(answer, getTry())
	counter := 1
	for try != 0 {
		fmt.Println(hint, answer)
		hint, try = getHint(answer, getTry())
		counter++
	}
	fmt.Println(hint, "guessed in", counter, "tries")

}

func readLine(prompt ...string) string {
	if len(prompt) > 0 {
		fmt.Print(prompt[0])
	}
	var input string
	fmt.Scanln(&input)
	return input
}

func getTry() int {
	a := readLine(" Guess number 0-100: ")
	c, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return c
}

func getHint(answer, try int) (string, int) {
	if try < 0 {
		return "Incoorect value", -1
	}
	if try > answer {
		return "Too hight", 1
	}
	if try < answer {
		return "Too low", -1
	} else {
		return "Correct!", 0
	}
}
func generateNumber(max int) int {
	return rand.Intn(100)
}
