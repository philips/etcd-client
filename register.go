package main

import (
	"io/ioutil"
	"io"
	"fmt"
	"crypto/md5"
)

func init() {
	cmds["register"] = cmd{register, " --ns=<nameSpace> --script=<scriptPath> --log==<logPath>", "register"}
	cmdHelp["register"] = `#####`	
}

func register(args []string) {
	c := dial()

	rev, err := c.Rev()

	if err != nil {
		panic(err)
	}

	ns := getArg(args[1], "--ns")

	if ns == "" {
		return
	}

	script := getArg(args[2], "--script")

	if script == "" {
		return
	}
 
	log := getArg(args[3], "--log")

	if log =="" {
		return
	}

	mid, err := ioutil.ReadFile("/etc/machine-id")

		// is new line char at the end?
	if mid[len(mid) - 1] == 10 {
		mid = mid[:len(mid) - 1]
	} 

	if err != nil {
		panic(err)
	}

	redoScriptStr := "false"
	if *redoScript {
		redoScriptStr = "true"
	}

	h := md5.New()
	io.WriteString(h, script)

	newRev, err := c.Set("/" + string(mid) + "/" + ns + "/" + fmt.Sprintf("%x", h.Sum(nil)), rev, 
		[]byte(script + "," + log + "," + redoScriptStr))

	if err != nil {
		bail(err)
	}

	fmt.Println(newRev)

}
