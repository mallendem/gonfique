package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ufukty/gonfique/pkg"
)

type Args struct {
	In  string
	Out string
	Pkg string
}

func getArgs() Args {
	args := Args{}
	flag.StringVar(&args.In, "in", "", "input file path (yml or yaml)")
	flag.StringVar(&args.Out, "out", "", "output file path (go)")
	flag.StringVar(&args.Pkg, "pkg", "", "package name that will be inserted into the generated file")
	flag.Parse()
	return args
}

func checkMissingArgs(args Args) error {
	ms := []string{}
	if args.In == "" {
		ms = append(ms, "-in")
	}
	if args.Out == "" {
		ms = append(ms, "-out")
	}
	if args.Pkg == "" {
		ms = append(ms, "-pkg")
	}
	if len(ms) > 0 {
		return fmt.Errorf("some arguments are missing: %s", strings.Join(ms, ", "))
	}
	return nil
}

func main() {
	args := getArgs()
	if err := checkMissingArgs(args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cfgts, err := pkg.ReadConfigYaml(args.In)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := pkg.WriteConfigGo(args.Out, cfgts, args.Pkg); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
