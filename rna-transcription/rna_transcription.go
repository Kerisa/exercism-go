package strand

// ToRNA simulate RNA Transcription
func ToRNA(dna string) string {
    rna := ""
    for _,ch := range(dna) {
        switch ch {
        case 'G':
            rna += "C"
        case 'C':
            rna += "G"
        case 'T':
            rna += "A"
        case 'A':
            rna += "U"
        default:
            panic("unknown dna nucleotides")
        }
    }
    return rna
}
