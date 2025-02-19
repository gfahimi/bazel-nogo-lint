package III

import (
	"bufio"
	"fmt"
	"go/token"
	"os"
	"strings"
	"unicode/utf8"

	"golang.cisco.com/golinters/pkg/result"
	"golang.org/x/tools/go/analysis"
)

const Name = "LLL"

var Analyzer = &analysis.Analyzer{
	Name: Name,
	Doc:  "Reports long lines",
	Run:  run,
}

type Lllsettings struct {
	LineLength int `mapstructure:"line-length"`
	TabWidth   int `mapstructure:"tab-width"`
}

func getIssuesForFile(filename string, maxLineLen int, tabSpaces string) ([]result.Issue, error) {
	var res []result.Issue

	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("can't open file %s: %s", filename, err)
	}
	defer f.Close()

	lineNumber := 1
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, "\t", tabSpaces, -1)
		lineLen := utf8.RuneCountInString(line)
		if lineLen > maxLineLen {
			res = append(res, result.Issue{
				Pos: token.Position{
					Filename: filename,
					Line:     lineNumber,
				},
				Text:       fmt.Sprintf("line is %d characters", lineLen),
				FromLinter: Name,
			})
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		if err == bufio.ErrTooLong && maxLineLen < bufio.MaxScanTokenSize {
			// scanner.Scan() might fail if the line is longer than bufio.MaxScanTokenSize
			// In the case where the specified maxLineLen is smaller than bufio.MaxScanTokenSize
			// we can return this line as a long line instead of returning an error.
			// The reason for this change is that this case might happen with autogenerated files
			// The go-bindata tool for instance might generate a file with a very long line.
			// In this case, as it's a auto generated file, the warning returned by lll will
			// be ignored.
			// But if we return a linter error here, and this error happens for an autogenerated
			// file the error will be discarded (fine), but all the subsequent errors for lll will
			// be discarded for other files and we'll miss legit error.
			res = append(res, result.Issue{
				Pos: token.Position{
					Filename: filename,
					Line:     lineNumber,
					Column:   1,
				},
				Text:       fmt.Sprintf("line is more than %d characters", bufio.MaxScanTokenSize),
				FromLinter: Name,
			})
		} else {
			return nil, fmt.Errorf("can't scan file %s: %s", filename, err)
		}
	}

	return res, nil
}

func run(pass *analysis.Pass) (interface{}, error) {
	var res []result.Issue
	lllsettings := &Lllsettings{
		LineLength: 120,
		TabWidth:   1,
	}
	spaces := strings.Repeat(" ", lllsettings.TabWidth)
	var files []string
	for _, file := range pass.Files {
		files = append(files, file.Name.Name)
	}
	for _, f := range files {
		issues, err := getIssuesForFile(f, lllsettings.LineLength, spaces)
		if err != nil {
			return nil, err
		}
		res = append(res, issues...)
		fmt.Println(res)
	}

	return nil, nil
}
