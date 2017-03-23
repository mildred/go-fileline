package fileline

import (
	"regexp"
)

type Modifications []Modification

type Modification struct {
	SelectRegion *SelectRegion
	DeleteLines  []DeleteLine
	InsertLines  []InsertLine
	ReplaceLines []ReplaceLine
}

type SelectRegion struct {
	Start         *regexp.Regexp
	End           *regexp.Regexp
	StartIncluded bool
	EndIncluded   bool
}

type DeleteLine struct {
	Pattern *regexp.Regexp
}

type InsertLine struct {
	Lines   []string
	Pattern *regexp.Regexp
	Before  bool
	Last    bool
}

type ReplaceLine struct {
	Pattern     *regexp.Regexp
	Replacement string
	Last        bool
	All         bool
}
