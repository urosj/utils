// Copyright 2014 Canonical Ltd.
// Copyright 2014 Cloudbase Solutions SRL
// Licensed under the AGPLv3, see LICENCE file for details.

package series_test

import (
	"sort"

	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/utils/series"
)

type supportedSeriesWindowsSuite struct {
}

var _ = gc.Suite(&supportedSeriesWindowsSuite{})

func (s *supportedSeriesWindowsSuite) TestSeriesVersion(c *gc.C) {
	vers, err := series.SeriesVersion("win8")
	if err != nil {
		c.Assert(err, gc.Not(gc.ErrorMatches), `invalid series "win8"`, gc.Commentf(`unable to lookup series "win8"`))
	} else {
		c.Assert(err, jc.ErrorIsNil)
	}
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(vers, gc.Equals, "win8")
}

func (s *supportedSeriesWindowsSuite) TestSupportedSeries(c *gc.C) {
	expectedSeries := []string{
		"arch",
		"centos7",
		"precise",
		"quantal",
		"raring",
		"saucy",
		"trusty",
		"utopic",
		"vivid",
		"win10",
		"win2012",
		"win2012hv",
		"win2012hvr2",
		"win2012r2",
		"win7",
		"win8",
		"win81",
	}
	series := series.SupportedSeries()
	sort.Strings(series)
	c.Assert(series, gc.DeepEquals, expectedSeries)
}
