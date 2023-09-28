// Capture 1 frame from video
// Захватываем один кадр
// example: ./cap /var/tmp/in.mp4 /var/tmp/out.jpg

package capture

import (
	"fmt"
	"github.com/unixlinuxgeek/coreutil"
	"log"
	"os"
	"os/exec"
)

func OneFrame() {
	if len(os.Args) == 3 {
		if coreutil.Installed("ffmpeg1") {
			// /var/tmp/in.mp4
			vidName := os.Args[1]
			// /var/tmp/in.jpg
			imgName := os.Args[2]
			f, err := os.OpenFile(vidName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 777)
			err = f.Close()
			if err != nil {
				//fmt.Printf("%s File not found!!!\n", vidName)
				log.Fatal(err)
			}

			// ffmpeg -y -ss 00:01:00 -i in.mp4 -frames:v 1 -q:v 2 output.jpg
			app := "/usr/bin/ffmpeg"
			arg0 := "-y"
			arg1 := "-ss"
			arg2 := "00:01:00"
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
		}
	} else {
		fmt.Fprintf(os.Stderr, "%s\n", "Enter input video file and output image path: ./cap /var/tmp/in.mp4 /var/tmp/out.jpg")
	}
}
