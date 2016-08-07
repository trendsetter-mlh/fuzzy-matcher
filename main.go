package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	w := flag.String("word", "", "Word to match")
	flag.Parse()

	r := regexp.MustCompile(".*" + *w + ".*")

	b, _ := ioutil.ReadFile("./master.json")
	markov := make(map[string]struct {
		before interface{} `json:"before"`
		after  interface{} `json:"after"`
	})
	err := json.Unmarshal(b, &markov)
	if err != nil {
		panic(err)
	}

	fmt.Println(markov["('capabilities',)"].after.(map[string]int))

	fmarkov := make(map[string]struct {
		before map[string]int `json:"before"`
		after  map[string]int `json:"after"`
	})
	for k, v := range markov {
		fmarkov[k[2:len(k)-3]] = v
	}

	matched := []string{}
	for k, v := range fmarkov {
		if r.Match([]byte(k)) {
			for k, v := range v.after {
				if v >= 1 {
					matched = append(matched, k)
				}
			}
		}
	}

	for _, v := range matched {
		fmt.Println("Hello" + v)
	}
}
