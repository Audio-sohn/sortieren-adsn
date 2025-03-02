package mergesort

// Merge erwartet zwei sortierte Listen und gibt eine sortierte Liste zurück,
// die alle Elemente der beiden Eingabelisten enthält.
func Merge(list1, list2 []int) []int {

	if len(list1) == 0 {

		return list2

	} else if len(list2) == 0 {

		return list1

	}

	return nil

	//return append(Merge()
}
