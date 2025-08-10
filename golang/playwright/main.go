package main

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
)

func main() {
	if err := playwright.Install(); err == nil {
		if playwright, err := playwright.Run(); err == nil {
			browser, e := playwright.Firefox.Launch()
			if e == nil {
				page, _ := browser.NewPage()
				page.Goto("https://google.com")
			}
			defer browser.Close()
		}
	} else {
		fmt.Println("error installing dependencies: ", err.Error())
	}
}
