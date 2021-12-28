package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func parseOrderIds(orderIdsString string) ([]uint64, error) {
	orderIdsSlice := strings.Split(orderIdsString, ",")
	orderIds := make([]uint64, 0)
	for _, orderId := range orderIdsSlice {
		orderIdUint64, err := strconv.ParseUint(orderId, 10, 64)
		if err != nil {
			return nil, err
		}
		orderIds = append(orderIds, orderIdUint64)
	}
	return orderIds, nil
}

func main() {
	fmt.Println("Welcome to Hieu test")
	firstInteger := flag.Int("hieu", 0, "hieu is testing")
	orderIdsString := flag.String("order_ids", "", "list of order ids separated by comma, eg: 1,2,3")
	flag.Parse()
	fmt.Println(*orderIdsString)
	fmt.Printf("%T %v \n", *orderIdsString, *orderIdsString)
	orderIds, err := parseOrderIds(*orderIdsString)
	if err != nil {
		fmt.Println("Fail to parse order ids")
	}
	fmt.Println(*firstInteger)
	fmt.Println(orderIds)
	fmt.Printf("%T \n", orderIds)
}
