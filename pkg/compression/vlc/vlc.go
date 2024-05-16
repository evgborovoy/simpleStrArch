package vlc

import (
	"strings"
	"unicode"
)

type EncoderDecoder struct{}

func New() EncoderDecoder {
	return EncoderDecoder{}
}

func (_ EncoderDecoder) Encode(str string) []byte {
	res := prepareText(str)
	chuncks := splitByChunks(encodeBin(res), ChunkSize)
	return chuncks.Bytes()
}

func (_ EncoderDecoder) Decode(encData []byte) string {
	bChuks := NewBinChunks(encData).Join()
	dTree := getEncodingTable().DecodingTree()

	return exportText(dTree.Decode(bChuks))
}

// Replace !+<lower symbol> to upper. Opposite func prepareText
// !hello -> Hello
func exportText(str string) string {
	var buf strings.Builder
	var isCapital bool

	for _, ch := range str {
		if isCapital {
			buf.WriteRune(unicode.ToUpper(ch))
			isCapital = false
			continue
		}
		if ch == '!' {
			isCapital = true
			continue
		} else {
			buf.WriteRune(ch)
		}

	}
	return buf.String()
}

func prepareText(str string) string {
	var buf strings.Builder
	for _, ch := range str {
		if unicode.IsUpper(ch) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}

func encodeBin(str string) string {
	var buf strings.Builder
	for _, ch := range str {
		buf.WriteString(bin(ch))
	}
	return buf.String()
}

func bin(ch rune) string {
	table := getEncodingTable()
	res, ok := table[ch]
	if !ok {
		panic("unknowm character: " + string(ch))
	}
	return res
}

func getEncodingTable() encodingTable {

	return encodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'c': "000101",
		'm': "000011",
		'g': "0000100",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'q': "000000000001",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		'h': "0011",
		'l': "001001",
		'u': "00011",
		'f': "000100",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "000000001",
		'x': "00000000001",
		'z': "000000000000",
	}
}