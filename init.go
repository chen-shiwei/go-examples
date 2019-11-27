package main

var f = func(a int) {
	print("x")
}

func main() {

	f := func(i int) {
		print(i)
		if i > 0 {
			f(i - 1)
		}
	}
	f(10) // 10x

}
