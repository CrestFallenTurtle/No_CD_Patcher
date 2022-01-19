package patcher

import (
	"encoding/hex"
	"io"
	"os"
	"strings"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

type file_t struct {
	gut  []string // Contains everything in the target file
	name string   // File name
	path string   // File path
}

var c_file file_t

const (
	EXE_NAME   = "trespass.exe" // Target exe name
	EXE_BACKUP = EXE_NAME + "_backup"
)

func Begin_rollback(file_path string) {
	c_file.name = EXE_NAME
	c_file.path = file_path
	read_file()

	notify.Inform("Phase 1: Creating a backup of the target")
	create_backup()

	notify.Inform("Phase 2: Removes the no-cd patch")
	patch_out_cd(true)

	notify.Inform("Phase 3: Removes the patch for the error that occurs")
	patch_out_cd_error(true)

	notify.Inform("Phase 4: Saving edits to the binary file")
	save_exe()

	notify.Inform("Phase 5: Done. The result has been written to '" + EXE_NAME + "' and a backup, '" + EXE_BACKUP + "', has been created in the local directory if you would like to roll back the changes")
}

func Begin_patch(file_path string) {
	c_file.name = EXE_NAME
	c_file.path = file_path

	read_file()

	notify.Inform("Phase 1: Creating a backup of the target")
	create_backup()

	notify.Inform("Phase 2: Patching the requirement to have a cd inserted.")
	patch_out_cd(false)

	notify.Inform("Phase 3: Patching an error that would occur if the cd is not present.")
	patch_out_cd_error(false)

	notify.Inform("Phase 4: Saving edits to the binary file")
	save_exe()

	notify.Inform("Phase 5: Done. The result has been written to '" + EXE_NAME + "' and a backup, '" + EXE_BACKUP + "', has been created in the local directory if you would like to roll back the changes")
}

func read_file() {
	if strings.Contains(c_file.path, EXE_NAME) { // This is what the exe is called
		out, err := os.Open(c_file.path)
		if err != nil {
			notify.Error(err.Error(), "patcher.read_file()")
		}
		c_file.gut = make([]string, 0)
		local := make([]byte, 12)
		for {
			_, err = out.Read(local)
			if err == io.EOF {
				break
			}
			temp := hex.EncodeToString(local) // This makes it a hella lot easier to handle
			c_file.gut = append(c_file.gut, temp)
		}
		out.Close()
	} else {
		notify.Error("The provided file '"+c_file.path+"' doesn't seem to the target exe '"+EXE_NAME+"'", "patcher.read_file()")
	}
}

func create_backup() {
	dst, err := os.Create(EXE_BACKUP) // Will create it in the local directory
	if err != nil {
		notify.Error(err.Error(), "patcher.create_backup()")
	}
	src, err := os.Open(c_file.path) // Open the file
	if err != nil {
		notify.Error(err.Error(), "patcher.create_backup()")
	}
	io.Copy(dst, src) // Copy everything from the src to the dst
	src.Close()
	dst.Close()
}

func patch_out_cd(reverse_work bool) {
	if !reverse_work {
		for i, line := range c_file.gut {
			if line == "ff0f85a60000008b1db05064" {
				c_file.gut[i] = "ff0f84a60000008b1db05064" // Equals jump if equal
				break
			}
		}
	} else {
		for i, line := range c_file.gut {
			if line == "ff0f84a60000008b1db05064" {
				c_file.gut[i] = "ff0f85a60000008b1db05064"
				break
			}
		}
	}
}

func patch_out_cd_error(reverse_work bool) {
	if !reverse_work {

	} else {

	}
}

func save_exe() {
	dst, _ := os.Create(c_file.name) // Creates our modified exe
	for _, line := range c_file.gut {
		arr, _ := hex.DecodeString(line) // Write all the content
		dst.Write(arr)
	}
	dst.Close()
}
