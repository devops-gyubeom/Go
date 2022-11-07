package koreaPost

import "fmt"

type PostSender struct {
}

func (k *PostSender) Send(parcel string) {
	fmt.Printf("KoreaPost sends %v parcel\n", parcel)
}
