// ありつつも君をば待たむうち靡く我が黒髪に霜の置くまでに(『万葉集巻1 87』 磐姫皇后)

package irotsukegakari

type Color int

const (
	Black Color = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

var m = map[Color]string{
	Black:   "30",
	Red:     "31",
	Green:   "32",
	Yellow:  "33",
	Blue:    "34",
	Magenta: "35",
	Cyan:    "36",
	White:   "37",
}

func irotsuke(s string, c Color) string {

	if val, ok := m[c]; ok {
		return "\\e[" + val + s + "\\e[0m"
	}

	return s
}
