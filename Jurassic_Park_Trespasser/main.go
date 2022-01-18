package main

import (
	"os"

	arg "github.com/s9rA16Bf4/ArgumentParser/go/arguments"
	patcher "github.com/s9rA16Bf4/No_CD_Cracks/Jurassic_Park_Trespasser/utility/patcher"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func main() {
	arg.Argument_add("--help", "-h", false, "Shows all available arguments and their purpose", []string{"NULL"})
	arg.Argument_add("--exe", "-x", true, "Trespasser exe to patch [REQUIRED]", []string{"NULL"})
	arg.Argument_parse()

	var file = "temp"

	if len(os.Args) > 1 {
		if arg.Argument_check("-h") {
			arg.Argument_help()
		} else {
			if arg.Argument_check("-x") {
				file = arg.Argument_get("-x")
			}
			patcher.Begin(file)
		}

	} else {
		notify.Error("No argument was provided, run '--help'/'-h' to have a look at the arguments available", "main.main()")
	}
}
