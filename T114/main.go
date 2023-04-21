package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer, Winter semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-10", "Introduction to Programming in Go", "4"},
				course{"CSCI-20", "Introduction to Web Programming with Go", "4"},
				course{"CSCI-30", "Mobile Apps Using Go", "4"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"CSCI-40", "Advanced Go", "5"},
				course{"CSCI-50", "Advanced Web Programming with Go", "5"},
				course{"CSCI-60", "Advanced Mobile Apps With Go", "5"},
			},
		},
		Winter: semester{
			Term: "Winter",
			Courses: []course{
				course{"CSCI-70", "Advanced Go", "5"},
				course{"CSCI-80", "Advanced Web Programming with Go", "5"},
				course{"CSCI-90", "Advanced Mobile Apps With Go", "5"},
			},
		},
		Summer: semester{
			Term: "Summer",
			Courses: []course{
				course{"CSCI-11", "Advanced Go", "5"},
				course{"CSCI-22", "Advanced Web Programming with Go", "5"},
				course{"CSCI-33", "Advanced Mobile Apps With Go", "5"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}
