package mlog

import (
	"github.com/gin-gonic/gin"
	"os"
	"time"
	"io"
	"fmt"
	"log"
)

func init() {
	gin.DisableConsoleColor()
	f, _ := os.Create("cloud10_" + time.Now().Format("2006-01-02") + ".log")
	go func() {
		gin.DefaultWriter = io.MultiWriter(f)
		tick := time.Tick(time.Hour * 24 * 3)
		for {
			select {
			case m := <-tick:
				logTime := m.Format("2006-01-02")
				fmt.Println(logTime)
				f, _ := os.Create("cloud10_" + logTime + ".log")
				gin.DefaultWriter = io.MultiWriter(f)
			}
		}
	}()
	gin.SetMode(gin.DebugMode)
}

func createLogFile() {
	go func() {
		t := time.Tick(7 * time.Hour * 24)
		//t := time.Tick(5 * time.Second)
		for {
			select {
			case s := <-t:
				log.SetFlags(log.Lshortfile | log.Ltime)
				file, err := os.Open("cloud9.log")
				fmt.Println(err)
				if err != nil {
					file, err = os.Create("cloud9.log")
					fmt.Println(err)
				}
				log.SetOutput(file)
				fmt.Println(s.String())
				break
			}
		}
	}()
}
