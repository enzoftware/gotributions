package main

import "flag"

func main() {
	var folder string
	var email = "lizama.enzo@gmail.com"

	flag.StringVar(&folder, "add", "", "add a new folder to scan for git repositories")
	flag.StringVar(&email, "email", "lizama.enzo@gmail.com", "the email to scan")

	if folder != "" {
		scan(folder)
		return
	}
	stats(email)
}
