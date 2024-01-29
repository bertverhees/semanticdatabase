package interval

import (
	"github.com/bertverhees/interval"
	"semanticdatabase/foundation/primitives"
	"semanticdatabase/generics"
)

/*
_type representing a 'proper' Interval, i.e. any two-sided or one-sided interval.
Inv_not_point: lower /= upper
*/
type ProperInterval[T generics.Number] struct {
	primitives.Any
	interval.Interval[T]
}

/*Value equality: return True if this and other are attached to objects considered to be equal in value.
 */
func (p *ProperInterval[T]) IsEqual(b primitives.IAny) *primitives.Boolean {
	return primitives.NewBoolean(
		p.Lower() == b.(*ProperInterval[T]).Lower() &&
			p.Upper() == b.(*ProperInterval[T]).Upper() &&
			p.LowerIncluded() == b.(*ProperInterval[T]).LowerIncluded() &&
			p.UpperIncluded() == b.(*ProperInterval[T]).UpperIncluded() &&
			p.LowerUnbounded() == b.(*ProperInterval[T]).LowerUnbounded() &&
			p.UpperUnbounded() == b.(*ProperInterval[T]).UpperUnbounded())
}

/*Reference equality for reference types, value equality for value types.
 */
func (p *ProperInterval[T]) Equal(any primitives.IAny) *primitives.Boolean {
	return p.IsEqual(any)
}
