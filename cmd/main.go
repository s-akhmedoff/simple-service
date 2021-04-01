package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit
	log.Println("Interrupted")
}
