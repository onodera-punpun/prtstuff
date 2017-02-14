package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/go2c/optparse"
)

// pull pulls in ports.
func pull(args []string) {
	// Define valid arguments.
	o := optparse.New()
	argh := o.Bool("help", 'h', false)

	// Parse arguments.
	vals, err := o.Parse(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	// Print help.
	if *argh {
		fmt.Println("Usage: prt pull [arguments] [repos]")
		fmt.Println("")
		fmt.Println("arguments:")
		fmt.Println("  -h,   --help            print help and exit")
		os.Exit(0)
	}

	// Count total repos that need to be pulled.
	var t int
	if len(vals) == 0 {
		t = len(config.Pull)
	} else {
		t = len(vals)
	}

	// TODO: Actually learn gitand check if all these commands are needed.
	var i int
	for n, r := range config.Pull {
		// Skip repos if needed.
		if len(vals) != 0 {
			if !stringInList(n, vals) {
				continue
			}
		}
		i++

		fmt.Printf("Pulling in repo %d/%d, ", i, t)
		color.Set(config.LightColor)
		fmt.Printf(n)
		color.Unset()
		fmt.Println(".")

		l := portFullLoc(n)

		// Check if location exists, clone if it doesn't.
		if _, err := os.Stat(l); err != nil {
			err := gitClone(r.URL, r.Branch, l)
			if err != nil {
				printe(err.Error())
			}
			continue
		}

		if err := gitCheckout(r.Branch, l); err != nil {
			printe(err.Error())
			continue
		}
		if err := gitFetch(l); err != nil {
			printe(err.Error())
			continue
		}

		// Print changes.
		dl, err := gitDiff(r.Branch, l)
		if err != nil {
			printe(err.Error())
			continue
		}
		for _, d := range dl {
			color.Set(config.DarkColor)
			fmt.Print(config.IndentChar)
			color.Unset()
			fmt.Println(d)
		}

		if err := gitClean(l); err != nil {
			printe(err.Error())
			continue
		}
		if err := gitReset(r.Branch, l); err != nil {
			printe(err.Error())
			continue
		}
	}
}