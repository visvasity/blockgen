// Code generated by github.com/visvasity/blockgen. DO NOT EDIT.

package output

import (
	"fmt"
	"github.com/visvasity/blockgen/blockgen"
	"strings"
)

// Reader type defines accessor methods for read-only access.
type Region blockgen.BlockBytes

// Writer type extends the reader with mutable methods.
type RegionWriter struct{ Region }

// BlockBytes returns access to the underlying byte slice.
func (v Region) BlockBytes() blockgen.BlockBytes {
	return blockgen.BlockBytes(v)
}

// Writer returns the Region writer for read-write access to it's fields.
func (v Region) Writer() RegionWriter {
	return RegionWriter{v}
}

// Reader returns the Region reader with read-only access to it's fields.
func (v RegionWriter) Reader() Region {
	return v.Region
}

func (v Region) IsZero() bool {
	return blockgen.IsZero(v[:16])
}

func (v RegionWriter) SetZero() {
	blockgen.SetZero(v.BlockBytes()[:16])
}

func (v Region) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "BeginPBA=%d", v.BeginPBA())
	fmt.Fprintf(&sb, " ")
	fmt.Fprintf(&sb, "EndPBA=%d", v.EndPBA())
	return sb.String()
}

func (v Region) BeginPBA() PBA {
	return PBA(v.BlockBytes().Uint64At(0))
}

func (v RegionWriter) SetBeginPBA(x PBA) {
	v.BlockBytes().SetUint64At(0, uint64(x))
}

func (v Region) EndPBA() PBA {
	return PBA(v.BlockBytes().Uint64At(8))
}

func (v RegionWriter) SetEndPBA(x PBA) {
	v.BlockBytes().SetUint64At(8, uint64(x))
}
