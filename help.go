//
// Superior - Toolkit for the Go programming language
// Available at http://github.com/liamka/Printm
//
// Copyright © Kirill Kotikov <liamka@me.com>.
// Superior is open-sourced software licensed under the [MIT license](http://opensource.org/licenses/MIT)
// See README.md for details.
//

package Superior

import (
	"fmt"
)

func help() {

	fmt.Fprintf(out, "Normal:\n")
	fmt.Fprintf(out, "\x1b[0;31m%s\x1b[0m\n", "Superior.Print('HELLO', 'normal', 'red')")
	fmt.Fprintf(out, "\x1b[0;32m%s\x1b[0m\n", "Superior.Print('HELLO', 'normal', 'green')")
	fmt.Fprintf(out, "\x1b[0;33m%s\x1b[0m\n", "Superior.Print('HELLO', 'normal', 'yellow')")
	fmt.Fprintf(out, "\x1b[0;34m%s\x1b[0m\n", "Superior.Print('HELLO', 'normal', 'blue')")
	fmt.Fprintf(out, "\x1b[0;35m%s\x1b[0m\n", "Superior.Print('HELLO', 'normal', 'purple')")
	fmt.Fprintf(out, "\x1b[0;36m%s\x1b[0m\n", "Superior.Print('HELLO', 'normal', 'cyan')")
	fmt.Fprintf(out, "\x1b[0;37m%s\x1b[0m\n", "Superior.Print('HELLO', 'normal', 'white')")
	fmt.Fprintf(out, "\n")

	fmt.Fprintf(out, "Bold:\n")
	fmt.Fprintf(out, "\x1b[1;31m%s\x1b[0m\n", "Superior.Print('HELLO', 'bold', 'red')")
	fmt.Fprintf(out, "\x1b[1;32m%s\x1b[0m\n", "Superior.Print('HELLO', 'bold', 'green')")
	fmt.Fprintf(out, "\x1b[1;33m%s\x1b[0m\n", "Superior.Print('HELLO', 'bold', 'yellow')")
	fmt.Fprintf(out, "\x1b[1;34m%s\x1b[0m\n", "Superior.Print('HELLO', 'bold', 'blue')")
	fmt.Fprintf(out, "\x1b[1;35m%s\x1b[0m\n", "Superior.Print('HELLO', 'bold', 'purple')")
	fmt.Fprintf(out, "\x1b[1;36m%s\x1b[0m\n", "Superior.Print('HELLO', 'bold', 'cyan')")
	fmt.Fprintf(out, "\x1b[1;37m%s\x1b[0m\n", "Superior.Print('HELLO', 'bold', 'white')")
	fmt.Fprintf(out, "\n")

	fmt.Fprintf(out, "Under:\n")
	fmt.Fprintf(out, "\x1b[4;31m%s\x1b[0m\n", "Superior.Print('HELLO', 'under', 'red')")
	fmt.Fprintf(out, "\x1b[4;32m%s\x1b[0m\n", "Superior.Print('HELLO', 'under', 'green')")
	fmt.Fprintf(out, "\x1b[4;33m%s\x1b[0m\n", "Superior.Print('HELLO', 'under', 'yellow')")
	fmt.Fprintf(out, "\x1b[4;34m%s\x1b[0m\n", "Superior.Print('HELLO', 'under', 'blue')")
	fmt.Fprintf(out, "\x1b[4;35m%s\x1b[0m\n", "Superior.Print('HELLO', 'under', 'purple')")
	fmt.Fprintf(out, "\x1b[4;36m%s\x1b[0m\n", "Superior.Print('HELLO', 'under', 'cyan')")
	fmt.Fprintf(out, "\x1b[4;37m%s\x1b[0m\n", "Superior.Print('HELLO', 'under', 'white')")
	fmt.Fprintf(out, "\n")

	fmt.Fprintf(out, "Background:\n")
	fmt.Fprintf(out, "\x1b[31m%s\x1b[0m\n", "Superior.Print('HELLO', 'background', 'red')")
	fmt.Fprintf(out, "\x1b[32m%s\x1b[0m\n", "Superior.Print('HELLO', 'background', 'green')")
	fmt.Fprintf(out, "\x1b[33m%s\x1b[0m\n", "Superior.Print('HELLO', 'background', 'yellow')")
	fmt.Fprintf(out, "\x1b[34m%s\x1b[0m\n", "Superior.Print('HELLO', 'background', 'blue')")
	fmt.Fprintf(out, "\x1b[35m%s\x1b[0m\n", "Superior.Print('HELLO', 'background', 'purple')")
	fmt.Fprintf(out, "\x1b[36m%s\x1b[0m\n", "Superior.Print('HELLO', 'background', 'cyan')")
	fmt.Fprintf(out, "\x1b[37m%s\x1b[0m\n", "Superior.Print('HELLO', 'background', 'white')")
	fmt.Fprintf(out, "\n")

}
