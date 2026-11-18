package cart

import (
	"errors"
	"math"
)

// Item represents a product added to the shopping cart.
type Item struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

// Cart represents the user shopping cart holding items.
type Cart struct {
	Items []Item `json:"items"`
}

// Summary represents the calculated cart totals including taxes and discounts.
type Summary struct {
	Subtotal float64 `json:"subtotal"`
	Discount float64 `json:"discount_amount"`
	Tax      float64 `json:"tax_amount"`
	Total    float64 `json:"total"`
}

// NewCart initializes an empty Cart.
func NewCart() *Cart {
	return &Cart{
		Items: []Item{},
	}
}

// AddItem adds an item to the cart or increments quantity if it already exists.
func (c *Cart) AddItem(newItem Item) error {
	if newItem.Quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}
	if newItem.Price < 0 {
		return errors.New("price cannot be negative")
	}

	for i, item := range c.Items {
		if item.ID == newItem.ID {
			c.Items[i].Quantity += newItem.Quantity
			return nil
		}
	}

	c.Items = append(c.Items, newItem)
	return nil
}

// RemoveItem removes an item from the cart by its ID or returns an error if not found.
func (c *Cart) RemoveItem(id string) error {
	for i, item := range c.Items {
		if item.ID == id {
			c.Items = append(c.Items[:i], c.Items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}

// CalculateTotal computes subtotal, applies discount percentage, then applies taxRate.
func (c *Cart) CalculateTotal(discount float64, taxRate float64) float64 {
	subtotal := 0.0
	for _, item := range c.Items {
		subtotal += item.Price * float64(item.Quantity)
	}

	discountAmount := subtotal * discount
	if discountAmount < 0 {
		discountAmount = 0
	}
	discountedSubtotal := subtotal - discountAmount

	taxAmount := discountedSubtotal * taxRate
	if taxAmount < 0 {
		taxAmount = 0
	}

	total := discountedSubtotal + taxAmount
	return math.Round(total*100) / 100
}

// GetSummary returns detailed breakdown of the cart.
func (c *Cart) GetSummary(discount float64, taxRate float64) Summary {
	subtotal := 0.0
	for _, item := range c.Items {
		subtotal += item.Price * float64(item.Quantity)
	}

	discountAmount := math.Round((subtotal*discount)*100) / 100
	discountedSubtotal := subtotal - discountAmount
	taxAmount := math.Round((discountedSubtotal*taxRate)*100) / 100
	total := math.Round((discountedSubtotal+taxAmount)*100) / 100

	return Summary{
		Subtotal: math.Round(subtotal*100) / 100,
		Discount: discountAmount,
		Tax:      taxAmount,
		Total:    total,
	}
}
