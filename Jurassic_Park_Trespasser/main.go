package main

import (
	"os"

	arg "github.com/s9rA16Bf4/ArgumentParser/go/arguments"
	patcher "github.com/s9rA16Bf4/No_CD_Cracks/Jurassic_Park_Trespasser/utility/patcher"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func main() {
	arg.Argument_add("--help", "-h", false, "Shows all available arguments and their purpose", []string{"NULL"})
	arg.Argument_add("--exe", "-x", true, "Path to Trespasser exe [REQUIRED]", []string{"NULL"})
	arg.Argument_add("--smks", "-s", true, "Path to a folder containing the four smk's [REQUIRED]", []string{"NULL"})
	arg.Argument_parse()

	if len(os.Args) > 1 {
		if arg.Argument_check("-h") { // Print help screen
			arg.Argument_help()
		} else {
			var path_to_exe string
			var path_to_smks string

			if arg.Argument_check("-x") { // Patch exe
				path_to_exe = arg.Argument_get("-x")
			} else {
				notify.Error("No exe file was provided", "main.main()")
			}

			if arg.Argument_check("-s") { // Patch exe
				path_to_smks = arg.Argument_get("-s")
			} else {
				notify.Error("No folder was provided", "main.main()")
			}

			patcher.Begin_patch(path_to_exe, path_to_smks)
		}
	} else {
		notify.Error("No argument was provided, run '--help'/'-h' to have a look at the arguments available", "main.main()")
	}
}
