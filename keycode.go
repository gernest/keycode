package keycode

var codes = map[string]int{
	"backspace":     8,
	"tab":           9,
	"enter":         13,
	"shift":         16,
	"ctrl":          17,
	"alt":           18,
	"pause/break":   19,
	"caps lock":     20,
	"esc":           27,
	"space":         32,
	"page up":       33,
	"page down":     34,
	"end":           35,
	"home":          36,
	"left":          37,
	"up":            38,
	"right":         39,
	"down":          40,
	"insert":        45,
	"delete":        46,
	"command":       91,
	"left command":  91,
	"right command": 93,
	"numpad *":      106,
	"numpad +":      107,
	"numpad -":      109,
	"numpad .":      110,
	"numpad /":      111,
	"num lock":      144,
	"scroll lock":   145,
	"my computer":   182,
	"my calculator": 183,
	";":             186,
	"=":             187,
	",":             188,
	"-":             189,
	".":             190,
	"/":             191,
	"`":             192,
	"[":             219,
	"\\":            220,
	"]":             221,
	"'":             222,

	// aliases
	"windows": 91,
	"⇧":       16,
	"⌥":       18,
	"⌃":       17,
	"⌘":       91,
	"ctl":     17,
	"control": 17,
	"option":  18,
	"pause":   19,
	"break":   19,
	"caps":    20,
	"return":  13,
	"escape":  27,
	"spc":     32,
	"pgup":    33,
	"pgdn":    34,
	"ins":     45,
	"del":     46,
	"cmd":     91,

	//characters
	"a": 65,
	"b": 66,
	"c": 67,
	"d": 68,
	"e": 69,
	"f": 70,
	"g": 71,
	"h": 72,
	"i": 73,
	"j": 74,
	"k": 75,
	"l": 76,
	"m": 77,
	"n": 78,
	"o": 79,
	"p": 80,
	"q": 81,
	"r": 82,
	"s": 83,
	"t": 84,
	"u": 85,
	"v": 86,
	"w": 87,
	"x": 88,
	"y": 89,
	"z": 90,

	//numbers
	"0": 48,
	"1": 49,
	"2": 50,
	"3": 51,
	"4": 52,
	"5": 53,
	"6": 54,
	"7": 55,
	"8": 56,
	"9": 57,

	//function keys
	"f1":  112,
	"f2":  113,
	"f3":  114,
	"f4":  115,
	"f5":  116,
	"f6":  117,
	"f7":  118,
	"f8":  119,
	"f9":  120,
	"f10": 121,
	"f11": 122,
	"f12": 123,

	//numpad
	"numpad 0": 96,
	"numpad 1": 97,
	"numpad 2": 98,
	"numpad 3": 99,
	"numpad 4": 100,
	"numpad 5": 101,
	"numpad 6": 102,
	"numpad 7": 103,
	"numpad 8": 104,
	"numpad 9": 105,
}

var reverse map[int]string

func init() {
	for k, v := range codes {
		reverse[v] = k
	}
}

// Number returns keycode number from string.
func Number(key string) int {
	return codes[key]
}

// String returns keycode string from a nummber.
func String(key int) string {
	return reverse[key]
}
