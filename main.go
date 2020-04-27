package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func main() {
	var name string

	fmt.Print("Enter the username you'd like to check: ")
	if _, err := fmt.Scanln(&name); err != nil {
		log.Printf("[ error ] error scanning username: %s", err)
	}

	for {
		if name == "admin" || name == "Admin" {
			fmt.Println("Not a valid name, please choose another: ")
			if _, err := fmt.Scanln(&name); err != nil {
				log.Printf("[ error ] error scanning username: %s", err)
			}
		}

		urls := []string{
			"youtube.com/user/",
			"soundcloud.com/",
			"twitter.com/",
			"instagram.com/",
			"twitch.tv/",
		}

		fmt.Println()
		checkAvailibility(urls, name)
		name = ""
		fmt.Print("\nPress enter to exit or enter another name: ")
		fmt.Scanln(&name)

		if name == "" {
			os.Exit(0)
		}
	}
}

func checkAvailibility(urls []string, name string) {
	// i'm not sure if this only works on windows or not
	w := tabwriter.NewWriter(color.Output, 0, 0, 1, '.', tabwriter.TabIndent)
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()
	for _, url := range urls {
		res, err := http.Get("https://" + url + name)
		if err != nil {
			log.Fatalf("[error]: %s", err)
		}

		if res.StatusCode != 404 {
			fmt.Fprintln(w, url+name+"\t[   "+color.RedString("taken")+"   ]")
		} else {
			fmt.Fprintln(w, url+name+"\t[ "+color.GreenString("available")+" ]")
		}
	}
	s.Stop()
	w.Flush()
}
