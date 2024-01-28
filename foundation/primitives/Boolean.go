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

func (p *Boolean) setValue(value bool) {
	p.value = value
}

/*Value equality: return True if this and other are attached to objects considered to be equal in value.
 */
func (p *Boolean) IsEqual(any Any) *Boolean {
	b := any.(*Boolean)
	return NewBoolean(p.value == b.value)
}

/*Reference equality for reference types, value equality for value types.
 */
func (p *Boolean) Equal(any Any) *Boolean {
	return nil
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
Post_de_Morgan: Result = not (not self or not other)
Post_commutative: Result = (other and self)
*/
func (p *Boolean) Conjunction(other Boolean) *Boolean {
	return NewBoolean(p.value == other.Value())
}
