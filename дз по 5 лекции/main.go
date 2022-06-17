package main

import "fmt"

const (
	priceApple float32 = 5.99
	pricePear  float32  = 7.0
	myMoney    float32  = 26.0
)

func main() {
	amountApple:= 9 
	applesCost1:=float32(priceApple)*float32(amountApple)
	
	amountPear:= 8 
	pearsCost1:=float32(pricePear)*float32(amountPear)

	fmt.Println("стоимость заданного количества яблок -", applesCost1)
	fmt.Println("стоимость заданного количества груш -", pearsCost1)
	
	fmt.Println("сколько нужно потратить денег, чтобы купить заданное количество яблок и груш?-", applesCost1+pearsCost1)
	
	fmt.Println("сколько груш мы можем купить?-", myMoney/pricePear)
	
	fmt.Println("сколько яблок мы можем купить?-", myMoney/priceApple)

	amountApple= 2
	applesCost2:=float32(priceApple)*float32(amountApple)
	amountPear=2 
	pearsCost2 := float32(pricePear) * float32(amountPear)
	fmt.Println("сколько стоит заданное количество яблок?-", applesCost2)
	fmt.Println("сколько стоит заданное количество груш?-", pearsCost2)

	sum :=applesCost2 + pearsCost2
	ourPurchase := myMoney >= sum
	fmt.Println("можем ли мы купить заданное количество яблок и груш?-", ourPurchase)
 }
