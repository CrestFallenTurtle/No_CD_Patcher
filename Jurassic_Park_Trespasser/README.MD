### Process to check if a disk is inserted
The program tries to load a file from the disk with help of `fileapi.getFileAttributesA` the return value from the function if it fails is `INVALID_FILE_ATTRIBUTES`
and the program compares this to check if an actual disk was inserted.

### Process to patch out
![Picture showing disassembled code](https://github.com/s9rA16Bf4/No_CD_Cracks/blob/main/pictures/trespasser_pic_1.png)

By replacing the row `jne trespass.40DCE5` (OPCODE `0F 85`) to `je trespass.40DCE5` (OPCODE `OF 84`) we make sure that the program doesn't exit because of the disk
not being inserted.

But we got a new issue now, that is that the file that the program tried to load from the disk is the intro video `tpassintro.smk` which now means that the program 
boots to a black screen and you essentially softbricked your computer untill restart.
