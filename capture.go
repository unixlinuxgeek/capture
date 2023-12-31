// Capture 1 frame from video
// Захватываем один кадр
// example: ./cap /var/tmp/in.mp4 /var/tmp/out.jpg

package capture

import (
	"errors"
	"fmt"
	"github.com/unixlinuxgeek/coreutil"
	"log"
	"os"
	"os/exec"
	"time"
)

func OneFrame(pth string, out string, time time.Time) (string, error) {
	if coreutil.Installed("ffmpeg") {
		fmt.Println("pth: " + pth)
		fmt.Println("out: " + out)
		HH, MM, SS := time.Clock()

		vidName := pth
		imgName := out
		f, err := os.OpenFile(vidName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 777)
		err = f.Close()
		if err != nil {
			log.Fatal(err)
		}

		// ffmpeg -y -ss 00:01:00 -i in.mp4 -frames:v 1 -q:v 2 output.jpg
		app := "/usr/bin/ffmpeg"
		arg0 := "-y"
		arg1 := "-ss"
		arg2 := string(rune(HH)) + ":" + string(rune(MM)) + ":" + string(rune(SS)) //"00:00:01"
		arg3 := "-i"
		arg4 := vidName
		arg5 := "-vf"
		arg6 := "scale=200:200"

		arg7 := "-frames:v"
		arg8 := "1"
		arg9 := "-q:v"
		arg10 := "2"
		arg11 := imgName

		cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11)
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		return imgName, nil
	} else {
		fmt.Println("no")
		return "", errors.New("ffmpeg is not installed!!!.")
	}
	return "", errors.New("Error!!!")
}
