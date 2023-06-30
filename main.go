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
	output := new(bytes.Buffer)
	flags := flag.NewFlagSet("tempconv", flag.ContinueOnError)

	conf, err := cli.ParseArgs(output, os.Args[1:], flags)
	if err != nil {
		fmt.Fprintln(os.Stderr, output.String())
		os.Exit(2)
	}

	err = cli.Run(output, conf, flags, version)
	if err != nil {
		fmt.Fprintln(os.Stderr, output.String())
		os.Exit(2)
	}

	output.WriteByte('\n')
	fmt.Fprint(os.Stdout, output.String())
}
