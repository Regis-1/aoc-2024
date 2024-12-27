package utils

import "testing"

func TestStringToIntSliceBySep(t *testing.T) {
    expect := [][]int{
        {1, 15, -5, 16},
        nil,
    }

    testStrings := []string{
        "1 15 -5 16",
        "s 2i 44dsf 5",
    }

    for i, str := range testStrings {
        got := StringToIntSliceBySep(str, " ")

        if got == nil && expect[i] != nil {
            t.Errorf("got %v, expected %v for '%v'", got, expect[i], str)
            continue
        }

        for j := range got {
            if got[j] != expect[i][j] {
                t.Errorf("got %v, expected %v for '%v'", got, expect[i], str)
                break
            }
        }
    }
}
