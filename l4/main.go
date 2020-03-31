//Домашнее задание
//Частотный анализ
//Написать функцию, которая получает на вход текст и возвращает
//10 самых часто встречающихся слов без учета словоформ

package main

import (
	"fmt"
	"sort"
	"strings"
	"regexp"
)

type kv struct {
	Key   string
	Value int
}

func GetWordFrequency(string_ string, count int) []kv {
	re := regexp.MustCompile(`[.,—\s]`)
	slittedWords := re.Split(string_, -1)
	result := map[string]int{}

	for _, word := range slittedWords {
		if word != "" {
			result[strings.ToLower(word)] += 1
		}
	}

	var ss []kv
	for k, v := range result {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss[0:count]
}

func main() {
	string_ := "Даже шею, даже уши ты испачкал в черной туши. Становись скорей под душ. Смой с ушей под душем тушь. Смой и с шеи тушь под душем. После душа Вытрись суше. Шею суше, суше уши — и не пачкай больше уши."
	result := GetWordFrequency(string_, 10)
	fmt.Println(result)
}
