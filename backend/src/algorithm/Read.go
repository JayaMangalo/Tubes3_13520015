// Golang program to read and write the files
package algorithm

// importing the packages
import (
	"io/ioutil"
)

func ReadFile(text string) (string, bool) {
	fileName := "../test/" + text

	// The ioutil package contains inbuilt
	// methods like ReadFile that reads the
	// filename and returns the contents.

	data, err := ioutil.ReadFile(fileName)
	myString := string(data[:])
	if err == nil {
		return myString, true
	}
	return "", false

}
