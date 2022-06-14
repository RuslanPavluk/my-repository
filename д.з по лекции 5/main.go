package main

import "fmt"

var (
	priceApple float32 = 5.99
	pricePear  int     = 7.00
	myMoney    int     = 23.00
)

func main() {

	fmt.Println("стоимость 9 яблок-", priceApple*9)
	fmt.Println("стоимость 8 груш-", pricePear*8)
	fmt.Println("сколько нужно потратить денег, чтобы купить 9 яблок и 8 груш?-", float32(priceApple)*9+float32(pricePear*8))
	fmt.Println("сколько груш мы можем купить?-", myMoney/pricePear)
	fmt.Println("сколько яблок мы можем купить?-", int(myMoney)/int(priceApple))
	fmt.Println("сколько стоит два яблока?-", priceApple*2)
	fmt.Println("сколько стоит две груши?-", pricePear*2)
	price2apple := priceApple * 2
	price2pear := pricePear * 2
	sum := float32(price2apple) + float32(price2pear)
	ourPurchase := float32(myMoney) > float32(sum)
	fmt.Println("можем ли мы купить два яблока и две груши?-", ourPurchase)
}
