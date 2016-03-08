package main

import (
	"fmt"
	"time"
	//"math/rand"
	"os"
)

func main() {
	fmt.Println("Started the Crashinator™")
	//randor := rand.New(rand.NewSource(time.Now().UnixNano()))
	//dozetime := randor.Intn(5)
	dozetime := 11
	fmt.Printf("Waiting for %d seconds\n", dozetime)
	time.Sleep(time.Duration(dozetime) * time.Second)
	fmt.Println("CRASHING")
	os.Exit(100)
}


