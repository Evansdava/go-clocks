package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	scribble "github.com/nanobox-io/golang-scribble"
)

type Clock struct {
	name   string
	size   int
	filled int
}

var db, _ = scribble.New("./db", nil)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Menu:")
	fmt.Println("C: Create a clock")
	fmt.Println("S: Select a clock")
	fmt.Println("D: Delete a clock")
	fmt.Println("Q: Quit the program")

	for scanner.Scan() {
		switch strings.ToLower(scanner.Text()) {
		case "c":
			createClock()
		case "s":
			selectClock()
		case "d":
			deleteClock()
		case "q":
			return
		}

		fmt.Println("\nMenu:")
		fmt.Println("C: Create a clock")
		fmt.Println("S: Select a clock")
		fmt.Println("D: Delete a clock")
		fmt.Print("Q: Quit the program\n\n")
	}
}

func createClock() Clock {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter name and size (separated by a space): ")
	create, err := reader.ReadString('\n')
	fmt.Println(create)
	if err != nil {
		panic(err)
	}

	createSlice := strings.Split(create, " ")

	name := createSlice[0]

	var size string
	if createSlice[1] != "" {
		size = createSlice[1]
		fmt.Println(size)
	}
	sizeInt, _ := strconv.Atoi(size)
	fmt.Println(sizeInt)

	clock := Clock{name, sizeInt, 0}
	fmt.Println(clock)
	db.Write("clock", name, clock)
	fmt.Println("Clock created")

	return clock
}

func selectClock() Clock {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter clock name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	var clock Clock
	db.Read("clock", name, &clock)
	return clock
}

func useClock(clock Clock) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Press T to tick the clock, and R to untick it")
	fmt.Println("Press Q to return to main menu")
	fmt.Println(clock)

	for scanner.Scan() {
		switch strings.ToLower(scanner.Text()) {
		case "t":
			clock.filled += 1
		case "r":
			clock.filled -= 1
		case "q":
			break
		}

		fmt.Println(clock)
	}
}

func deleteClock() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter clock name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	db.Delete("clock", name)
	fmt.Printf("Successfully deleted clock %s", name)
}

func (c Clock) String() string {
	str := "]"
	for i := 0; i < c.size; i++ {
		if i < c.filled {
			str += string(9609) // Left 7/8 block
		} else {
			str += "-"
		}
		fmt.Println(str)
	}
	str += "]"
	return str
}
