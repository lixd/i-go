package algorithm

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestHyper(t *testing.T) {
	hyperLL := NewHyperLL()
	isExits := hyperLL.PFAdd("first")
	logrus.Info(isExits)
	isExits = hyperLL.PFAdd("first")
	logrus.Info(isExits)
}
