package main

import (
	"io/ioutil"
	"io"
	"crypto/md5"
	"fmt"
)

func init() {
	cmds["deregister"] = cmd{deregister, " --ns=<nameSpace> --script==<script>", "deregister"}
	cmdHelp["deregister"] = `#####`	
}

func deregister(args []string) {
	c := dial()

	rev, err := c.Rev()

	if err != nil {
		panic(err)
	}

	ns := getArg(args[1], "--ns")

	script := getArg(args[2], "--script")

	if ns == "" {
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

	h := md5.New()
	io.WriteString(h, script)

	c.Del("/" + string(mid) + "/" + ns + "/" + fmt.Sprintf("%x", h.Sum(nil)), rev)
}