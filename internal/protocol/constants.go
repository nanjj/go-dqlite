package protocol

// VersionOne is version 1 of the server protocol.
const VersionOne = uint64(1)

// VersionLegacy is the pre 1.0 dqlite server protocol version.
const VersionLegacy = uint64(0x86104dd760433fe5)

// SQLite datatype codes
const (
	Integer = 1
	Float   = 2
	Text    = 3
	Blob    = 4
	Null    = 5
)

// Special data types for time values.
const (
	UnixTime = 9
	ISO8601  = 10
	Boolean  = 11
)

// Request types.
const (
	RequestLeader    = 0
	RequestClient    = 1
	RequestHeartbeat = 2
	RequestOpen      = 3
	RequestPrepare   = 4
	RequestExec      = 5
	RequestQuery     = 6
	RequestFinalize  = 7
	RequestExecSQL   = 8
	RequestQuerySQL  = 9
	RequestInterrupt = 10
	RequestJoin      = 12
	RequestPromote   = 13
	RequestRemove    = 14
	RequestDump      = 15
	RequestCluster   = 16
)

// Response types.
const (
	ResponseFailure = 0
	ResponseServer  = 1
	ResponseWelcome = 2
	ResponseServers = 3
	ResponseDb      = 4
	ResponseStmt    = 5
	ResponseResult  = 6
	ResponseRows    = 7
	ResponseEmpty   = 8
	ResponseFiles   = 9
)
