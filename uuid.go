/*
Copyright (c) 2013 Ziyi Fan

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

+build linux
*/

package uuid

/*
#include <string.h>
#include <uuid/uuid.h>
*/
import "C"
import "unsafe"

type UUID struct {
	data [16]C.uchar
}
const uuid_len = 16
const uuid_str_len = 36

func (id *UUID) String() string {
	cstr := make([]C.char, uuid_str_len)
	C.uuid_unparse(&id.data[0], &cstr[0])
	return C.GoString(&cstr[0])
}

func (id *UUID) Bytes() []byte {
	ptr := unsafe.Pointer(&id.data[0])
	return C.GoBytes(ptr, uuid_len)
}

func New() *UUID {
	id := new(UUID)
	C.uuid_generate(&id.data[0])
	return id
}

func FromBytes(bytes []byte) *UUID {
	id := new(UUID)
	C.memcpy(unsafe.Pointer(&id.data[0]), unsafe.Pointer(&bytes[0]), uuid_len)
	return id
}
