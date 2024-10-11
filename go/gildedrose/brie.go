package gildedrose

type Brie struct {
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
