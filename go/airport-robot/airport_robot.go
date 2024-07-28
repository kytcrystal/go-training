package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	LanguageName(language string) string
	Greet(name string) string
}

func SayHello(name string, g Greeter) string {
	return g.Greet(name)
}