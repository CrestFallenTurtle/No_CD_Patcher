package main

import (
	argumentparser "github.com/CrestFallenTurtle/ArgumentParser"
	patcher "github.com/CrestFallenTurtle/No_CD_Cracks/Jurassic_Park_Trespasser/utility/patcher"
)

func main() {
	handler := argumentparser.Constructor(true)

	handler.Add("--exe", "-x", true, true, "Path to Trespasser exe")
	handler.Add("--smks", "-s", true, true, "Path to a folder containing the four smk's")
	handler.Add("--levels", "-l", true, true, "Path to all the different levels and other materials")
	parsed_flags := handler.Parse()

	path_to_exe := ""
	path_to_smks := ""
	path_to_lvl := ""

	for key, value := range parsed_flags {
		switch key {
		case "-x":
			path_to_exe = value

		case "-s":
			path_to_smks = value

		case "-l":
			path_to_lvl = value
		}
	}

	patcher.Begin_patch(path_to_exe, path_to_smks, path_to_lvl)
}
