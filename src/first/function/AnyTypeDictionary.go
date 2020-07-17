package function

type Dictionary struct {
	data map[interface{}]interface{}
}

func (d *Dictionary) Get(key interface{}) interface{} {
	return d.data[key]
}

func (d *Dictionary) Set(key, value interface{}) {
	d.data[key] = value
}

func (d *Dictionary) Visit(callback func(k, v interface{}) bool) {
	if callback == nil {
		return
	}

	for key, value := range d.data {
		if !callback(key, value) {
			return
		}
	}
}

func (d *Dictionary) Clear() {
	d.data = make(map[interface{}]interface{})
}

func NewDictionary() *Dictionary {
	d := &Dictionary{}
	d.Clear()
	return d
}
