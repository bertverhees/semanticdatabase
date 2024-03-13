package primitives

type INumeric interface {
	IAny
	//1..1 (abstract)
	//add alias "+"
	//Sum with other (commutative). Actual type of result depends on arithmetic balancing rules.
	Add(other INumeric) (INumeric, error)
	//1..1(abstract)
	//subtract alias "-" (
	//Result of subtracting other. Actual type of result depends on arithmetic balancing rules.
	Subtract(other INumeric) (INumeric, error)
	//1..1 (abstract)
	//multiply alias "*" (
	//Product by other. Actual type of result depends on arithmetic balancing rules.
	Multiply(other INumeric) (INumeric, error)
	//1..1 (abstract)
	//divide alias "/" (
	//Divide by`other`. Actual type of result depends on arithmetic balancing rules.
	Divide(other INumeric) (INumeric, error)
	//1..1 (abstract)
	//exponent alias "^" (
	//Expontiation of this by other.
	Exponent(other INumeric) (INumeric, error)
	//1..1 (abstract)
	//negative alias "-" (): Numeric
	//Generate negative of current value.
	Negative() (INumeric, error)
}
