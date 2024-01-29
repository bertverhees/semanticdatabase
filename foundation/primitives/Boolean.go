package primitives

/* Type representing minimal interface of built-in Boolean type.

 */

type Boolean struct {
	Any
	value bool
}

func NewBoolean(value bool) *Boolean {
	v := new(Boolean)
	v.value = value
	return v
}

func (p *Boolean) Value() bool {
	return p.value
}

func (p *Boolean) SetValue(value bool) {
	p.value = value
}

/*Value equality: return True if this and other are attached to objects considered to be equal in value.
 */
func (p *Boolean) IsEqual(b IAny) *Boolean {
	return NewBoolean(p.value == b.(*Boolean).value)
}

/*Reference equality for reference types, value equality for value types.
 */
func (p *Boolean) Equal(any IAny) *Boolean {
	return p.IsEqual(any)
}

/*Create new instance of a type.
 */
func (p *Boolean) InstanceOf(aType String) *Boolean {
	return nil
}

/*Type name of an object as a string. May include generic parameters, as in "Interval<Time>".
 */
func (p *Boolean) TypeOf(anObject Any) *String {
	return nil
}

/*True if current object not equal to other. Returns not equal().
 */
func (p *Boolean) NotEqual(other Ordered) *Boolean {
	return nil
}

/*
conjunction alias "and", "∧", "&" (other: Boolean[1]): Boolean
Post_de_Morgan: Result = not (not self or not other)
Post_commutative: Result = (other and self)
NB: the bitwise operator "&" is not supported on boolean, so is replaced by "AndThen"
*/
func (p *Boolean) And(other Boolean) *Boolean {
	return p.AndThen(other)
}

/*
semistrict_conjunction alias "and then", "&&" (other: Boolean[1]): Boolean
Post_de_Morgan: Result = not (not self or else not other)
*/
func (p *Boolean) AndThen(other Boolean) *Boolean {
	return NewBoolean(p.value && other.Value())
}

/*
disjunction alias "or", "∨", "|" (other: Boolean[1]): Boolean
Post_de_Morgan: Result = not (not self and not other)
Post_commutative: Result = (other or Current)
Post_consistent_with_semi_strict: Result implies (self or else other)
NB: the bitwise operator "|" is not supported on boolean, so is replaced by "OrElse"
*/
func (p *Boolean) Or(other Boolean) *Boolean {
	return p.OrElse(other)
}

/*
semistrict_disjunction alias "or else", "||" (other: Boolean[1]): Boolean
Post_de_Morgan: Result = not (not self and then not other)
*/
func (p *Boolean) OrElse(other Boolean) *Boolean {
	return NewBoolean(p.value || other.Value())
}

/*
exclusive_disjunction alias "xor", "⊻" (other: Boolean[1]): Boolean
Post_definition: Result = self or other) and not (self and other
*/
func (p *Boolean) Xor(other Boolean) *Boolean {
	return NewBoolean((p.value || other.Value()) && !(p.value && other.Value()))
}

/*
implication alias "implies", "⇒" (other: Boolean[1]): Boolean
Post_definition: Result = (not self or else other)
*/
func (p *Boolean) Implies(other Boolean) *Boolean {
	return NewBoolean(!p.value || other.Value())
}

/*
negation alias "not", "¬", "!" (): Boolean
*/
func (p *Boolean) Not() *Boolean {
	return NewBoolean(!p.value)
}
