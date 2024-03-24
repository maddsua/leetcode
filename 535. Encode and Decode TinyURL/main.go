package main

import (
	"strconv"
)

type Codec struct {
	data    map[int64]string
	counter int64
}

func Constructor() Codec {
	return Codec{
		data:    map[int64]string{},
		counter: 0,
	}
}

func (this *Codec) encode(longUrl string) string {

	this.counter++
	id := this.counter
	this.data[id] = longUrl

	return strconv.FormatInt(id, 10)
}

func (this *Codec) decode(shortUrl string) string {

	id, valid := strconv.ParseInt(shortUrl, 10, 64)
	if valid != nil {
		return "invalid id"
	}

	link, has := this.data[id]
	if !has {
		return "null"
	}

	return link
}

func main() {

	svc := Constructor()

	shortened := svc.encode("https://leetcode.com/problems/design-tinyurl")
	println(svc.decode(shortened))

}
