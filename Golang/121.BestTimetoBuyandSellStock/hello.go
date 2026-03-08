package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {

	// Sanity checking
	if len(prices) == 0 || len(prices) == 1{
		return 0
	}

	buy := math.MaxInt64
	buy_index := 0
	sell := math.MinInt64
	sell_index := 0
	profit := math.MinInt64

	for index, price := range prices{

		if price < buy{
			buy = price
			buy_index = index
		}

		if (price >= buy) && (index > buy_index){
			sell = price
			sell_index = index
		}

		if ((sell - buy) > profit) && (sell != math.MinInt64) && (buy != math.MaxInt64) && (sell_index > buy_index){
			profit = sell - buy
		}
		
	}

	if profit < 0{
		return 0
	} else {
		return profit
	}
}

func main(){

	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
	fmt.Println(maxProfit([]int{2,4,1}))
	fmt.Println(maxProfit([]int{2,1,2,1,0,1,2}))
	fmt.Println(maxProfit([]int{3,3,5,0,0,3,1,4}))
}
