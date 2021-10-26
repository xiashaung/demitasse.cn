# 运行

## 下载依赖包

    go mod download

## 更新数据库文件

    需要安装 https://github.com/smallnest/gen

    gen --sqltype=mysql \
    --connstr "cx_rw:cx_rw@2020@@tcp(172.17.19.131:3306)/yx" \
    --database yx  \
    --json \
    --gorm \
    --out ./example \
    --module github.com/xiashauang/demitasse.cn \
    --overwrite


### 配置文件在config目录下 
