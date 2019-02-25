package main

type Mock struct {
	logged []string
}

func (m *Mock) Println(v ...interface{}) {
	m.logged = append(m.logged, string(v[0].(string)))
}

func ExamplePrint() {
	entry := TimeEntry{Description: "Foobar", ID: 5, Duration: 1000000000}
	entry.Print()
	// Output:
	// ID: 		 5
	// Description: 	 Foobar
	// Duration: 	 1 second
}
