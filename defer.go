package main

import "fmt"

var isConnected bool = false

func main() {
	fmt.Printf("Connection open:%v\n", isConnected)
	databaseProcesing()
	fmt.Printf("Connection open:%v\n", isConnected)

}

func databaseProcesing() {
	connect()
	fmt.Println("Defering disconnect")
	defer disconnect()
	fmt.Printf("connection open :%v\n", isConnected)
	fmt.Println("Doing something")

}

func connect() {
	isConnected = true
	fmt.Println("Connected to database!")

}

func disconnect() {
	isConnected = false
	fmt.Println("Disconnected!")

}
