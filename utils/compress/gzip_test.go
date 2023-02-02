package compress

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"i-go/utils/convert"
)

const (
	LocusPoint = `[{"X": 162, "Y": 209}, {"X": 164, "Y": 198}, {"X": 165, "Y": 186}, {"X": 166, "Y": 175}, {"X": 167, "Y": 163}, {"X": 168, "Y": 152}, {"X": 169, "Y": 140}, {"X": 170, "Y": 129}, {"X": 172, "Y": 117}, {"X": 173, "Y": 106}, {"X": 175, "Y": 94}, {"X": 178, "Y": 83}, {"X": 181, "Y": 71}, {"X": 186, "Y": 60}, {"X": 194, "Y": 49}, {"X": 206, "Y": 42}, {"X": 217, "Y": 41}, {"X": 228, "Y": 43}, {"X": 240, "Y": 47}, {"X": 251, "Y": 54}, {"X": 263, "Y": 61}, {"X": 274, "Y": 69}, {"X": 286, "Y": 79}, {"X": 297, "Y": 90}, {"X": 307, "Y": 102}, {"X": 316, "Y": 113}, {"X": 323, "Y": 125}, {"X": 327, "Y": 136}, {"X": 328, "Y": 147}, {"X": 321, "Y": 159}, {"X": 311, "Y": 163}, {"X": 300, "Y": 165}, {"X": 288, "Y": 166}, {"X": 277, "Y": 166}, {"X": 265, "Y": 166}, {"X": 254, "Y": 164}, {"X": 243, "Y": 163}, {"X": 231, "Y": 162}, {"X": 220, "Y": 160}, {"X": 208, "Y": 158}, {"X": 197, "Y": 156}, {"X": 185, "Y": 154}, {"X": 174, "Y": 152}, {"X": 162, "Y": 149}, {"X": 151, "Y": 147}, {"X": 139, "Y": 145}, {"X": 128, "Y": 142}, {"X": 116, "Y": 140}, {"X": 105, "Y": 138}, {"X": 93, "Y": 136}, {"X": 82, "Y": 134}, {"X": 70, "Y": 132}]`
)

func TestCompress(t *testing.T) {
	compress := Compress(convert.String2Bytes(LocusPoint))
	create, err := os.Create("1.gz")
	if err != nil {
		return
	}
	defer create.Close()
	create.Write(compress)
	deCompress, err := DeCompress(compress)
	if err != nil {
		t.Fatal(err)
	}
	bytes2String := convert.Bytes2String(deCompress)
	fmt.Println(bytes2String == LocusPoint)
}

// 169877 ns/op 0.1ms
func BenchmarkCompress(b *testing.B) {
	origin := convert.String2Bytes(LocusPoint)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Compress(origin)
	}
}

// 16502 ns/op 0.01mns
func BenchmarkDeCompress(b *testing.B) {
	origin := convert.String2Bytes(LocusPoint)
	compress := Compress(origin)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = DeCompress(compress)
	}
}

// IsGzipFile check file is compressed by gzip.
func IsGzipFile(path string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}
	// https://www.ietf.org/rfc/rfc1952.txt
	// There is a magic number at the beginning of the file.
	// Just read the first two bytes and check if they are equal to 0x1f8b.
	magic := make([]byte, 2)
	read, err := file.Read(magic)
	if err != nil {
		return false, err
	}
	if read != 2 {
		return false, errors.New("read magic number failed")
	}
	return magic[0] == 31 && magic[1] == 139, nil
}

func TestA(t *testing.T) {
	path := "/Users/lixueduan/17x/projects/i-go/utils/compress/1.gz"
	//path := "/Users/lixueduan/17x/projects/i-go/utils/compress/Docker.dmg"
	ok, err := IsGzipFile(path)
	if err != nil {
		t.Log("check fail:", err)
		return
	}
	t.Log("is gzip file:", ok)
}
