package raindrops

import "strconv"

// RainDropData raindrops-speak data
type RainDropData struct {
   number int;
   speak string;
}

// Convert raindrop-speak
func Convert(number int) string {
    var s string

    dataMap := []RainDropData{
        RainDropData{3, "Pling" },
        RainDropData{5, "Plang" },
        RainDropData{7, "Plong" },
    }

    for _,m := range(dataMap) {
        if number % m.number == 0 {
            s += m.speak
        }
    }

    if s == "" {
        s = strconv.Itoa(number)
    }

    return s
}