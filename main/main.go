package main
import (
//	"flag"
	"io"
	"fmt"
	"unique7/breacher"
)


func main() {
	hb := breacher.Make_from("f", "b")
	breacher.Make_key(&hb)
	bt, _ := breacher.Write_key(&hb)
	bt.Seek(2, io.SeekStart)

	t := make([]byte, 4)
	bt.Read(t)
	breacher.Close(&hb)

	fmt.Println(breacher.To_s(t))
}
