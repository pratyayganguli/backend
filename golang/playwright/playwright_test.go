package playwright

import (
	"testing"

	"github.com/playwright-community/playwright-go"
)

func TestInstall(t *testing.T) {
	// install the dependencies
	// this test case does not work, please use this command
	// go run github.com/playwright-community/playwright-go/cmd/playwright@v0.xxxx.x install --with-deps
	if err := playwright.Install(); err != nil {
		t.Fatalf("could not install playwright dependencies %s", err.Error())
	}
}

func TestLaunchBrowser(t *testing.T) {
	if playwright, err := playwright.Run(); err == nil {
		if browser, e := playwright.Chromium.Launch(); e == nil {
			page, _ := browser.NewPage()
			page.Goto("https://google.com")
		}
	}
	// log the error later
}
