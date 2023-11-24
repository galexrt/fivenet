package centrumutils

import "github.com/galexrt/fivenet/gen/go/proto/resources/centrum"

func IsStatusDispatchComplete(in centrum.StatusDispatch) bool {
	return in == centrum.StatusDispatch_STATUS_DISPATCH_ARCHIVED ||
		in == centrum.StatusDispatch_STATUS_DISPATCH_CANCELLED ||
		in == centrum.StatusDispatch_STATUS_DISPATCH_COMPLETED
}

func IsStatusDispatchUnassigned(in centrum.StatusDispatch) bool {
	return in == centrum.StatusDispatch_STATUS_DISPATCH_UNSPECIFIED ||
		in == centrum.StatusDispatch_STATUS_DISPATCH_NEW ||
		in == centrum.StatusDispatch_STATUS_DISPATCH_UNASSIGNED
}

func IsDispatchUnassigned(in *centrum.Dispatch) bool {
	if in == nil {
		return false
	}

	// Ignore dispatches with no status
	if in.Status == nil {
		return false
	}

	// Ignore completed dispatches
	if IsStatusDispatchComplete(in.Status.Status) {
		return false
	}

	// Dispatch is "new" or unassgined
	if IsStatusDispatchUnassigned(in.Status.Status) {
		return true
	}

	return len(in.Units) == 0
}