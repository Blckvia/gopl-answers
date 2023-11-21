// ex1.3 compares string concatenation techniques using Benchmarks

package concat

func Concat(args []string) {
	r, sep := "", ""
	for _, a := range args {
		r += sep + a
		sep = " "
	}
}
