package cli  
  
import (  
	"flag"  
	"os"  
	"strconv"  
	"strings"  
)  
  
type FlagBase struct {  
	Name string  
	Aliases []string  
	Usage string  
	EnvVars []string  
	HasBeenSet bool  
}  
  
func lookupEnv(envVars []string) (string, bool) {  
	for _, e := range envVars {  
		if v, ok := os.LookupEnv(e); ok {  
			return v, true  
		}  
	}  
	return "", false  
} 
  
type Flag interface {  
	Apply(set *flag.FlagSet) error  
	Names() []string  
} 
