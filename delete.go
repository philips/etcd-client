package main


import (
	"fmt"
	)

func init() {
	cmds["delete"] = cmd{delete, "--name=<nameSpace> <key> ...", "Deletes keys from a nameSpace"}
	cmdHelp["delete"] = ` #### `	
}

func delete(args []string) {
	var key []string

	ns := getArg(args[1], "--ns")

	for i := 2; i < len(args); i++ {
		key = append(key, args[i])
	}

	fmt.Println(key)
	deleteValue(ns, key)
}


func deleteValue(nameSpcae string, key []string) {
	c := dial()

	rev, err := c.Rev()

	if err != nil {
		panic(err)
	}

	for i := 0; i < len(key); i++ {
		c.Del("/" + nameSpcae + "/" + key[i], rev)
    }

}