## 阿里云docker镜像仓库工具

```
brew tap ysicing/tap
brew install crtools
```

## 使用

> 示例: ns命名空间

```
crtools ns
+---+------------+----------+-------+
|   | 区域       | 命名空间 | 权限  |
+---+------------+----------+-------+
| 0 | cn-beijing | k7scn    | ADMIN |
+---+------------+----------+-------+
```

> 示例: repo

```
crtools repo --ns k7scn --region cn-beijing --tail 10
+---+------------+--------------------------------------------------------------------------+---------------------+
|   | 区域       | 镜像                                                                     | 更新时间            |
+---+------------+--------------------------------------------------------------------------+---------------------+
| 0 | cn-beijing | registry.cn-beijing.aliyuncs.com/k7scn/coredns:1.6.7                     | 2020-08-14 14:56:05 |
| 1 | cn-beijing | registry.cn-beijing.aliyuncs.com/k7scn/etcd:3.4.3-0                      | 2020-08-14 14:56:02 |
| 2 | cn-beijing | registry.cn-beijing.aliyuncs.com/k7scn/pause:3.2                         | 2020-08-14 14:56:01 |
| 3 | cn-beijing | registry.cn-beijing.aliyuncs.com/k7scn/kube-proxy:v1.18.8                | 2020-08-14 14:55:58 |
| 4 | cn-beijing | registry.cn-beijing.aliyuncs.com/k7scn/kube-scheduler:v1.18.8            | 2020-08-14 14:55:44 |
| 5 | cn-beijing | registry.cn-beijing.aliyuncs.com/k7scn/kube-controller-manager:v1.18.8   | 2020-08-14 14:55:35 |
| 6 | cn-beijing | registry.cn-beijing.aliyuncs.com/k7scn/kube-apiserver:v1.18.8            | 2020-08-14 14:55:22 |
| 7 | cn-beijing | registry.cn-beijing.aliyuncs.com/k7scn/calico-kube-controllers:v3.14.0   | 2020-08-12 21:09:40 |
| 8 | cn-beijing | registry.cn-beijing.aliyuncs.com/k7scn/calico-node:v3.14.0               | 2020-08-12 21:09:14 |
| 9 | cn-beijing | registry.cn-beijing.aliyuncs.com/k7scn/calico-pod2daemon-flexvol:v3.14.0 | 2020-08-12 21:08:34 |
+---+------------+--------------------------------------------------------------------------+---------------------+
```

> 示例： tags

```
crtools tags --ns k7scn --repo coredns --tail 10
+---+------------------------------------------------+-------+----------+---------------------+
|   | 镜像                                           | 版本  | 镜像大小 | 更新时间            |
+---+------------------------------------------------+-------+----------+---------------------+
| 0 | registry.cn-beijing.aliyuncs.com/k7scn/coredns | 1.6.7 | 13594230 | 2020-08-14 14:56:05 |
| 1 | registry.cn-beijing.aliyuncs.com/k7scn/coredns | 1.6.5 | 13235675 | 2020-08-03 14:11:17 |
| 2 | registry.cn-beijing.aliyuncs.com/k7scn/coredns | 1.6.2 | 14119735 | 2020-07-16 17:24:00 |
| 3 | registry.cn-beijing.aliyuncs.com/k7scn/coredns | 1.3.1 | 12300647 | 2020-06-27 22:17:49 |
+---+------------------------------------------------+-------+----------+---------------------+

```

> 示例: 搜索

```bash
crtools sn --sn kube-api
+---+------------+----------------------------------------------------------------+---------------------+
|   | 区域       | 镜像                                                           | 镜像更新时间        |
+---+------------+----------------------------------------------------------------+---------------------+
| 0 | cn-beijing | registry.cn-beijing.aliyuncs.com/k7scn/kube-apiserver:v1.18.8  | 2020-08-14 14:55:22 |
| 1 | cn-beijing | registry.cn-beijing.aliyuncs.com/k7scn/kube-apiserver:v1.17.9  | 2020-08-03 14:10:50 |
| 2 | cn-beijing | registry.cn-beijing.aliyuncs.com/k7scn/kube-apiserver:v1.18.6  | 2020-08-03 14:07:52 |
| 3 | cn-beijing | registry.cn-beijing.aliyuncs.com/k7scn/kube-apiserver:v1.16.13 | 2020-07-16 17:23:12 |
| 4 | cn-beijing | registry.cn-beijing.aliyuncs.com/k7scn/kube-apiserver:v1.18.5  | 2020-07-10 15:45:38 |
| 5 | cn-beijing | registry.cn-beijing.aliyuncs.com/k7scn/kube-apiserver:v1.17.8  | 2020-07-07 15:24:05 |
| 6 | cn-beijing | registry.cn-beijing.aliyuncs.com/k7scn/kube-apiserver:v1.16.12 | 2020-07-07 15:22:54 |
| 7 | cn-beijing | registry.cn-beijing.aliyuncs.com/k7scn/kube-apiserver:v1.15.12 | 2020-06-27 22:16:24 |
| 8 | cn-beijing | registry.cn-beijing.aliyuncs.com/k7scn/kube-apiserver:v1.16.11 | 2020-06-25 14:45:39 |
| 9 | cn-beijing | registry.cn-beijing.aliyuncs.com/k7scn/kube-apiserver:v1.18.4  | 2020-06-19 23:24:20 |
+---+------------+----------------------------------------------------------------+---------------------+

crtools sn --sn 1.18.8  
+---+------------+------------------------------------------------------------------------+---------------------+
|   | 区域       | 镜像                                                                   | 镜像更新时间        |
+---+------------+------------------------------------------------------------------------+---------------------+
| 0 | cn-beijing | registry.cn-beijing.aliyuncs.com/k7scn/kube-proxy:v1.18.8              | 2020-08-14 14:55:58 |
| 1 | cn-beijing | registry.cn-beijing.aliyuncs.com/k7scn/kube-scheduler:v1.18.8          | 2020-08-14 14:55:44 |
| 2 | cn-beijing | registry.cn-beijing.aliyuncs.com/k7scn/kube-controller-manager:v1.18.8 | 2020-08-14 14:55:35 |
| 3 | cn-beijing | registry.cn-beijing.aliyuncs.com/k7scn/kube-apiserver:v1.18.8          | 2020-08-14 14:55:22 |
+---+------------+------------------------------------------------------------------------+---------------------+

```

> 示例: upgrade升级crtools

```bash
crtools upgrade 
I0816 08:32:32.655453   83127 upgrade.go:23] ==> Upgrading 1 outdated package:
ysicing/tap/crtools 0.0.3 -> 0.0.6
==> Upgrading ysicing/tap/crtools 0.0.3 -> 0.0.6 
==> Downloading https://github.com/ysicing/crtools/releases/download/0.0.6/crtools_darwin_amd64
Already downloaded: /Users/ysicing/Library/Caches/Homebrew/downloads/5ebd0244579cc37e5247a1a605eef5cdb24cc8a86e66fc49bed42927db7d30e2--crtools_darwin_amd64
🍺  /usr/local/Cellar/crtools/0.0.6: 3 files, 14.6MB, built in 3 seconds
Removing: /usr/local/Cellar/crtools/0.0.3... (3 files, 14.6MB)
Removing: /Users/ysicing/Library/Caches/Homebrew/crtools--0.0.3... (14.6MB)
```

### 版本计划

- [x] 列出命名空间
- [x] 列出镜像
- [x] 列出标签
- [x] 支持搜素
- [ ] 支持删除