把需要过滤的文本放到txt.txt中

一行一个

之后运行 会生成 rs.txt



编译到mac



CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64  go build main.go