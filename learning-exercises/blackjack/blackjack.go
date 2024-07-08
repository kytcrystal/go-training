package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card{
	case "ace":
		return 11 
	case "two":
		return 2 
	case "three":
		return 3 
	case "four":
		return 4 
	case "five":
		return 5 
	case "six":
		return 6 
	case "seven":
		return 7 
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	const Stand = "S"
	const Hit = "H"
	const Split = "P"
	const Win = "W"
	card1Value := ParseCard(card1)
	card2Value := ParseCard(card2) 
	sumCardsValue := card1Value + card2Value
	switch {
	case card1 == "ace" && card2 == "ace":
		return Split
	case sumCardsValue == 21 :
		if (ParseCard(dealerCard) < 10) {
			return Win
		} else {
			return Stand
		}
	case sumCardsValue >= 17 && sumCardsValue <= 20:
		return Stand
	case sumCardsValue >= 12 && sumCardsValue <= 16:
		if ParseCard(dealerCard) < 7 {
			return Stand
		} else {
			return Hit
		}
	default:
		return Hit
	}
}
