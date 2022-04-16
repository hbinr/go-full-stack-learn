package err

import (
	"fmt"
	"os"
	"testing"

	"github.com/pkg/errors"
)

func TestFileNotFound(t *testing.T) {
	_, err := os.Open("/xx/github/go-demo/README.md")
	newError := errors.Wrap(err, "read file failed")
	fmt.Printf("newError: %+v\n", newError)
}
