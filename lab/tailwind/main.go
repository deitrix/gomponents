package main

import (
	"os"

	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func main() {
	btn := BtnPrimary(Class("p-4"),
		Href("/"),
		Text("Home"),
	)
	btn.Render(os.Stdout)
}

func Btn(children ...Node) Node {
	return A(Class("no-underline", "text-white text-sm", "py-0.5 px-2"), Group(children))
}

func BtnPrimary(children ...Node) Node {
	return Btn(Class("bg-blue-600 hover:bg-blue-500"), Group(children))
}
