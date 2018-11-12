package grep

import (
    "bufio"
    "os"
    "io"
    "strings"
    "strconv"
)

// Search document here
func Search(pattern string, flags, files []string) []string {
    caseInsensitive := false
    printLineNo     := false
    invertResult    := false
    printFilename   := false
    matchEntireLine := false

    for _,flag := range(flags) {
        switch flag {
        case "-i":
            caseInsensitive = true
        case "-n":
            printLineNo = true
        case "-v":
            invertResult = true
        case "-l":
            printFilename = true
        case "-x":
            matchEntireLine = true
        }
    }

    result := []string{}

    if (caseInsensitive) {
        pattern = strings.ToLower(pattern)
    }

    for _,filename := range(files) {
        file,err := os.Open(filename)
        if err != nil {
            continue
        }

        prefix := ""
        if len(files) > 1 {
            prefix = filename + ":"
        }

        br := bufio.NewReader(file)
        line := 1
        for {
            a,_,c := br.ReadLine()
            if c == io.EOF {
                break
            }

            data := string(a)
            if caseInsensitive {
                data = strings.ToLower(string(a))
            }

            match := false
            if matchEntireLine {
                match = data == pattern
            } else {
                match = strings.Contains(data, pattern)
            }

            if match && !invertResult || !match && invertResult {
                if printFilename {
                    result = append(result, filename)
                    break
                } else if printLineNo {
                    result = append(result, prefix + strconv.Itoa(line) + ":" + string(a))
                } else {
                    result = append(result, prefix + string(a))
                }
            }

            line++
        }
        file.Close()
    }

    return result
}