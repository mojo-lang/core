package core

type Options map[string]interface{}

func NewOptions(kvs ...interface{}) Options {
	options := make(Options)
	return options.SetValues(kvs...)
}

func (o Options) SetValue(key string, value interface{}) Options {
	o[key] = value
	return o
}

func (o Options) SetValues(kvs ...interface{}) Options {
	for i := 0; i < len(kvs)-1; i += 2 {
		o[kvs[i].(string)] = kvs[i+1]
	}
	return o
}

func (o Options) Merge(options Options) Options {
	for k, v := range options {
		o[k] = v
	}
	return o
}

func (o Options) KeyValues() []interface{} {
	var kvs []interface{}
	if o != nil {
		for k, v := range o {
			kvs = append(kvs, k)
			kvs = append(kvs, v)
		}
	}
	return kvs
}
