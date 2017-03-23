package fileline

func formatLine(line string) string {
	n := len(line)
	if n >= 1 && line[n-1] == '\r' {
		return line[0 : n-1]
	} else if n >= 1 && line[n-1] == '\n' {
		line = line[0 : n-1]
	}
	if n >= 2 && line[n-2] == '\r' {
		line = line[0 : n-2]
	}
	return line
}

func extractEOL(line string) string {
	n := len(line)
	if n >= 1 && line[n-1] == '\r' {
		return "\r"
	} else if n < 1 || line[n-1] != '\n' {
		return ""
	}
	if n >= 2 && line[n-2] == '\r' {
		return "\r\n"
	} else {
		return "\n"
	}
}

func addLinesEOL(lines []string, eol string) []string {
	var res []string
	for _, line := range lines {
		if line == "" {
			line = eol
		} else {
			c := line[len(line)-1]
			if c != '\r' && c != '\n' {
				line += eol
			}
		}
		res = append(res, line)
	}
	return res
}
func concat(slices ...[]string) []string {
	var res []string
	for _, s := range slices {
		res = append(res, s...)
	}
	return res
}
