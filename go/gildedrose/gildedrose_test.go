package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func createAllItems() []*gildedrose.Item {

	items := []*gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Conjured Mana Cake", 3, 6},
	}

	return items

}

func Test_Foo(t *testing.T) {
	var items = []*gildedrose.Item{
		{"foo", 0, 0},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Name != "foo" {
		t.Errorf("Name: Expected %s but got %s ", "fixme", items[0].Name)
	}
}

func Test_aged_brie_quality_increases_as_sell_in_decreases(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Aged Brie", 2, 0},
	}

	i := 0
	for i < 5 {
		gildedrose.UpdateQuality(items)
		i += 1
	}

	expected := 8
	got := items[0].Quality

	if items[0].Quality != expected {
		t.Errorf("Got %v but expected %v", got, expected)
	}

}

func Test_sellin_of_all_items_after_12_days(t *testing.T) {
	items := createAllItems()

	i := 0
	for i < 12 {
		gildedrose.UpdateQuality(items)
		i += 1
	}

	tables := []struct {
		itemName string
		got      int
		expected int
	}{
		{items[0].Name, items[0].SellIn, -2},
		{items[1].Name, items[1].SellIn, -10},
		{items[2].Name, items[2].SellIn, -7},
		{items[3].Name, items[3].SellIn, 0},
		{items[4].Name, items[4].SellIn, -1},
		{items[5].Name, items[5].SellIn, 3},
		{items[6].Name, items[6].SellIn, -2},
		{items[7].Name, items[7].SellIn, -7},
		{items[8].Name, items[8].SellIn, -9},
	}

	for _, field := range tables {
		if field.got != field.expected {
			t.Errorf("%s got %v but expected %v", field.itemName, field.got, field.expected)
		}
	}
}

func Test_quality_of_all_items_after_12_days(t *testing.T) {
	items := createAllItems()

	i := 0

	for i < 12 {
		gildedrose.UpdateQuality(items)
		i += 1
	}

	tables := []struct {
		itemName string
		got      int
		expected int
	}{
		{items[0].Name, items[0].Quality, 6},
		{items[1].Name, items[1].Quality, 22},
		{items[2].Name, items[2].Quality, 0},
		{items[3].Name, items[3].Quality, 80},
		{items[4].Name, items[4].Quality, 80},
		{items[5].Name, items[5].Quality, 41},
		{items[6].Name, items[6].Quality, 0},
		{items[7].Name, items[7].Quality, 0},
		{items[8].Name, items[8].Quality, 0},
	}

	for _, field := range tables {
		if field.got != field.expected {
			t.Errorf("%s got %v but expected %v", field.itemName, field.got, field.expected)
		}
	}
}
