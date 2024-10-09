# -*- coding: utf-8 -*-
import unittest

from gilded_rose import Item, GildedRose


class GildedRoseTest(unittest.TestCase):
    def test_foo(self):
        items = [Item("foo", 0, 0)]
        gilded_rose = GildedRose(items)
        gilded_rose.update_quality()
        self.assertEqual("foo", items[0].name)

    def test_default_sell_in_decreases_by_5_days(self):
        items = [Item("goo", 15, 20)]
        gilded_rose = GildedRose(items)
        days = 5
        for day in range(days):
            gilded_rose.update_quality()
        self.assertEqual(10, items[0].sell_in)

    def test_ticket_quality_increases_over_5_days(self):
        items = [Item(name="Backstage passes to a TAFKAL80ETC concert", sell_in=15, quality=20)]
        gilded_rose = GildedRose(items)
        days = 6
        for day in range(days):
            gilded_rose.update_quality()
        self.assertEqual(27, items[0].quality)
        self.assertEqual(9, items[0].sell_in)

if __name__ == '__main__':
    unittest.main()
