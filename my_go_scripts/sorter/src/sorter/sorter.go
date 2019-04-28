package main

import (
	"bufio"
	"flag"
	"fmt"
	"goTraining/my_go_scripts/sorter/src/algorithms/bubblesort"
	"goTraining/my_go_scripts/sorter/src/algorithms/qsort"
	"io"
	"os"
	"strconv"
	"time"
)

// flag包用于快速命令行参数
var infile *string = flag.String("i", "infile", "File contains values for sorting.")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values.")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm.")

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Printf("Failed to open the input file\"%s\".", infile)
	}

	defer file.Close() // defer进行清理工作，释放资源。

	br := bufio.NewReader(file)
	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line. seems unexpected.")
			return
		}

		str := string(line) // 转换字符数组为字符串

		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}

	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Printf("Failed to create the output file\"%s\".", outfile)
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}

	return nil
}

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}

	values, err := readValues(*infile)

	if err != nil {
		fmt.Println(err)
	} else {
		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
		}

		t2 := time.Now()

		fmt.Println("The sorting process costs", t2.Sub(t1), "to complete.")
		writeValues(values, *outfile)
	}
}
