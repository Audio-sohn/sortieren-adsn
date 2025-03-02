package selectionsort

// SelectionSort sortiert eine Liste von Zahlen mit dem SelectionSort-Algorithmus.
func SelectionSort(list []int) {

	if len(list) == 0 {

		return

	}

	SwapSmallest(list)
	SelectionSort(list[1:])

}

// swapsmallest tauscht das kleinste element einer liste an den anfang
func SwapSmallest(list []int) {

	temp := list[0]
	pos := SmallestPos(list)

	list[0] = list[pos]
	list[pos] = temp

}
