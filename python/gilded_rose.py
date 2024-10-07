# -*- coding: utf-8 -*-

class GildedRose(object):

    def __init__(self, items):
        self.items = items

    def decrease_sell_in(self, item, n):
        item.sell_in -= n

    def decrease_quality(self, item, n):
        if item.quality < 0:
            item.quality = 0
            return
        item.quality -= n

    def increase_quality(self, item, n):
        if item.quality > 50:
            item.quality = 50
            return   
        item.quality += n




    def update_quality(self):
        for item in self.items:

            if item.name == 'Aged Brie':
                self.decrease_sell_in(item, 1)

                if item.sell_in < 0:
                    self.increase_quality(item, 2)
                else:
                    self.increase_quality(item, 1)
            
            elif item.name == 'Sulfuras, Hand of Ragnaros':
                pass

            elif item.name == 'Backstage passes to a TAFKAL80ETC concert':
                self.decrease_sell_in(item, 1)
                if item.sell_in <= 10:
                    self.increase_quality(item, 2)
                
                if item.sell_in <=5:
                    self.increase_quality(item, 3)

                if item.sell_in < 0:
                    self.decrease_quality(item, item.quality)
    
            else:
                self.decrease_sell_in(item, 1)
                
                if item.sell_in < 0:
                    self.decrease_quality(item, 2)
                else:
                    self.decrease_quality(item, 1)

                


            # if item.name != "Aged Brie" and item.name != "Backstage passes to a TAFKAL80ETC concert":
            #     if item.quality > 0:
            #         if item.name != "Sulfuras, Hand of Ragnaros":
            #             item.quality = item.quality - 1
            # else:
            #     if item.quality < 50:
            #         item.quality = item.quality + 1
            #         if item.name == "Backstage passes to a TAFKAL80ETC concert":
            #             if item.sell_in < 11:
            #                 if item.quality < 50:
            #                     item.quality = item.quality + 1
            #             if item.sell_in < 6:
            #                 if item.quality < 50:
            #                     item.quality = item.quality + 1
            # if item.name != "Sulfuras, Hand of Ragnaros":
            #     item.sell_in = item.sell_in - 1
            # if item.sell_in < 0:
            #     if item.name != "Aged Brie":
            #         if item.name != "Backstage passes to a TAFKAL80ETC concert":
            #             if item.quality > 0:
            #                 if item.name != "Sulfuras, Hand of Ragnaros":
            #                     item.quality = item.quality - 1
            #         else:
            #             item.quality = item.quality - item.quality
            #     else:
            #         if item.quality < 50:
            #             item.quality = item.quality + 1


class Item:
    def __init__(self, name, sell_in, quality):
        self.name = name
        self.sell_in = sell_in
        self.quality = quality

    def __repr__(self):
        return "%s, %s, %s" % (self.name, self.sell_in, self.quality)
