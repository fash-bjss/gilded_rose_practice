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

// func NewItem(name string, sellIn int, quality int) *Item {
// 	return &Item{
// 		Name:    name,
// 		SellIn:  sellIn,
// 		Quality: quality,
// 	}
// }

type Brie struct {
	Item
}

type Ticket struct {
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
			item.decreaseSellin(1)

			if item.isPastSellByDate() {
				item.decreaseQuality(2)
			} else {
				item.decreaseQuality(1)
			}
		}

	}

	// for i := 0; i < len(items); i++ {

	// 	if items[i].Name != "Aged Brie" && items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
	// 		if items[i].Quality > 0 {
	// 			if items[i].Name != "Sulfuras, Hand of Ragnaros" {
	// 				items[i].Quality = items[i].Quality - 1
	// 			}
	// 		}
	// 	} else {
	// 		if items[i].Quality < 50 {
	// 			items[i].Quality = items[i].Quality + 1
	// 			if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
	// 				if items[i].SellIn < 11 {
	// 					if items[i].Quality < 50 {
	// 						items[i].Quality = items[i].Quality + 1
	// 					}
	// 				}
	// 				if items[i].SellIn < 6 {
	// 					if items[i].Quality < 50 {
	// 						items[i].Quality = items[i].Quality + 1
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}

	// 	if items[i].Name != "Sulfuras, Hand of Ragnaros" {
	// 		items[i].SellIn = items[i].SellIn - 1
	// 	}

	// 	if items[i].SellIn < 0 {
	// 		if items[i].Name != "Aged Brie" {
	// 			if items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
	// 				if items[i].Quality > 0 {
	// 					if items[i].Name != "Sulfuras, Hand of Ragnaros" {
	// 						items[i].Quality = items[i].Quality - 1
	// 					}
	// 				}
	// 			} else {
	// 				items[i].Quality = items[i].Quality - items[i].Quality
	// 			}
	// 		} else {
	// 			if items[i].Quality < 50 {
	// 				items[i].Quality = items[i].Quality + 1
	// 			}
	// 		}
	// 	}
	// }

}
