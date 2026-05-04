package mobile

import (
	"fmt"
	"os"
	"runtime/debug"
	"sync"

	"github.com/hajimehoshi/ebiten/v2/mobile"
	"github.com/tinne26/imeglitch/game"
)

// make clean build
// make install

var once sync.Once

// called by apk-ebiten-builder
func SetAndroidID(id int) {
	once.Do(startGame)
}

func startGame() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("[PANIC] %s: %s\n", r, debug.Stack())
				os.Exit(1)
			}
		}()

		var game game.Game
		mobile.SetGame(&game)
	}()
}

type IMEBridge interface {
	Show(int32, int32)
	Hide()
}

// called by apk-ebiten-builder
func RegisterIMEBridge(ime IMEBridge) {
	game.IME = ime
}
