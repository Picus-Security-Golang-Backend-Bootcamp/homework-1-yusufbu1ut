package helper

import (
	"fmt"
	"strings"
)

func Search(input string, list []string) {
	fmt.Printf("\n'%s' Içeren Kitaplar\n----------------------------------------\n", input)
	for _, v := range list { //kitaplar içerisinde girilen değeri içeren kitaplar çıktı olarak yazdırılır
		if strings.Contains(strings.ToLower(v), strings.ToLower(input)) || strings.Contains(strings.ToUpper(v), strings.ToUpper(input)) {
			//karşılaştırma girilen değer ve kitap adaları küçük,büyük harfe döndürülerek  karşılaştırmalar gerçekleştirilir
			println(v)
		}
	}
}

func List(list []string) { //parametresi slice olan girdi içerisinde gezilerek çıktı verilir
	fmt.Printf("\nKitaplık\n----------------------------------------\n")
	for _, v := range list {
		fmt.Println(v)
	}
}
