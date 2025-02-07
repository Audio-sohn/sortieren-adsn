package bubblesort

// BubbleUp implementiert eine Iteration des Bubble-Sort-Algorithmus.
// Es iteriert über die Liste und vergleicht jedes Element mit dem nächsten.
// Wenn das Element größer als das nächste ist, werden sie vertauscht.
// Die Funktion gibt true zurück, wenn mindestens ein Tausch durchgeführt wurde.
func BubbleUp(list []int) bool {
	swapped := false

	for i, speci := range list {

		// is specimen bigger than next entry?
		if i+1 < len(list) && speci > list[i+1] {

			// if so, swap with next entry
			temp := list[i+1]
			list[i+1] = speci
			list[i] = temp

			// set swapped to true
			swapped = true
		}

	}

	return swapped
}
