package patcher

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func Begin(file string) {
	out := read_file(file)
	i := strings.LastIndex(file, "/")

	file_backup := file[i+1:] + "_backup"

	notify.Inform("Phase 1: Creating a backup of the target")
	create_backup(out, file_backup)

	notify.Inform("Phase 2: Patching the requirement to have a cd inserted.")
	patch_out_cd(out)

	notify.Inform("Phase 3: Patching an error that would occur if the cd is not present.")
	patch_out_cd_error(out)

	notify.Inform("Phase 4: Done. The result has been written to '" + file + "' and a backup, '" + file_backup + "', has been created if you would like to roll back the changes")
	out.Close()
}

func read_file(file string) *os.File {
	if strings.Contains(file, ".exe") {
		out, err := os.Open(file)
		if err != nil {
			notify.Error(err.Error(), "patcher.read_file()")
		}
		return out

	} else {
		notify.Error("The provided file '"+file+"' doesn't seem to be an .exe", "patcher.read_file()")
		return nil
	}
}

func create_backup(out *os.File, backup_file_name string) {
	in, err := os.Create(backup_file_name)
	if err != nil {
		notify.Error(err.Error(), "patcher.create_backup()")
	}
	io.Copy(in, out)
}

func patch_out_cd(out *os.File) {
	// We need to replace `OF 85` to `OF 84`, issue is that this is most likely not the only case of this opcode in the exe file
	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func patch_out_cd_error(out *os.File) {

}
