# BENZHI_README

## 项目说明

- 项目：11DingKing/rider-rights-task-06
- 项目用途：RiderGuard 面向新就业群体权益保护检察服务站。骑手或检察联络员登记权益诉求并关联证据后，系统依据生效的分派规则确定承办机构与协作机构，支持法律咨询、心理辅导、劳动报酬、职业安全和平台争议等诉求的受理、转办、会商、升级、回访与审计。超期诉求自动升级，结案结论不可被后续规则覆盖，重复线索按幂等键合并。
- Go 工具链：`golang:1.26`
- 前端工具链：无

## 标准构建、运行和测试命令

进入容器后执行：

```bash
# 编译
cd '/app' && GOTOOLCHAIN=local go build ./...

# 启动
cd '/app' && GOTOOLCHAIN=local go run ./cmd/riderctl
cd '/app' && GOTOOLCHAIN=local go run ./cmd/riderguard

# 测试
cd '/app' && GOTOOLCHAIN=local go test ./...
```

## Docker 构建和进入容器

```bash
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh benzhi-task-7-amd64 linux/amd64
./build_benzhi_docker.sh benzhi-task-7-arm64 linux/arm64
docker run -it benzhi-task-7-amd64:latest
docker run -it --platform linux/arm64 benzhi-task-7-arm64:latest
```

## 题目验证命令

1. 预期退出码 0：`go test ./internal/dispatch -run "^TestActiveDefaultRuleFallback$" -count=1`
2. 预期退出码 0：`go test -buildvcs=false -count=1 ./...`
3. 预期退出码 0：`GOTOOLCHAIN=local go build -buildvcs=false ./... && GOTOOLCHAIN=local go vet ./...`

## Bug 复现

Bug 现象、触发步骤和完整错误信息见 `BUG_REPRO.md`。
