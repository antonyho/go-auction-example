package auction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHouse_Add(t *testing.T) {
	house := NewHouse()

	item := NewItem("Mona Lisa")
	house.Add(item)

	assert.Len(t, house.list, 1)

	// Test item duplication result
	duplicate := NewItem("Mona Lisa")
	house.Add(duplicate)

	assert.Len(t, house.list, 1)
}

func TestHouse_Close(t *testing.T) {
	house := NewHouse()

	item := NewItem("Mona Lisa")
	house.Add(item)
	err := house.Close("Mona Lisa")
	assert.NoError(t, err)

	err = house.Close("Mona Lisa")
	assert.Error(t, err)
}

func TestHouse_Bid(t *testing.T) {
	house := NewHouse()

	house.Add(NewItem("Mona Lisa"))
	house.Add(NewItem("The Creation of Adam"))

	assert.Len(t, house.list, 2)

	bidResult, err := house.Bid("LeonardoDaVinci", "Mona Lisa", 10.0)
	assert.NoError(t, err)
	assert.True(t, bidResult)
	assert.Len(t, house.list["Mona Lisa"].Bids, 1)
	assert.Len(t, house.list["The Creation of Adam"].Bids, 0)

	bidResult, err = house.Bid("Antony", "MonaLisa", 11.0)
	assert.Error(t, err)
	assert.False(t, bidResult)
	assert.Len(t, house.list["Mona Lisa"].Bids, 1)

	bidResult, err = house.Bid("Antony", "Mona Lisa", 11.0)
	assert.NoError(t, err)
	assert.True(t, bidResult)
	assert.Len(t, house.list["Mona Lisa"].Bids, 2)

	bidResult, err = house.Bid("LeonardoDaVinci", "Mona Lisa", 11.0)
	assert.NoError(t, err)
	assert.False(t, bidResult)
	assert.Len(t, house.list["Mona Lisa"].Bids, 2)

	err = house.Close("Mona Lisa")
	assert.NoError(t, err)

	bidResult, err = house.Bid("Antony", "Mona Lisa", 100.0)
	assert.Error(t, err)
	assert.False(t, bidResult)
	assert.Len(t, house.list["Mona Lisa"].Bids, 2)
}

func TestHouse_GetLeading(t *testing.T) {
	house := NewHouse()

	house.Add(NewItem("Mona Lisa"))
	house.Add(NewItem("The Creation of Adam"))

	winningBid, err := house.GetLeading("The Creation of Adam")
	assert.NoError(t, err)
	assert.Nil(t, winningBid)

	winningBid, err = house.GetLeading("TheCreationofAdam")
	assert.Error(t, err)
	assert.Nil(t, winningBid)

	_, err = house.Bid("Michelangelo", "The Creation of Adam", 10.0)
	assert.NoError(t, err)

	winningBid, err = house.GetLeading("The Creation of Adam")
	assert.NoError(t, err)
	assert.EqualValues(t, 10.0, winningBid.Price)
	assert.EqualValues(t, "Michelangelo", winningBid.User)

	_, err = house.Bid("Antony", "The Creation of Adam", 9.0)
	assert.NoError(t, err)

	winningBid, err = house.GetLeading("The Creation of Adam")
	assert.NoError(t, err)
	assert.EqualValues(t, 10.0, winningBid.Price)
	assert.EqualValues(t, "Michelangelo", winningBid.User)

	_, err = house.Bid("Antony", "The Creation of Adam", 20.0)
	assert.NoError(t, err)

	winningBid, err = house.GetLeading("The Creation of Adam")
	assert.NoError(t, err)
	assert.EqualValues(t, 20.0, winningBid.Price)
	assert.EqualValues(t, "Antony", winningBid.User)
}

func TestHouse_GetBids(t *testing.T) {
	house := NewHouse()

	house.Add(NewItem("Mona Lisa"))
	house.Add(NewItem("The Creation of Adam"))

	_, err := house.Bid("Michelangelo", "The Creation of Adam", 10.0)
	assert.NoError(t, err)

	_, err = house.Bid("Antony", "The Creation of Adam", 9.0)
	assert.NoError(t, err)

	_, err = house.Bid("Antony", "The Creation of Adam", 20.0)
	assert.NoError(t, err)

	bids, err := house.GetBids("The Creation of Adam")
	assert.NoError(t, err)
	assert.Len(t, bids, 2)

	bids, err = house.GetBids("TheCreationofAdam")
	assert.Error(t, err)

	bids, err = house.GetBids("Mona Lisa")
	assert.NoError(t, err)
	assert.Len(t, bids, 0)
}

func TestHouse_GetUserActivities(t *testing.T) {
	house := NewHouse()

	house.Add(NewItem("Mona Lisa"))
	house.Add(NewItem("The Creation of Adam"))

	_, err := house.Bid("Michelangelo", "The Creation of Adam", 10.0)
	assert.NoError(t, err)

	_, err = house.Bid("Antony", "The Creation of Adam", 9.0)
	assert.NoError(t, err)

	_, err = house.Bid("Antony", "The Creation of Adam", 20.0)
	assert.NoError(t, err)

	_, err = house.Bid("Michelangelo", "Mona Lisa", 10.0)
	assert.NoError(t, err)

	_, err = house.Bid("Michelangelo", "The Creation of Adam", 20.01)
	assert.NoError(t, err)

	ma, err := house.GetUserActivities("Michelangelo")
	assert.NoError(t, err)
	assert.Len(t, ma, 2)
	assert.Len(t, ma["The Creation of Adam"], 2)
	assert.Len(t, ma["Mona Lisa"], 1)

	aa, err := house.GetUserActivities("Antony")
	assert.NoError(t, err)
	assert.Len(t, aa, 1)
	assert.Len(t, aa["The Creation of Adam"], 1)
	assert.Nil(t, aa["Mona Lisa"])

	na, err := house.GetUserActivities("Nobody")
	assert.Error(t, err)
	assert.Nil(t, na)
}
