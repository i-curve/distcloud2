# clouddist网盘

## 简介

是一个前后端分离的个人网盘存储系统，后端用go编写，前端用vue3编写。这里是后端部分。前段部分参考项目[clouddist前端](https://github.com/i-curve/distcloud.git)

## 内容

<!-- - 用户管理
- 日志显示
- 文件存储下载 -->
包含用户上传文件,下载文件

## 安装

clone it

```bash
# 1. 克隆本项项目
git clone https://github.com/i-curve/distcloud2.git

# 2. 进入conf目录修改配置文件
cd distcloud/conf
# 修改app.ini文件
# 需要修改的有 datavase数据库用户信息
# 如果是调试的话把server里面的runmode改为debug

# 3. 使用mysql目录下的table.sql创建数据库及数据表

# 4. 回到项目根目录用go构建主程序
go build main.go

# 5. 运行程序
./clouddist
```

## License
...