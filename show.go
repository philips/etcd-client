package main


import (
	"fmt"
	)

func init() {
	cmds["show"] = cmd{show, "--ns=<nameSpace> <key>...", "Shows values of key in the nameSpace"}
	cmdHelp["show"] = ` #### `	
}

func show(args []string) {
	var key []string

	ns := getArg(args[1], "--ns")

	for i := 2; i < len(args); i++ {
		key = append(key, args[i])
	}

	showValue(ns, key)
}


func showValue(nameSpcae string, key []string) {
	c := dial()

	for i := 0; i < len(key); i++ {
		
		b, _, err := c.Get(pathPrefix + nameSpcae + "/" + key[i], nil)
		if err != nil {
			bail(err)
		}

		if len(b) == 0 {
			fmt.Println(key[i], " Not Found")
		} else {
			fmt.Println(key[i], string(b))
		}
    }

}