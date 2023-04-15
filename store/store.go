package store

type Store map[string]any

func New() Store {
	return Store{}
}

func (s Store) Set(name string, value any) {
	s[name] = value
}

func (s Store) Get(name string) any {
	return s[name]
}

func (s Store) Delete(name string) {
	delete(s, name)
}
