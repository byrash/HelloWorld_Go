package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#ff0000",
		"blue":  "#ff0000",
	}
	// var colors map[string]string
	// colors := make(map[string]string)
	colors["white"] = "#ff0000"
	// fmt.Println(colors["white"])
	delete(colors, "white")
	// fmt.Println(colors)
	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("Key [", key, "] --> Value [", value, "]")
	}
}
