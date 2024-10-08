# -*- coding: utf-8 -*-

class GildedRose(object):

    def __init__(self, items):
        self.items = items

    def decrease_sell_in(self, item, amount):
        item.sell_in -= amount

    def decrease_quality(self, item, amount):
        min_quality = 0
        item.quality -= amount
        if item.quality < min_quality:
            item.quality = min_quality
        

    def increase_quality(self, item, amount):
        max_quality = 50
        if item.quality < max_quality:
            item.quality += amount
            return
        item.quality = max_quality

    def is_past_sell_by_date(self, date):
        final_day = 0
        return date < final_day

    def process_brie(self, item):
        self.decrease_sell_in(item, 1)

        if self.is_past_sell_by_date(item.sell_in):
            self.increase_quality(item, 2)
        else:
            self.increase_quality(item, 1) 

    def process_sulfuras(self, item):
        return 
    
    def process_ticket(self, item):
        thresholds = {
            'first': 11,
            'last': 6
        }
        self.increase_quality(item, 1)           

        if item.sell_in < thresholds['first']:
            self.increase_quality(item, 1)

        if item.sell_in < thresholds['last']:
            self.increase_quality(item, 1)

        self.decrease_sell_in(item, 1)
        if self.is_past_sell_by_date(item.sell_in):
            self.decrease_quality(item, item.quality)


    def process_conjured(self, item):
        self.decrease_sell_in(item, 1)
        self.decrease_quality(item, 2)

    def process_default(self, item):
        self.decrease_sell_in(item, 1)
        
        if self.is_past_sell_by_date(item.sell_in):
            self.decrease_quality(item, 2)
        else:
            self.decrease_quality(item, 1)

    def process_items(self):
        items_functions = {
            'Aged Brie': self.process_brie,
            'Sulfuras, Hand of Ragnaros': self.process_sulfuras,
            'Backstage passes to a TAFKAL80ETC concert': self.process_ticket,
            'Conjured Mana Cake': self.process_conjured,
            'default': self.process_default,
        }

        for item in self.items:
            items_functions.get(item.name, items_functions['default'])(item)

    def update_quality(self):
        self.process_items()


class Item:
    def __init__(self, name, sell_in, quality):
        self.name = name
        self.sell_in = sell_in
        self.quality = quality

    def __repr__(self):
        return "%s, %s, %s" % (self.name, self.sell_in, self.quality)
