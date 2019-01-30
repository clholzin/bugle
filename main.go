package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	fmt.Println("Bugle")

	netArg := "-an"
	interval := 2
	var err error
	var search []string
	var hasSearch bool
	args := os.Args
	args = args[1:]
	tableHeader := []string{`Active`, `Registered`, `Address`, `Socket`, `kctlref`, `Proto`, `pcb`}
	argsCheck := len(args) >= 2
	if argsCheck {
		if len(args) >= 1 {
			netArg = args[0]
		}
		if len(args) >= 2 {
			interval, err = strconv.Atoi(args[1])
			if err != nil {
				interval = 2
			}
		}
		if len(args) >= 3 {
			search = args[2:] //strings.Join(args[2:], " ")
			fmt.Println("Searching ...", search)
			hasSearch = !hasSearch
		} else {
			fmt.Println("Showing All ", args)
			search = make([]string, 0)
		}
	} else {
		if len(args) > 1 && args[1] == "-h" {
			fmt.Printf("Arguments...\n" +
				"-an : netstat arguments\n" +
				"2 : integer for time in seconds to refresh\n" +
				"80-ESTABLISHED *.80-LISTEN : Search delimited by spaces\n")
			return
		}
		fmt.Println("Missing arguments ", args)
		return
	}
	var captureOut strings.Builder
	width, _, _ := terminal.GetSize(0)
	var prevRet int
	for {

		cmd := exec.Command("netstat", netArg, ``)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}
		var ret int
		if argsCheck {
			paddit := func(value string) {
				diff := width - len(value)
				if diff < 0 {
					value = value[:width-(width/10)] + "..."
				}
				if len(value) > (width - (width / 10)) {
					value = value[:len(value)-(width/10)] + "..."
				}
				captureOut.WriteString(value + "\n")
				ret++
			}
			scan := bufio.NewScanner(stdout)
			for scan.Scan() {
				vals := scan.Text()
				var headerFound bool
				if len(vals) >= 1 {

					fields := strings.Fields(vals)
					for _, field := range fields {
						for _, header := range tableHeader {
							if strings.Contains(field, header) {
								headerFound = true
								break
							}
						}
						if headerFound {
							break
						}

					}
					if headerFound {
						paddit(vals)
						continue
					}

				}
				if !headerFound && !hasSearch {
					paddit(vals)
					continue
				}
				for _, find := range search {

					if strings.Index(find, "-") > -1 {
						se := strings.Split(find, "-")

						if strings.Contains(vals, se[0]) && strings.Contains(vals, se[1]) {
							paddit(vals)
						}
					} else {
						if strings.Contains(vals, find) {
							paddit(vals)
							break
						}

					}

				}
			}

		}

		if err := cmd.Wait(); err != nil {
			log.Fatal(err)
		}
		for i := 0; i < prevRet; i++ {
			fmt.Printf("\033[K")
			fmt.Printf("\033[1A")
		}
		if captureOut.Len() == 0 {
			ret = 1
			fmt.Printf("Empty...Try again\n")
		} else {
			fmt.Printf("%s", captureOut.String())
		}
		time.Sleep(time.Duration(interval) * time.Second)

		captureOut.Reset()
		prevRet = ret
		width, _, _ = terminal.GetSize(0)

	}
}
