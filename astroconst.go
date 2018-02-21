// Package astroconst is blablabla
package astroconst

import (
	"math"
)

const (
	// Constants for conversions

	// Pi is a constant representing the value of π (exported vars/const/methods should have a comment in the form // NAME ...)
	Pi float64 = math.Pi
	// TwoPi is a constant representing the value of 2π
	TwoPi float64 = 2.0 * Pi
	// Deg2Rad is the conversion factor from degrees to radiants: RAD = DEG * Deg2Rad
	Deg2Rad float64 = Pi / 180.0
	// Rad2Deg is the conversion factor from radiants to degrees: DEG = RAD * Deg2Rad
	Rad2Deg float64 = 180.0 / Pi
	// Day2Sec is the conversion factor from days to seconds
	Day2Sec float64 = 86400.
	// Day2Min is the conversion factor from days to minutes
	Day2Min float64 = 1440.

	// Conversion factors from/to MSun

	// MSun2Mer is the conversion factor from M☉ (solar masses) to M☿ (Mercury masses)
	MSun2Mer = 6.0236e6
	// MMer2Sun is ...
	MMer2Sun = 1.0 / MSun2Mer
	// MSun2Ven ...
	MSun2Ven = 4.08523719e5
	// MVen2Sun  ...
	MVen2Sun = 1.0 / MSun2Ven
	// MSun2Ear ...
	MSun2Ear = 332946.0487
	// MEar2Sun ...
	MEar2Sun = 1.0 / MSun2Ear
	// MSun2Mar ...
	MSun2Mar = 3.09870359e6
	// MMar2Sun ...
	MMar2Sun = 1.0 / MSun2Mar
	// MSun2Jup ...
	MSun2Jup = 1.047348644e3
	// MJup2Sun ...
	MJup2Sun = 1.0 / MSun2Jup
	// MSun2Sat ...
	MSun2Sat = 3.4979018e3
	// MSat2Sun ...
	MSat2Sun = 1.0 / MSun2Sat
	// MSun2Ura ...
	MSun2Ura = 2.290298e4
	// MUra2Sun ...
	MUra2Sun = 1.0 / MSun2Ura
	// MSun2Nep ...
	MSun2Nep = 1.941226e4
	// MNep2Sun ...
	MNep2Sun = 1.0 / MSun2Nep

	// Masses of Solar System objects

	// MSun is the Sun mass in kg, that is the conversion factor from solar masses to kg
	MSun = 1.9884e30
	// MMer is the Mercury mass in kg
	MMer = MSun * MMer2Sun //  Mercury mass in kg
	// MVen is the Venus mass in kg
	MVen = MSun * MVen2Sun //  Venus mass in kg
	//MEar ...
	MEar = 5.9722e24 // Earth mass in kg
	//MMar ...
	MMar = MSun * MMar2Sun //  Mars mass in kg
	//MJup ...
	MJup = MSun * MJup2Sun //  Jupiter mass in kg
	//MSat ...
	MSat = MSun * MSat2Sun //  Saturn mass in kg
	//MUra ...
	MUra = MSun * MUra2Sun //  Uranus mass in kg
	//MNep ...
	MNep = MSun * MNep2Sun //  Neptune mass in kg

	// Convert masses between planets - not all

	//MJup2Ear ...
	MJup2Ear = MJup / MEar
	//MEar2Jup ...
	MEar2Jup = MEar / MJup
	//MNep2Ear ...
	MNep2Ear = MNep / MEar
	//MEar2Nep ...
	MEar2Nep = MEar / MNep
)

// PlanetsMassConverter converts masses from a planet mass unit to another
// e.g. how to convert 5 Earth masses to Jupiter masses:
// 5 [M🜨] -> [M♃]: PlanetsMassConverter(5, MEar, MJup) = 0.015731799191836177 [M♃]
func PlanetsMassConverter(massValue float64, planetMassFrom float64, planetMassTo float64) float64 {
	return massValue * planetMassFrom / planetMassTo
}
