package main

import (
	"fmt"
)

// 使用函数实现一个简单的图书管理系统。
// 每本书有书名，作者，价格，上架信息，用户可以在控制台添加书籍，修改书籍信息，打印所有的书籍列表。
type bookItem struct {
	title   string
	author  string
	price   float32
	publish bool // false表示未上架
}

func createBook(bookItems []bookItem) []bookItem {
	var title string
	var author string
	var price float32
	var publish bool

	fmt.Print("请输入图书标题：")
	fmt.Scanln(&title)
	if title == "" {
		fmt.Println()
		fmt.Printf("【ERROR】输入的图书名为空，则无法添加，已返回顶层菜单。\n")
		fmt.Println()
		return bookItems
	}

	bookInstance := new(bookItem) // 初始化值类型后，得到的是bookInstance的内存地址。
	for _, v := range bookItems {
		if title == v.title {
			fmt.Println()
			fmt.Printf("【ERROR】输入的图书名\"%s\"已存在，无法添加，已返回顶层菜单。\n", title)
			fmt.Println()
			return bookItems
		}
	}

	fmt.Print("请输入图书作者：")
	fmt.Scanln(&author)
	fmt.Print("请输入图书价格：")
	fmt.Scanln(&price)
	fmt.Print("请输入图书上架状态：")
	fmt.Scanln(&publish)

	bookInstance.title = title
	bookInstance.author = author
	bookInstance.price = price
	bookInstance.publish = publish

	bookItems = append(bookItems, *bookInstance)
	fmt.Printf("【INFO】创建图书信息完成。详细信息 —— 图书标题：%s，图书作者：%s，图书价格：%.2f，图书上架状态：%v。\n",
		title, author, price, publish)
	fmt.Println()
	return bookItems
}

func readBook(bookItems []bookItem) {
	if len(bookItems) == 0 {
		fmt.Println("【INFO】查询结果：当前图书馆无任何图书信息。")
		return
	}

	fmt.Println("【INFO】查询结果：")
	for k, v := range bookItems {
		fmt.Printf("图书序号：%03d，图书标题：%s，图书作者：%s，图书价格：%.2f，图书上架状态：%v。\n",
			k, v.title, v.author, v.price, v.publish)
	}
}

func deleteBook(bookItems []bookItem, title string) []bookItem {
	for k, v := range bookItems {
		if title == v.title {
			fmt.Printf("【INFO】删除图书信息完成。详细信息 —— 图书标题：%s，图书作者：%s，图书价格：%.2f，"+
				"图书上架状态：%v。\n", bookItems[k].title, bookItems[k].author, bookItems[k].price, bookItems[k].publish)

			if len(bookItems) == 2 { // 当元素只有两个时
				if bookItems[0].title == title {
					bookItems = bookItems[1:] // 元素存在第一个索引处，返回后一个。
				} else {
					bookItems = bookItems[:1] // 元素存在第二个索引处，返回前一个。
				}
			} else if len(bookItems) >= 3 {
				if k == 0 {
					bookItems = bookItems[1:] // 元素排在第一位
				} else {
					tmpBookItems1 := make([]bookItem, 0, 100) // 申请临时变量来存放被切割的元素。
					tmpBookItems2 := make([]bookItem, 0, 100)
					tmpBookItems1 = append(tmpBookItems1, bookItems[:k]...)
					tmpBookItems2 = append(tmpBookItems2, bookItems[k+1:]...)
					bookItems = append(tmpBookItems1, tmpBookItems2[:]...)
				}
			} else {
				bookItems = make([]bookItem, 0, 100) // 当只存在一个元素时，并且要删除它时，只需重新初始化变量即可。
			}
			return bookItems
		}
	}
	fmt.Printf("【ERROR】删除图书信息失败。详细信息 —— 图书标题：%s，不存在图书管信息管理系统中。\n", title)
	return bookItems
}

func deleteAllBooks(bookItems []bookItem) []bookItem {
	bookItems = make([]bookItem, 0, 100)
	fmt.Println("【INFO】删除结果：当前图书馆所有图书信息已被清空。")
	return bookItems
}

func updateBook(bookItems []bookItem) []bookItem {
	var title string
	var author string
	var price float32
	var publish bool

	fmt.Print("请输入图书标题：")
	fmt.Scanln(&title)
	if title == "" {
		fmt.Println()
		fmt.Printf("【ERROR】输入的图书名为空，则无法修改，已返回顶层菜单。\n")
		fmt.Println()
		return bookItems
	}

	for k, v := range bookItems {
		if title != v.title {
			fmt.Println()
			fmt.Printf("【ERROR】输入的图书名\"%s\"不已存在，无法进行修改，已返回顶层菜单。\n", title)
			fmt.Println()
			return bookItems
		} else {
			fmt.Print("请输入图书作者：")
			fmt.Scanln(&author)
			fmt.Print("请输入图书价格：")
			fmt.Scanln(&price)
			fmt.Print("请输入图书上架状态：")
			fmt.Scanln(&publish)

			bookItems[k].author = author
			bookItems[k].price = price
			bookItems[k].publish = publish
			fmt.Printf("【INFO】修改图书信息完成。详细信息 —— 图书标题：%s，图书作者：%s，图书价格：%.2f，"+
				"图书上架状态：%v。\n", title, author, price, publish)
		}
	}
	return bookItems
}

func main() {
	// 初始化一个切片来存储图书信息，元素类型为创建的struct类型。
	bookItems := make([]bookItem, 0, 100)

	for {
		fmt.Println()
		fmt.Println("========================= 图书管理系统 =========================")
		fmt.Println("1. 查询所有图书详细列表。")
		fmt.Println("2. 新增图书信息。")
		fmt.Println("3. 删除图书信息（按书名删除）。")
		fmt.Println("4. 删除所有图书信息（慎用）。")
		fmt.Println("5. 修改图书信息（按书名修改，其书名可以被修改）。")
		fmt.Println("6. 退出系统。")

		var choice int
		fmt.Print("-- 输入标题对应的数字进入相应的菜单：")
		fmt.Scanln(&choice)
		fmt.Println()

		switch choice {
		case 1:
			readBook(bookItems)
		case 2:
			bookItems = createBook(bookItems)
		case 3:
			var title string
			fmt.Print("请输入要删除的图书标题：")
			fmt.Scanln(&title)
			bookItems = deleteBook(bookItems, title)
		case 4:
			bookItems = deleteAllBooks(bookItems)
		case 5:
			bookItems = updateBook(bookItems)
		case 6:
			return
		}
	}
}
