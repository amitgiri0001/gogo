package quiz

func main() {
	c := make(chan string)
	c <- "Hi there!" // Will throw -> fatal error: all goroutines are asleep - deadlock!
	// There is no goroutine to handle the channel response
}
