package db

type Repo interface {
	getOne(id string) (interface{}, error)
	getAll() ([]interface{}, error)
	addOne(item interface{}) (interface{}, error)
	update(id string, item interface{}) (interface{}, error)
}

type Doc struct{}

func (d Doc) getOne(id string) (interface{}, error) {
	// TODO
	return nil, nil
}

func (d Doc) getAll() (interface{}, error) {
	// TODO
	return nil, nil
}

func (d Doc) addOne(item interface{}) (interface{}, error) {
	// TODO
	return nil, nil
}

func (d Doc) update(id string, item interface{}) (interface{}, error) {
	// TODO
	return nil, nil
}
