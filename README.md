# go-demo
go restful demo

一、安装步骤：
1、安装hombrew
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

2、安装启动PostgreSQL
brew install postgresql
brew services start postgresql

3、建立数据库相关
建立database，执行sql脚本（文件init.sql，—U 参数为当前用户）
createdb user_relas_data -U user -E UTF8
psql -d user_relas_data -f init.sql

4、安装golang
brew install go

5、安装go依赖代码
go get -u github.com/gorilla/mux
go get -u github.com/go-pg/pg
go get -u ithub.com/kylelemons/go-gypsy

6、运行代码
（1）、将压缩包中src目录下代码移到到go env指令结果显示的GOPATH/src路径下。
（2）、修改配置文件GOPATH/src/new-task/config.yml，将sqlUser改为当前用户
（3）、在此（GOPATH/src）路径下，执行命令“go run main.go”


二、验证：
单元测试代码在test文件夹内，可进入该文件夹，执行go test -v进行测试。
交叉执行下面的指令进行四种请求的验证。
curl -XGET "http://localhost:8080/users"
curl -XPOST -d '{"name":"Alice"}' "http://localhost:8080/users"
curl -XGET "http://localhost:8080/users/1/relationships"
curl -XPUT -d '{"state":"liked"}' "http://localhost:8080/users/1/relationships/2"
