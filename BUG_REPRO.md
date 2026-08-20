# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

有些诉求既没有命中类别，也没有命中关键词，按服务站制度应由当前生效的默认规则接收，接口却返回没有匹配规则。请修复默认规则的兜底流程，并保证专项规则命中时它不会抢先生效；本题只改实现，不要修改测试文件。

## 含 Bug 版本

- 仓库：11DingKing/rider-rights-task-06
- 仓库地址：https://github.com/11DingKing/rider-rights-task-06.git
- parent SHA：2c42f72e51c6f7fa2bcc0a73ae05d13f3bca6c01

## 复现步骤

```bash
git clone -- https://github.com/11DingKing/rider-rights-task-06.git bug-repro
cd bug-repro
git checkout --detach 2c42f72e51c6f7fa2bcc0a73ae05d13f3bca6c01
go test ./internal/dispatch -run "^TestActiveDefaultRuleFallback$" -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/dispatch -run "^TestActiveDefaultRuleFallback$" -count=1
--- FAIL: TestActiveDefaultRuleFallback (0.01s)
    task06_test.go:20: default rule was not used: ref=<nil> err=no matching dispatch rule for item case-06: no matching dispatch rule
FAIL
FAIL	riderguard/internal/dispatch	0.053s
FAIL

```

stderr：

```text
(empty)
```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/dispatch -run "^TestActiveDefaultRuleFallback$" -count=1
--- FAIL: TestActiveDefaultRuleFallback (0.00s)
    task06_test.go:20: default rule was not used: ref=<nil> err=no matching dispatch rule for item case-06: no matching dispatch rule
FAIL
FAIL	riderguard/internal/dispatch	0.005s
FAIL

```

stderr：

```text
(empty)
```

## 通过条件

修复后，题面中的触发条件应得到正确业务结果，原始异常不再出现；定向验证、相关包测试和仓库全量回归测试必须通过，不得通过删除或跳过测试、降低断言强度或绕过目标逻辑使验证转绿。
