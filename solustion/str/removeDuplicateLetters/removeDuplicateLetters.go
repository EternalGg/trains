package removeDuplicateLetters

//
//func RemoveDuplicateLetters(s string) string {
//	bList := []byte(s)
//	bMap := make(map[byte]int)
//	for i := 0; i < len(bList); i++ {
//		bMap[bList[i]]++
//	}
//	result, key := "", byte('a')
//	for i := 0; i < len(bList); i++ {
//		switch bMap[bList[i]] {
//		case 1:
//			key--
//			if bList[i] > key {
//				for j := i; j >= 0; j-- {
//					if key!= {
//
//					}
//				}
//			} else if bList[i] == key {
//				result += string(bList[i])
//				bMap[bList[i]] = 0
//			}
//		case 0:
//			continue
//		default:
//			if bList[i] == key {
//				result += string(bList[i])
//				bMap[bList[i]] = 0
//				for j := key; j < 122; j++ {
//					if bMap[j] != 0 {
//						key = j
//					}
//				}
//			} else {
//				bMap[bList[i]]--
//			}
//
//		}
//	}
//	return result
//}
