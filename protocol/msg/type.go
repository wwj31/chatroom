package msg

//go install github.com/wwj31/spawner@v0.0.7
//go:generate spawner

func (c CHATROOM) Int32() int32 {
	return int32(c)
}
