package main

import (
	. "fmt"
	. "strings"
)

var keyStates map[int]bool
var apps []App

type App struct {
	title string
	down  []*KFC
	up    []*KFC
}

type KFC struct {
	inKey      int
	funcType   int
	funcParams []int
	funcs      []func(...interface{})
}

func (kfc *KFC) addKFC(f func(...interface{}), params int) {
	kfc.funcs = append(kfc.funcs, f)
	kfc.funcParams = append(kfc.funcParams, params)
}

func (app *App) addDownKFC(inKey int) *KFC {
	keyStates[inKey] = false
	kfc := new(KFC)
	kfc.inKey = inKey
	app.down = append(app.down, kfc)
	return kfc
}

func (app *App) addUpKFC(inKey int) *KFC {
	kfc := new(KFC)
	kfc.inKey = inKey
	app.up = append(app.up, kfc)
	return kfc
}

func onKeyDown(keyCode int) int {
	if keyStates[keyCode] {
		return 1
	}
	Println("[onKeyDown]", keyCode)
	find := 0
	title := getActTitle()

	var kfc *KFC
	for _, app := range apps {
		if Contains(title, app.title) {
			for _, k := range app.down {
				kfc = k
				if kfc.inKey == keyCode {
					for i, f := range kfc.funcs {
						f(kfc.funcParams[i])
					}
					find = 1
					break
				}
			}
			break
		}
	}
	if find == 1 {
		keyStates[keyCode] = true
	}
	return find
}
func onKeyUp(keyCode int) int {
	keyStates[keyCode] = false
	Println("[onKeyUp]", keyCode)
	var kfc *KFC
	find := 0
	title := getActTitle()
	for _, app := range apps {
		Contains(title, app.title)
		if Contains(title, app.title) {
			for _, k := range app.up {
				kfc = k
				if kfc.inKey == keyCode {
					for i, f := range kfc.funcs {
						f(kfc.funcParams[i])
					}
					find = 1
					break
				}
			}
			break
		}
	}
	return find
}
