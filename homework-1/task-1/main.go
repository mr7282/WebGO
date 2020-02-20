package main

import (
	"fmt"
)

func main() {
	a := GetRequest()
	b := GetReferenceSite()
	fmt.Println(a)
	fmt.Println(b)
}

//GetRequest - получает строку поискового запроса
func GetRequest() string {
	var StrRequest string
	fmt.Println("Введите поисковый запрос:")
	fmt.Scan(&StrRequest)
	return StrRequest
}

//GetReferenceSite - получает сайты на которых необходимо искать
func GetReferenceSite() []string {
	var ArrReference []string
	var GetSingleRef string
	fmt.Println("введите названия сайтов на которых поискать")
	fmt.Scanln(&GetSingleRef)
	ArrReference = append(ArrReference, GetSingleRef)
	// 		for i := 0; i > 3; i++ {}
	return ArrReference

}
