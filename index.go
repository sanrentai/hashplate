package hashpalte

import (
	"fmt"
)

func Hashplate(value string, hasEmoji ...bool) string {
	seed := getSeedFromString(value)
	rng := splitMix32(seed)

	// 京A 123TV
	province := []rune("黑吉辽京津晋冀鲁豫蒙沪渝苏浙皖闽湘赣鄂桂琼川贵云藏陕甘宁青新粤")
	alphabet := "ABCDEFGHJKLMNPQRSTUVWXYZ"               // removed "I" and "O"
	alphanumeric := "ABCDEFGHJKLMNPQRSTUVWXYZ0123456789" // removed "I" and "O"

	rand_p := int(rng() * float64(len(province)))
	rand_alphabet := int(rng() * float64(len(alphabet)))
	randArr := make([]int, 5)
	for i := 0; i < 5; i++ {
		randArr[i] = int(rng() * float64(len(alphanumeric)))
	}

	result_no_emoji := fmt.Sprintf("%c%c·%c%c%c%c%c",
		province[rand_p],
		alphabet[rand_alphabet],
		alphanumeric[randArr[0]],
		alphanumeric[randArr[1]],
		alphanumeric[randArr[2]],
		alphanumeric[randArr[3]],
		alphanumeric[randArr[4]])

	if len(hasEmoji) > 0 && hasEmoji[0] {
		result_emoji := fmt.Sprintf("%s %s %s",
			emojiDictionary[int(rng()*float64(len(emojiDictionary)))],
			result_no_emoji,
			emojiDictionary[int(rng()*float64(len(emojiDictionary)))])
		return result_emoji
	}

	return result_no_emoji
}
