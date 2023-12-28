package concurrency

func Example() {
	//println(runtime.NumCPU())
	//defer printNumbers()
	//println("Exit")
	deferValues()
}

func printNumbers() {
	for i := 0; i < 100; i++ {
		println(i)
	}
}

func deferValues() {
	for i := 0; i < 10; i++ {
		defer println("first", i)
	}

	// WRONG!
	//for i := 0; i < 10; i++ {
	//	defer func() {
	//		println("second", i)
	//	}()
	//}
}
