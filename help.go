package main

import "fmt"

func init() {
	cmds["help"] = cmd{help, "<command>", "provide detailed help on a command"}
}

func help(command []string) {
	c := cmds[command[0]]
	fmt.Printf("Use: %s [options] %s [options] %s\n\n", selfName, command[0], c.a)
	fmt.Print(cmdHelp[command[0]])
}
