package auction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuctioneer_List(t *testing.T) {
	house := NewHouse()
	auctioneer := NewAuctioneer(house)

	items := []string{
		"Mona Lisa",
		"Mona Lisa", // test duplication
		"The Creation of Adam",
		"Mona Lisa", // test duplication
	}

	for _, name := range items {
		auctioneer.List(name)
	}

	assert.Len(t, house.list, 2)
}

func TestAuctioneer_Close(t *testing.T) {
	house := NewHouse()
	auctioneer := NewAuctioneer(house)

	monalisa := "Mona Lisa"
	creationOfAdam := "The Creation of Adam"
	sistineMadonna := "Sistine Madonna"

	auctioneer.List(monalisa)
	auctioneer.List(creationOfAdam)
	auctioneer.List(sistineMadonna)

	closingItems := []string{
		monalisa,
		monalisa, // test duplication
		creationOfAdam,
		monalisa, // test duplication
	}

	for _, i := range closingItems {
		auctioneer.Close(i)
	}

	monalisaAuction, _ := house.getItem(monalisa)
	assert.True(t, monalisaAuction.IsClosed())

	creationOfAdamAuction, _ := house.getItem(creationOfAdam)
	assert.True(t, creationOfAdamAuction.IsClosed())

	sistineMadonnaAuction, _ := house.getItem(sistineMadonna)
	assert.False(t, sistineMadonnaAuction.IsClosed())
}
