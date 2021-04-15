package search

// es 文档类型
const (
	TypeDoc = "_doc"
)

// Index ES索引结构 index+type
type Index struct {
	Index string // 索引名
	Type  string // 文档类型 新版固定为`doc`
}

// ESite es中网址的具体结构
type ESite struct {
	ID       string   `json:"id"`       // _id
	Keywords []string `json:"keywords"` // 关键字
	Host     string   `json:"host"`     // 域名
}
