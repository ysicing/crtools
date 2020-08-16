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

### 版本计划

- [x] 列出命名空间
- [x] 列出镜像
- [x] 列出标签
- [ ] 支持搜素
- [ ] 支持删除