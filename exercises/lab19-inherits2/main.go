package main

import "fmt"

type A struct {
	Name string
	Age  int
}

type B struct {
	Name  string
	Score float64
}

type C struct {
	A
	B
	//Name string
}

type D struct {
	a A
}

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TVR struct {
	*Goods
	*Brand
}

type BasicType struct {
	Goods
	int
}

func main() {
	var c C
	// ambiguous name
	// c.Name = "tom"
	// close assign rule, as C have field of 'Name', use c.Name
	// c.Name = "Jack"
	// or C hasn't 'Name', specify c.A.Name
	c.A.Name = "tom"

	var d D
	// undefined except anonymously
	//d.Name = "Jason"
	d.a.Name = "jason"

	// TV
	tv1 := TV{
		Goods{"tv01", 500.99},
		Brand{"sony", "JP"},
	}

	tv2 := TV{
		Goods{
			Name:  "tv02",
			Price: 399.99,
		},
		Brand{
			"sharp",
			"JP",
		},
	}

	tvr := TVR{
		&Goods{
			Name:  "tvr",
			Price: 200.99,
		},
		&Brand{
			"samsung",
			"KO",
		},
	}

	fmt.Println("tv1 = ", tv1)
	fmt.Println("tv2 = ", tv2)
	fmt.Println("tvr = ", *tvr.Goods, *tvr.Brand)

	// anonymous value type
	basic := BasicType{
		Goods{Name: "cake", Price: 0.99},
		99,
	}

	fmt.Println("basic = ", basic)
}
