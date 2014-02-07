package slug

import (
	"unicode"
)

const (
	DASH       rune = '-'
	UNDERSCORE rune = '_'
)

type w []rune

var (
	WORDMAP = map[rune]w{
		'&': w("and"), '♥': w("love"),
		'Æ': w("AE"), 'ẞ': w("ss"),
		'þ': w("th"), 'æ': w("ae"),
		'ß': w("ss"), 'Þ': w("th"),
	}

	CHARMAP = map[rune]rune{
		'À': 'a', 'Á': 'a', 'Â': 'a', 'Ã': 'a', 'Ä': 'a', 'Å': 'a', 'Ç': 'c',
		'È': 'e', 'É': 'e', 'Ê': 'e', 'Ë': 'e', 'Ì': 'i', 'Í': 'i', 'Î': 'i',
		'Ï': 'i', 'Ð': 'd', 'Ñ': 'n', 'Ò': 'o', 'Ó': 'o', 'Ô': 'o', 'Õ': 'o',
		'Ö': 'o', 'Ő': 'o', 'Ø': 'o', 'Ù': 'u', 'Ú': 'u', 'Û': 'u', 'Ü': 'u',
		'Ű': 'u', 'Ý': 'y', 'à': 'a', 'á': 'a', 'â': 'a', 'ã': 'a', 'ä': 'a',
		'å': 'a', 'ç': 'c', 'è': 'e', 'é': 'e', 'ê': 'e', 'ë': 'e', 'ì': 'i',
		'í': 'i', 'î': 'i', 'ï': 'i', 'ð': 'd', 'ñ': 'n', 'ò': 'o', 'ó': 'o',
		'ô': 'o', 'õ': 'o', 'ö': 'o', 'ő': 'o', 'ø': 'o', 'ù': 'u', 'ú': 'u',
		'û': 'u', 'ü': 'u', 'ű': 'u', 'ý': 'y', 'ÿ': 'y',
	}
)

func Slug(str string) string {
	var (
		sep    rune
		result []rune
	)

	for _, v := range str {

		var replacement []rune

		switch {
		case v == DASH, v == UNDERSCORE:
			if len(result) > 0 && sep == 0 {
				sep = v
			}
		case v >= '0' && v <= '9', v >= 'a' && v <= 'z':
			replacement = []rune{v}
		case v >= 'A' && v <= 'Z':
			replacement = []rune{unicode.ToLower(v)}
		default:
			if c, ok := CHARMAP[v]; ok {
				replacement = []rune{c}
			} else if word, ok := WORDMAP[v]; ok {

				if sep == 0 {
					sep = DASH
				}

				replacement = word
			}
		}
		switch ln := len(replacement); ln {
		case 0:
			sep = DASH
		default:
			if len(result) > 0 && sep > 0 {
				result = append(result, sep)
				if ln == 1 {
					sep = 0
				}
			}
			result = append(result, replacement...)
		}

	}

	return string(result)
}
