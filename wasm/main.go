package main

import (
	inserter "pokeball-inserter/lib"
	"syscall/js"
)

func main() {
	ch := make(chan struct{}, 0)
	js.Global().Set("drawImages", JSDrawImages())
	<-ch
}

func JSDrawImages() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		img1 := args[0].String()
		img2 := args[1].String()

		return inserter.DrawImages(img1, img2)
	})
}
