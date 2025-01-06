package day4

import (
    "os"
    "bufio"
    "fmt"
)

type Coords struct {
    Y int
    X int
}

func SolvePartOne(inputFilePath string) uint32 {
    wordMap, err := LoadXmasWordMap(inputFilePath)
    if err != nil {
        fmt.Printf("ERROR: Error while loading file %v:\n%v\n", inputFilePath, err.Error())
        return 0
    } else if len(wordMap) == 0 {
        fmt.Printf("ERROR: Empty word map in file: %v:\n", inputFilePath)
        return 0
    }

    return CountXmasOccurrences(&wordMap)
}

func SolvePartTwo(inputFilePath string) uint32 {
    wordMap, err := LoadXmasWordMap(inputFilePath)
    if err != nil {
        fmt.Printf("ERROR: Error while loading file %v:\n%v\n", inputFilePath, err.Error())
        return 0
    } else if len(wordMap) == 0 {
        fmt.Printf("ERROR: Empty word map in file: %v:\n", inputFilePath)
        return 0
    }

    return CountCrossMasOccurrences(&wordMap)
}

func CountXmasOccurrences(wordMap *[]string) uint32 {
    var xmasOccurrences uint32 = 0
    for i, line := range(*wordMap) {
        for j, c := range(line) {
            if c == 'X' {
                xmasOccurrences += searchForXmases(i, j, wordMap)
            }
        }
    }

    return xmasOccurrences
}

func CountCrossMasOccurrences(wordMap *[]string) uint32 {
    /*
        Algorith "plan":
        1. create 4 masks that will cover all possibilities of X-MASes
        2. iterate through whole world map with those maps and try to
        find all occurrences
    */
}

func searchForXmases(y int, x int, wordMap *[]string) uint32 {
    sizeY := len(*wordMap)
    sizeX := len((*wordMap)[0])

    doAbove := y - 1 >= 0
    doBelow := y + 1 < sizeY
    doLeft := x - 1 >= 0
    doRight := x + 1 < sizeX

    var totalXmases uint32 = 0

    // line above
    if doAbove {
        if doLeft {
            if (*wordMap)[y-1][x-1] == 'M' {
                totalXmases += finishSearch(Coords{y-2, x-2}, Coords{y-3, x-3}, wordMap)
            }
        }
        if (*wordMap)[y-1][x] == 'M' {
            totalXmases += finishSearch(Coords{y-2, x}, Coords{y-3, x}, wordMap)
        }
        if doRight {
            if (*wordMap)[y-1][x+1] == 'M' {
                totalXmases += finishSearch(Coords{y-2, x+2}, Coords{y-3, x+3}, wordMap)
            }
        }
    }
    // current line
    if doLeft {
        if (*wordMap)[y][x-1] == 'M' {
            totalXmases += finishSearch(Coords{y, x-2}, Coords{y, x-3}, wordMap)
        }
    }
    if doRight {
        if (*wordMap)[y][x+1] == 'M' {
            totalXmases += finishSearch(Coords{y, x+2}, Coords{y, x+3}, wordMap)
        }
    }
    // line below
    if doBelow {
        if doLeft {
            if (*wordMap)[y+1][x-1] == 'M' {
                totalXmases += finishSearch(Coords{y+2, x-2}, Coords{y+3, x-3}, wordMap)
            }
        }
        if (*wordMap)[y+1][x] == 'M' {
            totalXmases += finishSearch(Coords{y+2, x}, Coords{y+3, x}, wordMap)
        }
        if doRight {
            if (*wordMap)[y+1][x+1] == 'M' {
                totalXmases += finishSearch(Coords{y+2, x+2}, Coords{y+3, x+3}, wordMap)
            }
        }
    }

    return totalXmases
}

func finishSearch(c1 Coords, c2 Coords, wordMap *[]string) uint32 {
    sizeY := len(*wordMap)
    sizeX := len((*wordMap)[0])

    if c1.Y < 0 || c1.Y >= sizeY || c1.X < 0 || c1.X >= sizeX {
        return 0
    } else if c2.Y < 0 || c2.Y >= sizeY || c2.X < 0 || c2.X >= sizeX {
        return 0
    }

    if (*wordMap)[c1.Y][c1.X] == 'A' && (*wordMap)[c2.Y][c2.X] == 'S' {
        return 1
    }

    return 0
}

func LoadXmasWordMap(filepath string) ([]string, error) {
    fileHandle, err := os.Open(filepath)
    if err != nil {
        return nil, err
    }
    defer fileHandle.Close()

    var wordMap []string
    scanner := bufio.NewScanner(fileHandle)
    for scanner.Scan() {
        wordMap = append(wordMap, scanner.Text())
    }

    return wordMap, nil
}
