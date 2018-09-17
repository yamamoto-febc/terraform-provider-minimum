package state

type SchemaValue struct {
	Default             string
	DefaultFunc         string
	DefaultFuncRequired string
	DiffSuppressFunc    string
	ForceNew            string
	StateFunc           string
	Sensitive           string
	ValidateFunc        string
	ConflictsWith1      string
	ConflictsWith2      string
	Deprecated          string
	Removed             string
}
