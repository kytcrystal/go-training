// twofer contains ShareWith function
package twofer

import "fmt"

// ShareWith takes in a string (name) and returns a sentence using the name
// Returns sentence using "you", if name is blank
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %v, one for me.", name)
}
