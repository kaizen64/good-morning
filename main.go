package main

import (
	"log"
	"os"
	"os/exec"
)

func execute_this(command []string) {

	cmd := exec.Command(command[0], command[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
	
}

func main() {
	
	// Warning: Make sure you alredy are inside a tmux session
	
	main_file := "workingmemory"
	sub_file := "todos"
	jarvis := "spd-say 'Welcome back, master'"

	tmux_split_vertical := []string{"tmux", "split-window", "-h"}
	tmux_split_horizontal := []string{"tmux", "split-window", "-v", "-t", "1"}
	big_window := []string{"tmux", "send-keys", "-t", "0", main_file, "C-m"}
	little_up_window := []string{"tmux", "send-keys", "-t", "1", jarvis, "C-m"}
	little_bottom_window := []string{"tmux", "send-keys", "-t", "2", sub_file, "C-m"}

	execute_this(tmux_split_vertical)
	execute_this(tmux_split_horizontal)
	execute_this(big_window)
	execute_this(little_up_window)
	execute_this(little_bottom_window)
	
}
