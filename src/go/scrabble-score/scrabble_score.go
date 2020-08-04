package scrabble

// import "unicode"

// var onePoints = []rune{'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'}
// var twoPoints = []rune{'D', 'G'}
// var threePoints = []rune{'B', 'C', 'M', 'P'}
// var fourPoints = []rune{'F', 'H', 'V', 'W', 'Y'}
// var fivePoints = []rune{'K'}
// var eightPoints = []rune{'J', 'X'}
// var tenPoints = []rune{'Q', 'Z'}
// var pointLists = [][]rune{onePoints, twoPoints, threePoints, fourPoints, fivePoints, eightPoints, tenPoints}
// var pointIndexes = []int{1, 2, 3, 4, 5, 8, 10}

// // Score returns the scrabble score of a word
// func Score(word string) int {
// 	var lookup = createMap()
// 	var score = 0

// 	for _, ru := range word {
// 		var r = unicode.ToUpper(ru)
// 		score += lookup[r]
// 	}
// 	return score
// }

// func createMap() map[rune]int {
// 	elementMap := make(map[rune]int)

// 	for i, list := range pointLists {
// 		for _, r := range list {
// 			elementMap[r] = pointIndexes[i]
// 		}
// 	}
// 	return elementMap
// }
