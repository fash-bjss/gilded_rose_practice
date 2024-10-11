package gildedrose

type ItemVariant interface {
	ProcessItem()
}

type Item struct {
	Name            string
	SellIn, Quality int
}

func (i *Item) decreaseSellin(amount int) {
	i.SellIn -= amount
}

func (i *Item) decreaseQuality(amount int) {
	i.Quality -= amount
	if i.Quality < 0 {
		i.Quality = 0
	}
}

func (i *Item) increaseQuality(amount int) {
	if i.Quality < 50 {
		i.Quality += amount
		return
	}
	i.Quality = 50
}

func (i *Item) isPastSellByDate() bool {
	return i.SellIn < 0
}

// This function must be declared on the "parent" struct first
func (i *Item) ProcessItem() {
	// This handles the default behaviour
	i.decreaseSellin(1)

	if i.isPastSellByDate() {
		i.decreaseQuality(2)
	} else {
		i.decreaseQuality(1)
	}
}

type Brie struct {
	*Item
}

type Ticket struct {
	*Item
}

type Sulfuras struct {
	*Item
}

func (b *Brie) ProcessItem() {
	b.decreaseSellin(1)

	if b.isPastSellByDate() {
		b.increaseQuality(2)
	} else {
		b.increaseQuality(1)
	}

}

func (t *Ticket) ProcessItem() {

	type threshold struct {
		first int
		last  int
	}

	thresholds := threshold{
		first: 11,
		last:  6,
	}

	t.increaseQuality(1)

	if t.SellIn < thresholds.first {
		t.increaseQuality(1)
	}

	if t.SellIn < thresholds.last {
		t.increaseQuality(1)
	}

	t.decreaseSellin(1)
	if t.isPastSellByDate() {
		t.decreaseQuality(t.Quality)
	}

}

func (s *Sulfuras) ProcessItem() {

}

// Returning the interface will allow you to return multiple structs that share the interface
func ChooseAndCreateItem(item_variant *Item) ItemVariant {

	switch item_variant.Name {
	case "Aged Brie":
		return &Brie{item_variant}

	case "Sulfuras, Hand of Ragnaros":
		return &Sulfuras{item_variant}

	case "Backstage passes to a TAFKAL80ETC concert":
		return &Ticket{item_variant}

	default:
		return item_variant
	}
}

func RunEndOfDayForItems(items []*Item) {

	for _, item := range items {

		chosenStruct := ChooseAndCreateItem(item)
		chosenStruct.ProcessItem()
	}
}

func UpdateQuality(items []*Item) {

	RunEndOfDayForItems(items)

}
