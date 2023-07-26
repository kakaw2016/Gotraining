package dataYear

type Course struct {
	Number string
	Name   string
	Units  string
}

type Semester struct {
	Term    string
	Courses []Course
}

type Year struct {
	AcaYear string
	Fall    Semester
	Spring  Semester
	Summer  Semester
}
