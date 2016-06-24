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
	if _, ok := n.scopes[scopeName]; !ok {
		n.scopes[scopeName] = scope{
			name:              scopeName,
			namespace:         n,
			UseDefaultGetRule: true,
		}
		return true, n.scopes[scopeName]
	}
	return false, scope{}
}

func (n *namespace) Use(scopeName string) scope {
	if _, ok := n.scopes[scopeName]; !ok {
		n.scopes[scopeName] = scope{
			name:              scopeName,
			namespace:         n,
			UseDefaultGetRule: true,
		}
	}
	return n.scopes[scopeName]
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
	if s, ok := n.scopes[scope]; ok {
		return s.getRule(scope, name)
	} else {
		return n.getRule(scope, name)
	}
}

func New() namespace {
	return namespace{
		scopes: make(map[string]scope),
		getRule: func(scope, name string) string {
			return scope + "." + name
		},
	}
}
