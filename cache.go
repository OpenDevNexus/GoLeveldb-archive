// A fork of https://github.com/jmhodges/levigo
/*
Copyright (c) 2012 Jeffrey M Hodges

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package GoLevelDb

/*
#cgo CFLAGS: -I./inc/leveldb-compiled/headers
#cgo windows LDFLAGS: -L./inc/leveldb-compiled/lib-windows -lleveldb -lstdc++
#cgo linux LDFLAGS: -L./inc/leveldb-compiled/lib-linux -lleveldb -lstdc++
#include <stdint.h>
#include "leveldb/c.h"
*/
import "C"

// Cache is a cache used to store data read from data in memory.
//
// Typically, NewLRUCache is all you will need, but advanced users may
// implement their own *C.leveldb_cache_t and create a Cache.
//
// To prevent memory leaks, a Cache must have Close called on it when it is
// no longer needed by the program. Note: if the process is shutting down,
// this may not be necessary and could be avoided to shorten shutdown time.
type Cache struct {
	Cache *C.leveldb_cache_t
}

// NewLRUCache creates a new Cache object with the capacity given.
//
// To prevent memory leaks, Close should be called on the Cache when the
// program no longer needs it. Note: if the process is shutting down, this may
// not be necessary and could be avoided to shorten shutdown time.
func NewLRUCache(capacity int) *Cache {
	return &Cache{C.leveldb_cache_create_lru(C.size_t(capacity))}
}

// Close deallocates the underlying memory of the Cache object.
func (c *Cache) Close() {
	C.leveldb_cache_destroy(c.Cache)
}
