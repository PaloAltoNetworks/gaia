package constants

// EntityStatus defines the status of an entity
type EntityStatus string

const (
	// Active defines that an entity is active
	Active EntityStatus = "Active"

	// Disabled defines that an entity is disabled
	Disabled EntityStatus = "Disabled"

	// Candidate defines that an entity is a candidate and could be either activated or deleted or disabled
	Candidate EntityStatus = "Candidate"
)

// Vulnerability defines the security vulnerability of an image or a layer
type Vulnerability int

const (
	// VulnerabilityNone defines that security vulnerability is None
	VulnerabilityNone Vulnerability = iota

	// VulnerabilityUnknown defines that security vulnerability is Unknown
	VulnerabilityUnknown

	// VulnerabilityNegligible defines that security vulnerability is Negligible
	VulnerabilityNegligible

	// VulnerabilityLow defines that security vulnerability is Low
	VulnerabilityLow

	// VulnerabilityMedium defines that security vulnerability is Medium
	VulnerabilityMedium

	// VulnerabilityHigh defines that security vulnerability is High
	VulnerabilityHigh

	// VulnerabilityCritical defines that security vulnerability is Critical
	VulnerabilityCritical

	// VulnerabilityDefcon1 defines that security vulnerability is Defcon1
	VulnerabilityDefcon1
)

// EntityStatusToString converts EntityStatusToString to its corresponding string value
func EntityStatusToString(entityStatus EntityStatus) string {

	switch entityStatus {
	case Active:
		return "Active"
	case Disabled:
		return "Disabled"
	case Candidate:
		return "Candidate"
	default:
		return ""
	}
}

const (

	// StatsTagKeyDestinationID represents the destination ID tag.
	StatsTagKeyDestinationID = "@destid"

	// StatsTagKeySourceID represents the source ID tag.
	StatsTagKeySourceID = "@srcid"

	// StatsTagKeyDestinationPort represents the destination port tag.
	StatsTagKeyDestinationPort = "@destport"

	// StatsTagKeyProcessingUnitID represents the processing unit ID tag.
	StatsTagKeyProcessingUnitID = "$id"

	// StatsTagKeyIP represents the IP tag.
	StatsTagKeyIP = "@ip"

	// StatsTagKeyAction represents the action tag.
	StatsTagKeyAction = "@action"

	// StatsTagKeyReason represents the reason tag.
	StatsTagKeyReason = "@reason"

	// StatsTagKeyEvent represents the event tag.
	StatsTagKeyEvent = "@event"

	// StatsTagKeyImage represents the image tag.
	StatsTagKeyImage = "@image"

	// StatsTagKeyServerID represents the server ID tag.
	StatsTagKeyServerID = "@serverid"

	// StatsTagKeyNamespace represents the namespace tag.
	StatsTagKeyNamespace = "$namespace"

	// StatsTagKeyAporetoContextID represents the aporeto context ID.
	StatsTagKeyAporetoContextID = "@squallid"

	// StatsTagKeyExternal represents the if the flow is coming from an uncontrolled network.
	StatsTagKeyExternal = "@external"

	// StatsTagKeyFileAccessPath represents the file path.
	StatsTagKeyFileAccessPath = "@path"

	// StatsTagKeyFileAccessHost represents the host where the file is stored.
	StatsTagKeyFileAccessHost = "@host"

	// StatsTagKeyFileAccessProtocol represents the protocol used to access the file.
	StatsTagKeyFileAccessProtocol = "@protocol"

	// StatsTagKeyFileAccessMode represents the mode of file.
	StatsTagKeyFileAccessMode = "@mode"

	// StatsTagKeySyscallAccessName represents the number of the syscall.
	StatsTagKeySyscallAccessName = "@syscall"

	// StatsTagKeySyscallAccessPID represents the pid of the related process.
	StatsTagKeySyscallAccessPID = "@pid"

	// StatsTagKeySyscallAccessProcessName represents the name of the related process.
	StatsTagKeySyscallAccessProcessName = "@processname"
)
