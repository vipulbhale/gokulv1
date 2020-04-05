package main

import "github.com/vipulbhale/gokulv1/internal"

// Version of the application
var VERSION string = "0.0.1"

func init() {
}

func main() {
	internal.Execute(VERSION)
}
