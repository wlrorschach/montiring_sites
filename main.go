package main

import (
	"bufio"
	"fmt" // IO
	"io"
	"net/http"
	"os" // We use this package to exit the app and say to OS the status of the aplication execution
	"strconv"
	"strings"
	"time"
)

const (
	monitoring = 3
	delay      = 5
	sitesPath  = "./files/sites.txt"
	logsPath   = "./files/logs/monitoring_sites.log.txt"
)

func main() {

	showIntroduction()

	for {
		showMenu()

		switch readCommand() {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Command not recognized")
			os.Exit(-1)
		}
	}
}

func showMenu() {
	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit Application")
}

func showIntroduction() {
	nome := "Duda da Silva"

	fmt.Println("Hello world, ", nome)
}

func readCommand() int {
	var commandRead int

	fmt.Scan(&commandRead)
	fmt.Println("The command that you chose was: ", commandRead)

	return commandRead
}

func startMonitoring() {
	fmt.Println("Monitoring...")

	sites := getSitesFromAFile()

	for i := 0; i < monitoring; i++ {
		for _, site := range sites {
			makeTheCall(site)
		}
		fmt.Println()
		time.Sleep(delay * time.Second)
	}
}

func makeTheCall(site string) {
	res, err := http.Get(site)

	if err != nil {
		fmt.Println("Error on verifying the site:", err)
		return
	}

	if res.StatusCode == 200 {
		registerLog(site, true)
		fmt.Println("Site ", site, " was loaded with success!")
		return
	}

	registerLog(site, false)
	fmt.Println("Site: ", site, " has a problem. Status Code: ", res.StatusCode)
}

func showLogs() {
	fmt.Println("Showing logs...")

	logs := readFile(logsPath)

	for _, line := range logs {
		fmt.Println(line)
	}

}

func getSitesFromAFile() []string {
	var slice []string
	slice = readFile(sitesPath)
	return slice
}

/**
Read a file and return a slice where each element would be a line form the file
*/
func readFile(path string) []string {
	var slice []string

	file, err := os.Open(path)
	reader := bufio.NewReader(file)

	if err != nil {
		fmt.Println("An error occurs:", err)
	}

	for {
		line, err2 := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		slice = append(slice, line)

		if err2 == io.EOF {
			break
		}
	}
	file.Close()
	return slice
}

func registerLog(site string, status bool) {
	file, err := os.OpenFile("./files/logs/monitoring_sites.log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("An error occurs at the reading of the file. Error:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}
