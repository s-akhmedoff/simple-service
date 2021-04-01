package views

type Price struct {
	Sale     int `json:"sale"`
	Factory  int `json:"factory"`
	Discount int `json:"discount"`
}

type Product struct {
	SKU         string  `json:"sku"`
	Name        string  `json:"name"`
	URI         string  `json:"uri"`
	Description string  `json:"description"`
	Prices      []Price `json:"prices"`
}

type Products struct {
	Total int       `json:"total"`
	Items []Product `json:"items"`
}

type Shop struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Address      string      `json:"address"`
	Longitude    string      `json:"longitude"`
	Latitude     string      `json:"latitude"`
	WorkingHours interface{} `json:"working_hours"`
}

type Shops struct {
	Total int    `json:"total"`
	Items []Shop `json:"items"`
}
