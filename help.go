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

	fmt.Printf("Normal:\n")
	fmt.Printf("\x1b[0;31m%s\x1b[0m\n", "Superior.Print('HELLO', 'normal', 'red')")
	fmt.Printf("\x1b[0;32m%s\x1b[0m\n", "Superior.Print('HELLO', 'normal', 'green')")
	fmt.Printf("\x1b[0;33m%s\x1b[0m\n", "Superior.Print('HELLO', 'normal', 'yellow')")
	fmt.Printf("\x1b[0;34m%s\x1b[0m\n", "Superior.Print('HELLO', 'normal', 'blue')")
	fmt.Printf("\x1b[0;35m%s\x1b[0m\n", "Superior.Print('HELLO', 'normal', 'purple')")
	fmt.Printf("\x1b[0;36m%s\x1b[0m\n", "Superior.Print('HELLO', 'normal', 'cyan')")
	fmt.Printf("\x1b[0;37m%s\x1b[0m\n", "Superior.Print('HELLO', 'normal', 'white')")
	fmt.Printf("\n")

	fmt.Printf("Bold:\n")
	fmt.Printf("\x1b[1;31m%s\x1b[0m\n", "Superior.Print('HELLO', 'bold', 'red')")
	fmt.Printf("\x1b[1;32m%s\x1b[0m\n", "Superior.Print('HELLO', 'bold', 'green')")
	fmt.Printf("\x1b[1;33m%s\x1b[0m\n", "Superior.Print('HELLO', 'bold', 'yellow')")
	fmt.Printf("\x1b[1;34m%s\x1b[0m\n", "Superior.Print('HELLO', 'bold', 'blue')")
	fmt.Printf("\x1b[1;35m%s\x1b[0m\n", "Superior.Print('HELLO', 'bold', 'purple')")
	fmt.Printf("\x1b[1;36m%s\x1b[0m\n", "Superior.Print('HELLO', 'bold', 'cyan')")
	fmt.Printf("\x1b[1;37m%s\x1b[0m\n", "Superior.Print('HELLO', 'bold', 'white')")
	fmt.Printf("\n")

	fmt.Printf("Under:\n")
	fmt.Printf("\x1b[4;31m%s\x1b[0m\n", "Superior.Print('HELLO', 'under', 'red')")
	fmt.Printf("\x1b[4;32m%s\x1b[0m\n", "Superior.Print('HELLO', 'under', 'green')")
	fmt.Printf("\x1b[4;33m%s\x1b[0m\n", "Superior.Print('HELLO', 'under', 'yellow')")
	fmt.Printf("\x1b[4;34m%s\x1b[0m\n", "Superior.Print('HELLO', 'under', 'blue')")
	fmt.Printf("\x1b[4;35m%s\x1b[0m\n", "Superior.Print('HELLO', 'under', 'purple')")
	fmt.Printf("\x1b[4;36m%s\x1b[0m\n", "Superior.Print('HELLO', 'under', 'cyan')")
	fmt.Printf("\x1b[4;37m%s\x1b[0m\n", "Superior.Print('HELLO', 'under', 'white')")
	fmt.Printf("\n")

	fmt.Printf("Background:\n")
	fmt.Printf("\x1b[31m%s\x1b[0m\n", "Superior.Print('HELLO', 'background', 'red')")
	fmt.Printf("\x1b[32m%s\x1b[0m\n", "Superior.Print('HELLO', 'background', 'green')")
	fmt.Printf("\x1b[33m%s\x1b[0m\n", "Superior.Print('HELLO', 'background', 'yellow')")
	fmt.Printf("\x1b[34m%s\x1b[0m\n", "Superior.Print('HELLO', 'background', 'blue')")
	fmt.Printf("\x1b[35m%s\x1b[0m\n", "Superior.Print('HELLO', 'background', 'purple')")
	fmt.Printf("\x1b[36m%s\x1b[0m\n", "Superior.Print('HELLO', 'background', 'cyan')")
	fmt.Printf("\x1b[37m%s\x1b[0m\n", "Superior.Print('HELLO', 'background', 'white')")
	fmt.Printf("\n")

}
