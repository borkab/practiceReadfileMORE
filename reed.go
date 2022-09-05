package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	file, err := os.Create("bee.txt")
	check(err)

	fmt.Println("bee.txt was successfully created")

	_, err = file.WriteString("Here is the beehive but where are all the bees?\nHiding away where nobody sees.\nHere they come flying out of their hive:\none, two, three, four, five\n:)")
	check(err)
	fmt.Println("You have successfully written to your bee.txt")

	//we open our file
	file, err = os.Open("bee.txt")
	check(err)
	defer file.Close()

	//read file by words
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {

		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}
