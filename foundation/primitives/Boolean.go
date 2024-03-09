package primitives

/* Type representing minimal interface of built-in Boolean type.

 */

type Boolean struct {
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
	v := ConvertToBooleanFromIAny(b)
	return NewBoolean(p.value == v.value)
}

func (p *Boolean) NotEqual(b IAny) *Boolean {
	v := ConvertToBooleanFromIAny(b)
	return NewBoolean(p.value != v.value)
}

/*Create new instance of a type.
 */
func (p *Boolean) InstanceOf(aType String) *Boolean {
	return NewBoolean(false)
}

/*
conjunction alias "and", "∧", "&" (other: Boolean[1]): Boolean
Post_de_Morgan: Result = not (not self or not other)
Post_commutative: Result = (other and self)
NB: the bitwise operator "&" is not supported on boolean, so is replaced by "AndThen"
*/
func (p *Boolean) And(other *Boolean) *Boolean {
	return p.AndThen(other)
}

/*
semistrict_conjunction alias "and then", "&&" (other: Boolean[1]): Boolean
Post_de_Morgan: Result = not (not self or else not other)
*/
func (p *Boolean) AndThen(other *Boolean) *Boolean {
	return NewBoolean(p.value && other.Value())
}

/*
disjunction alias "or", "∨", "|" (other: Boolean[1]): Boolean
Post_de_Morgan: Result = not (not self and not other)
Post_commutative: Result = (other or Current)
Post_consistent_with_semi_strict: Result implies (self or else other)
NB: the bitwise operator "|" is not supported on boolean, so is replaced by "OrElse"
*/
func (p *Boolean) Or(other *Boolean) *Boolean {
	return p.OrElse(other)
}

/*
semistrict_disjunction alias "or else", "||" (other: Boolean[1]): Boolean
Post_de_Morgan: Result = not (not self and then not other)
*/
func (p *Boolean) OrElse(other *Boolean) *Boolean {
	return NewBoolean(p.value || other.Value())
}

/*
exclusive_disjunction alias "xor", "⊻" (other: Boolean[1]): Boolean
Post_definition: Result = self or other) and not (self and other
*/
func (p *Boolean) Xor(other *Boolean) *Boolean {
	return NewBoolean((p.value || other.Value()) && !(p.value && other.Value()))
}

/*
implication alias "implies", "⇒" (other: Boolean[1]): Boolean
Post_definition: Result = (not self or else other)
*/
func (p *Boolean) Implies(other *Boolean) *Boolean {
	return NewBoolean(!p.value || other.Value())
}

/*
The logical NNOR (“Neither Nor”) is an operation on two logical values, typically the values of two propositions,
that produces a value of true if and only if both of its operands are false.
*/
func (p *Boolean) NeitherNor(other *Boolean) *Boolean {
	return NewBoolean(!p.value && !other.Value())
}

/*
The logical NAND is an operation on two logical values, typically the values of two propositions,
that produces a value of false if and only if both of its operands are true.
*/
func (p *Boolean) NeitherAnd(other *Boolean) *Boolean {
	return NewBoolean(!(p.value && other.Value()))
}

/*
negation alias "not", "¬", "!" (): Boolean
*/
func (p *Boolean) Not() *Boolean {
	return NewBoolean(!p.value)
}

/*
Logical equality is an operation on two logical values, typically the values of two propositions,
that produces a value of true if and only if both operands are false or both operands are true.
*/
func (p *Boolean) Equal(other *Boolean) *Boolean {
	return NewBoolean(p.value == other.Value())
}
