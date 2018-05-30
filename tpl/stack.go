package tpl

// StackT stores template names in nesting order
type StackT []string

// Push stores a template name at the end
func (sp *StackT) Push(pushee string) {
	s := *sp
	if s == nil {
		s = []string{}
	}
	s = append(s, pushee)
	*sp = s // without this , no assigment to sp
}

// Pop removes a template name from the end
func (sp *StackT) Pop() string {
	s := *sp
	if s == nil {
		return ""
	}
	l := len(s)
	if l == 0 {
		return ""
	}
	el := s[l-1]
	*sp = s[:l-1]
	return el
}

// Unshift removes a template name from the top
func (sp *StackT) Unshift() string {
	s := *sp
	if s == nil {
		return ""
	}
	l := len(s)
	if l == 0 {
		return ""
	}
	el := s[0]
	*sp = s[1:]
	return el
}
