package others

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestHyper(t *testing.T) {
	hyperLL := NewHyperLL()
	isExits := hyperLL.PFAdd("first")
	logrus.Info(isExits)
	isExits = hyperLL.PFAdd("first")
	logrus.Info(isExits)
}
