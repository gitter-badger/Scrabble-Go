package main

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

func main() {
	// fmt.Println("Hello, playground")
	js.Global.Call("alert", "Hello, JavaScript")
	println("Hello, JS")
	verify()
}

func verify() {
	document := dom.GetWindow().Document()
	login := document.GetElementByID("login").(*dom.HTMLDivElement)
	value := login.GetAttribute("value")
	// document := js.Global.Get("document")
	// login := document.Call("getElementById", "login")
	// document := dom.GetWindow().Document()
	// login := document.GetElementByID("login")
	// value := login.TextContent()
	// println("After ID")
	// fmt.Println(value)

	// document := dom.GetWindow().Document()
	// login := document.GetElementByID("login").(dom.HTMLDivElement).InnerHTML()
	// fmt.Println(login.GetAttribute("value"))
	// el := dom.GetWindow().Document().QuerySelector("#login")
	// htmlEl := el.(dom.HTMLElement)
	// pEl := el.(*dom.HTMLDataElement)
	// println(htmlEl, pEl)
	// println(login)
}
