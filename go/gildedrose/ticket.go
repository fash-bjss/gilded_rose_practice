package gildedrose

type Ticket struct {
	*Item
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
