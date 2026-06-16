package cli_test  
  
import (  
	"os"  
	"testing"  
  
	"github.com/thaohuynh14zc/urfave-cli/cli"  
)  
  
func TestStringSliceEnvVarOverridesDefault(t *testing.T) {  
	os.Setenv("TEST_SLICE_ENV", "env1,env2")  
	defer os.Unsetenv("TEST_SLICE_ENV")  
  
	app := &cli.App{  
		Name: "test",  
		Flags: []cli.Flag{  
			&cli.StringSliceFlag{  
				Name: "test-slice",  
				EnvVars: []string{"TEST_SLICE_ENV"},  
				Default: []string{"default1", "default2"},  
			},  
		},  
		Action: func(c *cli.Context) error {  
			vals := c.StringSlice("test-slice")  
			if len(vals) != 2 {  
				t.Errorf("Expected 2 values, got %%d", len(vals))  
			}  
			if vals[0] != "env1" {  
				t.Errorf("Expected vals[0]=env1, got %%s", vals[0])  
			}  
			if vals[1] != "env2" {  
				t.Errorf("Expected vals[1]=env2, got %%s", vals[1])  
			}  
			return nil  
		},  
	}  
  
	if err := app.Run([]string{"app"}); err != nil {  
		t.Fatalf("App run failed: %%v", err)  
	}  
} 
