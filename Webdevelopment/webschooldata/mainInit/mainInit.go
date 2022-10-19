package mainInit

import "text/template"

var tpl *template.Template

func Init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
