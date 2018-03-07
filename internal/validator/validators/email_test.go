package validators_test

import (
	"testing"

	"github.com/friasdesign/xii-simposio-infra/internal/test"
	"github.com/friasdesign/xii-simposio-infra/internal/validator/validators"
)

func TestEmail_ValidatesCorrectly(t *testing.T) {
	fix := []test.Case{
		test.Case{
			Input:    "pepe_pepe@pepe.com",
			Expected: true,
		},
		test.Case{
			Input:    "pepe_@pepe.com",
			Expected: true,
		},
		test.Case{
			Input:    "",
			Expected: false,
		},
		test.Case{
			Input:    "asd",
			Expected: false,
		},
		test.Case{
			Input:    "mario\\marielo@pepe.com",
			Expected: false,
		},
	}

	for _, val := range fix {
		s, ok := val.Input.(string)
		if ok == false {
			t.Fatal("Invalid type for Test Case", val.Input)
		}
		o := validators.Email(s)
		if o != val.Expected {
			t.Fatal("Wrong validation!", s)
		}
	}
}
