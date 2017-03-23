package fileline

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type File []string

// ScanLines is a split function for a Scanner that returns each line of
// text, stripped of any trailing end-of-line marker. The returned line may
// be empty. The end-of-line marker is one optional carriage return followed
// by one mandatory newline. In regular expression notation, it is `\r?\n`.
// The last non-empty line of input will be returned even if it has no
// newline.
func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, data[0 : i+1], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

func (f *File) String() string {
	return strings.Join(*f, "")
}

func (f *File) Bytes() []byte {
	return []byte(strings.Join(*f, ""))
}

func ReadFile(r io.Reader) (File, error) {
	var res []string
	scan := bufio.NewScanner(r)
	scan.Split(ScanLines)
	for scan.Scan() {
		res = append(res, scan.Text())
	}
	return res, scan.Err()
}

func ReadBytes(b []byte) File {
	f, _ := ReadFile(bytes.NewReader(b))
	return f
}

func ReadString(s string) File {
	f, _ := ReadFile(bytes.NewReader([]byte(s)))
	return f
}
