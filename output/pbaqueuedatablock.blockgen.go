// Code generated by github.com/visvasity/blockgen. DO NOT EDIT.

package output

import (
	"fmt"
	"github.com/visvasity/blockgen/common"
	"iter"
	"sort"
	"strings"
)

// Reader type defines accessor methods for read-only access.
type PBAQueueDataBlock common.BlockBytes

// Writer type extends the reader with mutable methods.
type PBAQueueDataBlockWriter struct{ PBAQueueDataBlock }

// BlockBytes returns access to the underlying byte slice.
func (v PBAQueueDataBlock) BlockBytes() common.BlockBytes {
	return common.BlockBytes(v)
}

// Writer returns the PBAQueueDataBlock writer for read-write access to it's fields.
func (v PBAQueueDataBlock) Writer() PBAQueueDataBlockWriter {
	return PBAQueueDataBlockWriter{v}
}

// Reader returns the PBAQueueDataBlock reader with read-only access to it's fields.
func (v PBAQueueDataBlockWriter) Reader() PBAQueueDataBlock {
	return v.PBAQueueDataBlock
}

func (v PBAQueueDataBlock) IsZero() bool {
	return common.IsZero(v[:40])
}

func (v PBAQueueDataBlockWriter) SetZero() {
	common.SetZero(v.BlockBytes()[:40])
}

func (v PBAQueueDataBlock) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "NextLBA=%d", v.NextLBA())
	fmt.Fprintf(&sb, " ")
	fmt.Fprintf(&sb, "TestNumericValue=%d", v.TestNumericValue())
	fmt.Fprintf(&sb, " ")
	fmt.Fprintf(&sb, "ValuesSlice=[:%d:%d]{", v.ValuesSliceLen(), v.ValuesSliceCap())
	for i := 0; i < v.ValuesSliceLen(); i++ {
		if i == 0 {
			fmt.Fprintf(&sb, "%d", v.ValuesSliceItemAt(i))
		} else {
			fmt.Fprintf(&sb, " %d", v.ValuesSliceItemAt(i))
		}
	}
	fmt.Fprintf(&sb, "}")
	return sb.String()
}

func (v PBAQueueDataBlock) NextLBA() LBA {
	return LBA(v.BlockBytes().Uint64At(0))
}

func (v PBAQueueDataBlockWriter) SetNextLBA(x LBA) {
	v.BlockBytes().SetUint64At(0, uint64(x))
}

func (v PBAQueueDataBlock) TestNumericValue() PBA {
	return PBA(v.BlockBytes().Uint64At(8))
}

func (v PBAQueueDataBlockWriter) SetTestNumericValue(x PBA) {
	v.BlockBytes().SetUint64At(8, uint64(x))
}

func (v PBAQueueDataBlock) ValuesSliceLen() int {
	return int(v.BlockBytes().Int64At(24))
}

func (v PBAQueueDataBlock) ValuesSliceCap() int {
	return int(v.BlockBytes().Int64At(32))
}

func (v PBAQueueDataBlockWriter) internalSetValuesSliceLen(x int) {
	v.BlockBytes().SetInt64At(24, int64(x))
}

func (v PBAQueueDataBlockWriter) internalSetValuesSliceCap(x int) {
	v.BlockBytes().SetInt64At(32, int64(x))
}

func (v PBAQueueDataBlock) ValuesSliceItemAt(i int) PBA {
	if i < 0 || i >= v.ValuesSliceLen() {
		panic(fmt.Sprintf("slice index %d is out of range [0:%d:%d]", i, v.ValuesSliceLen(), v.ValuesSliceCap()))
	}
	return PBA(v.BlockBytes().Uint64At(40 + i*8))
}

func (v PBAQueueDataBlockWriter) SetValuesSliceItemAt(i int, x PBA) {
	if i < 0 || i >= v.ValuesSliceLen() {
		panic(fmt.Sprintf("slice index %d is out of range [0:%d:%d]", i, v.ValuesSliceLen(), v.ValuesSliceCap()))
	}
	v.BlockBytes().SetUint64At(40+i*8, uint64(x))
}

func (v PBAQueueDataBlockWriter) AppendValuesSlice(x PBA) {
	n := v.ValuesSliceLen()
	if n == v.ValuesSliceCap() {
		panic(fmt.Sprintf("slice is already full with %d items", n))
	}
	v.internalSetValuesSliceLen(n + 1)
	v.BlockBytes().SetUint64At(40+n*8, uint64(x))
}

func (v PBAQueueDataBlock) CopyValuesSliceTo(xs []PBA) {
	n := min(len(xs), v.ValuesSliceLen())
	for i := 0; i < n; i++ {
		xs[i] = v.ValuesSliceItemAt(i)
	}
}

