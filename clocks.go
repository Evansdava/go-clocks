package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	scribble "github.com/nanobox-io/golang-scribble"
)

// Main struct for working with clocks
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
	fmt.Print("Q: Quit the program\n\n")

	for scanner.Scan() {
		switch strings.ToLower(scanner.Text()) {
		case "c":
			UseClock(CreateClock())
		case "s":
			UseClock(SelectClock())
		case "d":
			DeleteClock()
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

// Create a new clock
func CreateClock() Clock {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter name and size (separated by a space): ")
	create, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	createSlice := strings.Split(create, " ")

	name := createSlice[0]

	var size string
	if createSlice[1] != "" {
		size = createSlice[1]
	}
	sizeInt, _ := strconv.Atoi(strings.Trim(size, "\n"))

	clock := Clock{name, sizeInt, 0}
	fmt.Println(clock)
	err = db.Write("clock", name, Clock{name, sizeInt, 0})
	if err != nil {
		panic(err)
	}
	fmt.Println("Clock created")

	return clock
}

// Choose a clock to start using
func SelectClock() Clock {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter clock name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	var clock Clock
	db.Read("clock", name, &clock)
	fmt.Println(clock)
	return clock
}

// Bring up the interface for ticking or unticking a clock
func UseClock(clock Clock) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\nPress T to tick the clock, and R to untick it")
	fmt.Print("Press Q to return to main menu\n\n")
	fmt.Println(clock)

	for scanner.Scan() {
		switch strings.ToLower(scanner.Text()) {
		case "t":
			if clock.filled < clock.size {
				clock.filled += 1
			}
		case "r":
			if clock.filled > 0 {
				clock.filled -= 1
			}
		case "q":
			return
		}

		fmt.Println(clock)
	}
}

// Delete a clock from the database
func DeleteClock() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter clock name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	db.Delete("clock", name)
	fmt.Printf("Successfully deleted clock %s", name)
}

// Stringer to print clocks out
func (c Clock) String() string {
	str := "["
	for i := 0; i < c.size; i++ {
		if i < c.filled {
			str += string(9609) // Left 7/8 block
		} else {
			str += "-"
		}
		// fmt.Println(str)
	}
	str += "]"
	return str
}
