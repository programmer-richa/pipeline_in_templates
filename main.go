package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"dbl":    double,
	"sq":     sqRoot,
	"square": square,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.gohtml"))
}

func double(x float64) float64 {
	return x * 2
}

func square(x float64) float64 {
	return math.Pow(x, 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}
func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", float64(3))
	if err != nil {
		log.Fatal(err)
	}
}
