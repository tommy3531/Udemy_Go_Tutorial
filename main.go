package main


import 
( 
	// client "political/controller"
	test "political/functions"
	"fmt"
)

func main() {
	// client.PropublicaTest()
	// client.GetLegislators()
	// a := test.IntToString("42")
	// fmt.Println(a)
	
	// names := []string{"tom", "marler"}	
	// test.Multiparas(names...)
	// twittername := test.PrintTwitterName("Follow me on Twitter: @tommarler3")
	// fmt.Println(twittername)
	// test.FullName("Mike", "Bob", "franklin")
	_, twitter, _ := test.DiscardValue()
	fmt.Println(twitter)
}