package admin

import "errors"

// step 1: define admin mode type and constants
type AdminMode int

const (
	Manager AdminMode = iota + 1
	Editor
	Viewer
)

var (
	ErrInvalidAdminMode   = errors.New("invalid admin mode")
	ErrExceedsMaxAdmins   = errors.New("exceeds maximum number of admins")
	ErrExceedsMaxManagers = errors.New("exceeds maximum number of managers")
	ErrExceedsMaxEditors  = errors.New("exceeds maximum number of editors")
	ErrExceedsMaxViewers  = errors.New("exceeds maximum number of viewers")
)

const (
	MaxAdmins   = 0
	MaxManagers = 1
	MaxEditors  = 3
	MaxViewers  = 5
)

type AdminCount struct {
	Total    int
	Managers int
	Editors  int
	Viewers  int
}

var adminCount = AdminCount{}

func validateAndAddAdminMode(mode AdminMode) error {
	switch mode {
	case Manager:
		if adminCount.Managers >= MaxManagers {
			return ErrExceedsMaxManagers
		}
		adminCount.Managers++
	case Editor:
		if adminCount.Editors >= MaxEditors {
			return ErrExceedsMaxEditors
		}
		adminCount.Editors++
	case Viewer:
		if adminCount.Viewers >= MaxViewers {
			return ErrExceedsMaxViewers
		}
		adminCount.Viewers++
	default:
		return ErrInvalidAdminMode
	}

	adminCount.Total++
	if adminCount.Total > MaxAdmins {
		return ErrExceedsMaxAdmins
	}
	return nil
}
