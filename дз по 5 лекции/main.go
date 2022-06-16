package main

import "fmt"

var (
	priceApple float32 = 5.99
	pricePear  float32  = 7.0
	myMoney    float32  = 26.0
)

func main() {

	fmt.Println("стоимость 9 яблок-", priceApple*9)
	fmt.Println("стоимость 8 груш-", pricePear*8)
	price9apple:=priceApple*9
	price8pear:=pricePear*8
	fmt.Println("сколько нужно потратить денег, чтобы купить 9 яблок и 8 груш?-", price9apple+price8pear)
	fmt.Println("сколько груш мы можем купить?-", myMoney/pricePear)
	fmt.Println("сколько яблок мы можем купить?-", myMoney/priceApple)
	fmt.Println("сколько стоит два яблока?-", priceApple*2)
	fmt.Println("сколько стоит две груши?-", pricePear*2)
	price2apple := priceApple * 2
	price2pear := pricePear * 2
	sum := price2apple + price2pear
	ourPurchase := myMoney >= sum
	fmt.Println("можем ли мы купить два яблока и две груши?-", ourPurchase)
}
