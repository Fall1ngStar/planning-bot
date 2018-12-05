package main

import "io/ioutil"

func main () {
	raw, err := ioutil.ReadFile("planning")
	if err != nil {
		panic(err)
	}
	data := string(raw)
	print(data)
}