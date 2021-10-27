package imp

type Store interface {

}

type memStore struct {
	imps map [string] string
}
