package concurrency

import "fmt"

// mutex и канал работабт одинаково, однако есил есть буфер, может быть разница
func ChannelExample() {
	//nilChannel()
	unbufferedChannel()
}

func nilChannel() {
	var nilChannel chan int
	fmt.Printf("Len: %d Cap: %d \n", len(nilChannel), cap(nilChannel))

	nilChannel <- 1 // запись 1 в канал
	<-nilChannel    // чтение из канала
	fmt.Printf("Len: %d Cap: %d \n", len(nilChannel), cap(nilChannel))
}

func unbufferedChannel() {
	unbufferedChannel := make(chan int)
	fmt.Printf("Len: %d Cap: %d \n", len(unbufferedChannel), cap(unbufferedChannel))

	unbufferedChannel <- 1 // запись 1 в канал
	<-unbufferedChannel    // чтение из канала
	fmt.Printf("Len: %d Cap: %d \n", len(unbufferedChannel), cap(unbufferedChannel))
}
