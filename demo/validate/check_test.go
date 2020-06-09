package validate

import (
	"fmt"
	"testing"
)

func TestURL(t *testing.T) {
	fmt.Println(IsNumbers("https://studygolang.com/topics/8696/comment/27119"))
	fmt.Println(IsNumbers("1234567890"))
}

func TestParseItemId(t *testing.T) {
	URL := "https: //detail.tmall.com/item.htm?spm=a219t.11817213.0.d7ec8dedc.5f9a6a15bbbH8B&id=614391759359&scm=1007.15348.109552.0_28030&pvid=0e68e24d-89be-43b0-bfff-0f611e14c788&app_pvid=59590_11.89.150.229_530_1591670505710&ptl=floorId:28030;originalFloorId:28030;pvid:0e68e24d-89be-43b0-bfff-0f611e14c788;app_pvid:59590_11.89.150.229_530_1591670505710&union_lens=lensId:OPT@1591670505@0e68e24d-89be-43b0-bfff-0f611e14c788_614391759359@1"
	fmt.Println(ParseItemId(URL))
}
