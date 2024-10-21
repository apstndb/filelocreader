package filelocreader

import (
	"fmt"
	"go/token"
	"os"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`^(?P<filename>[^:].*):(?P<line>[0-9]*):(?P<startCol>[0-9]*)-(?P<endCol>[0-9]*)$`)

func ParseLoc(s string) (*Location, error) {
	matches := re.FindStringSubmatch(s)
	filename := matches[re.SubexpIndex("filename")]

	line, err := strconv.Atoi(matches[re.SubexpIndex("line")])
	if err != nil {
		return nil, err
	}

	startCol, err := strconv.Atoi(matches[re.SubexpIndex("startCol")])
	if err != nil {
		return nil, err
	}

	endCol, err := strconv.Atoi(matches[re.SubexpIndex("endCol")])
	if err != nil {
		return nil, err
	}

	return &Location{
		Filename: filename,
		Line:     line,
		StartCol: startCol,
		EndCol:   endCol,
	}, nil
}

type Location struct {
	Filename               string
	Line, StartCol, EndCol int
}

func ExtractLoc(loc *Location) ([]byte, error) {
	if loc == nil {
		return nil, fmt.Errorf("loc is nil")
	}
	b, err := os.ReadFile(loc.Filename)
	if err != nil {
		return nil, err
	}

	var fs token.FileSet
	file := fs.AddFile(loc.Filename, 0, len(b))

	for i, r := range string(b) {
		if r == '\n' {
			file.AddLine(i)
		}
	}

	linePos := int(file.LineStart(loc.Line))
	return b[linePos+loc.StartCol : linePos+loc.EndCol], nil
}
