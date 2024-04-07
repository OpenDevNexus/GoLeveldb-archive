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
#include "leveldb/c.h"
*/
import "C"

// GetLevelDBMajorVersion returns the underlying LevelDB implementation's major
// version.
func GetLevelDBMajorVersion() int {
	return int(C.leveldb_major_version())
}

// GetLevelDBMinorVersion returns the underlying LevelDB implementation's minor
// version.
func GetLevelDBMinorVersion() int {
	return int(C.leveldb_minor_version())
}
