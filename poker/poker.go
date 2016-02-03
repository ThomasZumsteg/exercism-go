package poker

import (
	"fmt"
	"sort"
	"strings"
)

//TestVersion the version of the unit test that this will pass
const TestVersion = 1

//ordered rank of poker hands
const (
	highCard = iota
	onePair
	twoPair
	threeOfAKind
	straight
	flush
	fullHouse
	fourOfAKind
	straightFlush
)

/*BestHand finds the winning poker hand.
Reports an error if any of the hands aren'y properly formatted.
If two hand tie, both hands are returned.*/
func BestHand(hands []string) ([]string, error) {
	var bestHand []string

	//first int is type of the hand, remainer are ordered card ranks
	var bestHandRank [6]int
	for _, hand := range hands {
		ranks, suits, err := getCards(hand)
		if err != nil {
			return nil, err
		}

		if handRank := rankHand(ranks, suits); handRank == bestHandRank {
			bestHand = append(bestHand, hand)
		} else if less(bestHandRank, handRank) {
			bestHandRank = handRank
			bestHand = []string{hand}
		}
	}
	return bestHand, nil
}

/*less determines if the first poker hand is lower rank than the second.*/
func less(firstHand, secondHand [6]int) bool {
	for i, v := range firstHand {
		if v != secondHand[i] {
			return v < secondHand[i]
		}
	}
	return false
}

/*rankHand scores a poker hand.
First by determining the type of hand,
Then by sorting the card ranks into their order of comparison.*/
func rankHand(ranks [5]int, suits string) [6]int {
	cardGroups := groupCards(&ranks)

	// Change ranks if ace is played low
	if ranks == [5]int{14, 5, 4, 3, 2} {
		ranks = [5]int{5, 4, 3, 2, 1}
	}
	isFlush, isStraight := isFlush(suits), isStraight(ranks)

	var handRank [6]int
	switch {
	case isFlush && isStraight:
		handRank[0] = straightFlush
	case cardGroups == "41":
		handRank[0] = fourOfAKind
	case cardGroups == "32":
		handRank[0] = fullHouse
	case isFlush:
		// Reorder so that flush can have multiple cards of the same value
		sort.Sort(sort.Reverse(sort.IntSlice(ranks[:])))
		handRank[0] = flush
	case isStraight:
		handRank[0] = straight
	case cardGroups == "311":
		handRank[0] = threeOfAKind
	case cardGroups == "221":
		handRank[0] = twoPair
	case cardGroups == "2111":
		handRank[0] = onePair
	default:
		handRank[0] = highCard
	}
	copy(handRank[1:6], ranks[:])
	return handRank
}

/*groupCards orders ranks to match their importance.
Largest matches come first then order by value.*/
func groupCards(ranks *[5]int) string {
	//Group and count the ranks
	rankGroups := make(map[int]int)
	for _, r := range ranks {
		rankGroups[r]++
	}

	//Intermediate sortable data structure
	var rankList [][2]int
	for k, v := range rankGroups {
		rankList = append(rankList, [2]int{v, k})
	}

	sort.Sort(byGroupAndRank(rankList))

	//Copy to ranks and record groupings
	var groupSize string
	i := 0
	for _, group := range rankList {
		groupSize += fmt.Sprintf("%1d", group[0])
		stop := i + group[0]
		for ; i < stop; i++ {
			ranks[i] = group[1]
		}
	}
	return groupSize
}

//byGroupAndRank sorts by group size than rank
type byGroupAndRank [][2]int

func (by byGroupAndRank) Len() int      { return len(by) }
func (by byGroupAndRank) Swap(i, j int) { by[i], by[j] = by[j], by[i] }
func (by byGroupAndRank) Less(i, j int) bool {
	if by[i][0] != by[j][0] {
		return by[i][0] > by[j][0]
	}
	return by[i][1] > by[j][1]
}

/*isFlush determines if all the suits are the same.*/
func isFlush(suits string) bool {
	suit := ' '
	for _, s := range suits {
		if suit == ' ' || s == suit {
			suit = s
		} else {
			return false
		}
	}
	return true
}

/*isStraight determines if the cards are sequential.*/
func isStraight(ranks [5]int) bool {
	rank := -1
	for _, r := range ranks {
		if rank == -1 || rank-1 == r {
			rank = r
		} else {
			return false
		}
	}
	return true
}

/*getCards validates a hand and splits the cards into suits and ranks.*/
func getCards(hand string) ([5]int, string, error) {
	cards := strings.Fields(hand)
	if len(cards) != 5 {
		return [5]int{}, "", fmt.Errorf("Not a valid hand: %s", hand)
	}
	var ranks [5]int
	var suits string

	for i, card := range cards {
		cardRunes := strings.Split(card, "")
		var rank int
		var suit string
		if l := len(cardRunes); l == 2 {
			rank, suit = getRank(cardRunes[0]), cardRunes[1]
		} else if l == 3 {
			rank, suit = getRank(strings.Join(cardRunes[:2], "")), cardRunes[2]
		} else {
			return [5]int{}, "", fmt.Errorf("Not a valid card: %s", card)
		}

		if rank == -1 || !strings.Contains("♡♢♧♤", suit) {
			return [5]int{}, "", fmt.Errorf("Not a valid card: %s", card)
		}
		suits += string(suit)
		ranks[i] = rank
	}
	return ranks, suits, nil
}

/*getRank converts a cards string rank into an integer rank*/
func getRank(rank string) int {
	switch rank {
	case "A":
		return 14
	case "2", "3", "4", "5", "6", "7", "8", "9":
		return int(rank[0] - '0')
	case "10":
		return 10
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	default:
		return -1
	}
}
