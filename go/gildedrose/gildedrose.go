package gildedrose

type Items interface {
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

type Brie struct {
	Item
}

type Ticket struct {
	Item
}

type Default struct {
	Item
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

func (item *Default) ProcessItem() {

	item.decreaseSellin(1)

	if item.isPastSellByDate() {
		item.decreaseQuality(2)
	} else {
		item.decreaseQuality(1)
	}

}

func UpdateQuality(items []*Item) {

	for _, item := range items {

		switch item.Name {
		case "Aged Brie":
			useBrieStruct := Brie{*item}
			useBrieStruct.ProcessItem()

			item.SellIn = useBrieStruct.SellIn
			item.Quality = useBrieStruct.Quality

		case "Sulfuras, Hand of Ragnaros":

		case "Backstage passes to a TAFKAL80ETC concert":
			useTicketStruct := Ticket{*item}
			useTicketStruct.ProcessItem()

			item.SellIn = useTicketStruct.SellIn
			item.Quality = useTicketStruct.Quality

		default:
			useDefaultStruct := Default{*item}
			useDefaultStruct.ProcessItem()

			item.SellIn = useDefaultStruct.SellIn
			item.Quality = useDefaultStruct.Quality
		}

	}

}
