package main

import (
	"fmt"
	"goTraining/my_go_scripts/list_prime_numbers/src/is_prime_number"
	"log"
	"os"
	"strconv"
)

var Usage = func() {
	fmt.Printf("Usage: %s [postive_integer_num1] [postive_integer_num2]\n", os.Args[0])
}

func main() {
	if os.Args == nil || len(os.Args) != 3 {
		Usage()
		return
	}

	numStart, errStart := strconv.Atoi(os.Args[1])
	numEnd, errEnd := strconv.Atoi(os.Args[2])

	if errStart != nil || errEnd != nil {
		log.Fatalf("输入的参数有误，必须是两个正整数。\n")
	}

	if numStart < 0 {
		log.Fatalf("当前不支持负整数。\n")
	}

	if numStart > numEnd {
		log.Fatalf("第一个整数必须小于第二个整数。")
	}

	chs := make([]chan int, numEnd-numStart+1)
	i := 0
	for num := numStart; num <= numEnd; num++ {
		chs[i] = make(chan int)
		go is_prime_number.IsPrimeNumber(chs[i], num)
		i++
	}

	var primeNumbers []int

	for _, ch := range chs {
		tmp := <-ch
		if tmp != 0 {
			primeNumbers = append(primeNumbers, tmp)
		}
	}

	fmt.Println(primeNumbers)
}

/*
# time go run list_prime_numbers.go 1 1000000	# 执行命令所花费的时间
real    5m28.729s
user    5m27.186s
sys     0m1.216s
*/

/*
# time go run list_prime_numbers.go 1 1000000	# 改版后的所消耗的时间
real    1m34.656s
user    11m28.946s
sys     0m3.801s
*/
