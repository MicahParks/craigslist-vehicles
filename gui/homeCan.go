package main

import (
	"strings"

	"fyne.io/fyne/widget"
)

func homeCan(o *orb) *widget.Form {
	o.username = "test" // TODO Delete this.
	if err := getUser(o); err != nil {
		o.l.Fatalln(err.Error())
	}
	presetBox := widget.NewButton("view", func() {
		o.canChan <- presetCan(o)
	})
	createPresetBox := widget.NewButton("new", func() {
		o.canChan <- presetCreationCan(o)
	})
	hPreset := widget.NewHBox(presetBox, createPresetBox)
	suffix := ", "
	domains := ""
	for _, d := range o.user.Domains {
		domains += d + suffix
	}
	domains = strings.TrimSuffix(domains, suffix)
	domainBox := widget.NewEntry()
	domainBox.SetText(domains)
	upDomainsBox := widget.NewButton("update", func() {
		domains := make([]string, 0)
		for _, d := range strings.Split(domainBox.Text, suffix) {
			d = strings.TrimSpace(d)
			if len(d) != 0 {
				domains = append(domains, d)
			}
			if len(domains) != 0 {
				o.user.Domains = domains
				if err := updateDomains(o); err != nil {
					o.l.Fatalln(err.Error())
				}
			}
		}
	})
	hDomain := widget.NewHBox(domainBox, upDomainsBox)
	loginBox := widget.NewButton("logout", func() {
		o.username = ""
		o.user = nil
		// TODO Other logout stuff?
		o.canChan <- loginCan(o)
	})
	return widget.NewForm(widget.NewFormItem("preset", hPreset), widget.NewFormItem("domains", hDomain),
		widget.NewFormItem("logout", loginBox))
}
