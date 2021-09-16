package asynchronous

func RunTasks() {
	ch := make(chan int, 3)

	coNum := 2
	done := make(chan bool, coNum)
	for i := 1; i <= coNum; i++ {
		go Consumer(i, ch, done)
	}

	go Producer(ch)
	for i := 1; i <= coNum; i++ {
		<-done
	}
}
