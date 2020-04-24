package tool

type ResultDTO struct {
	url string
	hash string
}

func (r ResultDTO) Render() {
	println(r.url + " " + r.hash)
}
