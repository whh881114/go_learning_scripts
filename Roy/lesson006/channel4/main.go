package main

import (
	"fmt"
	"math/rand"
)

// 生产者消费者模型

// 使用goroutine和channel实现一个简易的生产者消费者模型
// 生产者：产生随机数
// 消费者：计算每个随机数的每个位的数字的和 1234=10

// 1个生产者，20个消费者。

var chanProducer chan *item
var chanConsumer chan *result

type item struct {
	id  int64
	num int64
}

type result struct {
	item *item
	sum  int64
}

func producer(ch chan *item) {
	// 1- 生产随机数。
	var id int64
	for {
		id++
		number := rand.Int63()

		tmp := item{
			id:  id,
			num: number,
		}

		// 2- 把随机数发送到通道中。
		ch <- tmp
	}

}

func calc(num int64) int64 {
	var sum int64
	if num > 0 {
		sum += num % 10
		num = num / 10
	}
	return sum
}

func consumer(ch chan *item, resultChan chan *result) {
	for {
		// 1- 取值。
		tmp := <-ch
		sum := calc(tmp.num)
		retObj := &result{
			item: tmp,
			sum:  sum,
		}
		resultChan <- retObj
	}
}

func printResult(resultChan chan *result) {
	for ret := range resultChan {
		fmt.Printf("ID:%v, NUM:%v, SUM:%V\n", ret.item.id, ret.item.num, ret.sum)
	}
}

func startWorker(n int, ch chan *item, resultChan chan *result) {
	for i := 0; i < n; i++ {
		go consumer(ch, resultChan)
	}
}

func main() {
	itemChan := make(chan *item, 100)
	resultChan := make(chan *result, 100)
	go producer(itemChan)
	go consumer(itemChan, resultChan)

	startWorker(20, itemChan, resultChan)

	// 打印结果
	printResult(resultChan)

	// 随机数产生
	// rand.Seed(time.Now().UnixNano())
	// ret1 := rand.Int63()
	// fmt.Println(ret1)
	// ret2 := rand.Int63()
	// fmt.Println(ret2)
}
