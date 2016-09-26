package mbconv

import (
	"strings"
	"unicode"

	"golang.org/x/text/width"
)

const (
	hiraganaLo = 0x3041 // ぁ
	hiraganaHi = 0x3096 // ゖ

	katakanaLo = 0x30a1 // ァ
	katakanaHi = 0x30f6 // ヶ
)

var (
	jaCase = unicode.SpecialCase{
		unicode.CaseRange{
			// Hiragana To Katakana
			Lo: hiraganaLo, Hi: hiraganaHi,
			Delta: [unicode.MaxCase]rune{
				katakanaLo - hiraganaLo,
				0,
				0,
			},
		},
		unicode.CaseRange{
			// Katakana To Hiragana
			Lo: katakanaLo, Hi: katakanaHi,
			Delta: [unicode.MaxCase]rune{
				0,
				hiraganaLo - katakanaLo,
				0,
			},
		},
	}

	halfToFullReplacer = strings.NewReplacer(
		"ﾞ", "゛",
		"ﾟ", "゜",
	)

	fullToHalfReplacer = strings.NewReplacer(
		"ガ", "ｶﾞ",
		"ギ", "ｷﾞ",
		"グ", "ｸﾞ",
		"ゲ", "ｹﾞ",
		"ゴ", "ｺﾞ",
		"ザ", "ｻﾞ",
		"ジ", "ｼﾞ",
		"ズ", "ｽﾞ",
		"ゼ", "ｾﾞ",
		"ゾ", "ｿﾞ",
		"ダ", "ﾀﾞ",
		"ヂ", "ﾁﾞ",
		"ヅ", "ﾂﾞ",
		"デ", "ﾃﾞ",
		"ド", "ﾄﾞ",
		"バ", "ﾊﾞ",
		"パ", "ﾊﾟ",
		"ビ", "ﾋﾞ",
		"ピ", "ﾋﾟ",
		"ブ", "ﾌﾞ",
		"プ", "ﾌﾟ",
		"ベ", "ﾍﾞ",
		"ペ", "ﾍﾟ",
		"ボ", "ﾎﾞ",
		"ポ", "ﾎﾟ",
		"ヮ", "ﾜ",
		"ヰ", "ｲ",
		"ヱ", "ｴ",
		"ヴ", "ｳﾞ",
	)
)

// HiraganaToKatakana converts string.
func HiraganaToKatakana(str string) string {
	src := []rune(str)
	dst := make([]rune, len(src))
	for i, r := range src {
		switch {
		case unicode.In(r, unicode.Hiragana):
			dst[i] = jaCase.ToUpper(r)
		default:
			dst[i] = r
		}
	}
	return string(dst)
}

// KatakanaToHiragana converts string.
func KatakanaToHiragana(str string) string {
	src := []rune(str)
	dst := make([]rune, len(src))
	for i, r := range src {
		switch {
		case unicode.In(r, unicode.Katakana):
			dst[i] = jaCase.ToLower(r)
		default:
			dst[i] = r
		}
	}
	return string(dst)
}

// HalfWidthToFullWidth converts string.
func HalfWidthToFullWidth(str string) string {
	return width.Widen.String(halfToFullReplacer.Replace(str))
}

// FullWidthToHalfWidth converts string.
func FullWidthToHalfWidth(str string) string {
	return width.Narrow.String(fullToHalfReplacer.Replace(str))
}
