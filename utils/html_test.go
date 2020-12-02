package utils

import (
	"fmt"
	"testing"
)

const WithHtml = `<p>买了个正版的刺客信条英灵殿，但是一进入游戏3分钟左右就闪退，无论在线还是离线模式，而且各种贴图错误。CPU、显卡性能是没有问题的，1080ti。 然后我下了个破解版的刺客信条奥德赛就能正常玩，太讽刺了。</p><p>网上找了很多方法都不行，包括重装1903的系统，重装稳定版的显卡驱动等等。。。求靠谱的解决方案，百度复制粘贴的就免了。</p>`

func TestRemoveHtml(t *testing.T) {
	removeHtml := TrimHtml(WithHtml)
	fmt.Println(removeHtml)
}

func BenchmarkTrimHtml(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = TrimHtml(WithHtml)
	}
}
