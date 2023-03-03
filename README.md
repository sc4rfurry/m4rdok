<h1 align="center">
  m4rdok
</h1>

<h4 align="center">A Go Lang utility to play with running lolBins (linux) in memory.</h4>
<div style="text-align:center">
    <div style="align:center">
    <img src="https://img.shields.io/badge/Author-sc4rfurry-informational?style=flat-square&logo=github&logoColor=white&color=5194f0&bgcolor=110d17" alt="Author">
    <img src="https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square&logo=github&logoColor=white&color=5194f0&bgcolor=110d17" alt="Version">
    <img src="https://img.shields.io/badge/Go_Version-1.18.1-informational?style=flat-square&logo=Go&logoColor=cyan&color=5194f0&bgcolor=110d17" alt="Go Version">
    <img src="https://img.shields.io/badge/OS-Linux-informational?style=flat-square&logo=ubuntu&logoColor=green&color=5194f0&bgcolor=110d17" alt="OS">
    </div>
</div>

#

## Table of Contents

- [Installation](#installation)
- [Features](#features)
- [Running m4rdok](#running-m4rdok)
- [Options](#options)
- [Building m4rdok](#building-m4rdok)
- [Pre-Compiled Binaries](#pre-compiled-binaries)
- [References](#references)
- [Contributing](#contributing)
- [License](#license)


#

### 🔧 Technologies & Tools

![](https://img.shields.io/badge/Editor-VS_Code-informational?style=flat-square&logo=visual-studio&logoColor=blue&color=5194f0)
![](https://img.shields.io/badge/Language-Go-informational?style=flat-square&logo=Go&logoColor=cyan&color=5194f0&bgcolor=110d17)
![](https://img.shields.io/badge/Go_Version-1.18.1-informational?style=flat-square&logo=Go&logoColor=cyan&color=5194f0&bgcolor=110d17)

#

### 📚 Requirements
> - Go 18.1 linux/amd64

#
### Installation

- sudo apt-get update && sudo apt-get golang
- git clone https://github.com/sc4rfurry/m4rdok.git
- cd m4rdok
- go get .
- go get github.com/amenzhinsky/go-memexec  # for lolbin in-memory execution
    > You can also use `install_deps.sh` script to install dependencies.
- go build main.go
    - or use the `builder.sh` script to build the tool.


### Features

- Can load a local lolBin and run it in memory.
    > `--args` option can only parse one argument. 
#

## Running m4rdok
```sh
go run main.go --help
go run main.go lolBin --args <arg>
```
> `! Note !` - The lolBin must be in /tmp directory.

+ ### lolBin
    - `lolBin` - lolBin to run in memory.
    - `--args` - Argument to pass to the lolBin.
    - `-h/--help` - Show help menu.

</br>

+ ### Testing lolBins
    - try the following command on your terminal.
    ```sh
    sudo cp /usr/bin/dir /tmp
    chown $USER:$USER /tmp/dir
    chmod +x /tmp/dir
    ```


### Example

To run the tool on a target, just use the following command.
- using above lolBin `dir`.

```console
go run main.go dir --args -la
```
#

## Options
```sh
				~ Help Menu ~

	Usage: ./main  <lolBin> --args <lolBinArgs> 
___________________________________________________________________________________________________________

	Options: 
		lolBin 		system binary to run in memory
		--args 		arguments to pass to the binary
		-h/--help 	Show this help menu


	!Note:- Kindly copy the desired binary to the /tmp folder and make it executable.
```

## Building m4rdok
> To build the tool, you can use the following command.
```sh
env GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w -extldflags "-static"' -o m4rdok main.go
```

> You can also use the bultin Bash script to build the tool.

- Before running the script, make sure to give it execution permissions.
- The bash script can build both Linux and Windows binaries.
- Binaries will be Stripped and Compressed. (lolcat, strip and upx are required)
```sh
chmod +x builder.sh
./builder.sh main.go
```
#
## Pre-Compiled Binaries
<div>
<div style="text-align:center">
    <a href="https://github.com/sc4rfurry/m4rdok/releases/tag/v1.0.0">
    <img src="https://img.shields.io/badge/Download-v1.0.0-informational?style=flat-square&logo=github&logoColor=white&color=5194f0&bgcolor=110d17" alt="Download">
    </a>
<div style="text-align:center">
    <img src="https://img.shields.io/badge/Status-Active-informational?style=flat-square&logo=github&logoColor=white&color=5194f0&bgcolor=110d17" alt="Status">
</div>
</div>
</div>

#

## References
* Special thanks to the following projects (they are the base of this tool):
    - [go-memexec](https://github.com/amenzhinsky/go-memexec)
    - [go-color](https://github.com/TwiN/go-color)
    - [go-figure](https://github.com/common-nighthawk/go-figure)

#

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)

[def]: https://img.shields.io/badge/OS-Linux-informational?style=flat-square&logo=ubuntu&logoColor=green&color=5194f0&bgcolor=110d17