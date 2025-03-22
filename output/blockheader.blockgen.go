// Code generated by github.com/visvasity/blockgen. DO NOT EDIT.

package output

import (
	"fmt"
	"github.com/visvasity/blockgen/blockgen"
	"iter"
	"sort"
	"strings"
)

// Reader type defines accessor methods for read-only access.
type BlockHeader blockgen.BlockBytes

// Writer type extends the reader with mutable methods.
type BlockHeaderWriter struct{ BlockHeader }

// BlockBytes returns access to the underlying byte slice.
func (v BlockHeader) BlockBytes() blockgen.BlockBytes {
	return blockgen.BlockBytes(v)
}

// Writer returns the BlockHeader writer for read-write access to it's fields.
func (v BlockHeader) Writer() BlockHeaderWriter {
	return BlockHeaderWriter{v}
}

// Reader returns the BlockHeader reader with read-only access to it's fields.
func (v BlockHeaderWriter) Reader() BlockHeader {
	return v.BlockHeader
}

func (v BlockHeader) IsZero() bool {
	return blockgen.IsZero(v[:48])
}

func (v BlockHeaderWriter) SetZero() {
	blockgen.SetZero(v.BlockBytes()[:48])
}

func (v BlockHeader) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "Checksum=[%d]{%x}", v.ChecksumLen(), v.Checksum())
	fmt.Fprintf(&sb, " ")
	fmt.Fprintf(&sb, "BlockLSN=%d", v.BlockLSN())
	fmt.Fprintf(&sb, " ")
	fmt.Fprintf(&sb, "BlockTypeID=%d", v.BlockTypeID())
	return sb.String()
}

func (v BlockHeader) ChecksumLen() int {
	return int(32)
}

func (v BlockHeader) ChecksumItemAt(i int) uint8 {
	if i < 0 || i >= v.ChecksumLen() {
		panic(fmt.Sprintf("array index %d is out of range [0:%d]", i, v.ChecksumLen()))
	}
	return v.BlockBytes().Uint8At(0 + i*1)
}

func (v BlockHeaderWriter) SetChecksumItemAt(i int, x uint8) {
	if i < 0 || i >= v.ChecksumLen() {
		panic(fmt.Sprintf("array index %d is out of range [0:%d]", i, v.ChecksumLen()))
	}
	v.BlockBytes().SetUint8At(0+i*1, x)
}

func (v BlockHeader) Checksum() (xs [32]uint8) {
	for i := range xs {
		xs[i] = v.ChecksumItemAt(i)
	}
	return
}

func (v BlockHeaderWriter) SetChecksum(xs [32]uint8) {
	for i := range xs {
		v.SetChecksumItemAt(i, xs[i])
	}
	return
}

func (v BlockHeader) CopyChecksumTo(xs []uint8) {
	n := min(len(xs), v.ChecksumLen())
	for i := 0; i < n; i++ {
		xs[i] = v.ChecksumItemAt(i)
	}
}

func (v BlockHeaderWriter) CopyChecksumFrom(xs []uint8) {
	n := min(len(xs), v.ChecksumLen())
	for i := 0; i < n; i++ {
		v.SetChecksumItemAt(i, xs[i])
	}
}

func (v BlockHeader) AllChecksum() iter.Seq2[int, uint8] {
	return func(yield func(int, uint8) bool) {
		for i := 0; i < v.ChecksumLen(); i++ {
			if !yield(i, v.ChecksumItemAt(i)) {
				return
			}
		}
	}
}

func (v BlockHeaderWriter) SwapChecksumItems(i, j int) {
	tmp := make([]byte, 1)
	ioff := 0 + i*1
	joff := 0 + j*1
	copy(tmp, v.BlockBytes()[ioff:ioff+1])
	copy(v.BlockBytes()[ioff:ioff+1], v.BlockBytes()[joff:joff+1])
	copy(v.BlockBytes()[joff:joff+1], tmp)
}

func (v BlockHeaderWriter) SortChecksumFunc(cmp func(a, b uint8) int) {
	helper := blockgen.SortHelper{
		LenFunc:     v.ChecksumLen,
		SwapFunc:    v.SwapChecksumItems,
		CompareFunc: func(i, j int) int { return cmp(v.ChecksumItemAt(i), v.ChecksumItemAt(j)) },
	}
	sort.Sort(&helper)
}

func (v BlockHeader) FindChecksumFunc(cmp func(x uint8) int) (int, bool) {
	return sort.Find(v.ChecksumLen(), func(i int) int { return cmp(v.ChecksumItemAt(i)) })
}

func (v BlockHeader) BlockLSN() LSN {
	return LSN(v.BlockBytes().Int64At(32))
}

func (v BlockHeaderWriter) SetBlockLSN(x LSN) {
	v.BlockBytes().SetInt64At(32, int64(x))
}

func (v BlockHeader) BlockTypeID() BlockType {
	return BlockType(v.BlockBytes().Uint16At(40))
}

func (v BlockHeaderWriter) SetBlockTypeID(x BlockType) {
	v.BlockBytes().SetUint16At(40, uint16(x))
}
