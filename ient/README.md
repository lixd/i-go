当将代码转换为Markdown格式时，可以使用以下格式化方法：

markdown
Copy code
# 基本使用

## 创建表

```bash
ent init --target ent/schema User
ent init --target ent/schema Tag
ent init --target ent/schema TagAssociation
```
在执行上述命令后，会生成 User.go 文件。编辑该文件以添加表结构定义：

```go
package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
    ent.Schema
}

// Fields of the User. 用户字段
func (User) Fields() []ent.Field {
    return []ent.Field{
        // 年龄整数
        field.Int("age").Positive(),
        // 姓名为字符串
        field.String("name").Default("unknown"),
    }
}

// Edges of the User. 用户表关系
func (User) Edges() []ent.Edge {
    return []ent.Edge{}
}
```
定义好字段之后，生成 CRUD 代码：

```bash
ent generate ./ent/schema
```
运行
```go
package main

import (
    "context"
    "log"

    "tag-manager/pkg/ent"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    client, err := ent.Open("mysql", "root:your-sql-password@tcp(172.20.150.246:3306)/test?parseTime=true")
    if err != nil {
        log.Fatalf("failed opening connection to mysql: %v", err)
    }
    defer client.Close()

    // 运行自动迁移工具
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
}
```
以上代码会自动创建/更新相应的表。

其他
指定表名
```go
func (Tag) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "tags"},
    }
}
```
创建索引
```go
// TODO(user): add indexes here.
// func (Tag) Indexes() []ent.Index {
//  return []ent.Index{
//     index.Fields("name").Unique(),
//  }
// }
```
扩展字段
```go
func (Tag) Mixin() []ent.Mixin {
    return []ent.Mixin{
        AuditMixin{},
    }
}
```
具体用法可以参考这个：Entgo 实现 软删除（Soft Delete）

创建 gRPC 接口
具体参考这个：gRPC 接口创建
可以根据 Schema 定义直接生成 gRPC 接口






