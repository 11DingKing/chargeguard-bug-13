# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

查看站点最近十次巡检后，接口为了补充“当前状态”会改掉历史记录的最后一项，之后导出的台账也带着同样错误。请修复展示结果对已保存历史的污染，追加和排序不能反过来改变持久化数据。相关切片测试代码不得调整，也不能删去跨请求复查断言。

## 含 Bug 版本

- 仓库：11DingKing/chargeguard-bug-13
- 仓库地址：https://github.com/11DingKing/chargeguard-bug-13.git
- parent SHA：f095452f9fc9bd491f7c2e607b52ca44a9bd65de

## 复现步骤

```bash
git clone -- https://github.com/11DingKing/chargeguard-bug-13.git bug-repro
cd bug-repro
git checkout --detach f095452f9fc9bd491f7c2e607b52ca44a9bd65de
go test ./internal/httpapi -run TestTaskBehavior -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/httpapi -run TestTaskBehavior -count=1
--- FAIL: TestTaskBehavior (0.01s)
    task_behavior_test.go:16: history=["checked","rectified","closed"]
FAIL
FAIL	chargeguard/internal/httpapi	0.070s
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
$ go test ./internal/httpapi -run TestTaskBehavior -count=1
--- FAIL: TestTaskBehavior (0.00s)
    task_behavior_test.go:16: history=["checked","rectified","closed"]
FAIL
FAIL	chargeguard/internal/httpapi	0.003s
FAIL

```

stderr：

```text
(empty)
```

## 通过条件

修复后，在题面描述的触发条件下应得到预期业务结果且不再出现原始症状；定向验证命令修复前必须失败、应用修复后必须通过，相关回归和仓库全量测试必须通过；不得新增、删除或修改测试文件，不得跳过测试、降低断言或绕过目标逻辑。
