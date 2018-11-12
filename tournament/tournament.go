package tournament

import (
	"bytes"
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"	
)

type teamInfo struct {
	name string
	win,draw,loss int
}

func (t *teamInfo)Score() int {
	return t.win * 3 + t.draw
}

// Tally tallies the results of a small football competition.
func Tally(input io.Reader, out io.Writer) error {
	var teams = map[string]*teamInfo{}
	var reader = bufio.NewReader(input)
	for {
		line,_,err := reader.ReadLine()
		if err != nil {
			break
		}
		if line = bytes.Trim(line, " "); len(line) == 0 || bytes.HasPrefix(line, []byte("#")) {
			continue
		}

		parts := strings.Split(string(line), ";")
		if len(parts) != 3 {
			return errors.New("error line format")
		}

		if teams[parts[0]] == nil {
			teams[parts[0]] = &teamInfo{name:parts[0]}
		}
		if teams[parts[1]] == nil {
			teams[parts[1]] = &teamInfo{name:parts[1]}
		}

		switch parts[2] {
		case "win":
			teams[parts[0]].win++
			teams[parts[1]].loss++
		case "loss":
			teams[parts[0]].loss++
			teams[parts[1]].win++
		case "draw":
			teams[parts[0]].draw++
			teams[parts[1]].draw++
		default:
			return errors.New("unknown result")
		}
	}

	var arr = make([]*teamInfo, 0, len(teams))
	for _,v := range teams {
		arr = append(arr, v)
	}
	sort.Slice(arr, func(i, j int)bool {
		l,r := arr[i].Score(), arr[j].Score()
		return l > r || (l == r && arr[i].name < arr[j].name)
	})
	out.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
	for _,t := range arr {
		out.Write([]byte(fmt.Sprintf("%-31s| %2d | %2d | %2d | %2d | %2d\n", t.name, t.win+t.draw+t.loss, t.win, t.draw, t.loss, t.win*3+t.draw)))
	}
	return nil
}