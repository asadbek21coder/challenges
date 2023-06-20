package challenges

// Online Auction System:
// Create an Auction struct with fields Item, MinimumBid, and Bids. Implement methods to place bids, check if a bid is valid, and display the highest bid.

import "fmt"

type Bid struct {
	Bidder string
	Amount float64
}

type Auction struct {
	Item          string
	MinimumBid    float64
	Bids          []Bid
	HighestBid    float64
	HighestBidder string
}

func (a *Auction) PlaceBid(bid Bid) {
	if a.isValidBid(bid.Amount) {
		a.Bids = append(a.Bids, bid)
		if bid.Amount > a.HighestBid {
			a.HighestBid = bid.Amount
			a.HighestBidder = bid.Bidder
		}
		fmt.Println("Bid placed.")
	} else {
		fmt.Println("Invalid bid amount.")
	}
}

func (a Auction) isValidBid(amount float64) bool {
	return amount >= a.MinimumBid
}

func (a Auction) DisplayHighestBid() {
	fmt.Printf("Highest Bid: %.2f by %s\n", a.HighestBid, a.HighestBidder)
}

// func main() {
// 	auction := Auction{
// 		Item:       "Painting",
// 		MinimumBid: 100.0,
// 	}

// 	auction.PlaceBid(Bid{Bidder: "John", Amount: 150.0})
// 	auction.PlaceBid(Bid{Bidder: "Jane", Amount: 200.0})
// 	auction.PlaceBid(Bid{Bidder: "Alice", Amount: 180.0})

// 	auction.DisplayHighestBid()
// }
// Output:

// csharp
// Copy code
// Bid placed.
// Bid placed.
// Invalid bid amount.
// Highest Bid: 200.00 by Jane
