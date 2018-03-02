# go-shell_exec
A small golang program to execute a shell command. It makes use of the os/exec package in golang.

## How to use
### Using regular commands
1. Split the bash commands into words separated by ",".
2. Insert them to the line where the command is input.
3. Compile and run.

>Example
>
>        output, err := exec.Command(echo,"Hello World!").CombinedOutput()


### Using commands that read files from other directories
> The _cmdpath placeholder is used as a placeholder to hold the directory to the file that will be executed as part of the shell command

1. Set the __cmdpath_ directory to match the file to be executed.
2. Set the preceeding shell commands separated by ",".
3. Compile and run
> Example
>
>           _cmdpath := "filename.sh" 
>           output, err := exec.Command("bash", _cmdpath).CombinedOutput()

