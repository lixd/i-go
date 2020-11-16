package tire

import (
	"fmt"
	"testing"
)

func TestTire_Search(t *testing.T) {
	// var words=[]string{"Hello","Hi","Helm","Go","Golang","Gopher"}
	var words = []string{"三余无梦生", "三非焉罪,无梦至胜", "三非焉罪", "一线生", "一线生机"}
	tire := NewTire()
	tire.InsertMany(words)
	fmt.Println(tire.Search("三余无梦生"))
	fmt.Println(tire.Search("三非"))
}
