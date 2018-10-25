package functions

import "strconv"
import "fmt"

// Named func has the following
// 1. a name
// 2. declared at package level
// 3. declared outside the body of another func

// Takes a string and returns int
func IntToString(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		print("Error")
	}
	return value
}

func PrintTwitterName(twitterName string) string {
	// PrintTwitterName: is the function name
	// twitterName: is the input params
	// string: is the result type

	return twitterName
}

func FullName(firstName, middleName, lastName string) {
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}

func DiscardValue() (value int, twitter string, not bool){
	return 1, "Follow me on twitter @tommarler3", true
}