package main

import "log"

const numSig = 65

type handler struct {
	mask [3]uint32 // 3个元素的数组, 元素类型uint32位, 占4个字节
}

func (h *handler) want(sig int) bool {
	return (h.mask[sig/32]>>uint(sig&31))&1 != 0
}

func (h *handler) set(sig int) {
	// uint(sig&31)是整数截取后5位
	h.mask[sig/32] |= 1 << uint(sig&31)
}

func (h *handler) clear(sig int) {
	h.mask[sig/32] &^= 1 << uint(sig&31)
}

func main() {
	h := handler{}
	h.set(1)
	log.Printf("h=%#v", h.mask[0])
}
