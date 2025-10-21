package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int
    switch card{
        case "ace":
        	value=11
        case "king" , "queen", "jack":
        	value=10
        case "two":
        	value=2
        case "three":
        	value=3
        case "four":
        	value=4
        case "five":
        	value=5
        case "six":
        	value=6
        case "seven":
        	value=7
        case "eight":
        	value=8
        case "nine":
        	value=9
        case "ten":
        	value=10
        default:
        	value=0
    }
    return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    sum :=ParseCard(card1)+ParseCard(card2)
	switch {
        case sum==22:
        	return "P"
        case sum==21 && (dealerCard!="ace" && dealerCard!="king" && dealerCard!="queen" && dealerCard!="jack" && dealerCard!="ten"):
        	return "W"
        case sum==21 && (dealerCard =="ace" || dealerCard =="king" || dealerCard =="queen" || dealerCard =="jack" || dealerCard =="ten"):
        	return "S"
        case 17<=sum && sum<=20 :
        	return "S"
        case 12<=sum && sum<=16 && ParseCard(dealerCard)<7 :
        	return "S"
        case 12<=sum && sum<=16 && ParseCard(dealerCard)>=7 :
        	return "H"
        case sum<=11:
        	return "H"
        default :
        	return " "
    }
}
