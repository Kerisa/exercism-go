package romannumerals

import "errors"

// RomanNumData arabic-roman transform data
type RomanNumData struct {
    Arabic  int;
    Roman   string;
}

// ToRomanNumeral do arabic-roman transform
func ToRomanNumeral(arabic int) (string,error) {
    romanToArabic := []RomanNumData {
        RomanNumData{ 1000, "M"  },
        RomanNumData{  900, "CM" },
        RomanNumData{  500, "D"  },
        RomanNumData{  400, "CD" },
        RomanNumData{  100, "C"  },
        RomanNumData{   90, "XC" },
        RomanNumData{   50, "L"  },
        RomanNumData{   40, "XL" },
        RomanNumData{   10, "X"  },
        RomanNumData{    9, "IX" },
        RomanNumData{    5, "V"  },
        RomanNumData{    4, "IV" },
        RomanNumData{    1, "I"  },
    }

    if arabic <= 0 || arabic > 3000 {
        return "", errors.New("invalid number")
    }

    result := ""
    for i := 0; i < len(romanToArabic); i++ {
        for romanToArabic[i].Arabic <= arabic {
            result += romanToArabic[i].Roman
            arabic -= romanToArabic[i].Arabic
        }
    }
    return result, nil
}