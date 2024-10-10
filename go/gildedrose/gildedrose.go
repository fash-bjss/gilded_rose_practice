package gildedrose

type Process interface {
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

func (item *Brie) ProcessItem() {
	item.decreaseSellin(1)

	if item.isPastSellByDate() {
		item.increaseQuality(2)
	} else {
		item.increaseQuality(1)
	}

}

func (item *Ticket) ProcessItem() {

	type threshold struct {
		first int
		last  int
	}

	thresholds := threshold{
		first: 11,
		last:  6,
	}

	item.increaseQuality(1)

	if item.SellIn < thresholds.first {
		item.increaseQuality(1)
	}

	if item.SellIn < thresholds.last {
		item.increaseQuality(1)
	}

	item.decreaseSellin(1)
	if item.isPastSellByDate() {
		item.decreaseQuality(item.Quality)
	}

}

func (item *Sulfuras) ProcessItem() {

}

// Returning the interface will allow you to return multiple structs that share the interface
func ChooseAndCreateItem(item *Item) Process {

	switch item.Name {
	case "Aged Brie":
		return &Brie{item}

	case "Sulfuras, Hand of Ragnaros":
		return &Sulfuras{item}

	case "Backstage passes to a TAFKAL80ETC concert":
		return &Ticket{item}

	default:
		return item
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
