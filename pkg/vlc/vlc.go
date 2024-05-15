package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

const ChunkSize = 8

type encodingTable map[rune]string
type BinaryChunk string
type BinaryChunks []BinaryChunk
type HexChunck string
type HexChunks []HexChunck

func Encode(str string) string {
	res := prepareText(str)
	chuncks := splitByChunks(encodeBin(res), ChunkSize)
	return chuncks.ToHex().ToString()
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

func splitByChunks(bStr string, size int) BinaryChunks {
	strLen := utf8.RuneCountInString(bStr)
	chunksCount := strLen / size
	if strLen/size != 0 {
		chunksCount++
	}
	res := make(BinaryChunks, 0, chunksCount)
	var buf strings.Builder
	for i, ch := range bStr {
		buf.WriteString(string(ch))
		if (i+1)%ChunkSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}
	if buf.Len() != 0 {
		lastChunk := buf.String()
		lastChunk += strings.Repeat("0", size-len(lastChunk))
		res = append(res, BinaryChunk(lastChunk))
	}
	return res
}

func getEncodingTable() encodingTable {

	return encodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		'm': "000011",
		'i': "01001",
		'a': "011",
		'e': "101",
		's': "0101",
		'd': "00101",
		'!': "001000",
		'y': "0000001",
	}
}

func (bcs BinaryChunks) ToHex() HexChunks {
	res := make(HexChunks, 0, len(bcs))
	for _, chunk := range bcs {
		hexChunk := chunk.toHex()
		res = append(res, hexChunk)
	}
	return res
}
func (bch BinaryChunk) toHex() HexChunck {
	num, err := strconv.ParseUint(string(bch), 2, ChunkSize)
	if err != nil {
		panic("can't parse binary chunk" + err.Error())
	}
	res := strings.ToUpper(fmt.Sprintf("%x", num))
	if len(res) == 1 {
		res = "0" + res
	}
	return HexChunck(res)
}

func (hcs HexChunks) ToString() string {
	const sep = " "
	switch len(hcs) {
	case 0:
		return ""
	case 1:
		return string(hcs[0])
	}
	var buf strings.Builder
	buf.WriteString(string(hcs[0]))
	for _, ch := range hcs[1:] {
		buf.WriteString(sep)
		buf.WriteString(string(ch))
	}
	return buf.String()
}
