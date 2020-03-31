package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
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
	listCol   *mongo.Collection
	postsCol  *mongo.Collection
	presetCol *mongo.Collection
	user      *types.User
	userCol   *mongo.Collection
	username  string
}

type logWriter struct {
	w *widget.Entry
}

func (l logWriter) Write(p []byte) (int, error) {
	l.w.SetText(l.w.Text + "\n" + string(p))
	println(string(p))
	return len(p), nil
}

func main() {
	l := log.New(os.Stdout, "", log.LstdFlags|log.LUTC|log.Lshortfile)
	canLogin := true
	username := ""
	if userB, err := ioutil.ReadFile(diskAsset); err != nil {
		if !os.IsNotExist(err) {
			l.Fatalln(err.Error())
		}
		canLogin = true // TODO Make false.
	} else {
		username = strings.TrimSpace(string(userB))
	}
	listCol, err := mongodb.Init("List")
	if err != nil {
		l.Fatalln(err.Error())
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
		listCol:   listCol,
		postsCol:  postsCol,
		presetCol: presetCol,
		userCol:   userCol,
		username:  username,
	}
	a := app.New()
	w := a.NewWindow("cars")
	logs := widget.NewMultiLineEntry()
	lW := &logWriter{w: logs}
	l.SetOutput(lW)
	logs.Disable()
	scrollTab := widget.NewTabItem("logs", widget.NewScrollContainer(logs))
	programTab := widget.NewTabItem("program", widget.NewLabel(""))
	con := widget.NewTabContainer(programTab, scrollTab)
	if canLogin {
		programTab.Content = loginCan(o)
	} else {
		programTab.Content = registerCan(o)
	}
	w.SetContent(con)
	go func() {
		for {
			select {
			case <-o.death:
				a.Quit()
				return
			case can := <-o.canChan:
				programTab.Content = can
				con.Refresh()
			}
		}
	}()
	w.ShowAndRun()
}
