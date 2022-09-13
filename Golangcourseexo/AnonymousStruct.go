package main

import (
	"fmt"
)

func main() {
	s := struct {
		peopleage  int
		school     map[string]string
		activities []string
	}{
		peopleage: 18,
		school: map[string]string{
			"eco1": "gbegamey",
			"eco2": "cadjehoun",
			"eco3": "calavi",
		},
		activities: []string{"foot", "dance", "sing", "lecture"},
	}
	fmt.Println(s)

	for _, v := range s.school {
		fmt.Println(v)

	}

	for _, v := range s.activities {
		fmt.Println(v)

	}

}
