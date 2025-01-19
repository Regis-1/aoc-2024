package day4

import "testing"

func TestLoadXmasWordMap(t *testing.T) {
	gotFiles := [...]string{
		"../data/day4/invalid.txt",
		"../data/day4/test_input.txt",
	}

	expectedResults := []struct {
		out        []string
		errPresent bool
	}{
		{nil, true},
		{[]string{
			"MMMSXXMASM",
			"MSAMXMSMSA",
			"AMXSXMAAMM",
			"MSAMASMSMX",
			"XMASAMXAMM",
			"XXAMMXXAMA",
			"SMSMSASXSS",
			"SAXAMASAAA",
			"MAMMMXMMMM",
			"MXMXAXMASX"},
			false,
		},
	}

	for i, file := range gotFiles {
		content, err := LoadXmasWordMap(file)

		if (err != nil) != expectedResults[i].errPresent {
			t.Errorf("Wrong error output. File: %v", file)
		}

		if len(content) != len(expectedResults[i].out) {
			t.Errorf("Wrong number of lines when reading file: %v", file)
		}

		for j, line := range content {
			if line != expectedResults[i].out[j] {
				t.Errorf("Wrong line: %v. Should be: %v", line, expectedResults[i].out[j])
			}
		}
	}
}

func TestCountXmasOccurrences(t *testing.T) {
	gotMaps := [][]string{
		[]string{},
		[]string{
			"MXMASMSAMX",
		},
		[]string{
			"MMMSXXMASM",
			"MSAMXMSMSA",
			"AMXSXMAAMM",
			"MSAMASMSMX",
			"XMASAMXAMM",
			"XXAMMXXAMA",
			"SMSMSASXSS",
			"SAXAMASAAA",
			"MAMMMXMMMM",
			"MXMXAXMASX",
		},
	}

	expectedResults := []uint32{0, 2, 18}

	for i, m := range gotMaps {
		result := CountXmasOccurrences(&m)
		if result != expectedResults[i] {
			t.Errorf("Wrong occurrences count: %v. Should be %v.", result, expectedResults[i])
		}
	}
}

func TestCountCrossMasOccurrences(t *testing.T) {
	gotMap := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}

	var expectedResult uint32 = 9

	result := CountCrossMasOccurrences(&gotMap)
	if result != expectedResult {
		t.Errorf("Wrong occurrences count: %v. Should be %v.", result, expectedResult)
	}
}
