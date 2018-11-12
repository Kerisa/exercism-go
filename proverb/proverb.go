package proverb

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
    var proverb []string

    for i := 0; i < len(rhyme) - 1; i++ {
        proverb = append(proverb, "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost.")
    }

    if len(rhyme) > 0 {
        proverb = append(proverb, "And all for the want of a " + rhyme[0] + ".")
    }
	return proverb
}
