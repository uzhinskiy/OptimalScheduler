/*
OptimalScheduling(I)
	While (I != 0) do
		Из всего множества фильмов I выбираем фильм j с самым ранним окончанием съемок
			Удаляем фильм j и любой другой фильм, съемки которого пересекаются с фильмом j, из множества доступных фильмов I
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sort"
	"strings"
	"text/template"
)

type PlotData struct {
	DATA string
}

type events struct {
	event    string
	position int
	start    int
	stop     int
	distance int
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var iEvents, uEvents []events
	var J events

	for {
		str, err := r.ReadString(10)                          // 0x0A separator = newline
		fields := strings.Split(strings.TrimSpace(str), "\t") // делим строки на фрагменты
		if len(fields) == 4 {
			iEvents = append(iEvents, events{fields[3], _atoi(fields[0]), _atoi(fields[1]), _atoi(fields[2]), _atoi(fields[2]) - _atoi(fields[1])})
		}
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		}
	}

	for k := 0; k <= len(iEvents); k++ {
		// сортировка массива по полю "конец отрезка"
		sort.Slice(iEvents[:], func(i, j int) bool {
			return iEvents[i].stop < iEvents[j].stop
		})
		// Выделяем отрезок, который заканчивается раньше всех и заносим его в итоговый массив
		J = iEvents[0]
		uEvents = append(uEvents, J)
		// Удаляем этот отрезок из исходного массива
		iEvents = iEvents[1:]

		// Перебираем остаток и выкидываем все пересечения с J
		for x := 0; x < len(iEvents); x++ {
			//fmt.Println("Now testing", J.event, " and ", iEvents[x].event)
			if _intersect(J.start, J.stop, iEvents[x].start, iEvents[x].stop) == true {
				//fmt.Println(J.event, " is Intersect with ", iEvents[x].event)
				iEvents = append(iEvents[:x], iEvents[x+1:]...)
				x = x - 1
			}
		}
		//fmt.Println("AFTER", iEvents, k, len(iEvents))
		k = 0
	}
	fmt.Println("Total: ", uEvents)
	_plot(uEvents)
}

// вспомогательная функция ковертации строки в INT
func _atoi(s string) int {
	var (
		n uint64
		i int
		v byte
	)
	for ; i < len(s); i++ {
		d := s[i]
		if '0' <= d && d <= '9' {
			v = d - '0'
		} else if 'a' <= d && d <= 'z' {
			v = d - 'a' + 10
		} else if 'A' <= d && d <= 'Z' {
			v = d - 'A' + 10
		} else {
			n = 0
			break
		}
		n *= uint64(10)
		n += uint64(v)
	}
	return int(n)
}

func _intersect(B0, E0, B1, E1 int) bool {
	// учитывай дистанцию!!!!!
	intersect := false

	if E0 <= B1 {
		intersect = false
	} else if B0 > B1 && E0 < E1 {
		intersect = true
	} else if B0 < B1 && E0 < E1 {
		intersect = true
	} else if B0 == B1 && E0 < E1 {
		intersect = true
	} else if B0 == B1 && E0 > E1 {
		intersect = true
	} else if B0 == B1 && E0 == E1 {
		intersect = true
	} else if B0 > B1 && E0 == E1 {
		intersect = true
	} else if B0 < B1 && E0 == E1 {
		intersect = true
	} else if B0 > B1 && E0 > E1 {
		intersect = true
	} else if B0 < B1 && E0 > E1 {
		intersect = true
	}
	return intersect
}

func _plot(data []events) {
	str := ""
	sort.Slice(data[:], func(i, j int) bool {
		return data[i].position < data[j].position
	})

	for _, v := range data {
		str += fmt.Sprintf("%d\t%d\t%d\t%s\n", v.position, v.start, v.stop, v.event)
	}
	d := PlotData{str}
	fn := fmt.Sprintf("/tmp/plot%d", os.Getpid())
	plotF, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE, 0600)
	defer plotF.Close()
	if err != nil {
		fmt.Println("Error while opening file. ", err)
	}

	plotS, err := template.ParseFiles("./plot.tmpl")
	if err != nil {
		fmt.Println("Error while parse file. ", err)
	}

	err = plotS.Execute(plotF, d)
	if err != nil {
		fmt.Println("Error while parse file. ", err)
	}

	out, err := exec.Command("/usr/bin/gnuplot", fn).Output()
	if err != nil {
		fmt.Println("Error ", err)
		fmt.Println("Output ", out)
	}

}
