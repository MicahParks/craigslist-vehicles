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
	canChan   chan fyne.CanvasObject
	death     chan struct{}
	l         *log.Logger
	postsCol  *mongo.Collection
	presetCol *mongo.Collection
	user      *types.User
	userCol   *mongo.Collection
	username  string
}

func main() {
	l := log.New(os.Stdout, "", log.LstdFlags|log.LUTC|log.Lshortfile)
	//canLogin := false
	username := ""
	if userB, err := ioutil.ReadFile(diskAsset); err != nil {
		if !os.IsNotExist(err) {
			l.Fatalln(err.Error())
		}
		//canLogin = true
	} else {
		username = strings.TrimSpace(string(userB))
	}
	userCol, err := mongodb.Init("User")
	if err != nil {
		l.Fatalln(err.Error())
	}
	postsCol, err := mongodb.Init("Posts")
	if err != nil {
		l.Fatalln(err.Error())
	}
	presetCol, err := mongodb.Init("Preset")
	if err != nil {
		l.Fatalln(err.Error())
	}
	o := &orb{
		canChan:   make(chan fyne.CanvasObject),
		death:     make(chan struct{}),
		l:         l,
		postsCol:  postsCol,
		presetCol: presetCol,
		userCol:   userCol,
		username:  username,
	}
	a := app.New()
	w := a.NewWindow("cars")
	w.SetContent(homeCan(o))
	//if canLogin {
	//	w.SetContent(loginCan(o))
	//} else {
	//	w.SetContent(registerCan(o))
	//}
	go func() {
		for {
			select {
			case <-o.death:
				a.Quit()
				return
			case can := <-o.canChan:
				w.SetContent(can)
			}
		}
	}()
	w.ShowAndRun()
}
