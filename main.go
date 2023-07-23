package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
)

type WordFreq struct {
	word string
	freq int
}

func (p WordFreq) String() string {
	return fmt.Sprintf("%s %d", p.word, p.freq)
}

func main() {
	filePath := "project\\word_count\\gen_ai.txt" 
	text, err := readFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	words := countWords(text)
	wordFreqs := convertToWordFreqs(words)
	sortWordFreqs(wordFreqs)

	printTopWords(wordFreqs, 10)
}

func readFile(filePath string) (string, error) {
	bs, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func countWords(text string) map[string]int {
	reg := regexp.MustCompile("[a-zA-Z']+")
	matches := reg.FindAllString(text, -1)

	words := make(map[string]int)
	for _, match := range matches {
		words[match]++
	}
	return words
}

func convertToWordFreqs(words map[string]int) []WordFreq {
	var wordFreqs []WordFreq
	for k, v := range words {
		wordFreqs = append(wordFreqs, WordFreq{k, v})
	}
	return wordFreqs
}

func sortWordFreqs(wordFreqs []WordFreq) {
	sort.Slice(wordFreqs, func(i, j int) bool {
		return wordFreqs[i].freq > wordFreqs[j].freq
	})
}

func printTopWords(wordFreqs []WordFreq, n int) {
	for i := 0; i < n && i < len(wordFreqs); i++ {
		fmt.Println(wordFreqs[i])
	}
}
