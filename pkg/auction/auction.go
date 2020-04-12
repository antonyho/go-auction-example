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
)

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
	Name string `json:"name"`
	Bids []*Bid `json:"bids"`
}

// ActivityList is a list of auction acitvities on an item done by specific user
type ActivityList map[string][]*Bid // item_name:bids

// NewItem instantiates a new Item pointer to struct
func NewItem(name string) *Item {
	return &Item{
		Name: name,
		Bids: make([]*Bid, 0, 0),
	}
}

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

// Bid an item
// returns whether the bid has been accepted
func (h *House) Bid(user string, itemName string, price float64) (bool, error) {
	item, exist := h.list[itemName]
	if !exist {
		return false, ErrItemNotExist
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
