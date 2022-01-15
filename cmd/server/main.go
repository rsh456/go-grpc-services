package main

import "log"

func Run() error {
	//Responsible for initializing and starting
	//gRPC server
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
