package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func main() {
	s := []string{"ドド", "スコ"}
	var res []string
	ans := []string{"ドド", "スコ", "スコ", "スコ", "ドド", "スコ", "スコ", "スコ", "ドド", "スコ", "スコ", "スコ"}
	for i := 0; ; i++ {
		out := s[rand.Intn(2)]
		fmt.Print(out)
		res = append(res, out)
		if i >= 12 {
			if reflect.DeepEqual(res[i-12:i], ans) {
				fmt.Println("ラブ注入")
				break
			}

		}
	}
}
