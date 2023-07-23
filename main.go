package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
)

type WordFreq struct {
	word  string
	count int
}

func main() {
	//read text from file
	bs, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(bs)

	//split the text into words
	reg := regexp.MustCompile("[a-zA-Z']+")
	matches := reg.FindAllString(text, -1)

	words := make(map[string]int)
	for _, match := range matches {
		words[match]++
	}

	var wordFreqs []WordFreq
	for k, v := range words {
		wordFreqs = append(wordFreqs, WordFreq{k, v})
	}

	sort.Slice(wordFreqs, func(i, j int) bool { return wordFreqs[i].count > wordFreqs[j].count })

	for i := 0; i < 10 && i < len(wordFreqs); i++ {
		fmt.Println(wordFreqs[i])
	}

}
