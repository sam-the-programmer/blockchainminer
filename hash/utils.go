// Basic utils for the hashes package.
package hash

// HashFunction is a function that returns a hashed version of a string
type HashFunction func(string) string

// DoubleHash returns a version of the inputted hash function
// that computes the hash twice, instead of once.
// (Input is hashed, then that hash is hashed).
func DoubleHash(f HashFunction) HashFunction {
	return MultiHash(f, 2)
}

// TripleHash returns a version of the inputted hash function
// that computes the hash 3 times, instead of once.
// (Input is hashed, then that hash is hashed, then that one...).
func TripleHash(f HashFunction) HashFunction {
	return MultiHash(f, 3)
}

// MultiHash returns a version of the inputted hash function
// that computes the hash i times, instead of once.
// (Input is hashed, then that hash is hashed, then that one...).
func MultiHash(f HashFunction, i uint) HashFunction {
	return func(s string) string {
		for x := uint(0); x < i; x++ {
			s = f(s)
		}

		return s
	}
}
