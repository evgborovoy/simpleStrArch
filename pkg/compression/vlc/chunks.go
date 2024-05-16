package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type encodingTable map[rune]string
type BinaryChunk string
type BinaryChunks []BinaryChunk

const ChunkSize = 8

func NewBinChunks(data []byte) BinaryChunks {
	res := make(BinaryChunks, 0, len(data))
	for _, p := range data {
		res = append(res, NewBinChunk(p))
	}
	return res
}

func NewBinChunk(code byte) BinaryChunk {
	return BinaryChunk(fmt.Sprintf("%08b", code))
}
func (bch BinaryChunks) Bytes() []byte {
	res := make([]byte, 0, len(bch))
	for _, ch := range bch {
		res = append(res, ch.Byte())

	}
	return res
}

func (bc BinaryChunk) Byte() byte {
	num, err := strconv.ParseUint(string(bc), 2, ChunkSize)
	if err != nil {
		panic("can't parse binary chunk" + err.Error())
	}
	return byte(num)
}

// Join chunks into one string and return result
func (bks BinaryChunks) Join() string {
	var buf strings.Builder
	for _, chunk := range bks {
		buf.WriteString(string(chunk))
	}
	return buf.String()
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
