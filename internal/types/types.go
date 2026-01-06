package types

type RunMode int

const (
	ModeNone RunMode = iota
	ModeSearch
	ModeBruteForce
	ModeCombined
)

type Options struct {
	Domain     string
	OutputFile string
	Mode       RunMode
}
