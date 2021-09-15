package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("FeatureA")
	fmt.Println("Hello World")
	fmt.Println("Hello world2")
	fmt.Println("Hello World3")
	fmt.Println("Hello WorldXX")
	fmt.Println("Hello World4")
	fmt.Println("Hello World5")
	fmt.Println("Hello World6")
	fmt.Println("Hello World7")
	fmt.Println(funWithAnagrams([]string{"code", "aaagmnrs", "anagrams", "doce"}))
	fmt.Println(funWithAnagrams([]string{"poke", "pkoe", "okpe", "ekop"}))
}

func funWithAnagrams(text []string) []string {
	// Write your code here
	var mapData = []string{}

	for _, v := range text {
		if !checkIsAnagram(mapData, v) {
			mapData = append(mapData, v)
		}
	}

	sort.Strings(mapData)
	return mapData
}

func checkIsAnagram(text []string, key string) bool {
	var mapData = map[string]int{}

	exist := false

	for _, v := range text {
		mapData = map[string]int{}

		if len(v) != len(key) {
			continue
		}

		for _, char := range v {
			mapData[string(char)]++
		}

		for _, charKey := range v {
			_, exist = mapData[string(charKey)]
			if exist {
				mapData[string(charKey)]--
			}
		}

		for _, valueMapData := range mapData {
			if valueMapData > 0 {
				return false
			}
		}

		return true
	}
	return false
}
