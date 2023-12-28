package concurrency

func PanicExample() {
	makePanic()
}

func makePanic() {
	defer func() {
		panicValue := recover()
		println(panicValue)
	}()
	panic("some panic")
	//println("unreachable code")
}
