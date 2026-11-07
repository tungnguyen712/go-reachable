// Sample app for the ground-truth tests. It pins vulnerable dependency
// versions ON PURPOSE. Never deploy it.
//
// Calls the vulnerable function of three fixture advisories. The "ref:" markers
// name the advisory a line is expected to be reported for.
package main

import (
	"fmt"

	gu "github.com/Masterminds/goutils" // renamed import
	. "golang.org/x/text/language"      // dot import
	"gopkg.in/yaml.v3"
)

func main() {
	tags, _, err := ParseAcceptLanguage("da, en-gb;q=0.8, en;q=0.7") // ref:GO-2022-1059
	if err != nil {
		fmt.Println("bad header:", err)
	}
	fmt.Println(tags)

	var cfg map[string]any
	if err := yaml.Unmarshal([]byte("name: demo\n"), &cfg); err != nil { // ref:GO-2022-0603
		fmt.Println("bad config:", err)
	}

	token, err := gu.RandomAlphaNumeric(16) // ref:GO-2022-0411
	if err != nil {
		fmt.Println("token:", err)
	}
	fmt.Println(cfg, token)
}
