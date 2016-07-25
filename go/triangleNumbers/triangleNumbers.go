package triangleNumbers

type triangles struct {
	Count, Value int
}

func New() triangles {
	return triangles{Count: 1, Value: 1}
}

func (t *triangles) Next() {
	t.Count++
	t.Value += t.Count
	return
}
