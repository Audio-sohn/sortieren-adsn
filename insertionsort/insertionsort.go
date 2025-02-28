package insertionsort

// InsertionSort sortiert die gegebene Liste mit dem Insertion-Sort-Algorithmus.
func InsertionSort(list []int) {

	// ab 2. element anfangen und durch liste iterieren
	for i := 1; i < len(list); i++ {

		pos := i
		// pro index so lange nach links tauschen bis nächstes element kleiner is
		for pos > 0 && list[pos-1] > list[pos] {

			MoveLeft(list, pos)
			pos--
		}
	}
}
