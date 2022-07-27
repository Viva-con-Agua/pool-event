package token

import (
	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcapool"
)

var cookie = vcago.AccessCookieMiddleware(&vcapool.AccessToken{})
