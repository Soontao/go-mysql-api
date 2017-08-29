package swagger

import (
	"fmt"
	"regexp"
	"testing"
)

func Test_getEnumIfItIs(t *testing.T) {
	re := regexp.MustCompile("\\'([\\w]+)\\'")
	enum := re.FindAllString("enum('circulated','reserved','finished')", -1)
	fmt.Println(len(enum))
	fmt.Println(enum)
}
