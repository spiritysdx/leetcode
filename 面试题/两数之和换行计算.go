package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    for {
        temp, _ := reader.ReadString('\n')
        temp = strings.ReplaceAll(temp, "\n", "")
        if strings.Count(temp, " ") == 1 {
            tempList := strings.Split(temp, " ")
            a, _ := strconv.Atoi(tempList[0])
            b, _ := strconv.Atoi(tempList[1])
            fmt.Printf("%d\n", a + b)
        } else if len(temp) == 0 {
            break
        }
    }
}
