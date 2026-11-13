package Advanced

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LineFilter() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(ucl)
	}
}
