package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/solbero/tempconv/cli"
)

var (
	version = "dev"
)

func main() {
	buff := new(bytes.Buffer)
	flags := flag.NewFlagSet("tempconv", flag.ContinueOnError)

	conf, err := cli.ParseArgs(buff, os.Args[1:], flags)
	if err != nil {
		fmt.Fprintln(os.Stderr, buff.String())
		os.Exit(2)
	}

	err = cli.Run(buff, conf, flags, version)
	if err != nil {
		fmt.Fprintln(os.Stderr, buff.String())
		os.Exit(2)
	}

	buff.WriteByte('\n')
	fmt.Fprint(os.Stdout, buff.String())
}
