package twofer

func ShareWith(name string) string {
	str := "One for "
	if name == "" {
        str = str + "you"
    } else {
        str = str + name
    }
    
    str = str + ", one for me."
    return str
}
