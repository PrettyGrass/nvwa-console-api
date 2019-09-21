module github.com/PrettyGrass/nvwa-console-api

go 1.12

replace gitee.com/y0/nvwa-sys => ../nvwa-sys

require (
	gitee.com/y0/nvwa-sys v0.0.0-20190920095048-14064d802a5c
	github.com/sirupsen/logrus v1.4.1
	github.com/wwcd/grpc-lb v0.0.0-20190626102234-3c50b6a555ae

	google.golang.org/grpc v1.21.1
)
