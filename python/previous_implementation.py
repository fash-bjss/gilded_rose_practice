self = {}
for item in self.items:
    """
    First Refactor
    """
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
        self.increase_quality(item, 1)

        if item.sell_in < 0:
            self.decrease_quality(item, item.quality)
        elif item.sell_in < 6:
            self.increase_quality(item, 1)
        elif item.sell_in < 11:
            self.increase_quality(item, 1)


    else:
        self.decrease_sell_in(item, 1)
        
        if item.sell_in < 0:
            self.decrease_quality(item, 2)
        else:
            self.decrease_quality(item, 1)

            

        """
        Original Code has a bug
        """

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
        #                     # bug this should be add 2
        #                     item.quality = item.quality + 1
        #             if item.sell_in < 6:
        #                 if item.quality < 50:
        #                     # bug this should be add 3
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