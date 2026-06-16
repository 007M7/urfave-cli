package main  
  
import (  
	"fmt"  
	"os"  
	"github.com/thaohuynh14zc/urfave-cli/cli"  
)  
  
func main() {  
	os.Setenv("TEST_SLICE", "env1,env2")  
	defer os.Unsetenv("TEST_SLICE")  
	app := &cli.App{  
		Name: "test",  
		Flags: []cli.Flag{  
			&cli.StringSliceFlag{  
				Name: "test-slice",  
				EnvVars: []string{"TEST_SLICE"},  
				Default: []string{"default1", "default2"},  
			},  
		},  
		Action: func(c *cli.Context) error {  
			vals := c.StringSlice("test-slice")  
			fmt.Printf("Slice values: %%v", vals)  
			fmt.Println()  
			if len(vals) == 2 && vals[0] == "env1" && vals[1] == "env2" {  
				fmt.Println("PASS: Env var overrides default")  
			} else {  
				fmt.Printf("FAIL: Expected [env1 env2] but got %%v", vals)  
			}  
			return nil  
		},  
	}  
	app.Run(os.Args)  
} 
