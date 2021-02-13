package split

import (
	"reflect"
	"testing"
)

// 定义一个切割字符串的包
// a:b:cd --> [a, b, cd]

func TestSplit(t *testing.T) {
	got := Split("a:b:cd", ":")
	want := []string{"a", "b", "cd"}
	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("期望得到：%v，实际得到：%v\n", want, got)
	}

	got = Split("a:b:cd:x;;y;", ":")
	want = []string{"a", "b", "cd", "x;;y;"}
	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("期望得到：%v，实际得到：%v\n", want, got)
	}

	got = Split("你好，世界！", "，")
	want = []string{"你好", "世界！"}
	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("期望得到：%v，实际得到：%v\n", want, got)
	}

	got = Split("xabyabz;", "ab")
	want = []string{"x", "y", "z;"}
	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("期望得到：%v，实际得到：%v\n", want, got)
	}
}
