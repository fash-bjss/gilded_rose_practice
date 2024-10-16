# -*- coding: utf-8 -*-

class GildedRose(object):

    def __init__(self, items):
        self.items = items

    def decrease_sell_in(self, amount):
        self.sell_in -= amount

    def decrease_quality(self, amount):
        min_quality = 0
        self.quality -= amount
        if self.quality < min_quality:
            self.quality = min_quality
        

    def increase_quality(self, amount):
        max_quality = 50
        if self.quality < max_quality:
            self.quality += amount
            return
        self.quality = max_quality

    def is_past_sell_by_date(self):
        final_day = 0
        return self.sell_in < final_day

    def choose_and_convert_item(self, item):

        item_classes = {
            "Aged Brie": Brie,
            "Sulfuras, Hand of Ragnaros": Sulfaras,
            "Backstage passes to a TAFKAL80ETC concert": Ticket,
            "Conjured Mana Cake": Conjured,
            "Default": Default
        }
        chosen_class = item_classes.get(item.name, item_classes["Default"])

        return chosen_class(item.name, item.sell_in, item.quality)

    def process_all_items(self):
        for item in self.items:
            converted_item = self.choose_and_convert_item(item)
            converted_item.process_item()
            
            index_of_item = self.items.index(item)
            self.items[index_of_item] = converted_item

    def update_quality(self):
        self.process_all_items()


class Item:
    def __init__(self, name, sell_in, quality):
        self.name = name
        self.sell_in = sell_in
        self.quality = quality

    def __repr__(self):
        return "%s, %s, %s" % (self.name, self.sell_in, self.quality)

class Brie(Item, GildedRose):
    def __init__(self, name, sell_in, quality):
        super().__init__(name, sell_in, quality)

    def process_item(self):
        self.decrease_sell_in(1)

        if  self.is_past_sell_by_date():
            self.increase_quality(2)
        else:
            self.increase_quality(1)

class Ticket(Item, GildedRose):
    def __init__(self, name, sell_in, quality):
        super().__init__(name, sell_in, quality)

    def process_item(self):
        thresholds = {
            'first': 11,
            'last': 6
        }
        self.increase_quality(1)           

        if self.sell_in < thresholds['first']:
            self.increase_quality(1)

        if self.sell_in < thresholds['last']:
            self.increase_quality(1)

        self.decrease_sell_in(1)
        if self.is_past_sell_by_date():
            self.decrease_quality(self.quality)


class Sulfaras(Item, GildedRose):
    def __init__(self, name, sell_in, quality):
        super().__init__(name, sell_in, quality)

    def process_item(self):
        pass

class Conjured(Item, GildedRose):
    def __init__(self, name, sell_in, quality):
        super().__init__(name, sell_in, quality)

    def process_item(self):
        self.decrease_sell_in(1)
        self.decrease_quality(2)

class Default(Item, GildedRose):
    def __init__(self, name, sell_in, quality):
        super().__init__(name, sell_in, quality)

    def process_item(self):
        self.decrease_sell_in(1)
        
        if self.is_past_sell_by_date():
            self.decrease_quality(2)
        else:
            self.decrease_quality(1)
