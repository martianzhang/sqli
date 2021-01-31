package sqli

/*
#cgo CFLAGS: -I./libinjection
#cgo LDFLAGS: -L./libinjection -linjection
#include "libinjection.h"
#include "libinjection_sqli.h"
*/
import "C"
import (
	"bytes"
	"fmt"
	"unsafe"
)

// SQLi check args or sql is sqli
func SQLi(sql string) (fingerprint string, is bool) {
	var out [8]C.char
	pointer := (*C.char)(unsafe.Pointer(&out[0]))
	if found := C.libinjection_sqli(C.CString(sql), C.size_t(len(sql)), pointer); found == 1 {
		is = true
		output := C.GoBytes(unsafe.Pointer(&out[0]), 8)
		fingerprint = fmt.Sprint(string(output[:bytes.Index(output, []byte{0})]))
	}
	return fingerprint, is
}

