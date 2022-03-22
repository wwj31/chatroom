package login

import (
	"fmt"
	"time"
)

func (s *Login) InitCmd() {
	s.RegistCmd("stats", s.stats, "stats <userName> 打印玩家信息")
	s.RegistCmd("popular", s.popular, "popular <roomId> 打印房间10分钟内发送频率最高的词")
}

func (s *Login) stats(arg ...string) {
	if len(arg) != 1 {
		fmt.Println("stats 指令参数数量错误")
		return
	}
	name := arg[0]

	info, ok := s.users[name]
	if !ok {
		fmt.Println("找不到玩家或者玩家未上线")
		return
	}
	t := time.UnixMilli(info.loginAt)
	sec := (time.Now().UnixMilli() - info.loginAt) / 1000
	fmt.Println(fmt.Sprintf("登录时间:%v 在线时长:%v分 %v秒，所在房间:%v", t.Format("2006-01-02 15:04:05"), sec/60, sec%60, info.roomId))
}
func (s *Login) popular(arg ...string) {
	if len(arg) != 1 {
		fmt.Println("stats 指令参数数量错误")
		return
	}
	id := arg[0]
	_ = s.Send("room"+id, "popular")
}
