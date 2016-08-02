package consistent

import (
	"fmt"
	//"log"
	"testing"
)

func TestNew(t *testing.T) {
	c := NewRing()
	c.AddNode("cacheA")
	c.AddNode("cacheB")
	c.AddNode("cacheC")
	users := []string{"user_mcnulty", "user_bunk", "user_omar", "user_bunny", "user_stringer"}
	for _, u := range users {
		server := c.Get(u)
		fmt.Printf("%s => %s\n", u, server)
	}
}

func Test_ExampleAdd(t *testing.T) {
	c := NewRing()
	c.AddNode("cacheA")
	c.AddNode("cacheB")
	c.AddNode("cacheC")
	users := []string{"user_mcnulty", "user_bunk", "user_omar", "user_bunny", "user_stringer"}
	fmt.Println("initial state [A, B, C]")
	for _, u := range users {
		server := c.Get(u)
		fmt.Printf("%s => %s\n", u, server)
	}
	c.AddNode("cacheD")
	c.AddNode("cacheE")
	fmt.Println("\nwith cacheD, cacheE [A, B, C, D, E]")
	for _, u := range users {
		server := c.Get(u)
		fmt.Printf("%s => %s\n", u, server)
	}
	// Output:
	// initial state [A, B, C]
	// user_mcnulty => cacheA
	// user_bunk => cacheA
	// user_omar => cacheA
	// user_bunny => cacheC
	// user_stringer => cacheC
	//
	// with cacheD, cacheE [A, B, C, D, E]
	// user_mcnulty => cacheE
	// user_bunk => cacheA
	// user_omar => cacheA
	// user_bunny => cacheE
	// user_stringer => cacheE
}

func Test_ExampleRemove(t *testing.T) {
	c := NewRing()
	c.AddNode("cacheA")
	c.AddNode("cacheB")
	c.AddNode("cacheC")
	users := []string{"user_mcnulty", "user_bunk", "user_omar", "user_bunny", "user_stringer"}
	fmt.Println("initial state [A, B, C]")
	for _, u := range users {
		server := c.Get(u)
		fmt.Printf("%s => %s\n", u, server)
	}
	c.RemoveNode("cacheC")
	fmt.Println("\ncacheC removed [A, B]")
	for _, u := range users {
		server := c.Get(u)
		fmt.Printf("%s => %s\n", u, server)
	}
	// Output:
	// initial state [A, B, C]
	// user_mcnulty => cacheC
	// user_bunk => cacheA
	// user_omar => cacheA
	// user_bunny => cacheB
	// user_stringer => cacheB

	// cacheC removed [A, B]
	// user_mcnulty => cacheB
	// user_bunk => cacheA
	// user_omar => cacheA
	// user_bunny => cacheB
	// user_stringer => cacheB
}
