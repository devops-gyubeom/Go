package main

import (
	"github.com/devops-gyubeom/Go-Tucker/ch20/fedex"
	"github.com/devops-gyubeom/Go-Tucker/ch20/koreaPost"
)

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	sender := &koreaPost.PostSender
	SendBook("어린 왕자", sender)
	SendBook("그리스인 조르바", sender)
}
