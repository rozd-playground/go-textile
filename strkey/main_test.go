package strkey

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	cases := []struct {
		Name                string
		Address             string
		ExpectedVersionByte VersionByte
	}{
		{
			Name:                "AccountID",
			Address:             "TBU2RRGLXH3E5CQHTD3ODLDF2BWDCYUSSBLLZ5GNW7JXHDIYKXZWHXL7",
			ExpectedVersionByte: VersionByteAccountID,
		},
		{
			Name:                "Seed",
			Address:             "SBU2RRGLXH3E5CQHTD3ODLDF2BWDCYUSSBLLZ5GNW7JXHDIYKXZWHOKR",
			ExpectedVersionByte: VersionByteSeed,
		},
		{
			Name:                "Other (0x60)",
			Address:             "MBU2RRGLXH3E5CQHTD3ODLDF2BWDCYUSSBLLZ5GNW7JXHDIYKXZWGTOG",
			ExpectedVersionByte: VersionByte(0x60),
		},
	}

	for _, kase := range cases {
		actual, err := Version(kase.Address)
		if assert.NoError(t, err, "An error occured decoding case %s", kase.Name) {
			assert.Equal(t, kase.ExpectedVersionByte, actual, "Output mismatch in case %s", kase.Name)
		}
	}
}
