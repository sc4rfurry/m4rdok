#!/bin/bash

set -e

ascii_art='\r\n
███    ███ ██   ██ ██████  ██████   ██████  ██   ██ 
████  ████ ██   ██ ██   ██ ██   ██ ██    ██ ██  ██  
██ ████ ██ ███████ ██████  ██   ██ ██    ██ █████   
██  ██  ██      ██ ██   ██ ██   ██ ██    ██ ██  ██  
██      ██      ██ ██   ██ ██████   ██████  ██   ██ 
                                                    
'

desc="--> m4rdok - A Go Lang utility to play with running lolBins (linux) in memory."
banner="$ascii_art \r\n $desc \r \n"

if ! command -v lolcat &>/dev/null; then
    echo "lolcat could not be found"
    echo "Installing lolcat"
    sudo apt-get install lolcat
fi

echo -e "$banner" | lolcat -a -d 5

echo -e "\r\n [+] Checking for Go installation \r\n" | lolcat -a -d 5
# check for go installation
if ! command -v go &>/dev/null; then
    echo "go could not be found"
    echo "Installing go"
    sudo apt-get install golang
fi
echo -e "\r\n [+] Go installation found \r\n" | lolcat -a -d 5 || handle_error "Failed to install dependencies. Please check your Go installation."

echo -e "\r\n [+] Installing dependencies \r\n" | lolcat -a -d 5
# install go dependencies
go get . && go mod tidy && go get github.com/amenzhinsky/go-memexec &>/dev/null || handle_error "Failed to install dependencies. Please check your Go installation."

echo -e "\r\n [+] Dependencies installed \r\n" | lolcat -a -d 5

echo -e "----------------------------------------------------------------------------------------------------\r \n" | lolcat -a -d 22
