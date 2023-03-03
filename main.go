package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/TwiN/go-color"
	"github.com/common-nighthawk/go-figure"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

var Bold = "\033[1m"
var Dim = "\033[2m"
var Underlined = "\033[4m"
var Blink = "\033[5m"
var Reverse = "\033[7m"
var Hidden = "\033[8m"

func init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}
	if len(os.Args) == 1 {
		help()
	}
	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		help()
	}
	if len(os.Args) < 2 {
		help()
	} else {
		tmp_dir := "/tmp/" + os.Args[1]
		if _, err := os.Stat(tmp_dir); os.IsNotExist(err) {
			fmt.Println(color.Bold + color.Red + "Error: " + color.Reset + "Binary not found in /tmp folder. Kindly copy the binary to the /tmp folder and make it executable.")
			os.Exit(1)
		}
	}
}

func banner() {
	var author string = "sc4rfurry"
	var version string = "0.1.0"
	var go_version string = "1.18.1 or higher"
	var github string = "https://github.com/sc4rfurry"
	var description string = "A Go Lang utility to play with running lolBins (linux) in memory." + color.Bold + color.Cyan + "(! Alpha !)." + color.Reset
	banner := figure.NewColorFigure("m4rdok", "", "cyan", true)
	banner.Print()
	println("\n")
	fmt.Println(color.Bold + color.Green + "  " + description + "  " + color.Reset + "\n")
	fmt.Println(color.Bold + color.Purple + "\t  Author: " + color.Reset + author)
	fmt.Println(color.Bold + color.Purple + "\t  Version: " + color.Reset + version)
	fmt.Println(color.Bold + color.Purple + "\t  Go Version: " + color.Reset + go_version)
	fmt.Println(color.Bold + color.Purple + "\t  Github: " + color.Reset + github)
	fmt.Println(color.Bold + color.Purple + "\t  Date: " + color.Reset + time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("-------------------------------------------------------------------------------->")
	println("\n")

}

func help() {
	println(color.Bold + color.Green + "\n\t\t\t~ Help Menu ~" + color.Reset)
	println("\n\t" + color.Bold + color.Cyan + "Usage: " + color.Reset + "./main  <lolBin> --args <lolBinArgs> ")
	println(color.Bold + color.Gray + "___________________________________________________________________________________________________________" + color.Reset)
	println("\n\t" + color.Bold + color.Green + "Options: " + color.Reset)
	println("\t\t" + color.Bold + color.Yellow + "lolBin " + color.Reset + "\t\tsystem binary to run in memory")
	println("\t\t" + color.Bold + color.Yellow + "--args " + color.Reset + "\t\targuments to pass to the binary")
	println("\t\t" + color.Bold + color.Yellow + "-h/--help " + color.Reset + "\tShow this help menu" + "\n")
	println("\n\t" + color.Bold + color.Cyan + "!Note:-" + color.Reset + color.Bold + color.Yellow + " Kindly copy the desired binary to the /tmp folder and make it executable." + color.Reset + "\n")
	os.Exit(0)
}

func force_clean(file string) {
	println(color.Bold + color.Green + "\n\t\t\t~ Force Cleaning Up ~" + color.Reset)
	time.Sleep(1 * time.Second)
	println(color.Bold + color.Green + "\n\t\t\t~ Removing Temp Agenda ~" + color.Reset)
	time.Sleep(1 * time.Second)
	if err := os.Remove(file); err != nil {
		println("\n\t\t\t" + color.Bold + color.Red + "[~] Fatel: " + color.Reset + color.Bold + color.Cyan + err.Error() + color.Reset)
		os.Exit(1)
	}
	println(color.Bold + color.Green + "\n\t\t\t~ Agenda Not Accomplished ...!" + color.Reset)
	time.Sleep(1 * time.Second)
	os.Exit(0)
}

func tempo(code string) {
	println(color.Bold + color.Green + "\n\t\t\t~ Writing Temp Agenda ~" + color.Reset)
	file, err := ioutil.TempFile("", "*.go")
	if err != nil {
		println("\n\t\t\t" + color.Bold + color.Red + "[~] Fatel: " + color.Reset + color.Bold + color.Cyan + err.Error() + color.Reset)
		os.Exit(1)
	}

	defer os.Remove(file.Name())
	defer file.Close()

	if _, err := file.WriteString(code); err != nil {
		println("\n\t\t\t" + color.Bold + color.Red + "[~] Fatel: " + color.Reset + color.Bold + color.Cyan + err.Error() + color.Reset)
		os.Exit(1)
	}
	println(color.Bold + color.Green + "\n\t\t\t~ Compiling Temp Agenda ~" + color.Reset)
	time.Sleep(1 * time.Second)
	println(color.Bold + color.Green + "\n\t\t\t~ Fire in the |.-.| Meme32  |.-.|~" + color.Reset)
	time.Sleep(1 * time.Second)
	cmd := exec.Command("go", "run", file.Name())
	out, err := cmd.CombinedOutput()
	if err != nil {
		println("\n\t\t\t" + color.Bold + color.Red + "[~] Fatel: " + color.Reset + color.Bold + color.Cyan + err.Error() + color.Reset)
		force_clean(file.Name())
		os.Exit(1)
	}
	println(color.Bold + color.Green + "\n\t\t\t~ Agenda Accomplished ...!" + color.Reset)
	println("\n\t\t\t" + color.Bold + color.Cyan + "Output: " + color.Reset + "\nx	" + string(out) + "\n")
	time.Sleep(1 * time.Second)
	println(color.Bold + color.Green + "\n\t\t\t~ Cleaning Up ~" + color.Reset)
	if err := os.Remove(file.Name()); err != nil {
		println("\t\t\t" + color.Bold + color.Red + "[~] Fatel: " + color.Reset + color.Bold + color.Cyan + err.Error() + color.Reset)
		os.Exit(1)
	}
	println(color.Bold + color.Green + "\t\t\t~ All Done ...!" + color.Reset)
}

func custom_code_template(bin string, argz string) {
	println(color.Bold + color.Green + "\n\t\t\t~ Loading Code Template ~" + color.Reset)
	code := `package main

	import (
		_ "embed"
		"fmt"
		"os"
	
		"github.com/amenzhinsky/go-memexec"
	)
	
	//go:embed PLACEHOLDER
	var binex []byte
	
	func run() {
		exe, err := memexec.New(binex)
		if err != nil {
			fmt.Printf("error creating memexec.Executable: %v", err)
			os.Exit(1)
		}
		defer exe.Close()
	
		cmd := exe.Command("ARGZ")
		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("error executing command: %v", err)
			os.Exit(1)
		}
		fmt.Printf("%s", output)
	}
	
	func main() {
		run()
		println("\n")

	}`
	println(color.Bold + color.Green + "\n\t\t\t~ Code Template Loaded ...!" + color.Reset)
	time.Sleep(1 * time.Second)
	code = strings.Replace(code, "PLACEHOLDER", bin, -1)
	if argz != "" {
		code = strings.Replace(code, "ARGZ", argz, -1)
	} else {
		code = strings.Replace(code, "ARGZ", "--help", -1)
	}
	println(color.Bold + color.Green + "\n\t\t\t~ Loading Exec Code ...!" + color.Reset)
	tempo(code)
}

func main() {
	if len(os.Args) == 0 {
		help()
	}
	if len(os.Args) == 2 {
		banner()
		bin := strings.TrimSpace(os.Args[1])
		argz := ""
		custom_code_template(bin, argz)
	} else if len(os.Args) == 4 {
		banner()
		bin := strings.TrimSpace(os.Args[1])
		argz := strings.TrimSpace(os.Args[3])
		custom_code_template(bin, argz)

	} else if os.Args[1] == "-h" || os.Args[1] == "--help" {
		help()
	} else {
		help()
	}
}
