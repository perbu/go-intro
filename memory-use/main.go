package main

type person struct {
	name string
	phones []string
	pois map[string]poi
	poip map[string]*poi
}

type poi struct {
	name string

}
