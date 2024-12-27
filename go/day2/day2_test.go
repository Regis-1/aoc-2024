package day2

import "testing"

func TestIsDecOrIncWithinRange(t *testing.T) {
    testArray := [...]int{1, 2, 2, 6, 5}
    validRange := [...]float64{1, 3}

    results := [...]struct{
        flag bool
        relation LevelRelation
    }{
        {true, Increase},
        {false, NotChanging},
        {false, Increase},
        {true, Decrease},
    }

    for i := 0; i < len(testArray) - 2; i++ {
        flag, relation := IsDecOrIncWithinRange(testArray[i:i+2], validRange)

        if flag != results[i].flag || relation != results[i].relation {
            t.Errorf("got (%v, %d) wanted (%v, %d)",
                flag,
                relation,
                results[i].flag,
                results[i].relation,
            )
        }
    }
}

func TestIsValidReport(t *testing.T) {
    testReports := [...][5]int{
        {1, 2, 3, 4, 5},
        {7, 5, 4, 2, 1},
        {2, 3, 5, 5, 6},
        {1, 3, 5, 4, 7},
    }
    validRange := [...]float64{1, 3}

    expected := [...]bool{true, true, false, false}

    for i, report := range(testReports) {
        got := IsValidReport(report[:], validRange)
        if got != expected[i] {
            t.Errorf("got %v, expected %v", got, expected[i])
        }
    }
}

func TestIsValidReportWithOneLvlTolerance(t *testing.T) {
    testReports := [...][5]int{
        {7,6,4,2,1},
        {1,2,7,8,9},
        {9,7,6,2,1},
        {1,3,2,4,5},
        {8,6,4,4,1},
        {1,3,6,7,9},
    }
    validRange := [...]float64{1, 3}

    expected := [...]bool{true, false, false, true, true, true}

    for i, report := range(testReports) {
        got := IsValidReportWithOneLvlTolerance(report[:], validRange)
        if got != expected[i] {
            t.Errorf("got %v, expected %v", got, expected[i])
        }
    }
}
