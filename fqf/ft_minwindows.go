package tpmongorinluisjy

// import (
// 	"math"
// )

// func Ft_min_window(s string, t string) string {
// 	if len(s) < len(t) || len(t) == 0 {
// 		return ""
// 	}

// 	tFreq := make(map[byte]int)
// 	for i := 0; i < len(t); i++ {
// 		tFreq[t[i]]++
// 	}

// 	windowFreq := make(map[byte]int)
// 	left, right, formed := 0, 0, 0
// 	required := len(tFreq)

// 	minLen := math.MaxInt32
// 	minLeft := 0

// 	for right < len(s) {
// 		charRight := s[right]
// 		windowFreq[charRight]++

// 		if count, found := tFreq[charRight]; found && windowFreq[charRight] == count {
// 			formed++
// 		}

// 		for left <= right && formed == required {
// 			charLeft := s[left]

// 			if right-left+1 < minLen {
// 				minLen = right - left + 1
// 				minLeft = left
// 			}

// 			windowFreq[charLeft]--
// 			if count, found := tFreq[charLeft]; found && windowFreq[charLeft] < count {
// 				formed--
// 			}

// 			left++
// 		}

// 		right++
// 	}

// 	if minLen == math.MaxInt32 {
// 		return ""
// 	}

// 	return s[minLeft : minLeft+minLen]
// }
