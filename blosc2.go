// Package blosc2 wraps blosc2 for compressing numbers.
package blosc2

/*
#cgo LDFLAGS: -lpthread /usr/local/lib/libblosc2.a -ldl -lm
#include "blosc2_include.h"
*/
import "C"
import (
	"reflect"
	"unsafe"
)

// to use shared library use LDFLAGS: -lpthread -lblosc2

func init() {
	C.blosc2_init()
}

// Compress takes a slice of numbers and compresses according to level and shuffle.
func Compress(level int, shuffle bool, slice interface{}) []byte {

	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		panic("blosc: expected slice to Compress")
	}
	l := rv.Len()
	size := int(rv.Index(0).Type().Size())
	ptr := unsafe.Pointer(rv.Pointer())

	s := 1
	if !shuffle {
		s = 0
	}
	compressed := make([]byte, l*size+C.BLOSC2_MAX_OVERHEAD)
	// BLOSC_EXPORT int blosc2_compress(int clevel, int doshuffle, int32_t typesize,
	// 	const void* src, int32_t srcsize, void* dest,
	//	int32_t destsize);
	csize := C.blosc2_compress(C.int(level), C.int(s), C.int32_t(size),
		ptr, C.int32_t(l*size),
		unsafe.Pointer(&compressed[0]), C.int32_t(len(compressed)))
	return compressed[:csize]
}

type typed []byte

const maxSize = 1<<36 - 1

func (c typed) Uint16s() []uint16 {
	n := len(c) / int(unsafe.Sizeof(uint16(0)))
	return (*[maxSize]uint16)(unsafe.Pointer(&c[0]))[:n]
}

// Decompress takes a byte of compressed data and returns the uncompressed data.
func Decompress(compressed []byte) typed {

	nbytes := C.int32_t(0)
	cbytes := C.int32_t(0)
	blksz := C.int32_t(0)

	// BLOSC_EXPORT int blosc2_cbuffer_sizes(const void* cbuffer, int32_t* nbytes,
	//	int32_t* cbytes, int32_t* blocksize);
	C.blosc2_cbuffer_sizes(unsafe.Pointer(&compressed[0]), &nbytes, &cbytes, &blksz)

	data := make([]byte, int(nbytes))
	// BLOSC_EXPORT int blosc2_decompress(const void* src, int32_t srcsize,
	//	void* dest, int32_t destsize);
	C.blosc2_decompress(unsafe.Pointer(&compressed[0]), cbytes, unsafe.Pointer(&data[0]), nbytes)
	return data
}
