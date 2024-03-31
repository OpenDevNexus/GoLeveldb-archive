// A fork of https://github.com/jmhodges/levigo
/*
Copyright (c) 2012 Jeffrey M Hodges

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package GoLevelDb

// #cgo CFLAGS: -I inc/leveldb/include/leveldb
// #cgo LDFLAGS: -L inc/leveldb/build -lleveldb
// #include <stdlib.h>
// #include "leveldb/c.h"
import "C"

// FilterPolicy is a factory type that allows the LevelDB database to create a
// filter, such as a bloom filter, that is stored in the sstables and used by
// DB.Get to reduce reads.
//
// An instance of this struct may be supplied to Options when opening a
// DB. Typical usage is to call NewBloomFilter to get an instance.
//
// To prevent memory leaks, a FilterPolicy must have Close called on it when
// it is no longer needed by the program.
type FilterPolicy struct {
	Policy *C.leveldb_filterpolicy_t
}

// NewBloomFilter creates a filter policy that will create a bloom filter when
// necessary with the given number of bits per key.
//
// See the FilterPolicy documentation for more.
func NewBloomFilter(bitsPerKey int) *FilterPolicy {
	policy := C.leveldb_filterpolicy_create_bloom(C.int(bitsPerKey))
	return &FilterPolicy{policy}
}

// Close reaps the resources associated with this FilterPolicy.
func (fp *FilterPolicy) Close() {
	C.leveldb_filterpolicy_destroy(fp.Policy)
}
