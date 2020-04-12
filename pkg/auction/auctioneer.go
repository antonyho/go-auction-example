package auction

// Auctioneer is a manager at auction house which handles concurrent bids
type Auctioneer struct {
	house    *House
	bidChans map[string]chan Offer // item_name:Offer
}

// NewAuctioneer instantiates a new auctioneer for an auction house
func NewAuctioneer(house *House) *Auctioneer {
	return &Auctioneer{
		house:    house,
		bidChans: make(map[string]chan Offer, 0),
	}
}

// List a new item
func (a *Auctioneer) List(newItemName string) {
	_, exist := a.bidChans[newItemName]
	if !exist { // if item exists, new listing will be ignored.
		a.bidChans[newItemName] = make(chan Offer) // create a non-block channel
		newItem := NewItem(newItemName)
		a.house.Add(newItem)
	}

	go a.chant(newItemName)
}

// Hear a bid on specific item
// Returns if the bid is valid
// Returns failure with reason if any
func (a *Auctioneer) Hear(itemName string, offer *Offer) (bool, error) {
	item, err := a.house.getItem(itemName)
	if err == ErrItemNotExist {
		return false, err
	}
	if item.IsClosed() {
		return false, ErrAuctionClose
	}
	bidChan, exist := a.bidChans[itemName]
	if exist {
		offer.ReplyChan = make(chan BidResult, 1)
		bidChan <- *offer
	} else {
		return false, ErrAuctionClose // A safeguard on closed auction
	}

	result := <-offer.ReplyChan

	var rejectReason error
	switch result.RejectReason {
	case ErrItemNotExist.Error():
		rejectReason = ErrItemNotExist
	case ErrAuctionClose.Error():
		rejectReason = ErrAuctionClose
	default:
		rejectReason = nil
	}

	return result.Accepted, rejectReason
}

// chant on bidding
// An indefinite executing process until an item auction is closed
func (a *Auctioneer) chant(itemName string) {
	c := a.bidChans[itemName]
	for offer := range c {
		accepted, err := a.house.Bid(offer.Bidder, offer.Item, offer.Price)
		var result *BidResult
		if err != nil {
			result = NewBidResult(accepted, err.Error())
		} else {
			result = NewBidResult(accepted, "")
		}
		offer.ReplyChan <- *result
	}
	// This channel will be closed when auction has ended for this item
	// This indefinite loop will end
}

// Close auctioning on an item
// This function is not protected against concurrency.
// Therefore, it is DDoS vulnerable.
func (a *Auctioneer) Close(itemName string) {
	item, err := a.house.getItem(itemName)
	if err != ErrItemNotExist && !item.IsClosed() {
		auctionChan, exist := a.bidChans[itemName]
		if exist {
			close(auctionChan)
			delete(a.bidChans, itemName)
			a.house.Close(itemName)
		}
		item.Close()
	}
}
