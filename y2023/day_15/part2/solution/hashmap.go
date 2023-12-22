package solution

import "fmt"

type Value struct {
	Label    string
	FocalLen int
}

type HashMap struct {
	data [256][]*Value
}

func (h *HashMap) Get(k string) (*Value, bool) {
	i := CustomHash(k)
	return h.at(i, k)
}

func (h *HashMap) at(i int, k string) (*Value, bool) {
	for _, value := range h.data[i] {
		if value.Label == k {
			return value, true
		}
	}
	return nil, false
}

func (h *HashMap) Remove(k string) {
	i := CustomHash(k)
	l := h.data[i]
	for j, value := range l {
		if value.Label == k {
			h.data[i] = append(l[:j], l[j+1:]...)
		}
	}
}

func (h *HashMap) Set(k string, v int) {
	i := CustomHash(k)
	if oldValue, ok := h.at(i, k); ok {
		oldValue.FocalLen = v
	} else {
		h.data[i] = append(h.data[i], &Value{Label: k, FocalLen: v})
	}
}

func (h *HashMap) Print() {
	for i, datum := range h.data {
		if len(datum) > 0 {
			fmt.Println("box: ", i)
			for _, value := range datum {
				fmt.Print(value, "\t")
			}
			fmt.Print("\t")
			fmt.Println()
		}
	}
}

func NewHashMap() *HashMap {
	return &HashMap{}
}
