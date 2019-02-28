package main

import "time"

type Mock struct {
	logged []string
}

func (m *Mock) Println(v ...interface{}) {
	m.logged = append(m.logged, string(v[0].(string)))
}

func ExamplePrint() {
	entry := TimeEntry{Description: "Foobar", ID: 5, Duration: time.Minute}
	entry.Print()
	// Output:
	// ID: 		 5
	// Description: 	 Foobar
	// Duration: 	 1 minutes
}
