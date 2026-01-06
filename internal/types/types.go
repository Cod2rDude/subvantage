package types

// Public Types
type RunMode int

// Public Constants
const (
	ModeNone RunMode = iota
	ModeSearch
	ModeBruteForce
	ModeCombined
)

// Public Structs
type Options struct {
	Domain     string
	OutputFile string
	Mode       RunMode
}

type CRTResult struct {
	NameValue string `json:"name_value"`
}
