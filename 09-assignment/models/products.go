package models

type Products []Product

func (products *Products) AddProduct(product Product) {
	*products = append(*products, product)
}

func (products Products) IndexOf(product Product) int {
	for idx, prod := range products {
		if prod == product {
			return idx
		}
	}
	return -1
}

func (products Products) Includes(product Product) bool {
	return products.IndexOf(product) != -1
}

func (products Products) Any(predicate func(Product) bool) bool {
	for _, prod := range products {
		if predicate(prod) {
			return true
		}
	}
	return false
}

func (products Products) All(predicate func(Product) bool) bool {
	for _, prod := range products {
		if !predicate(prod) {
			return false
		}
	}
	return true
}

func (products Products) Format() string {
	result := ""
	for _, product := range products {
		result += product.Format()
	}
	return result
}

func (products Products) Filter(predicate func(Product) bool) *Products {
	result := &Products{}
	for _, prod := range products {
		if predicate(prod) {
			*result = append(*result, prod)
		}
	}
	return result
}
