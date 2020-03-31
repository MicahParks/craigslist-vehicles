package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

var (
	domains = map[string]bool{
		"richmond":     true,
		"washingtondc": false,
	}
)

func domainsCan(o *orb) *fyne.Container {
	con := fyne.NewContainerWithLayout(layout.NewGridLayout(3))
	for k, v := range domains {
		con.AddObject(widget.NewLabel(k))
		check := widget.NewCheck("", func(_ bool) {})
		for _, d := range o.user.Domains {
			if d == k {
				check.Checked = true
				break
			}
		}
		check.Disable()
		label := "switch"
		if v {
			label = "required"
		}
		con.AddObject(widget.NewButton(label, func() {
			if v {
				return
			}
			if check.Checked {
				for i := 0; i < len(o.user.Domains); i++ {
					if o.user.Domains[i] == k {
						o.user.Domains = append(o.user.Domains[:i], o.user.Domains[i+1:]...)
						break
					}
				}
			} else {
				o.user.Domains = append(o.user.Domains, k)
			}
			check.Checked = !check.Checked
			check.Refresh()
			if err := updateDomains(o); err != nil {
				o.l.Fatalln(err.Error())
			}
		}))
		con.AddObject(check)
	}
	return con
}
