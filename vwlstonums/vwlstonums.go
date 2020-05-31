//Package vwlstonums encodes/decodes vowels with numbers
//  'a': '1'
// 	'e': '2'
// 	'i': '3'
// 	'o': '4'
// 	'u': '5'
package vwlstonums

var vowelsToNumbers = map[rune]rune{
	'a': '1',
	'e': '2',
	'i': '3',
	'o': '4',
	'u': '5',
	'A': '1',
	'E': '2',
	'I': '3',
	'O': '4',
	'U': '5',
}

var numbersToVowels = map[rune]rune{
	'1': 'a',
	'2': 'e',
	'3': 'i',
	'4': 'o',
	'5': 'u',
}

// Encode replaces vowels with numbers
//  'a': '1'
// 	'e': '2'
// 	'i': '3'
// 	'o': '4'
// 	'u': '5'
func Encode(s string) string {
	if len(s) < 1 {
		return s
	}

	res := ""
	for _, char := range s {
		if rchar, ok := vowelsToNumbers[char]; ok {
			res += string(rchar)
		} else {
			res += string(char)
		}
	}

	return res
}

// Decode replaces vowels with numbers
//  '1': 'a'
// 	'2': 'e'
// 	'3': 'i'
// 	'4': 'o'
// 	'5': 'u'
func Decode(s string) string {
	if len(s) < 1 {
		return s
	}

	res := ""
	for _, char := range s {
		if rchar, ok := numbersToVowels[char]; ok {
			res += string(rchar)
		} else {
			res += string(char)
		}
	}

	return res
}
