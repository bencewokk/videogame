package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func input() []int {
	var input string
	var output []int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	inputSplit := strings.Split(input, " ")

	for i := 0; i < len(inputSplit); i++ {
		charInt, _ := strconv.Atoi(inputSplit[i])
		output = append(output, charInt)
	}

	return output
}

func main() {
	input := input()
	var aBreaks, yBreaks []int
	var boretime, boretimeNow, aLenght, yLenght, output int

	fmt.Println(yLenght)

	//input apprehension
	for i := 0; i < len(input); i++ {
		if i == 0 {
			boretime = input[i]
		} else if i == 1 {
			aLenght = input[i] * 2
		} else if i < aLenght+2 {
			aBreaks = append(aBreaks, input[i])
		} else if i == aLenght+2 {
			yLenght = input[i]
		} else {
			yBreaks = append(yBreaks, input[i])
		}

	}

	//chart creation
	var max int

	for i := 1; i < len(yBreaks); i++ {
		if max < yBreaks[i] {
			max = yBreaks[i]
		}
	}

	for i := 1; i < len(aBreaks); i++ {
		if max < aBreaks[i] {
			max = aBreaks[i]
		}
	}
	charta := make([]int, max)
	charty := make([]int, max)

	//fmt.Println(aBreaks, yBreaks)

	/*
		//chart modification (a)
		swi := 1

		swi = 1
		for i := 0; i < len(aBreaks); i++ {

			charta[aBreaks[i]-1] = charta[aBreaks[i]-1] + swi

			if swi == 1 {
				swi = swi - 2
			} else {
				swi = swi + 2
			}

			fmt.Println(charta)

		}

		//chart modification (y)
		swi = 1
		for i := 0; i < len(yBreaks); i++ {

			charty[yBreaks[i]-1] = charty[yBreaks[i]-1] + swi

			if swi == 1 {
				swi = swi - 2
			} else {
				swi = swi + 2
			}

			fmt.Println(charty)
		}

	*/

	//fmt.Println(charty)

	q, u := 1, 0
	for {

		if q > len(aBreaks) {
			break
		}

		charta[aBreaks[q]-u-1] = 1

		if aBreaks[q]-u == aBreaks[q-1] {
			u = 0
			q = q + 2
		} else {
			u++
		}
		//fmt.Println(charta)
	}

	q, u = 1, 0
	for {
		if q > len(yBreaks) {
			break
		}

		charty[yBreaks[q]-u-1] = 1

		if yBreaks[q]-u == yBreaks[q-1] {
			u = 0
			q = q + 2
		} else {
			u++
		}
		//fmt.Println(charty)
	}

	var index int

	//fmt.Println(charta)
	//fmt.Println(charty)

	for i := 0; i < len(charta); i++ {
		if boretime == boretimeNow {
			index = i
			for charta[index] != 0 {
				charta[index] = 0
				index++
			}
			boretimeNow = 0
		}

		if charta[i] == 1 && charta[i] != charty[i] {
			boretimeNow++
		}

	}

	for i := 0; i < len(charta); i++ {
		if charta[i] == 1 && charta[i] == charty[i] {
			output++
		}
	}

	fmt.Println(charta)
	fmt.Println(charty)

	fmt.Println(output)
}
