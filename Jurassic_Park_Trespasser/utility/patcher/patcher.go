package patcher

import (
	"encoding/hex"
	"io"
	"os"
	"strings"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

var file_content []string

func Begin(file string) {
	read_file(file)
	i := strings.LastIndex(file, "/")
	file_backup := file[i+1:] + "_backup"

	notify.Inform("Phase 1: Creating a backup of the target")
	create_backup(file, file_backup)

	notify.Inform("Phase 2: Patching the requirement to have a cd inserted.")
	patch_out_cd()

	notify.Inform("Phase 3: Patching an error that would occur if the cd is not present.")
	patch_out_cd_error()

	notify.Inform("Phase 4: Saving edits to the binary file")
	save_exe(file)

	notify.Inform("Phase 5: Done. The result has been written to '" + file + "' and a backup, '" + file_backup + "', has been created if you would like to roll back the changes")
}

func read_file(file string) {
	if strings.Contains(file, ".exe") {
		out, err := os.Open(file)
		if err != nil {
			notify.Error(err.Error(), "patcher.read_file()")
		}
		file_content = make([]string, 0)
		local := make([]byte, 12)
		for err != io.EOF {
			_, err = out.Read(local)
			temp := hex.EncodeToString(local) // This makes it a hella lot easier to handle
			file_content = append(file_content, temp)
		}
		out.Close()
	} else {
		notify.Error("The provided file '"+file+"' doesn't seem to be an .exe", "patcher.read_file()")
	}
}

func create_backup(file string, backup_file_name string) {
	out, err := os.Create(backup_file_name)
	if err != nil {
		notify.Error(err.Error(), "patcher.create_backup()")
	}
	in, err := os.Open(file)
	if err != nil {
		notify.Error(err.Error(), "patcher.create_backup()")
	}
	io.Copy(out, in)
	in.Close()
	out.Close()
}

func patch_out_cd() {
	// We need to replace `OF 85` to `OF 84`, issue is that this is most likely not the only case of this opcode in the exe file
	for i, line := range file_content {
		if line == "ff0f85a60000008b1db05064" {
			file_content[i] = "ff0f84a60000008b1db05064" // Equals jump if equal
			break
		}
	}
}

func patch_out_cd_error() {

}

func save_exe(file string) {
	out, _ := os.Create(file)
	for _, line := range file_content {
		arr, _ := hex.DecodeString(line)
		out.Write(arr)
	}
	out.Close()
}
