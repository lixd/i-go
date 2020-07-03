package crc32

import (
	"hash/crc32"
	"hash/crc64"
)

func HashCRC32(strKey string) uint32 {
	table := crc32.MakeTable(crc32.IEEE)
	ret := crc32.Checksum([]byte(strKey), table)
	return ret
}

func HashCRC64(strKey string) uint64 {
	table := crc64.MakeTable(crc64.ISO)
	ret := crc64.Checksum([]byte(strKey), table)
	return ret
}
