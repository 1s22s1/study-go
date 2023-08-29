// ありつつも君をば待たむうち靡く我が黒髪に霜の置くまでに(『万葉集巻1 87』 磐姫皇后)

package irotsukegakari

type Color int

const (
	Black Color = iota
	Red
)

func irotsuke(s string, c Color) string {
	m := map[Color]string{
		Black: "30",
		Red:   "31",
	}

	if val, ok := m[c]; ok {
		return "\\e[" + val + s + "\\e[0m"
	}

	return s
}
