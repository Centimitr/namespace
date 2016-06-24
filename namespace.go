package namespace

type namespace struct {
	scopes  map[string]scope
	getRule func(scope, name string) string
}

type scope struct {
	// scope's name
	name string
	// the namespace it belongs to
	namespace         *namespace
	getRule           func(scope, name string) string
	UseDefaultGetRule bool
}

func (n *namespace) Apply(scopeName string) (ok bool, _ scope) {
	if ok, _ := n.scopes[scopeName]; !ok {
		n.scopes[scopeName] = scope{
			name:              scopeName,
			namespace:         n,
			UseDefaultGetRule: true,
		}
		return true, n.scopes[scope]
	}
	return false, nil
}

func (n *namespace) Use(scopeName string) scope {
	if ok, _ := n.scopes[scopeName]; !ok {
		n.scopes[scopeName] = scope{
			name:              scopeName,
			namespace:         n,
			UseDefaultGetRule: true,
		}
	}
	return n.scopes[scope]
}

func (n *namespace) SetGetRule(fn func(string, string) string) {
	n.getRule = fn
}

func (s *scope) SetGetRule(fn func(string, string) string) {
	s.UseDefaultGetRule = false
	s.getRule = fn
}

func (s *scope) Get(name string) string {
	if s.UseDefaultGetRule {
		return s.namespace.getRule(s.name, name)
	} else {
		return s.getRule(s.name, name)
	}
}

func (n *namespace) DirectGet(scope, name string) string {
	if ok, s := n.scopes[scope]; ok {
		return s.getRule(scope, name)
	} else {
		return n.getRule(scope, name)
	}
}

func New() {
	return namespace{
		scopes: make(map[string]scope),
		getRule: func(string, string) string {
			return string + '.' + string
		},
	}
}
