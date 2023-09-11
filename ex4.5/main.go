// ex4.5 delete adjacent dublicates.
package main

func deleteAdjacentDub(strings []string) []string {
	idx := 0 
	
	for _, v := range strings {
		if strings[idx] == v {
			continue
		}
		idx++
		strings[idx] = v
	}
	return strings[:idx+1]
}