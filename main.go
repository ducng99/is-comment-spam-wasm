package main

import (
	"syscall/js"

	spam "is-comment-spam-wasm/iscommentspam"
)

func main() {
	js.Global().Set("isCommentSpam", js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return false
		}

		return spam.IsSpam(args[0].String())
	}))

	<-make(chan bool)
}
