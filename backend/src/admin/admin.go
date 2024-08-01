package admin

import (
	"errors"
)

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
	MaxManagers = 1
	MaxEditors  = 3
	MaxViewers  = 5
	MaxAdmins   = MaxManagers + MaxEditors + MaxViewers
)

type AdminCount struct {
	Total    int
	Managers int
	Editors  int
	Viewers  int
}

var adminCount = AdminCount{}

func ValidateAndAddAdminMode(mode AdminMode, count AdminCount) error {
	switch mode {
	case Manager:
		if count.Managers >= MaxManagers {
			return ErrExceedsMaxManagers
		}
		count.Managers++
	case Editor:
		if count.Editors >= MaxEditors {
			return ErrExceedsMaxEditors
		}
		count.Editors++
	case Viewer:
		if count.Viewers >= MaxViewers {
			return ErrExceedsMaxViewers
		}
		count.Viewers++
	default:
		return ErrInvalidAdminMode
	}

	count.Total++
	if count.Total > MaxAdmins {
		return ErrExceedsMaxAdmins
	}
	return nil
}

func ValidateAndRemoveAdminMode(mode AdminMode, count AdminCount) {
	switch mode {
	case Manager:
		count.Managers--
	case Editor:
		count.Editors--
	case Viewer:
		count.Viewers--
	}

	count.Total--
}
