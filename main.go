package main

import "github.com/TheAlgorithms/Go/strings/genetic"

func main() {
	rune := []rune("Hello, World!")
	res, _ := genetic.GeneticString("Hello, World!", rune, &genetic.Conf{})
	println(res)
}
