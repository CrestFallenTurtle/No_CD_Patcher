package patcher

import (
	"crypto/sha512"
	"encoding/hex"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
	"golang.org/x/sys/windows/registry"
)

type file_t struct {
	name               string // File name
	path_to_exe        string // File path
	path_to_smks       string // Folder path
	path_to_credits    string
	path_to_newgame    string
	path_to_tpassintro string
	path_to_win        string
	path_to_other      string // We need to copy about 50 other files..
}

var c_file file_t

const (
	EXE_NAME        = "trespass.exe" // Target exe name
	EXE_HASH        = "4da71e51e08a7ebea1ecd2caf72482c54a8c07f09af1d44330731f3c554e4f5c6a40eab90856da90518a5735be4783f8344ec263eaf7072aa98ee7c05fa3b08f"
	CREDITS_HASH    = "e775e63a2874edf2f463d46c22216e48d498cb3c2e8d1af212b0d759e653d350e139cb827c97692de142dd179a6cedeba15f420f4b09cfdc918d354da77565db"
	NEWGAME_HASH    = "3d5ec239a026ae10713ebf600ad1ef9e1a647f041d7ea004a3e7a8ed977a09d56b10613de58bf8c2b839ea6dec2e0730acb103071e999ca00492d0b34c9eef2f"
	TPASSINTRO_HASH = "fef933abe8af3ce32a0c859bb7e75c782a8790df89635a796039f6b6cc94b9a70170edf86637babf89dcfc5c218319b239c05b668c64d46806bcd957c8b3443d"
	WIN_HASH        = "2f0fb246948b1cbeec818459ba3fbc4ea942a2450a09f23f3bd6ea10ee4f1aa0d764e2afa561343cc1f7c7031e0ebe61e5bdbfe984d4b5f9992d0f6f5ba11499"
)

func Begin_patch(file_path_exe string, file_path_smks string, file_path_other string) {
	if runtime.GOOS != "windows" {
		notify.Error("This tool works only on windows", "patcher.Begin_patch")
	}
	c_file.name = EXE_NAME
	c_file.path_to_exe = file_path_exe // This should be be the local path to the installation
	c_file.path_to_smks = file_path_smks
	c_file.path_to_other = file_path_other

	notify.Inform("[Phase 1] Checking if target exe is correct")
	check_exe()

	notify.Inform("[Phase 2] Checking if the provided smks are correct")
	check_smk()

	notify.Inform("[Phase 3] Creating the directories")
	create_dir()

	notify.Inform("[Phase 4] Copying 'credits.smk'")
	copy_credits()

	notify.Inform("[Phase 5] Copying 'newgame.smk'")
	copy_newgame()

	notify.Inform("[Phase 6] Copying 'tpassintro.smk'")
	copy_tpassintro()

	notify.Inform("[Phase 7] Copying 'win.smk'")
	copy_win()

	notify.Inform("[Phase 8] Copying other necessary files")
	copy_other()

	notify.Inform("[Phase 9] Changing key value in regedit")
	regedit()

	notify.Inform("All done over here, enjoy over a hot cup of coffee or tea")
}

func internal_hash(path_to_file string) string {
	exe, err := os.Open(path_to_file)
	if err != nil {
		notify.Error(err.Error(), "patcher.internal_hash")
	}
	exe_size, err := exe.Stat()
	if err != nil {
		notify.Error(err.Error(), "patcher.internal_hash")
	}
	buf := make([]byte, exe_size.Size())
	hash := sha512.New()
	for {
		roof, err := exe.Read(buf)
		if err == io.EOF || err != nil {
			break
		} else {
			_, err := hash.Write(buf[:roof])
			if err != nil {
				notify.Error(err.Error(), "patcher.internal_hash")
			}
		}
	}
	return hex.EncodeToString(hash.Sum(nil))
}

func check_exe() {
	if strings.Contains(c_file.path_to_exe, EXE_NAME) { // This is what the exe is called
		hash_sum := internal_hash(c_file.path_to_exe)
		if hash_sum != EXE_HASH {
			notify.Error("Hash values for the exe didn't match. Assuming it's not the correct exe.", "patcher.check_exe")
		}
	} else {
		notify.Error("The provided file '"+c_file.path_to_exe+"' doesn't seem to the target exe '"+EXE_NAME+"'", "patcher.read_file()")
	}
}

