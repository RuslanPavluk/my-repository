package main

import "fmt"

const (
	priceApple float64 = 5.99
	pricePear  float64 = 7.0
	myMoney    float64 = 26.0
)

func main() {
	amountApple := 9
	applesCost := priceApple * float64(amountApple)

	amountPear := 8
	pearsCost := pricePear * float64(amountPear)

	fmt.Println("стоимость заданного количества яблок -", applesCost)
	fmt.Println("стоимость заданного количества груш -", pearsCost)

	fmt.Println("сколько нужно потратить денег, чтобы купить заданное количество яблок и груш?-", applesCost+pearsCost)

	applesCanBuy := myMoney / priceApple
	pearsCanBuy := myMoney / pricePear
	fmt.Println("сколько груш мы можем купить?-", int(applesCanBuy))
	fmt.Println("сколько яблок мы можем купить?-", int(pearsCanBuy))

	amountApple = 2
	applesCost = priceApple * float64(amountApple)
	amountPear = 2
	pearsCost = pricePear * float64(amountPear)
	fmt.Println("сколько стоит заданное количество яблок?-", applesCost)
	fmt.Println("сколько стоит заданное количество груш?-", pearsCost)

	sum := applesCost + pearsCost
	ourPurchase := myMoney >= sum
	fmt.Println("можем ли мы купить заданное количество яблок и груш?-", ourPurchase)
}
