package encoding

import "encoding/base32"

const alphabet = "0123456789abcdefghjkmnpqrstvwxyz"

var Base32 = base32.NewEncoding(alphabet).WithPadding(base32.NoPadding)
