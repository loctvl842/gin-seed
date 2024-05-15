package gosdk

import (
	"flag"
	"fmt"
	"reflect"
	"strings"

	"github.com/facebookgo/flagenv"
)

// isZeroValue guesses whether the string represents the zero
// value for a flag. It is not accurate but in practice works OK.
func isZeroValue(f *flag.Flag, value string) bool {
	// Build a zero value of the flag's Value type, and see if the
	// result of calling its String method equals the value passed in.
	// This works unless the Value type is itself an interface type.
	typ := reflect.TypeOf(f.Value)
	var z reflect.Value
	if typ.Kind() == reflect.Ptr {
		z = reflect.New(typ.Elem())
	} else {
		z = reflect.Zero(typ)
	}
	if value == z.Interface().(flag.Value).String() {
		return true
	}

	switch value {
	case "false":
		return true
	case "":
		return true
	case "0":
		return true
	}
	return false
}

func getEnvName(name string) string {
	name = strings.ReplaceAll(name, ".", "_")
	name = strings.ReplaceAll(name, "-", "_")
	if flagenv.Prefix != "" {
		name = flagenv.Prefix + name
	}
	return strings.ToUpper(name)
}

type AppFlagSet struct {
	*flag.FlagSet
}

func newFlagSet(_ string, fs *flag.FlagSet) *AppFlagSet {
	fSet := &AppFlagSet{fs}
	return fSet
}

func (f *AppFlagSet) Parse(args []string) {
	flagenv.Parse()
	f.FlagSet.Parse(args)
}

func (f *AppFlagSet) GetSampleEnvs() {
	f.VisitAll(func(f *flag.Flag) {
		s := fmt.Sprintf("# %s (-%s)\n", f.Usage, f.Name)
		s += fmt.Sprintf("%s=", getEnvName(f.Name))

		if !isZeroValue(f, f.DefValue) {
			t := fmt.Sprintf("%T", f.Value)
			if t == "*flag.stringValue" {
				// put quotes on the value
				s += fmt.Sprintf("%q", f.DefValue)
			} else {
				s += fmt.Sprintf("%v", f.DefValue)
			}
		}
		fmt.Print(s, "\n\n")
	})
}
