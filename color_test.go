package color

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestPrintToFile(t *testing.T) {
	s := New(FgBlue, BgGreen).Sprint(time.Now().Format("2006-01-02 15:04:05"))
	f , err := os.OpenFile("./color.log", os.O_CREATE|os.O_RDWR, 660)
	if err != nil {
		fmt.Println(New(FgRed).Sprint(err.Error()))
		return
	}
	defer f.Close()
	f.Write([]byte(s))
}

func TestPrint(t *testing.T){
	fmt.Println(New(FgBlue, BgGreen).Sprint(time.Now().Format("2006-01-02 15:04:05")))

	for i := 0; i < 100; i++{
		fmt.Println(strconv.Itoa(i), New(i).Sprint(time.Now().Format("2006-01-02 15:04:05")))
	}
}

func BenchmarkSprint(b *testing.B) {
	c := FgRed
	str := "adaldjadjandasd"

	b.Run("print", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := "\033["+strconv.Itoa(int(c))+"m["+str+"]\033[0m"
			_ = s
		}
	})
	b.Run("printf", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := fmt.Sprintf("\033[%dm[", c)+str+fmt.Sprintf("]\033[0m")
			_ = s
		}
	})
}
