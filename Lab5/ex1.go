package main

import (
	"fmt"
	"math"
	"strconv"
)

type Currency struct {
	name   string
	exRate float64
}

func (p *Currency) SetName(name string) {
	p.name = name
}
func (p Currency) Name() string {
	return p.name
}
func (p *Currency) SetExRate(exRate float64) {
	if exRate > 0 {
		p.exRate = exRate
	}
}
func (p Currency) ExRate() float64 {
	return p.exRate
}

type Product struct {
	name     string
	price    float64
	cost     Currency
	quantity int
	producer string
	weight   float64
}

func (p *Product) SetName(name string) {
	p.name = name
}
func (p Product) Name() string {
	return p.name
}
func (p *Product) SetPrice(price float64) bool {
	if price > 0 {
		p.price = price
		return true
	}
	return false
}
func (p Product) Price() float64 {
	return p.price
}
func (p *Product) SetCost(cost Currency) {
	p.cost = cost
}
func (p Product) Cost() Currency {
	return p.cost
}
func (p *Product) SetQuantity(quantity int) bool {
	if quantity > 0 {
		p.quantity = quantity
		return true
	}
	return false
}
func (p Product) Quantity() int {
	return p.quantity
}
func (p *Product) SetProducer(producer string) {
	p.producer = producer
}
func (p Product) Producer() string {
	return p.producer
}
func (p *Product) SetWeight(weight float64) bool {
	if weight > 0 {
		p.weight = weight
		return true
	}
	return false
}
func (p Product) Weight() float64 {
	return p.weight
}
func (p Product) GetPriceIn() float64 {

	return p.Price()
}
func (p Product) GetTotalPrice() float64 {
	return p.Price() * float64(p.Quantity())
}
func (p Product) GetTotalWeight() float64 {
	return float64(p.Quantity()) * p.weight
}

func ReadProductsArray() []Product {
	var products []Product
	for true {
		product := Product{}
		fmt.Print("Товар №", len(products)+1, ":\n")
		fmt.Print("Введіть назву: ")

		fmt.Scan(&product.name)
		var temp string
		for true {
			fmt.Print("Введіть ціну: ")
			fmt.Scan(&temp)
			if s, err := strconv.ParseFloat(temp, 32); err == nil {
				product.SetPrice(s)
				break
			} else {
				fmt.Print("Помилка\n")
			}
		}
		fmt.Print("Введіть валюту: ")

		fmt.Scan(&product.cost.name)
		for true {
			fmt.Print("Введіть курс: ")
			fmt.Scan(&temp)
			if s, err := strconv.ParseFloat(temp, 32); err == nil {
				product.cost.SetExRate(s)
				break
			} else {
				fmt.Print("Помилка\n")
			}
		}

		for true {
			fmt.Print("Введіть к-сть товару: ")
			fmt.Scan(&temp)
			if s, err := strconv.ParseInt(temp, 10, 32); err == nil {
				product.SetQuantity(int(s))
				break
			} else {
				fmt.Print("Помилка\n")
			}
		}

		fmt.Print("Введіть назву компанії-виробника: ")
		fmt.Scan(&product.producer)

		for true {
			fmt.Print("Введіть вагу одиниці товару: ")
			fmt.Scan(&temp)
			if s, err := strconv.ParseFloat(temp, 32); err == nil {
				product.SetWeight(s)
				break
			} else {
				fmt.Print("Помилка\n")
			}
		}

		products = append(products, product)
		fmt.Print("Додати ще?\n")
		fmt.Print("1.Так\n2.Ні")
		var tmp int
		fmt.Print("\nВиберіть:")
		fmt.Scan(&tmp)
		if tmp == 1 {
			continue
		} else {
			break
		}
	}
	return products
}
func PrintProduct(product Product) {
	fmt.Printf("%-10v%v(%v грн.)%-10v%-10v%-10v%-10v\n", product.Name(), product.Price(), math.Round(product.Price()*product.Cost().ExRate()), product.Cost().name, product.Quantity(), product.Producer(), product.Weight())
}

func PrintProducts(products []Product) {
	fmt.Print("Назва\tЦіна\t\tВалюта\t\tК-сть\tВиробник\t\tВага\n")
	for _, product := range products {
		PrintProduct(product)
	}

}

func GetProductsInfo(products []Product) (Product, Product) {
	min := 0.0
	max := 0.0
	minI := 0
	maxI := 0
	for i := 0; i < len(products); i++ {
		price := products[i].Price() * products[i].Cost().ExRate()
		if price < min {
			minI = i
		}
		if price > max {
			maxI = i
		}

	}
	return products[minI], products[maxI]
}

func main() {
	var products []Product
	for true {
		fmt.Print("Меню:\n1.Додати товари\n2.Вивести продукт\n3.Вивести всі продукти\n4.Вивести найдешевший та найдорожчий")
		menu := 0
		fmt.Print("\nВиберіть:")
		fmt.Scan(&menu)

		switch menu {
		case 1:
			{
				products = append(products, ReadProductsArray()...)
				PrintProducts(products)
				break
			}
		case 2:
			{

				fmt.Print("Введіть номер:")
				num := -1
				fmt.Scanf("%d", &num)
				num -= 1
				if num >= 0 && num < len(products) {
					fmt.Print("Назва\tЦіна\t\tВалюта\t\tК-сть\tВиробник\t\tВага\n")
					PrintProduct(products[num])
				} else {
					fmt.Print("Такого товару не існує\n")
				}
				break
			}

		case 3:
			{

				PrintProducts(products)
				break
			}
		case 4:
			{
				min, max := GetProductsInfo(products)
				fmt.Print("Мінімальна ціна:\n")
				fmt.Print("Назва\tЦіна\t\tВалюта\t\tК-сть\tВиробник\t\tВага\n")
				PrintProduct(min)
				fmt.Print("Максимальна ціна:\n")
				fmt.Print("Назва\tЦіна\t\tВалюта\t\tК-сть\tВиробник\t\tВага\n")
				PrintProduct(max)
				break
			}
		default:
			{
			}
		}

	}
}