func check_smk() {
	folder, err := os.Open(c_file.path_to_smks)
	if err != nil {
		notify.Error(err.Error(), "patcher.check_smk()")
	}
	files, err := folder.ReadDir(0) // Reads all the contents
	if err != nil {
		notify.Error(err.Error(), "patcher.check_smk()")
	}
	hash_smk := make(map[string]string)
	hash_smk["credits.smk"] = "NULL"
	hash_smk["newgame.smk"] = "NULL"
	hash_smk["tpassintro.smk"] = "NULL"
	hash_smk["win.smk"] = "NULL"

	for _, file := range files {
		switch file.Name() {
		case "credits.smk":
			c_file.path_to_credits = c_file.path_to_smks + "credits.smk"
			hash_smk["credits.smk"] = internal_hash(c_file.path_to_smks + file.Name())
		case "newgame.smk":
			c_file.path_to_newgame = c_file.path_to_smks + "newgame.smk"
			hash_smk["newgame.smk"] = internal_hash(c_file.path_to_smks + file.Name())
		case "tpassintro.smk":
			c_file.path_to_tpassintro = c_file.path_to_smks + "tpassintro.smk"
			hash_smk["tpassintro.smk"] = internal_hash(c_file.path_to_smks + file.Name())
		case "win.smk":
			c_file.path_to_win = c_file.path_to_smks + "win.smk"
			hash_smk["win.smk"] = internal_hash(c_file.path_to_smks + file.Name())
		}
	}
	if hash_smk["credits.smk"] != CREDITS_HASH ||
		hash_smk["newgame.smk"] != NEWGAME_HASH ||
		hash_smk["tpassintro.smk"] != TPASSINTRO_HASH ||
		hash_smk["win.smk"] != WIN_HASH {
		notify.Error("One or more hashes for the smks was wrong", "patcher.check_smk()")
	}
}

func create_dir() {
	err := os.MkdirAll("C:\\Program Files\\DreamWorks Interactive\\Trespasser\\data\\menu\\", 0700)
	if err != nil {
		notify.Error(err.Error(), "patcher.create_dir()")
	}
}

func copy_credits() {
	dst, err := os.Create("C:\\Program Files\\DreamWorks Interactive\\Trespasser\\data\\menu\\credits.smk")
	if err != nil {
		notify.Error(err.Error(), "patcher.copy_credits()")
	}
	src, err := os.Open(c_file.path_to_credits)
	if err != nil {
		notify.Error(err.Error(), "patcher.copy_credits()")
	}
	io.Copy(dst, src)
	dst.Close()
	src.Close()
}
func copy_newgame() {
	dst, err := os.Create("C:\\Program Files\\DreamWorks Interactive\\Trespasser\\data\\menu\\newgame.smk")
	if err != nil {
		notify.Error(err.Error(), "patcher.copy_newgame()")
	}
	src, err := os.Open(c_file.path_to_newgame)
	if err != nil {
		notify.Error(err.Error(), "patcher.copy_newgame()")
	}
	io.Copy(dst, src)
	dst.Close()
	src.Close()
}
func copy_tpassintro() {
	dst, err := os.Create("C:\\Program Files\\DreamWorks Interactive\\Trespasser\\data\\menu\\tpassintro.smk")
	if err != nil {
		notify.Error(err.Error(), "patcher.copy_tpassintro()")
	}
	src, err := os.Open(c_file.path_to_tpassintro)
	if err != nil {
		notify.Error(err.Error(), "patcher.copy_tpassintro()")
	}
	io.Copy(dst, src)
	dst.Close()
	src.Close()
}
func copy_win() {
	dst, err := os.Create("C:\\Program Files\\DreamWorks Interactive\\Trespasser\\data\\menu\\win.smk")
	if err != nil {
		notify.Error(err.Error(), "patcher.copy_win()")
	}
	src, err := os.Open(c_file.path_to_win)
	if err != nil {
		notify.Error(err.Error(), "patcher.copy_win()")
	}
	io.Copy(dst, src)
	dst.Close()
	src.Close()
}

func copy_other() {
	folder, err := os.Open(c_file.path_to_other)
	if err != nil {
		notify.Error(err.Error(), "patcher.copy_other()")
	}

	files, err := folder.Readdir(0)
	if err != nil {
		notify.Error(err.Error(), "patcher.copy_other()")
	}

	if len(files) < 53 {
		notify.Warning("Uncorrect amount of files. Cannot guarantee that the game will work with less files.")
	}
	if c_file.path_to_other[len(c_file.path_to_other)-1] != '/' {
		c_file.path_to_other += "/"
	}
	for _, file := range files {
		notify.Inform("Sub-phase] Copying '" + file.Name() + "'")

		dst, err := os.Create("C:\\Program Files\\DreamWorks Interactive\\Trespasser\\data\\" + file.Name())
		if err != nil {
			notify.Error(err.Error(), "patcher.copy_other()")
		}

		src, err := os.Open(c_file.path_to_other + file.Name())
		if err != nil {
			notify.Error(err.Error(), "patcher.copy_other()")
		}
		io.Copy(dst, src)
		dst.Close()
		src.Close()
	}
}

func regedit() {
	reg, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\\DreamWorks Interactive\\Trespasser`, registry.WRITE)

	if err != nil {
		notify.Error(err.Error(), "patcher.regedit")
	}

	err = reg.SetStringValue("Data Drive", "C:\\Program Files\\DreamWorks Interactive\\Trespasser\\")
	if err != nil {
		notify.Error(err.Error(), "patcher.regedit")
	}
	reg.Close()
}
