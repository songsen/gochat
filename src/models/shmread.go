package models  
/* 
#cgo LDFLAGS: -lrt
#include <stdlib.h>  
#include "shmread.h"
*/
import "C"  
import "unsafe"  
//import "fmt"  

func Readmm(filename string) string {  
    f := C.CString(filename)  
    defer C.free(unsafe.Pointer(f))  
    s := C.readshmm(f)   //	shmm.Read("mqtt_shm")
    defer C.free(unsafe.Pointer(s))  
    return C.GoString(s)  
}  
