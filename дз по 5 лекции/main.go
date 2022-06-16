package main

import "fmt"

var (
	priceApple float32 = 5.99
	pricePear  float32  = 7.0
	myMoney    float32  = 26.0
)

func main() {
	//задаем переменную для количества яблок
	var amountApple int = 9.0
	//задаем переменную "цена заданного количества яблок". для этого цену одного яблока умножаем на количество 
	priceAnyApple:=float32(priceApple)*float32(amountApple)
	
	//задаем переменную для количества груш
	var amountPear int = 8.0
	//задаем переменную "цена заданного количества груш". для этого цену одной груши умножаем на количество 
	priceAnyPear:=float32(pricePear)*float32(amountPear)

	fmt.Println("стоимость заданного количества яблок -", priceAnyApple)
	fmt.Println("стоимость заданного количества груш -", priceAnyPear)
	
	fmt.Println("сколько нужно потратить денег, чтобы купить заданное количество яблок и груш?-", priceAnyApple+priceAnyPear)
	
	fmt.Println("сколько груш мы можем купить?-", int(myMoney)/int(pricePear))
	
	fmt.Println("сколько яблок мы можем купить?-", int(myMoney)/int(priceApple))

	//задаем переменную для количества яблок
	var quantityApple int = 2.0
	//задаем переменную "цена заданного количества яблок". для этого цену одного яблока умножаем на количество 
	priceAnyQuantityApple := float32(priceApple) * float32(quantityApple)
	
	//задаем переменую для количества груш
	var quantityPear int = 2.0
	//задаем переменную "цена заданного количества груш". для этого цену одной груши умножаем на количество 
	priceAnyQuantityPear := float32(pricePear) * float32(quantityPear)
		
	fmt.Println("сколько стоит два яблока?-", priceAnyQuantityApple)
	
	fmt.Println("сколько стоит две груши?-", priceAnyQuantityPear)

	//нам нужна суммарная стоимость 2 яблок и 2 груш. для этого складываем стоимость 2 яблок и стоимость 2 груш
	sum := priceAnyQuantityApple + priceAnyQuantityPear
	
	//нам нужно понимать, можем ли мы купить, за указанную у нас сумму денег, 2 яблока и две груши. для этого мы сравниваем нашу указанную сумму денег и сумму стоимости 2 яблок и 2 груш. если наше условие выполняется, то вернется "true". если же нет-вернется "falce"
	ourPurchase := myMoney >= sum
	fmt.Println("можем ли мы купить два яблока и две груши?-", ourPurchase)
}
