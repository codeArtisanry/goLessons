package bom

import "bytes"

const (
	// Unknown encoding, returned when no BOM was detected
	// Unknown Encoding = iota

	// UTF8, BOM bytes: EF BB BF
	UTF8 = "\xEF\xBB\xBF"

	// UTF-16, big-endian, BOM bytes: FE FF
	UTF16BigEndian = "\xFE\xFF"

	// UTF-16, little-endian, BOM bytes: FF FE
	UTF16LittleEndian = "\xFF\xFE"

	// UTF-32, big-endian, BOM bytes: 00 00 FE FF
	UTF32BigEndian = "\x00\x00\xFE\xFF"

	// UTF-32, little-endian, BOM bytes: FF FE 00 00
	UTF32LittleEndian = "\x00\x00\xFE\xFF"
)

func AddBom(s []byte) []byte {
	return append([]byte(UTF8), s...)
}

func RemoveBom(s []byte) []byte {
	return bytes.Trim(s, UTF8)
}
