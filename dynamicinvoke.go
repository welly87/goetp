package main

import "fmt"

type Hello struct {
	Who string
}

func tell(msg interface{}) {
	switch msg := msg.(type) {
	case Hello:
		fmt.Printf("Hello %v\n", msg.Who)
	}
}

func main() {
	tell(Hello{Who: "Roger"})
}
