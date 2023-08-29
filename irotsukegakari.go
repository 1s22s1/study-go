// ありつつも君をば待たむうち靡く我が黒髪に霜の置くまでに(『万葉集巻1 87』 磐姫皇后)

package irotsukegakari

type Color int

const (
	Black Color = iota
	Red
)

func irotsuke(s string, c Color) string {
	switch c {
	case Black:
		return "\\e[30" + s + "\\e[0m"
	case Red:
		return "\\e[31" + s + "\\e[0m"
	default:
		return s
	}
}
