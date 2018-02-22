package astroconst_test

import (
	"testing"

	"fmt"

	"github.com/brunetto/astroconst"
	"github.com/stretchr/testify/assert"
)

func Test_Import(t *testing.T) {
	assert.NotEmpty(t, astroconst.Pi)
	assert.NotEmpty(t, astroconst.TwoPi)
	assert.NotEmpty(t, astroconst.Deg2Rad)
	assert.NotEmpty(t, astroconst.Rad2Deg)
	assert.NotEmpty(t, astroconst.Day2Sec)
	assert.NotEmpty(t, astroconst.Day2Min)
	assert.NotEmpty(t, astroconst.MSun2Mer)
	assert.NotEmpty(t, astroconst.MMer2Sun)
	assert.NotEmpty(t, astroconst.MSun2Ven)
	assert.NotEmpty(t, astroconst.MVen2Sun)
	assert.NotEmpty(t, astroconst.MSun2Ear)
	assert.NotEmpty(t, astroconst.MEar2Sun)
	assert.NotEmpty(t, astroconst.MSun2Mar)
	assert.NotEmpty(t, astroconst.MMar2Sun)
	assert.NotEmpty(t, astroconst.MSun2Jup)
	assert.NotEmpty(t, astroconst.MJup2Sun)
	assert.NotEmpty(t, astroconst.MSun2Sat)
	assert.NotEmpty(t, astroconst.MSat2Sun)
	assert.NotEmpty(t, astroconst.MSun2Ura)
	assert.NotEmpty(t, astroconst.MUra2Sun)
	assert.NotEmpty(t, astroconst.MSun2Nep)
	assert.NotEmpty(t, astroconst.MNep2Sun)
	assert.NotEmpty(t, astroconst.MSun)
	assert.NotEmpty(t, astroconst.MMer)
	assert.NotEmpty(t, astroconst.MVen)
	assert.NotEmpty(t, astroconst.MEar)
	assert.NotEmpty(t, astroconst.MMar)
	assert.NotEmpty(t, astroconst.MJup)
	assert.NotEmpty(t, astroconst.MSat)
	assert.NotEmpty(t, astroconst.MUra)
	assert.NotEmpty(t, astroconst.MNep)
	assert.NotEmpty(t, astroconst.MJup2Ear)
	assert.NotEmpty(t, astroconst.MEar2Jup)
	assert.NotEmpty(t, astroconst.MNep2Ear)
	assert.NotEmpty(t, astroconst.MEar2Nep)
}

func TestPlanetsMassConverter(t *testing.T) {
	assert.Equal(t, "0.0157", fmt.Sprintf("%1.4f", astroconst.PlanetsMassConverter(5, astroconst.MEar, astroconst.MJup)))
}
