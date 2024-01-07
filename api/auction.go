package stalcraftgo

import (
	stalcraftgo "github.com/re-incarnation/stalcraftgo/object"
)

func (exbo *EXBO) ItemPriceHistory(Region string, Item string, params Params) (response stalcraftgo.ItemPriceHistoryResponse, err error) {
	err = exbo.RequestUnmarshalAuctionItemPriceHistory(Region, Item, params, &response)
	return
}

func (exbo *EXBO) ActiveItemLots(Region string, Item string, params Params) (response stalcraftgo.ActiveItemLotsResponse, err error) {
	err = exbo.RequestUnmarshalAuctionActiveItemLots(Region, Item, params, &response)
	return
}
