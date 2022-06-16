package main

import "fmt"

const (
	priceApple float32 = 5.99
	pricePear  float32  = 7.0
	myMoney    float32  = 26.0
)

func main() {

	fmt.Println("стоимость 9 яблок-", priceApple*9)
	fmt.Println("стоимость 8 груш-", pricePear*8)

	//задаем переменную для количества яблок
	var amountApple float32 = 9.0
	//задаем переменную "цена 9 яблок". для этого цену одного яблока умножаем на количество 
	price9apple:=priceApple*amountApple

	//задаем переменную для количества груш
	var amountPear float32 = 8.0
	//задаем переменную "цена 8 груш". для этого цену одной груши умножаем на количество 
	price8pear:=pricePear*amountPear

	
	fmt.Println("сколько нужно потратить денег, чтобы купить 9 яблок и 8 груш?-", price9apple+price8pear)
	
	
	fmt.Println("сколько груш мы можем купить?-", myMoney/pricePear)
	
	
	fmt.Println("сколько яблок мы можем купить?-", myMoney/priceApple)


	//задаем переменную для количества яблок
	var quantityApple float32 = 2.0
	//задаем переменную "цена 2 яблок". для этого цену одного яблока умножаем на количество 
	price2apple := priceApple * quantityApple
	
	//задаем переменую для количества груш
	var quantityPear float32 = 2.0
	//задаем переменную "цена 2 груш". для этого цену одного яблока умножаем на количество 
	price2pear := pricePear * quantityPear
		
	
	
	fmt.Println("сколько стоит два яблока?-", price2apple)
	
	
	fmt.Println("сколько стоит две груши?-", price2pear)



	
	//нам нужна суммарная стоимость 2 яблок и 2 груш. для этого складываем стоимость 2 яблок и стоимость 2 груш
	sum := price2apple + price2pear
	
	
	//нам нужно понимать, можем ли мы купить, за указанную у нас сумму денег, 2 яблока и две груши. для этого мы сравниваем нашу указанную сумму денег и сумму стоимости 2 яблок и 2 груш. если наше условие выполняется, то вернется "true". если же нет-вернется "falce"
	ourPurchase := myMoney >= sum
	fmt.Println("можем ли мы купить два яблока и две груши?-", ourPurchase)
}
