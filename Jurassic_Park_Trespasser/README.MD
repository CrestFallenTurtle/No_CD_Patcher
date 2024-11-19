![The games art box](../pictures/trespasser/Jp-trespasser-cover.png)

### Process to check if a disk is inserted
The program tries to load a file from the disk with help of `fileapi.getFileAttributesA` the return value from the function if it fails is `INVALID_FILE_ATTRIBUTES`
and the program compares this to check if an actual disk was inserted.

### Process to patch out
![Picture showing disassembled code](../pictures/trespasser/pic1.png)

By replacing the row `jne trespass.40DCE5` (OPCODE `0F 85`) to `je trespass.40DCE5` (OPCODE `OF 84`) we make sure that the program doesn't exit because of the disk
not being inserted.

But we got a new issue now, that is that the file that the program tried to load from the disk is the intro video `tpassintro.smk` which now means that the program 
boots to a black screen and you essentially softbricked your computer untill restart.

By continuing to observe the code we can see that the program tries to load `tpassintro.smk`, `credits.smk`, `newgame.smk` and `win.smk` from the disk.<br>
![Picture showing the necessary files](../pictures/trespasser/pic4.png)

My observation leads to that if we would try to replace all the codes were the program tries to load one out of the four `.smk` files with a `nop` instruction we essentially would make the game unstable to run and/or stopping it from working entirerly. 

A lot simpler solution was found by observing the games registry, located at `HKEY_LOCAL_MACHINE\SOFTWARE\DreamWorks Interactive\Trespasser`<br>
![Picture showing windows regedit](../pictures/trespasser/pic2.png)

What we see in the keyvalue `Data Drive` is simply where the system will start to look for the folders `data/menu` which would contain the four `.smk`, and this is the final part of the puzzel to crack for a no-cd requirement. The solution that I'll go with and the program aswell is to replace the the keyvalue from `D:\` to `C:\Program Files\DreamWorks Interactive\Trespasser\` but also creating the directories `data/menu` in the Trespasser directory which in turn contains the necessary files.

![Picture showing windows regedit but with modified values](../pictures/trespasser/pic3.png)<br>

### For the program to work
For the program to work, you will need to pass a valid .exe file to verify ownage of the game (how you get it is up to you) but you will also be required to pass a folder containing the necessary `.smk` files. The program will then handle everything else when it comes to modify the keyvalue in regedit. It will also be required to run the program with escalated privileges.

### Dependencies
The dependencies for this project is 
1. Registry - `go get "golang.org/x/sys/windows/registry"`
2. Notify - `go get "github.com/CrestFallenTurtle/notify_handler/go/notify"`

All of the dependencies can be installed by running `make install_dependencies`

### PlayOnLinux/Wine
<b>Note:</b> The program will still be needed to compiled against a windows architecture, but we will instead apply it with wine

If you would like to apply this patch on Playonlinux, my suggestion is to ...
1. Install it under wine version `7.0-staging`
2. Go to the settings in playonlinux and select `Miscellaneous`
3. Click on `Open a shell`
4. Once the shell is open you can locate the patcher and continue as normal.
