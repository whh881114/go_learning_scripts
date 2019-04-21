package main

func main() {
	defer func() {
		panic("defer panic")
	}()

	panic("test panic")
}
