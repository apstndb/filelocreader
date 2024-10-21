package filelocreader

import (
	"fmt"
	"go/token"
	"os"
	"regexp"
	"strconv"
)

type Location struct {
	Line, StartCol, EndCol int
}

type FileLocation struct {
	Filename string
	Loc      Location
}

var re = regexp.MustCompile(`^(?P<filename>[^:].*):(?P<line>[0-9]*):(?P<startCol>[0-9]*)-(?P<endCol>[0-9]*)$`)

func ParseFileLocation(s string) (*FileLocation, error) {
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

	return &FileLocation{
		Filename: filename,
		Loc: Location{
			Line:     line,
			StartCol: startCol,
			EndCol:   endCol,
		},
	}, nil
}

func ExtractFileLocation(floc *FileLocation) ([]byte, error) {
	if floc == nil {
		return nil, fmt.Errorf("floc is nil")
	}
	b, err := os.ReadFile(floc.Filename)
	if err != nil {
		return nil, err
	}

	return ExtractLocation(b, floc)
}

func ExtractLocation(b []byte, floc *FileLocation) ([]byte, error) {
	var fs token.FileSet
	file := fs.AddFile(floc.Filename, 0, len(b))

	for i, r := range string(b) {
		if r == '\n' {
			file.AddLine(i)
		}
	}

	linePos := int(file.LineStart(floc.Loc.Line))
	return b[linePos+floc.Loc.StartCol : linePos+floc.Loc.EndCol], nil
}
