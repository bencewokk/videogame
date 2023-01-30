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
    var aBreaks, yBreaks []int
    var boretime, boretimeNow, aLenght, yLenght, output int
 
    boretime = input()[0]
 
    aLenght = input()[0]
 
    for i := 0; i < aLenght; i++ {
        ln := input()
        aBreaks = append(aBreaks, ln[0])
        aBreaks = append(aBreaks, ln[1])
    }
 
    yLenght = input()[0]
 
    for i := 0; i < yLenght; i++ {
        ln := input()
        yBreaks = append(yBreaks, ln[0])
        yBreaks = append(yBreaks, ln[1])
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
 
    //fmt.Println(charta)
    //fmt.Println(charty)
 
    fmt.Println(output)
}
