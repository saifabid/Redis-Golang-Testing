package main

import (
	"bufio"
	"fmt"
	"github.com/fzzy/radix/redis"
	"os"
	"strings"
)

const ( // conncections command
	network = "tcp"
	addr    = "127.0.0.1:6379"
)
const ( // command constants

	get    = "get"
	set    = "set"
	del    = "del"
	incr   = "incr"
	expire = "expire"
	setex  = "setex"
	ttl    = "ttl"
)

func main() {
	go welcome() // by the time this executes, connections and connection related errors can be delt with
	done := make(chan bool)
	c, err := redis.Dial(network, addr) // client connects to server
	if err != nil {
		fmt.Println(err)
		return
	} // makes sure redis server was running and no errors occured

	defer func() {

		c.Close()
	}() // clean up needed for exiting

	go getInput(done, c) // gets input and exceutes command

	<-done // when hang around this point until 'exit' is inputted

	fmt.Println("Thanks for using Redis console client") // closing good

}
func welcome() {
	fmt.Println()
	fmt.Println("Welcome to the console application for redis")
	fmt.Println("---------------------------------------------")
	fmt.Println()
	fmt.Println("Currently we support the following commands: ")
	fmt.Println("1. set")
	fmt.Println("2. get")
	fmt.Println("3. del")
	fmt.Println("4. incr")
	fmt.Println("5. setex")
	fmt.Println("6. expire")
	fmt.Println("7. ttl")
	fmt.Println("8. exit")

	fmt.Println()
	fmt.Println("Please enter your commands choice in lower case")
	fmt.Println("---------------------------------------------")
	fmt.Println()
}
func getInput(done chan bool, c *redis.Client) {

	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Command ( ex. set john 17 ): ")
		text, _ := reader.ReadString('\n')         //read intil new line
		clean(&text)                               //removes the delimter
		textSlice := strings.Split(text, " ")      // tokenizes string
		cmd, arg, err, exit := parseCmd(textSlice) // some manual parsing for error for fun
		if exit {                                  // exit application if user inputs "exit"

			done <- true
			break
		} else if err {
			fmt.Println("Invalid command, try again")
		} else {
			excecuteCmd(cmd, arg, c) // exceute the user asked command
		}

	}

}

func clean(text *string) { // simple function to remove newline after reading from console
	*text = strings.TrimSuffix(*text, "\n")
	*text = strings.TrimSpace(*text)
}

func parseCmd(test []string) (cmd string, arg []string, err bool, exit bool) {

	length := len(test)

	if length == 1 {
		if test[0] == "exit" {
			exit = true

		} else {
			err = true
			exit = false

		}
	} else if length < 2 {
		err = true
		exit = false

	} else if length == 2 {
		if test[0] == get || test[0] == del || test[0] == incr || test[0] == ttl { // has to be get or del if length is 2
			err = false
			cmd = test[0]
			arg = test[1:]
			exit = false

		} else {
			err = true
			exit = false

		}

	} else if length >= 3 {
		if test[0] == set || test[0] == expire || test[0] == setex {
			err = false
			exit = false
			cmd = test[0]
			arg = test[1:]
		} else {
			err = true
			exit = false
		}
	} else {
		err = true
		exit = false

	}
	return
}

func excecuteCmd(cmd string, arg []string, c *redis.Client) {

	switch cmd {
	case get:
		rep := c.Cmd(cmd, arg)
		fmt.Println("Got: ", rep)
	case set:
		rep := c.Cmd(cmd, arg)
		fmt.Println("Set: ", rep)
	case del:
		rep := c.Cmd(cmd, arg)
		fmt.Println("deleted: ", rep)
	case incr:
		rep := c.Cmd(cmd, arg)
		fmt.Println("incremented: ", rep)
	case expire:
		rep := c.Cmd(cmd, arg)
		fmt.Println("expire set: ", rep)
	case setex:
		rep := c.Cmd(cmd, arg)
		fmt.Println("Set  & expire : ", rep)
	case ttl:
		rep := c.Cmd(cmd, arg)
		fmt.Println("Time left: ", rep)
	default:
		fmt.Println("nil")

	}

}
