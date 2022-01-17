package patcher

import (
	"io"
	"os"
	"strings"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func read_file(file string) {
	if strings.Contains(file, ".exe") {
		file_backup := file + "_backup"
		out, err := os.Open(file)
		if err != nil {
			notify.Error(err.Error(), "patcher.read_file()")
		}
		notify.Inform("Phase 1: Creating a backup of the target")
		create_backup(out)

		notify.Inform("Phase 2: Patching the requirement to have a cd inserted.")
		patch_out_cd(out)

		notify.Inform("Phase 3: Patching an error that would occur if the cd is not present.")
		patch_out_cd_error(out)

		notify.Inform("Phase 4: Done. The result has been written to '" + file + "' and a backup, '" + file_backup + "', has been created if you would like to roll back the changes")
		out.Close()

	} else {
		notify.Error("The provided file '"+file+"' doesn't seem to be an .exe", "patcher.read_file()")
	}
}

func create_backup(out *os.File) {
	in, err := os.Open(out.Name() + "_backup")
	if err != nil {
		notify.Error(err.Error(), "patcher.create_backup()")
	}
	io.Copy(in, out)
}

func patch_out_cd(out *os.File) {
	// We need to replace `OF 85` to `OF 84`, issue is that this is most likely not the only case of this opcode in the exe file
}

func patch_out_cd_error(out *os.File) {

}
