package cli  
  
import (  
	"flag"  
	"os"  
	"strings"  
)  
  
type StringSlice []string  
  
func NewStringSlice(defaults ...string) *StringSlice {  
	s := make(StringSlice, len(defaults))  
	copy(s, defaults)  
	return &s  
}  
  
func (s *StringSlice) Set(v string) error {  
	*s = append(*s, v)  
	return nil  
}  
  
func (s *StringSlice) String() string {  
	return strings.Join(*s, ", ")  
} 
  
type StringSliceFlag struct {  
	FlagBase  
	Value *StringSlice  
	Default []string  
} 
  
func (f *StringSliceFlag) Apply(set *flag.FlagSet) error {  
	// THE FIX: Check env vars first, before defaults  
	if envVal, ok := lookupEnv(f.EnvVars); ok {  
		parts := strings.Split(envVal, ",")  
		sv := make(StringSlice, len(parts))  
		for i, p := range parts {  
			sv[i] = strings.TrimSpace(p)  
		}  
		f.Value = &sv  
		f.HasBeenSet = true  
	} else if f.Default != nil {  
		f.Value = NewStringSlice(f.Default...)  
	}  
	if f.Value == nil {  
		f.Value = new(StringSlice)  
	}  
	set.Var(f.Value, f.Name, f.Usage)  
	for _, alias := range f.Aliases {  
		set.Var(f.Value, alias, f.Usage)  
	}  
	return nil  
}  
  
func (f *StringSliceFlag) Names() []string {  
	return append(f.Aliases, f.Name)  
} 
