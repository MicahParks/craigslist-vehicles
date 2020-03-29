package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"go.mongodb.org/mongo-driver/mongo"

	"gitlab.com/MicahParks/cano-cars/mongodb"
	"gitlab.com/MicahParks/cano-cars/types"
)

const (
	diskAsset = "carsLogin.username"
)

type orb struct {
	canvasC  chan fyne.CanvasObject
	death    chan struct{}
	l        *log.Logger
	user     *types.User
	userCol  *mongo.Collection
	username string
}

func main() {
	l := log.New(os.Stdout, "", log.LstdFlags|log.LUTC|log.Lshortfile)
	canLogin := false
	username := ""
	if userB, err := ioutil.ReadFile(diskAsset); err != nil {
		if !os.IsNotExist(err) {
			l.Fatalln(err.Error())
		}
		canLogin = true
	} else {
		username = strings.TrimSpace(string(userB))
	}
	userCol, err := mongodb.Init("User")
	if err != nil {
		l.Fatalln(err.Error())
	}
	orb := &orb{
		canvasC:  make(chan fyne.CanvasObject),
		death:    make(chan struct{}),
		l:        l,
		userCol:  userCol,
		username: username,
	}
	a := app.New()
	w := a.NewWindow("cars")
	if canLogin {
		w.SetContent(loginCanvas(orb))
	} else {
		w.SetContent(registerCanvas(orb))
	}
	go func() {
		for {
			select {
			case <-orb.death:
				a.Quit()
				return
			case canv := <-orb.canvasC:
				w.SetContent(canv)
			}
		}
	}()
	w.ShowAndRun()
}
