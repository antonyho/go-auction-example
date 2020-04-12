package auction

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	// ErrItemNotExist is an error when an item is not found by the given name
	ErrItemNotExist = errors.New("item does not exist")
	// ErrUserNotExist is an error when a user is not found by the given name
	ErrUserNotExist = errors.New("User does not exist")
	// ErrAuctionClose is an error when bidding on a closed auction
	ErrAuctionClose = errors.New("auction has closed")
)

// Offer by bidder to an auction item
type Offer struct {
	Item      string
	Bidder    string
	Price     float64
	ReplyChan chan BidResult
}

// NewOffer instantiates a new offer pointer to struct
func NewOffer(itemName string, bidder string, price float64) *Offer {
	return &Offer{
		Item:      itemName,
		Bidder:    bidder,
		Price:     price,
		ReplyChan: nil,
	}
}

// BidResult is the outcome of an offer
type BidResult struct {
	Accepted     bool
	RejectReason string
}

// NewBidResult instantiates a new BidResult pointer to struct
func NewBidResult(isAccepted bool, rejectReason string) *BidResult {
	return &BidResult{
		Accepted:     isAccepted,
		RejectReason: rejectReason,
	}
}

// Bid is the event when a user place a bid to an auction item
type Bid struct {
	ID        string    `json:"id"`
	User      string    `json:"user"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

// NewBid instantiates a new Bid pointer to struct
func NewBid(user string, price float64) *Bid {
	return &Bid{
		ID:        uuid.New().String(),
		User:      user,
		Price:     price,
		CreatedAt: time.Now(),
	}
}

// Item is the object which is listed in an auction
type Item struct {
	Name     string     `json:"name"`
	Bids     []*Bid     `json:"bids"`
	ClosedAt *time.Time `json:"closed_at"`
}

// NewItem instantiates a new Item pointer to struct
func NewItem(name string) *Item {
	return &Item{
		Name:     name,
		Bids:     make([]*Bid, 0, 0), // This auction doesn't implement opening bid
		ClosedAt: nil,
	}
}

// Close the auction on an item
func (i *Item) Close() {
	if i.ClosedAt == nil {
		now := time.Now()
		i.ClosedAt = &now
	}
}

// IsClosed returns auction closure status
func (i *Item) IsClosed() bool {
	if i.ClosedAt != nil {
		return true
	}
	return false
}

// ActivityList is a list of auction acitvities on an item done by specific user
type ActivityList map[string][]*Bid // item_name:bids

// House is an auction house system to all items which are put into auction
type House struct {
	list           map[string]*Item        // item_name:item
	userActivities map[string]ActivityList // user_id:ActivityList
}

// NewHouse instantiates a new House pointer to struct
func NewHouse() *House {
	return &House{
		list:           make(map[string]*Item, 0),
		userActivities: make(map[string]ActivityList, 0),
	}
}

// Add new item into the auction list
func (h *House) Add(item *Item) {
	if _, exist := h.list[item.Name]; !exist {
		h.list[item.Name] = item
	}
}

// Close an item auction
func (h *House) Close(itemName string) error {
	item, exist := h.list[itemName]
	if !exist {
		return ErrItemNotExist
	} else if item.IsClosed() {
		return ErrAuctionClose
	}

	item.Close()

	return nil
}

// Bid an item
// Returns whether the bid has been accepted
func (h *House) Bid(user string, itemName string, price float64) (bool, error) {
	item, exist := h.list[itemName]
	if !exist {
		return false, ErrItemNotExist
	}
	if item.IsClosed() {
		return false, ErrAuctionClose
	}
	bid := NewBid(user, price)
	if len(item.Bids) > 0 {
		lastBid := item.Bids[len(item.Bids)-1]
		if price <= lastBid.Price {
			return false, nil
		}
	}

	item.Bids = append(item.Bids, bid)
	activities, exist := h.userActivities[user]
	if !exist {
		activities = make(ActivityList, 1)
		activities[itemName] = []*Bid{bid}
		h.userActivities[user] = activities
	} else {
		bids, bidExist := activities[itemName]
		if !bidExist {
			activities[itemName] = []*Bid{bid}
		} else {
			activities[itemName] = append(bids, bid)
		}
	}

	return true, nil
}

// GetLeading gets the winning bid on given item
func (h *House) GetLeading(itemName string) (*Bid, error) {
	item, exist := h.list[itemName]
	if !exist {
		return nil, ErrItemNotExist
	}

	if len(item.Bids) == 0 {
		return nil, nil
	}

	return item.Bids[len(item.Bids)-1], nil
}

// GetBids gets all bids on given item
func (h *House) GetBids(itemName string) ([]*Bid, error) {
	item, exist := h.list[itemName]
	if !exist {
		return nil, ErrItemNotExist
	}

	return item.Bids, nil
}

// GetUserActivities gets all bidding activities by given user
func (h *House) GetUserActivities(user string) (ActivityList, error) {
	activityList, exist := h.userActivities[user]
	if !exist {
		return nil, ErrUserNotExist
	}

	return activityList, nil
}

// getItem returns item internal-only
func (h *House) getItem(itemName string) (*Item, error) {
	item, exist := h.list[itemName]
	if !exist {
		return nil, ErrItemNotExist
	}

	return item, nil
}
