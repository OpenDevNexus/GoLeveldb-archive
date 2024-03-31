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
// #include "leveldb/c.h"
import "C"

// DestroyComparator deallocates a *C.leveldb_comparator_t.
//
// This is provided as a convienience to advanced users that have implemented
// their own comparators in C in their own code.
func DestroyComparator(cmp *C.leveldb_comparator_t) {
	C.leveldb_comparator_destroy(cmp)
}
