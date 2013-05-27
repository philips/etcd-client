package main

import (
	"flag"
	"fmt"
	"github.com/ha/doozer"
	"os"
	"sort"
	"strings"
	"strconv"
)

const pathPrefix = "/etcd/"

var (
	uri         = flag.String("a", "doozer:?ca=127.0.0.1:8046", "the address to bind to")
	buri        = flag.String("b", "", "the DzNS uri")
	rrev        = flag.Int64("r", -1, "request rev")
	showHelp    = flag.Bool("h", false, "show help")
	showVersion = flag.Bool("v", false, "print version string")
	reWatch		= flag.Bool("reWatch", false, "whether to watch the existing service in the watchedService list")
	redoScript	= flag.Bool("redo", false, "whether to redo the script for existing entries in the watchedService")
)

type cmd struct {
	f func(args []string)
	a string // args
	d string // short description
}

var (
	selfName = os.Args[0]
	cmds     = map[string]cmd{}
	cmdHelp  = map[string]string{}
)

const (
	usage1 = `
Each command takes zero or more options and zero or more arguments.
In addition, there are some global options that can be used with any command.
The exit status is 0 on success, 1 for a rev mismatch, and 2 otherwise.

Global Options:
`
	usage2 = `Environment:

  DOOZER_URI - The doozer cluster to bind to; overriden by -a.

  DOOZER_BOOT_URI - The DzNS to lookup address in; overriden by -b.

Commands:
`
)

func usage() {
	fmt.Fprintf(os.Stderr, "Use: %s [options] <command> [options] [args]\n", selfName)
	fmt.Fprint(os.Stderr, usage1)
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr)

	fmt.Fprint(os.Stderr, usage2)
	var max int
	var names []string
	us := make(map[string]string)
	for k := range cmds {
		u := k + " " + cmds[k].a
		if len(u) > max {
			max = len(u)
		}
		names = append(names, k)
		us[k] = u
	}
	sort.Strings(names)
	for _, k := range names {
		fmt.Fprintf(os.Stderr, "  %-*s - %s\n", max, us[k], cmds[k].d)
	}

}

func bail(e error) {
	fmt.Fprintln(os.Stderr, "Error:", e)
	if e == doozer.ErrOldRev {
		os.Exit(1)
	}
	os.Exit(2)
}

func mustAtoi64(arg string) int64 {
	n, err := strconv.ParseInt(arg, 10, 64)
	if err != nil {
		bail(err)
	}
	return n
}

func dial() *doozer.Conn {
	c, err := doozer.DialUri(*uri, *buri)
	if err != nil {
		bail(err)
	}
	return c
}


func getArg(str string, arg string) string {
	namePairs := strings.Split(str, "=")
	if len(namePairs) != 2 {
		fmt.Println("Wrong argument with " + str)
		return ""
	}
	if namePairs[0] != arg {
		fmt.Println("undefined nameSpace")
		return ""
	}

	nameSpace := namePairs[1]
	return nameSpace
}

func main() {
	flag.Parse()
	cmd := flag.Arg(0)

	c, ok := cmds[cmd]
	if !ok {
		fmt.Fprintln(os.Stderr, "Unknown command:", cmd)
		usage()
		os.Exit(127)
	}

	args := flag.Args()

	c.f(args)

}
