# shahaohuo
啥好货是一个小众产品分享平台，使用以下语言工具等开发：
- golang
- mysql
- s3
- k3s
- github action

## 如何使用
修改 `dev` or `prod` 配置，指定 mysql 和 s3 相关地址，使用你喜欢的开发工具进行调试。

### docker 
项目支持多步构建，直接 build 容器镜像即可。详细信息请查看 `dockerfile` 和 `workflows` 的 CI / CD 流程。

# 服务端
## 后端集群
https://k3s.io/

## 其他
### 主题颜色
https://colorhunt.co/**palette/179398
https://colorhunt.co/palette/181112
### 相关图片版权
https://pixabay.com/zh/photos/mars-mars-rover-space-travel-robot-67522/

### 相关BUG
当下 contained 拉取 github 镜像存在问题，更改为阿里云
