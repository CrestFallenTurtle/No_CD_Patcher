package main

import (
	"os"

	arg "github.com/s9rA16Bf4/ArgumentParser/go/arguments"
	patcher "github.com/s9rA16Bf4/No_CD_Cracks/Jurassic_Park_Trespasser/utility/patcher"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func main() {
	arg.Argument_add("--help", "-h", false, "Shows all available arguments and their purpose", []string{"NULL"})
	arg.Argument_add("--exe", "-x", true, "Path to Trespasser exe to patch [REQUIRED]", []string{"NULL"})
	arg.Argument_add("--reverse", "-r", true, "Path to Trespasser exe to reverse any changes made", []string{"NULL"})
	arg.Argument_parse()

	if len(os.Args) > 1 {
		if arg.Argument_check("-h") { // Print help screen
			arg.Argument_help()
		} else if arg.Argument_check("-r") { // Reverse changes
			patcher.Begin_rollback(arg.Argument_get("-r"))
		} else {
			if arg.Argument_check("-x") { // Patch exe
				patcher.Begin_patch(arg.Argument_get("-x"))
			} else {
				notify.Error("No exe file was provided", "main.main()")
			}
		}
	} else {
		notify.Error("No argument was provided, run '--help'/'-h' to have a look at the arguments available", "main.main()")
	}
}
