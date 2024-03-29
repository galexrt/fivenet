// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/jobs/conduct.proto
// source: services/jobs/jobs.proto
// source: services/jobs/qualifications.proto
// source: services/jobs/timeclock.proto

package permsjobs

import (
	"github.com/galexrt/fivenet/pkg/perms"
)

const (
	JobsConductServicePerm        perms.Category = "JobsConductService"
	JobsQualificationsServicePerm perms.Category = "JobsQualificationsService"
	JobsServicePerm               perms.Category = "JobsService"
	JobsTimeclockServicePerm      perms.Category = "JobsTimeclockService"

	JobsConductServiceCreateConductEntryPerm            perms.Name = "CreateConductEntry"
	JobsConductServiceDeleteConductEntryPerm            perms.Name = "DeleteConductEntry"
	JobsConductServiceListConductEntriesPerm            perms.Name = "ListConductEntries"
	JobsConductServiceListConductEntriesAccessPermField perms.Key  = "Access"
	JobsConductServiceUpdateConductEntryPerm            perms.Name = "UpdateConductEntry"
	JobsServiceGetColleaguePerm                         perms.Name = "GetColleague"
	JobsServiceGetColleagueAccessPermField              perms.Key  = "Access"
	JobsServiceListColleaguesPerm                       perms.Name = "ListColleagues"
	JobsServiceSetJobsUserPropsPerm                     perms.Name = "SetJobsUserProps"
	JobsServiceSetJobsUserPropsAccessPermField          perms.Key  = "Access"
	JobsServiceSetMOTDPerm                              perms.Name = "SetMOTD"
	JobsTimeclockServiceListInactiveEmployeesPerm       perms.Name = "ListInactiveEmployees"
	JobsTimeclockServiceListTimeclockPerm               perms.Name = "ListTimeclock"
	JobsTimeclockServiceListTimeclockAccessPermField    perms.Key  = "Access"
)
