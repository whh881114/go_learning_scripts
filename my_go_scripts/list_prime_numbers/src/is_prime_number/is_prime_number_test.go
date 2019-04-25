package is_prime_number

import (
	"fmt"
	"testing"
)

func TestIsPrimeNumber(t *testing.T) {
	chs := make([]chan int, 2)
	r, s := IsPrimeNumber(chs[0], 5)
	if s == true {
		t.Logf("测试实例正确：输入的是%d，返回状态是%v, 它是素数，与期望相同。", r, s)
	} else {
		t.Errorf("测试实例失败：输入的是%d，返回状态是%v，它本来是素数，测试结果为非素数。", r, s)
	}

	r, s = IsPrimeNumber(chs[1], 10)
	if s == false {
		t.Logf("测试实例正确：输入的是%d，返回状态是%v, 它是不是素数，与期望相同。", r, s)
	} else {
		t.Errorf("测试实例失败：输入的是%d，返回状态是%v，它本来不是素数，测试结果为素数。", r, s)
	}

	for _, ch := range chs {
		fmt.Println(<-ch)
	}
}

/*

[Kanon@MacBookPro  ~/go/src/goTraining/my_go_scripts/list_prime_numbers/src/is_prime_number 23:25]$ 228>go test -v -run=IsPrimeNumber
=== RUN   TestIsPrimeNumber
--- PASS: TestIsPrimeNumber (0.00s)
    is_prime_number_test.go:10: 测试实例正确：输入的是5，返回状态是true, 它是素数，与期望相同。
    is_prime_number_test.go:17: 测试实例正确：输入的是10，返回状态是false, 它是不是素数，与期望相同。
PASS
ok      goTraining/my_go_scripts/list_prime_numbers/src/is_prime_number 0.005s
[Kanon@MacBookPro  ~/go/src/goTraining/my_go_scripts/list_prime_numbers/src/is_prime_number 23:26]$ 229>

*/
