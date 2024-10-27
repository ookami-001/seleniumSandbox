package selenium_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/aleksslitvinovs/go-selenium"
	"github.com/aleksslitvinovs/go-selenium/keys"
)

func Test(t *testing.T) {
	selenium.SetTest(MyTest)

	selenium.Run()
}

func MyTest(s *selenium.Session) {
	s.OpenURL("https://duckduckgo.com/")

	s.NewElement("#searchbox_input").
		WaitFor(10 * time.Second).UntilIsVisible().
		SendKeys("theRealAlpaca/go-selenium").
		SendKeys(keys.Enter)

	result := s.NewElement("#r1-0 [data-testid=result-title-a]").
		WaitFor(10 * time.Second).UntilIsVisible().
		GetText()

	fmt.Printf("DuckDuckGo result: %s\n", result)
}
