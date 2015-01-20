package main

/*
#ifndef MAINC
#define MAINC
#cgo LDFLAGS:-lstdc++
#cgo CFLAGS:-w -s
#include <stdlib.h>
#include <windows.h>
#include "hook.c"
extern void hook();
extern char * getActTitle();
extern void keyEvent(int,int);
extern void mouseEvent(int);
#endif
*/
import "C"
import (
	"fmt"
	. "time"
)

//export log1
func log1(str *C.char) {
	println(C.GoString(str))
}

var press = false

//export onKey
func onKey(state C.int, keyCode C.int) int {
	switch state {
	case 1: //key down
		return onKeyDown(int(keyCode))
	case 2: //key up
		return onKeyUp(int(keyCode))
	}
	return 0
}

func getActTitle() string {
	title := C.getActTitle()
	return string(C.GoString(title))
}

func keyDown(keyCode ...interface{}) {
	C.keyEvent(0, C.int(keyCode[0].(int)))
}
func keyUp(keyCode ...interface{}) {
	C.keyEvent(2, C.int(keyCode[0].(int)))
}
func lmbDown(delay ...interface{}) {
	C.mouseEvent((C.MOUSEEVENTF_LEFTDOWN))
	Sleep(20 * Millisecond)
}

func lmbUp(delay ...interface{}) {
	C.mouseEvent((C.MOUSEEVENTF_LEFTUP))
	Sleep(20 * Millisecond)
}
func mmbDown(delay ...interface{}) {
	C.mouseEvent((C.MOUSEEVENTF_MIDDLEDOWN))
	Sleep(20 * Millisecond)
}

func mmbUp(delay ...interface{}) {
	C.mouseEvent((C.MOUSEEVENTF_MIDDLEUP))
	Sleep(20 * Millisecond)
}
func rmbDown(delay ...interface{}) {
	C.mouseEvent((C.MOUSEEVENTF_RIGHTDOWN))
	Sleep(20 * Millisecond)
}

func rmbUp(delay ...interface{}) {
	C.mouseEvent((C.MOUSEEVENTF_RIGHTUP))
	Sleep(20 * Millisecond)
}

func main() {
	setup()
	setupCinema4d()
	C.hook()
	fmt.Println("end")
}
