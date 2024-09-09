package main

import {
	"fmt"
}

const {
	todofile = ".todos.json"
}

func main() {
	add := flag.Bool("add", false, "Add a new task")
	flag.Parse()
	todos := &todo.Todos{}
	if err := tods.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit()
	}
}