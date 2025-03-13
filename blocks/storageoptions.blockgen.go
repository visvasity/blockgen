// Code generated by github.com/visvasity/blockgen. DO NOT EDIT.

package blocks

import (
	"fmt"
	"github.com/visvasity/blockgen/common"
	"strings"
)

// Reader type defines accessor methods for read-only access.
type StorageOptions common.BlockBytes

// Writer type extends the reader with mutable methods.
type StorageOptionsWriter struct{ StorageOptions }

// BlockBytes returns access to the underlying byte slice.
func (v StorageOptions) BlockBytes() common.BlockBytes {
	return common.BlockBytes(v)
}

// Writer returns the StorageOptions writer for read-write access to it's fields.
func (v StorageOptions) Writer() StorageOptionsWriter {
	return StorageOptionsWriter{v}
}

// Reader returns the StorageOptions reader with read-only access to it's fields.
func (v StorageOptionsWriter) Reader() StorageOptions {
	return v.StorageOptions
}

func (v StorageOptions) IsZero() bool {
	return common.IsZero(v[:36])
}

func (v StorageOptionsWriter) SetZero() {
	common.SetZero(v.BlockBytes()[:36])
}

func (v StorageOptions) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "BlockSize=%d", v.BlockSize())
	fmt.Fprintf(&sb, " ")
	fmt.Fprintf(&sb, "MaxLBAMetadataBlockItems=%d", v.MaxLBAMetadataBlockItems())
	fmt.Fprintf(&sb, " ")
	fmt.Fprintf(&sb, "MaxObjectListBlockItems=%d", v.MaxObjectListBlockItems())
	fmt.Fprintf(&sb, " ")
	fmt.Fprintf(&sb, "MaxFreeDBAListBlockItems=%d", v.MaxFreeDBAListBlockItems())
	fmt.Fprintf(&sb, " ")
	fmt.Fprintf(&sb, "MaxFreePBAListBlockItems=%d", v.MaxFreePBAListBlockItems())
	fmt.Fprintf(&sb, " ")
	fmt.Fprintf(&sb, "MaxFreeLBAListBlockItems=%d", v.MaxFreeLBAListBlockItems())
	fmt.Fprintf(&sb, " ")
	fmt.Fprintf(&sb, "MaxNonDataLBAListBlockItems=%d", v.MaxNonDataLBAListBlockItems())
	fmt.Fprintf(&sb, " ")
	fmt.Fprintf(&sb, "MaxFreeDataRegionListBlockItems=%d", v.MaxFreeDataRegionListBlockItems())
	fmt.Fprintf(&sb, " ")
	fmt.Fprintf(&sb, "MaxDeferredPBAListBlockItems=%d", v.MaxDeferredPBAListBlockItems())
	return sb.String()
}

func (v StorageOptions) BlockSize() uint32 {
	return v.BlockBytes().Uint32At(0)
}

func (v StorageOptionsWriter) SetBlockSize(x uint32) {
	v.BlockBytes().SetUint32At(0, x)
}

func (v StorageOptions) MaxLBAMetadataBlockItems() uint32 {
	return v.BlockBytes().Uint32At(4)
}

func (v StorageOptionsWriter) SetMaxLBAMetadataBlockItems(x uint32) {
	v.BlockBytes().SetUint32At(4, x)
}

func (v StorageOptions) MaxObjectListBlockItems() uint32 {
	return v.BlockBytes().Uint32At(8)
}

func (v StorageOptionsWriter) SetMaxObjectListBlockItems(x uint32) {
	v.BlockBytes().SetUint32At(8, x)
}

func (v StorageOptions) MaxFreeDBAListBlockItems() uint32 {
	return v.BlockBytes().Uint32At(12)
}

func (v StorageOptionsWriter) SetMaxFreeDBAListBlockItems(x uint32) {
	v.BlockBytes().SetUint32At(12, x)
}

func (v StorageOptions) MaxFreePBAListBlockItems() uint32 {
	return v.BlockBytes().Uint32At(16)
}

func (v StorageOptionsWriter) SetMaxFreePBAListBlockItems(x uint32) {
	v.BlockBytes().SetUint32At(16, x)
}

func (v StorageOptions) MaxFreeLBAListBlockItems() uint32 {
	return v.BlockBytes().Uint32At(20)
}

func (v StorageOptionsWriter) SetMaxFreeLBAListBlockItems(x uint32) {
	v.BlockBytes().SetUint32At(20, x)
}

func (v StorageOptions) MaxNonDataLBAListBlockItems() uint32 {
	return v.BlockBytes().Uint32At(24)
}

func (v StorageOptionsWriter) SetMaxNonDataLBAListBlockItems(x uint32) {
	v.BlockBytes().SetUint32At(24, x)
}

func (v StorageOptions) MaxFreeDataRegionListBlockItems() uint32 {
	return v.BlockBytes().Uint32At(28)
}

func (v StorageOptionsWriter) SetMaxFreeDataRegionListBlockItems(x uint32) {
	v.BlockBytes().SetUint32At(28, x)
}

func (v StorageOptions) MaxDeferredPBAListBlockItems() uint32 {
	return v.BlockBytes().Uint32At(32)
}

func (v StorageOptionsWriter) SetMaxDeferredPBAListBlockItems(x uint32) {
	v.BlockBytes().SetUint32At(32, x)
}
