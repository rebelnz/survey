package sessions

import (
	
	
)

var Store = sessions.NewCookieStore([]byte("something-secret"))
