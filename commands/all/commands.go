package all

import (
	"v2ray.com/core/commands/all/api"
	"v2ray.com/core/commands/all/tls"
	"v2ray.com/core/commands/base"
)

// go:generate go run v2ray.com/core/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		cmdConvert,
		cmdLove,
		tls.CmdTLS,
		cmdUUID,
		cmdVerify,

		// documents
		docFormat,
		docMerge,
	)
}
