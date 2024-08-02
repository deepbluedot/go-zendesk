package mock

import (
	"github.com/deepbluedot/go-zendesk/zendesk"
)

var _ zendesk.API = (*Client)(nil)
