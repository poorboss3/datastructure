package leetcode

type OrderedStream struct {
	Len    int
	Ptr    int
	Values []string
}

func Constructor(n int) OrderedStream {
	stream := OrderedStream{
		Len: n,
		Ptr: 1,
	}
	stream.Values = make([]string, n)
	return stream
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.Values[idKey] = value
	res := make([]string, 0)
	if this.Ptr == idKey {
		for ; this.Values[idKey] != "" && idKey <= this.Len; idKey++ {
			res = append(res, this.Values[idKey])
		}
		this.Ptr = idKey
	}
	return res
}
