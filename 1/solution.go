package main

import (
    "fmt"
    "log"
    "os"
    "bufio"
    "strconv"
    "sort"
)


func main() {

// start: https://golangdocs.com/golang-read-file-line-by-line

    content, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    fileScanner := bufio.NewScanner(content)
 
    fileScanner.Split(bufio.ScanLines)
    
  var calories []int
  var sum int
    for fileScanner.Scan() {
        line, _ := strconv.Atoi(fileScanner.Text())
    if(line != 0) {
        sum += line
        
    } else {
        calories = append(calories, sum)
        sum = 0   
    }
    
}


fmt.Println("part1: ", largest(calories))
fmt.Println("part2: ", sumArray(top(calories, 3)))
content.Close()
}

func largest(array []int) int {
    var max int
    for _, i := range array {
        if i > max {
            max = i
        }
    }
    return max
}

func top(array []int, i int) []int {
    // anti-pattern to modify original array. But who cares? I have dinner to cook.
    sort.Slice(array, func(i, j int) bool {
		return array[i] > array[j]
	})
    
    return array[0:i]
}

func sumArray(array []int) int {  
    result := 0  
    for _, v := range array {  
     result += v  
    }  
    return result  
   }  

