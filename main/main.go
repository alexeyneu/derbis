package main
import (
//	"flag"
	"io"
	"fmt"
	"github.com/alexeyneu/derbis/breacher"
)


func main() {
/*	env := flag.String("address", "zero", "")
	flag.Parse()
	if *env == "zero" {
		on_green.Make_c()

	} else {
		on_green.Made(*env)
	}
*/
//	w, wh := breacher.Make_c()
	hb := breacher.Make_from("f", "b")
	breacher.Make_key(&hb)
	breacher.Write_key(&hb)
	breacher.Key_Seek(&hb, 0, io.SeekStart)
	t := make([]byte, 4)
	breacher.Key_ReadFull(&hb, t)
	breacher.Close(&hb)

	fmt.Println(breacher.To_s(t))


}
