package main

import (
	_ "devops-super/internal/logic"
	_ "devops-super/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"devops-super/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
