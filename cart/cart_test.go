package cart

import (
	"testing"
)

func TestAddItem_NewItem(t *testing.T) {
	c := NewCart()
	item := Item{ID: "prod-1", Name: "Go Gopher Plushie", Price: 25.00, Quantity: 2}

	err := c.AddItem(item)
	if err != nil {
		t.Fatalf("unexpected error adding item: %v", err)
	}

	if len(c.Items) != 1 {
		t.Errorf("expected 1 item in cart, got %d", len(c.Items))
	}
	if c.Items[0].Quantity != 2 {
		t.Errorf("expected quantity 2, got %d", c.Items[0].Quantity)
	}
}

func TestAddItem_ExistingQuantityIncrement(t *testing.T) {
	c := NewCart()
	item1 := Item{ID: "prod-1", Name: "Go Gopher Plushie", Price: 25.00, Quantity: 2}
	item2 := Item{ID: "prod-1", Name: "Go Gopher Plushie", Price: 25.00, Quantity: 3}

	_ = c.AddItem(item1)
	_ = c.AddItem(item2)

	if len(c.Items) != 1 {
		t.Errorf("expected still 1 distinct item, got %d", len(c.Items))
	}
	if c.Items[0].Quantity != 5 {
		t.Errorf("expected total quantity 5, got %d", c.Items[0].Quantity)
	}
}

func TestRemoveItem_Success(t *testing.T) {
	c := NewCart()
	_ = c.AddItem(Item{ID: "prod-1", Name: "Item 1", Price: 10.0, Quantity: 1})
	_ = c.AddItem(Item{ID: "prod-2", Name: "Item 2", Price: 20.0, Quantity: 1})

	err := c.RemoveItem("prod-1")
	if err != nil {
		t.Fatalf("expected successful removal, got %v", err)
	}
	if len(c.Items) != 1 {
		t.Errorf("expected 1 remaining item, got %d", len(c.Items))
	}
	if c.Items[0].ID != "prod-2" {
		t.Errorf("expected remaining item prod-2, got %s", c.Items[0].ID)
	}
}

func TestRemoveItem_NotFound(t *testing.T) {
	c := NewCart()
	err := c.RemoveItem("non-existent-id")
	if err == nil {
		t.Fatal("expected error for non-existent item, got nil")
	}
	if err.Error() != "item not found" {
		t.Errorf("expected 'item not found' error message, got %s", err.Error())
	}
}

func TestCalculateTotal_DiscountAndTax(t *testing.T) {
	c := NewCart()
	_ = c.AddItem(Item{ID: "prod-1", Name: "Mechanical Keyboard", Price: 50.00, Quantity: 2})

	total := c.CalculateTotal(0.10, 0.19)
	expected := 107.10

	if total != expected {
		t.Errorf("expected total %.2f, got %.2f", expected, total)
	}
}
