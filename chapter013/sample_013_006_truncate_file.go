package main

import (
	"log"
	"os"
)

func main() {
	err := os.Truncate("/tmp/test2.txt", 100)
	if err != nil {
		log.Fatal(err)
	}
}

// 裁剪一个文件到100字节，如果文件本来就少于100字节，则原始内容得以保留，剩余的字节以null字节填充。
// 如果超过100字节，则超过的字切会被抛弃。

/*
[Kanon@MacBookPro  /tmp 14:58]$ 5>more /tmp/test2.txt
This is a test.
^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@^@
[Kanon@MacBookPro  /tmp 14:58]$ 6>
[Kanon@MacBookPro  /tmp 14:58]$ 6>ll /tmp/test2.txt
-rw-r--r--  1 Kanon  wheel  100  4 22 14:57 /tmp/test2.txt
[Kanon@MacBookPro  /tmp 14:58]$ 7>
*/