func (v PBAQueueDataBlockWriter) CopyValuesSliceFrom(xs []PBA) {
	n := min(len(xs), v.ValuesSliceLen())
	for i := 0; i < n; i++ {
		v.SetValuesSliceItemAt(i, xs[i])
	}
}

func (v PBAQueueDataBlockWriter) RemoveValuesSliceItemAt(i int) {
	n := v.ValuesSliceLen()
	if i < 0 || i >= n {
		panic(fmt.Sprintf("slice index %d is out of range [0:%d:%d]", i, v.ValuesSliceLen(), v.ValuesSliceCap()))
	}
	beg := 40 + i*8
	end := 40 + n*8
	copy(v.BlockBytes()[beg:], v.BlockBytes()[beg+8:end])
	common.SetZero(v.BlockBytes()[end-8 : end])
	v.internalSetValuesSliceLen(n - 1)
}

func (v PBAQueueDataBlockWriter) DeleteValuesSliceItems(i, j int) {
	n := v.ValuesSliceLen()
	if i < 0 || i >= n {
		panic(fmt.Sprintf("first slice index %d is out of range [0:%d:%d]", i, v.ValuesSliceLen(), v.ValuesSliceCap()))
	}
	if j < 0 || j >= n {
		panic(fmt.Sprintf("second slice index %d is out of range [0:%d:%d]", i, v.ValuesSliceLen(), v.ValuesSliceCap()))
	}
	if j < i {
		panic(fmt.Sprintf("invalid slice indices %d < %d", j, i))
	}
	if i == j {
		return
	}
	ioff := 40 + i*8
	joff := 40 + j*8
	end := 40 + n*8
	copy(v.BlockBytes()[ioff:end], v.BlockBytes()[joff:end])
	common.SetZero(v.BlockBytes()[end-(joff-ioff) : end])
	v.internalSetValuesSliceLen(n - (j - i))
}

func (v PBAQueueDataBlock) AllValuesSlice() iter.Seq2[int, PBA] {
	return func(yield func(int, PBA) bool) {
		for i := 0; i < v.ValuesSliceLen(); i++ {
			if !yield(i, v.ValuesSliceItemAt(i)) {
				return
			}
		}
	}
}

func (v PBAQueueDataBlockWriter) SwapValuesSliceItems(i, j int) {
	tmp := make([]byte, 8)
	ioff := 40 + i*8
	joff := 40 + j*8
	copy(tmp, v.BlockBytes()[ioff:ioff+8])
	copy(v.BlockBytes()[ioff:ioff+8], v.BlockBytes()[joff:joff+8])
	copy(v.BlockBytes()[joff:joff+8], tmp)
}

func (v PBAQueueDataBlockWriter) SortValuesSliceFunc(cmp func(a, b PBA) int) {
	helper := common.SortHelper{
		LenFunc:     v.ValuesSliceLen,
		SwapFunc:    v.SwapValuesSliceItems,
		CompareFunc: func(i, j int) int { return cmp(v.ValuesSliceItemAt(i), v.ValuesSliceItemAt(j)) },
	}
	sort.Sort(&helper)
}

func (v PBAQueueDataBlock) FindValuesSliceFunc(cmp func(x PBA) int) (int, bool) {
	return sort.Find(v.ValuesSliceLen(), func(i int) int { return cmp(v.ValuesSliceItemAt(i)) })
}

func PBAQueueDataBlockValuesSliceCapForNumBytes(nbytes int) int {
	return (nbytes - 40) / 8
}

// NewPBAQueueDataBlock creates a zero-initialized PBAQueueDataBlock. Returns nil if input block size is too small.
func NewPBAQueueDataBlock(block []byte) PBAQueueDataBlock {
	size := len(block)
	if size < 40 {
		return nil
	}
	common.SetZero(block)
	v := PBAQueueDataBlock(block)
	// PBAQueueDataBlock type has a slice field; we must set a cap on it.
	n := (size - 40) / 8
	v.Writer().internalSetValuesSliceCap(n)
	return v
}

func OpenPBAQueueDataBlock(block []byte) (PBAQueueDataBlock, error) {
	size := len(block)
	if size < 40 {
		return nil, fmt.Errorf("input size is too small")
	}
	v := PBAQueueDataBlock(block)
	// PBAQueueDataBlock type has a slice field; validate it's len and cap.
	n := (size - 40) / 8
	if x := v.ValuesSliceCap(); x != n {
		return nil, fmt.Errorf("slice field cap must be %d, found %d", n, x)
	}
	if x := v.ValuesSliceLen(); x < 0 || x > n {
		return nil, fmt.Errorf("slice field len is %d, must be between [%d-%d)", x, 0, n)
	}
	return v, nil
}
