package day2

import (
	"aoc-2024-go/utils"
	"bufio"
	"math"
	"os"
)

type LevelRelation = int

const (
    Increase LevelRelation = iota
    Decrease
    NotChanging
)

func IsDecOrIncWithinRange(frame []int, ran [2]float64) (bool, LevelRelation) {
    var relation LevelRelation
    diff := float64(frame[0]) - float64(frame[1])

    if diff == 0 {
        return false, NotChanging
    } else if diff > 0 {
        relation = Decrease
    } else {
        relation = Increase
    }

    if math.Abs(diff) < ran[0] || math.Abs(diff) > ran[1] {
        return false, relation
    }

    return true, relation
}

func IsValidReport(report []int, validRange [2]float64) bool {
    var definedRelation LevelRelation
    for i := 0; i < len(report) - 1; i++ {
        result, relation := IsDecOrIncWithinRange(report[i:i+2], validRange)

        if !result {
            return false
        }

        if i == 0 {
            definedRelation = relation
        }

        if definedRelation != relation {
            return false
        }
    }

    return true
}

func IsValidReportWithOneLvlTolerance(report []int, validRange [2]float64) bool {
    var definedRelation LevelRelation
    badLevelFound := false
    badLevelPairIdx := 0

    for i := 0; i < len(report) - 1; i++ {
        result, relation := IsDecOrIncWithinRange(report[i:i+2], validRange)

        if !result {
            badLevelFound = true
        }

        if i == 0 {
            definedRelation = relation
        }

        if definedRelation != relation {
            badLevelFound = true
        }

        if badLevelFound {
            badLevelPairIdx = i
            break
        }
    }

    if badLevelFound {
        if !IsValidReport(excludeIdxFromSlice(report, badLevelPairIdx), validRange) &&
        !IsValidReport(excludeIdxFromSlice(report, badLevelPairIdx+1), validRange) {
            if badLevelPairIdx == 1 {
                return IsValidReport(excludeIdxFromSlice(report, 0), validRange)
            }
            return false
        }
    }

    return true
}

func excludeIdxFromSlice(slice []int, idx int) []int {
    var outputSlice []int

    for i, num := range(slice) {
        if i != idx {
            outputSlice = append(outputSlice, num)
        }
    }

    return outputSlice
}

func extractReportsFromInput(input string) [][]int {
    file, err := os.Open(input)
    var reports [][]int

    if err != nil {
        return nil
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        reports = append(reports, utils.StringToIntSliceBySep(scanner.Text(), " "))
    }

    return reports
}

func SolvePartOne(inputPath string) int {
    validityRange := [2]float64{1, 3}
    reports := extractReportsFromInput(inputPath)

    totalValidReports := 0
    for _, r := range reports {
        if IsValidReport(r, validityRange) {
            totalValidReports++
        }
    }

    return totalValidReports
}

func SolvePartTwo(inputPath string) int {
    validityRange := [2]float64{1, 3}
    reports := extractReportsFromInput(inputPath)

    totalValidReports := 0
    for _, r := range reports {
        if IsValidReportWithOneLvlTolerance(r, validityRange) {
            totalValidReports++
        }
    }

    return totalValidReports
}
