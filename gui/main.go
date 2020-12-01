package main

import (
	"log"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/MicahParks/craigslist-vehicles/mongodb"
	"github.com/MicahParks/craigslist-vehicles/types"
)

type orb struct {
	current   fyne.Canvas
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
	l.w.SetText(string(p) + "\n" + l.w.Text)
	println(string(p))
	return len(p), nil
}

func main() {
	l := log.New(os.Stdout, "", log.LstdFlags|log.LUTC|log.Lshortfile)
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
	}
	a := app.New()
	w := a.NewWindow("Craigslist Vehicles")
	o.current = w.Canvas()
	logs := widget.NewMultiLineEntry()
	lW := &logWriter{w: logs}
	l.SetOutput(lW)
	logs.Disable()
	scrollTab := widget.NewTabItem("logs", widget.NewScrollContainer(logs))
	programTab := widget.NewTabItem("program", widget.NewLabel(""))
	con := widget.NewTabContainer(programTab, scrollTab)
	programTab.Content = loginCan(o)
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
				o.current = w.Canvas()
			}
		}
	}()
	w.ShowAndRun()
}
