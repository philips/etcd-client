package main


import (
	"strings"
	"fmt"
	)

func init() {
	cmds["set"] = cmd{set, "--ns=<nameSpace> <key>=<value> ...", "Adds values to a nameSpace"}
	cmdHelp["set"] = ` #### `	
}

func set(args []string) {
	var key []string
	var value []string

	ns := getArg(args[1], "--ns")

	for i := 2; i < len(args); i++ {
		pairs := strings.Split(args[i], "=")
		if len(pairs) != 2 {
			fmt.Println("Wrong argument with " + args[i])
			return
		}
		key = append(key, pairs[0])
		value = append(value, pairs[1])

	}

	fmt.Println(key, " ", value)
	addValue(ns, key, value)
}


func addValue(nameSpcae string, key []string, value []string) {
	c := dial()

	rev, err := c.Rev()

	if err != nil {
		panic(err)
	}

	for i := 0; i < len(key); i++ {
		fmt.Println(key[i], " ", value[i])
		c.Set("/" + nameSpcae + "/" + key[i], rev, []byte(value[i]))
    }

}

