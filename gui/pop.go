package main

import (
	"fyne.io/fyne/widget"
)

func userPop(o *orb, shares *[]string, callbacks ...func()) *widget.PopUp {
	usernames, err := allUsernames(o)
	if err != nil {
		o.l.Fatalln(err)
	}
	v := widget.NewVBox()
	for _, u := range usernames {
		name := u
		check := widget.NewCheck(name, func(b bool) {
			if b {
				for _, s := range *shares {
					if s == name {
						return
					}
				}
				*shares = append(*shares, name)
			} else {
				for i := 0; i < len(*shares); i++ {
					if name == (*shares)[i] {
						*shares = append((*shares)[:i], (*shares)[i+1:]...)
						break
					}
				}
			}
			for _, c := range callbacks {
				c() // Fookin lol.
			}
		})
		if name == o.username {
			check.Disable()
		} else {
			for _, s := range *shares {
				if u == s {
					check.SetChecked(true)
					break
				}
			}
		}
		v.Append(check)
	}
	return widget.NewPopUp(v, o.current)
}
