package options

import "bee-playground/utils"

// Workshop for Functional Options
func Workshop() {
	result := *NewHouse("55/55", "Rodo", "Nontuaburi", WithFloors(5), WithOwner("Bee"), WithTelephoneNo("5003949"))
	utils.Dump(result)
}

func WithTelephoneNo(telephoneNo string) Options {
	return func(v *House) {
		v.TelephoneNo = telephoneNo
	}
}

func WithOwner(owner string) Options {
	return func(v *House) {
		v.Owner = owner
	}
}

func WithFloors(floors int) Options {
	return func(v *House) {
		v.Floors = floors
	}
}

type House struct {
	HouseNo     string
	Road        string
	Floors      int
	Province    string
	Owner       string
	TelephoneNo string
}

type Options func(v *House)

func NewHouse(HouseNo string, Road string, Province string, options ...Options) *House {
	result := &House{HouseNo: HouseNo, Road: Road, Province: Province}
	for _, option := range options {
		option(result)
	}
	return result
}
