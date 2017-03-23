package fileline

func (f *File) Modify(mods Modifications) {
	for _, mod := range mods {
		f.ModifySingle(mod)
	}
}

func (f *File) ModifySingle(mod Modification) {
	var before, region, after []string
	region = *f
	if mod.SelectRegion != nil {
		before, region, after = f.ExtractRegion(*mod.SelectRegion)
	}

	reg := File(region)

	for _, dl := range mod.DeleteLines {
		reg.DeleteLines(dl)
	}

	for _, el := range mod.InsertLines {
		reg.InsertLines(el)
	}

	for _, rl := range mod.ReplaceLines {
		reg.ReplaceLines(rl)
	}

	*f = concat(before, reg, after)
}

func (f *File) DeleteLines(dl DeleteLine) {
	var res File
	for _, l := range *f {
		if !dl.Pattern.MatchString(formatLine(l)) {
			res = append(res, l)
		}
	}
	*f = res
}

func (f *File) InsertLines(il InsertLine) {
	var res File
	var n int
	var eol string
	for i, line := range *f {
		if il.Pattern.MatchString(formatLine(line)) {
			eol = extractEOL((*f)[i])
			n = i
			if !il.Last {
				break
			}
		}
	}
	if n < len(*f) {
		if !il.Before {
			res = append(res, (*f)[:n+1]...)
		} else if n > 0 {
			res = append(res, (*f)[:n]...)
		}
		res = append(res, addLinesEOL(il.Lines, eol)...)
		if il.Before {
			res = append(res, (*f)[n:]...)
		} else if n+1 < len(*f) {
			res = append(res, (*f)[n+1:]...)
		}
		*f = res
	}
}

func (f *File) ReplaceLines(rl ReplaceLine) {
	var res File
	var n int
	var eol string
	for i, line := range *f {
		eol = extractEOL(line)
		if rl.All {
			res = append(res, rl.Pattern.ReplaceAllString(formatLine(line), rl.Replacement)+eol)
		} else if rl.Pattern.MatchString(formatLine(line)) {
			n = i
			if !rl.Last {
				break
			}
		}
	}
	if rl.All {
		*f = res
		return
	} else if n < len(*f) {
		(*f)[n] = rl.Pattern.ReplaceAllString(formatLine((*f)[n]), rl.Replacement) + eol
	}
}

func (f *File) ExtractRegion(reg SelectRegion) ([]string, []string, []string) {
	var before, region, after []string
	var i int
	for i < len(*f) {
		line := (*f)[i]
		i += 1
		if reg.Start.MatchString(formatLine(line)) {
			if reg.StartIncluded {
				region = append(region, line)
			} else {
				before = append(before, line)
			}
			break
		}
		before = append(before, line)
	}
	for i < len(*f) {
		line := (*f)[i]
		i += 1
		if reg.End.MatchString(formatLine(line)) {
			if reg.EndIncluded {
				region = append(region, line)
			} else {
				after = append(after, line)
			}
			break
		}
		region = append(region, line)
	}
	if i < len(*f) {
		after = append(after, (*f)[i:]...)
	}
	return before, region, after
}
