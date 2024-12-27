package utils

import (
    "fmt"
    "strings"
    "strconv"
)

func StringToIntSliceBySep(line string, sep string) []int {
    var slice []int

    tokens := strings.Split(line, sep)
    for _, token := range tokens {
        num, err := strconv.Atoi(token)
        if err != nil {
            _ = fmt.Errorf("ERROR: Invalid line in data file. Couldn't convert token to a integer!")
            return nil
        }

        slice = append(slice, num)
    }

    return slice
}
