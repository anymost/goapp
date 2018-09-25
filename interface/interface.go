package main

import "fmt"

/**
	Golang中的面向对象
	1、封装：通过可见性原则来保证
    2、继承：基于组合来实现。内嵌结构体，内嵌结构体
    3、多态：基于接口来实现。松耦合的正交，隐式实现接口，鸭式辩形
 */

 type Seller interface {
	sell() string
 }

 type BusinessMan interface {
 	Seller
 	purchase() string
 }

 type FruitSeller struct {
 	fruitType string
 }

 type ShoeSeller struct {
 	shoeType string
}

func main()  {
		sellers := []Seller{FruitSeller{"apple"}, &ShoeSeller{"nike"}}
		for _, seller := range sellers{
			fmt.Println(seller.sell())
		}
}
func (fruit FruitSeller) sell() string {
	return fruit.fruitType
}

func (shoe *ShoeSeller) sell() string {
	return shoe.shoeType
}

