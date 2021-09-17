package jd_cookie

import (
	"github.com/cdle/sillyGirl/core"
	"github.com/cdle/sillyGirl/im"
)

func init() {
	core.AddCommand("jd", []core.Function{
		{
			Rules: []string{"raw ^jd update$"},
			Cron:  "41 * * * *",
			Admin: true,
			Handle: func(s im.Sender) interface{} {
				s.Disappear()
				s.Reply(name + "开始拉取代码。")
				need1, err := core.GitPull("")
				if err != nil {
					return err
				}
				need2, err := core.GitPull("/develop/jd_cookie")
				if err != nil {
					return err
				}
				if !need1 && !need2 {
					return name + "已是最新版。"
				}
				s.Reply(name + "代码拉取成功。")
				s.Reply(name + "正在编译程序。")
				if err := core.CompileCode(); err != nil {
					return err
				}
				s.Reply(name + "编译程序完毕。")
				core.Daemon()
				return nil
			},
		},
	})
}
