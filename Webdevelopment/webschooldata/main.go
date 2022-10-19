package main

import (
	"log"
	"os"

	"github.com/kakaw2016/Gotraining/Webdevelopment/webschooldata/dataYear"
	"github.com/kakaw2016/Gotraining/Webdevelopment/webschooldata/mainInit"
)

func main() {

	years := populateData()

	mainInit.Init()

	err := tpl.Execute(os.Stdout, years)
	if err != nil {
		log.Fatalln(err)
	}
}

func populateData() []dataYear.Year {
	return []dataYear.Year{
		dataYear.Year{
			AcaYear: "2020-2021",
			Fall: dataYear.Semester{
				Term: "Fall",
				dataYear.Courses: []dataYear.Course{
					dataYear.Course{"CSCI-40", "Introduction to Programming in Go", "4"},
					dataYear.Course{"CSCI-130", "Introduction to Web Programming with Go", "4"},
					dataYear.Course{"CSCI-140", "Mobile Apps Using Go", "4"},
				},
			},
			Spring: dataYear.Semester{
				Term: "Spring",
				dataYear.Courses: []dataYear.Course{
					dataYear.Course{"CSCI-50", "Advanced Go", "5"},
					dataYear.Course{"CSCI-190", "Advanced Web Programming with Go", "5"},
					dataYear.Course{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
				},
			},
		},
		dataYear.Year{
			AcaYear: "2021-2022",
			Fall: dataYear.Semester{
				Term: "Fall",
				dataYear.Courses: []dataYear.Course{
					dataYear.Course{"CSCI-40", "Introduction to Programming in Go", "4"},
					dataYear.Course{"CSCI-130", "Introduction to Web Programming with Go", "4"},
					dataYear.Course{"CSCI-140", "Mobile Apps Using Go", "4"},
				},
			},
			Spring: dataYear.Semester{
				Term: "Spring",
				dataYear.Courses: []dataYear.Course{
					dataYear.Course{"CSCI-50", "Advanced Go", "5"},
					dataYear.Course{"CSCI-190", "Advanced Web Programming with Go", "5"},
					dataYear.Course{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
				},
			},
		},
	}
}
