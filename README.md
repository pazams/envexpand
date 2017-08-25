# envexpand
Expand environment variables in strings

## Example

```golang
package main

import (
	"fmt"
	"github.com/pazams/envexpand"
	"os"
)

func main() {
	os.Setenv("ENVEXPAND_TEST", "google.com")
	defer os.Unsetenv("ENVEXPAND_TEST")

	input := "http://${ENVEXPAND_TEST}"
	output := envexpand.ExpandString(input)

	fmt.Println("input:", input)
	fmt.Println("ouput:", output)
}
```

## License
This package takes a few lines and inspiration from golang's regexp package [source](https://golang.org/src/regexp/regexp.go), namely from `func (re *Regexp) expand(dst []byte, template string, bsrc []byte, src string, match []int) []byte`. 

The use of BSD-3 license follows from that.
