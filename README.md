# i-go
golang learning  

Thanks  
***
[![JetBrains](./assets/jetbrains-variant.svg)](https://www.jetbrains.com/?from=i-go)






## type


feat: A new feature
fix: A bug fix
docs: Documentation only changes
style: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
refactor: A code change that neither fixes a bug nor adds a feature
perf: A code change that improves performance
test: Adding missing or correcting existing tests
chore: Changes to the build process or auxiliary tools and libraries such as documentation generation

build :改变了build工具 如 grunt换成了 npm

revert: 撤销上一次的 commit

scope :用来说明此次修改的影响范围 可以随便填写任何东西，commitizen也给出了几个 如：location 、browser、compile，不过我推荐使用

all ：表示影响面大 ，如修改了网络框架  会对真个程序产生影响

loation： 表示影响小，某个小小的功能

module：表示会影响某个模块 如登录模块、首页模块 、用户管理模块等等

subject: 用来简要描述本次改动，概述就好了

body:具体的修改信息 应该尽量详细

footer：放置写备注啥的，如果是 bug ，可以把bug id放入