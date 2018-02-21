// astroconst library
package astroconst

import (
	"math"
)

const (
	// constants for conversions
	Pi float64 = math.Pi
	TwoPi float64 = 2.0 * Pi
	Deg2Rad float64 = Pi / 180.0
	Rad2Deg float64 = 180.0 / Pi
	Day2Sec float64 = 86400.
	Day2Min float64 = 1440.
	
	// masses conversions from/to Msun
	Msun2mer = 6.0236e6      // Msun to Mmer
	Mmer2sun = 1.0 / Msun2mer  // Mmer to Msun
	Msun2ven = 4.08523719e5  // Msun to Mven
	Mven2sun = 1.0 / Msun2ven  //  Mven to Msun
	Msun2ear = 332946.0487   // Msun to Mear
	Mear2sun = 1.0 / Msun2ear  //  Mear to Msun
	Msun2mar = 3.09870359e6  // Msun to Mmar
	Mmar2sun = 1.0 / Msun2mar  //  Mmar to Msun
	Msun2jup = 1.047348644e3 // Msun to Mjup
	Mjup2sun = 1.0 / Msun2jup  //  Mjup to Msun
	Msun2sat = 3.4979018e3   // Msun to Msat
	Msat2sun = 1.0 / Msun2sat  //  Msat to Msun
	Msun2ura = 2.290298e4    // Msun to Mura
	Mura2sun = 1.0 / Msun2ura  //  Mura to Msun
	Msun2nep = 1.941226e4    // Msun to Mnep
	Mnep2sun = 1.0 / Msun2nep  //  Mnep to Msun
	// masses of Solar System objects
	Msun = 1.9884e30         // Sun mass in kg
	Mmer = Msun * Mmer2sun     //  Mercury mass in kg
	Mven = Msun * Mven2sun     //  Venus mass in kg
	Mear = 5.9722e24         // Earth mass in kg
	Mmar = Msun * Mmar2sun     //  Mars mass in kg
	Mjup = Msun * Mjup2sun     //  Jupiter mass in kg
	Msat = Msun * Msat2sun     //  Saturn mass in kg
	Mura = Msun * Mura2sun     //  Uranus mass in kg
	Mnep = Msun * Mnep2sun     //  Neptune mass in kg
	// convert masses between planets - not all
	Mjup2ear = Mjup / Mear
	Mear2jup = Mear / Mjup
	Mnep2ear = Mnep / Mear
	Mear2nep = Mear / Mnep
)
