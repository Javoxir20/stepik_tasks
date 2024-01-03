package main

func main() {
	done := make(chan struct{})
	go func(chan struct{}) {
		work()
		close(done)
	}(done)
	<-done
}
