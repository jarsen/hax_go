package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/jarsen/hax/repl"
)

const INTRO = `
 __  __     ______     __  __    
/\ \_\ \   /\  __ \   /\_\_\_\   
\ \  __ \  \ \  __ \  \/_/\_\/_  
 \ \_\ \_\  \ \_\ \_\   /\_\/\_\ 
  \/_/\/_/   \/_/\/_/   \/_/\/_/ 

`

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println(INTRO)
	fmt.Printf("Hello, %s. Welcome to HAX.\n\n", user.Name)
	repl.Start(os.Stdin, os.Stdout)
}
