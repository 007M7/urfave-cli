package cli  
  
import (  
	"flag"  
	"fmt"  
	"os"  
)  
  
type App struct {  
	Name string  
	Usage string  
	Flags []Flag  
	Action func(c *Context) error  
}  
  
func (a *App) Run(args []string) error {  
	set := flag.NewFlagSet(a.Name, flag.ContinueOnError)  
	for _, f := range a.Flags {  
		if err := f.Apply(set); err != nil {  
			return err  
		}  
	}  
	if err := set.Parse(args[1:]); err != nil {  
		return err  
	}  
	if a.Action != nil {  
		ctx := &Context{App: a, FlagSet: set}  
		return a.Action(ctx)  
	}  
	return nil  
}  
  
type Context struct {  
	App *App  
	FlagSet *flag.FlagSet  
}  
  
func (c *Context) StringSlice(name string) []string {  
	if c.FlagSet == nil {  
		return nil  
	}  
	f := c.FlagSet.Lookup(name)  
	if f == nil {  
		return nil  
	}  
	if sv, ok := f.Value.(*StringSlice); ok {  
		return []string(*sv)  
	}  
	return nil  
} 
