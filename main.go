package main

import (
	"fmt"

	"github.com/Ion-Mobility/can_parser_example/parser"
)

func main() {
	bleMsg := parser.Message{
		BikeID:    "d5efcd9e-ebf9-4b35-adca-c80784a7da43",
		Timestamp: "2023-08-08 11:14:14.23 UTC",
		CanID:     "18edf4fa",
		Length:    5,
		Data:      "2537013596",
	}

	encodedMsg, err := parser.Parse(bleMsg)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Encoded message: %+v\n", encodedMsg)
}
