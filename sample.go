// ありつつも君をば待たむうち靡く我が黒髪に霜の置くまでに(『万葉集巻1 87』 磐姫皇后)

package irotsukegakari

func irotsuke(s string) string {
	return "\\e[30" + s + "\\e[0m"
}
