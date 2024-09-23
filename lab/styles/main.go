package main

import (
	"os"

	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func main() {
	btn := BtnPrimary(Style("padding: 1rem"),
		Href("/"),
		Text("Home"),
	)
	btn.Render(os.Stdout)
}

func Btn(children ...Node) Node {
	return A(
		Style("text-decoration-line: none", "color: white", "font-size: 0.875rem", "padding: 0.125rem 0.5rem"),
		Group(children),
	)
}

func BtnPrimary(children ...Node) Node {
	return Btn(
		Style("background-color: blue"),
		Group(children),
	)
}
