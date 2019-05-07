package hlc

import "github.com/noxproject/nox/common/hash"

const (
	//128
	MaxBlockHeaderPayload = 4 + (hash.HashSize * 3) + 4 + 8 + 8 + 8
	//every mode position
	TIMESTART = 112
	TIMEEND = 120
	NONCESTART = 120
	NONCEEND = 128
	HEIGHTSTART = 104
	HEIGHTEND = 112
	NBITSTART = 100
	NBITEND = 104
	STATESTART = 68
	STATEEND = 100
	MERKLESTART = 36
	MERKLEEND = 68
	PRESTART = 4
	PREEND = 36
	VERSIONSTART = 0
	VERSIONEND = 4
)