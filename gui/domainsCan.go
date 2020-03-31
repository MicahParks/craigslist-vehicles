package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

var (
	domains = map[string]bool{
		"frederick":    false,
		"lynchburg":    false,
		"norfolk":      false,
		"richmond":     true,
		"washingtondc": false,
	}
)

func domainsCan(o *orb) *fyne.Container {
	con := fyne.NewContainerWithLayout(layout.NewGridLayout(2))
	for k, v := range domains {
		con.AddObject(widget.NewLabel(k))
		check := widget.NewCheck("", func(b bool) {
			if v {
				return
			}
			if !b {
				for i := 0; i < len(o.user.Domains); i++ {
					if o.user.Domains[i] == k {
						o.user.Domains = append(o.user.Domains[:i], o.user.Domains[i+1:]...)
						break
					}
				}
			} else {
				o.user.Domains = append(o.user.Domains, k)
			}
			if err := updateDomains(o); err != nil {
				o.l.Fatalln(err.Error())
			}
		})
		for _, d := range o.user.Domains {
			if d == k {
				check.Checked = true
				check.Refresh()
				break
			}
		}
		if v {
			check.Disable()
		}
		con.AddObject(check)
	}
	back := widget.NewButton("back", func() {
		o.canChan <- homeCan(o)
	})
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, back, nil, nil), back, con)
}
