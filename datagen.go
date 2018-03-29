package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func OptimalSchedule(N int) {
	f := "OptimalSchedule"
	d := ""

	for i := 0; i <= N; i++ {
		s1 := rand.NewSource(time.Now().UnixNano() + int64(i))
		r1 := rand.New(s1)
		t1 := r1.Intn(20)
		t2 := t1 + r1.Intn(10)
		if t2 > t1 {
			d += fmt.Sprintf("%d\t%d\t%d\tEvent%d\n", i+1, t1, t2, i)
		}

	}

	CreateDataFile(f, d)
}

func CreateDataFile(f string, data string) error {
	fname := "./" + f
	dataF, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE, 0644)
	defer dataF.Close()
	dataF.Truncate(0)
	dataF.Seek(0, 0)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = dataF.WriteString(data)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func main() {
	fmt.Println("OptimalSchedule")
	OptimalSchedule(20)
}
