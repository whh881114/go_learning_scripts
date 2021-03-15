# Go交叉（跨平台）编译

Golang 支持交叉编译，在一个平台上生成另一个平台的可执行程序。

## Mac下编译Linux和Windows 64位可执行程序
```shell script
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
  ​
  CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

## Linux下编译Mac和Windows 64位可执行程序
```shell script
  CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
  ​
  CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

## Windows下编译Mac和Linux 64位可执行程序
```shell script
  SET CGO_ENABLED=0
  SET GOOS=darwin
  SET GOARCH=amd64
  go build main.go
  ​
  SET CGO_ENABLED=0
  SET GOOS=linux
  SET GOARCH=amd64
  go build main.go
```

## 小结
GOOS：目标平台的操作系统（darwin、freebsd、linux、windows）。
GOARCH：目标平台的体系架构（386、amd64、arm），交叉编译不支持CGO所以要禁用它。