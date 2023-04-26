package rector

import (
	"context"

	database "github.com/galexrt/fivenet/proto/resources/common/database"
	"github.com/galexrt/fivenet/query/fivenet/table"
	jet "github.com/go-jet/jet/v2/mysql"
)

const AuditLogPageSize = 30

var (
	auditLog = table.FivenetAuditLog.AS("auditentry")
	user     = table.Users.AS("usershort")
)

func (s *Server) ViewAuditLog(ctx context.Context, req *ViewAuditLogRequest) (*ViewAuditLogResponse, error) {
	condition := jet.Bool(true)
	if len(req.UserIds) > 0 {
		ids := make([]jet.Expression, len(req.UserIds))
		for i := 0; i < len(req.UserIds); i++ {
			ids[i] = jet.Int32(req.UserIds[i])
		}
		condition = condition.AND(auditLog.UserID.IN(ids...))
	}
	if req.From != nil {
		condition = condition.AND(auditLog.CreatedAt.GT_EQ(
			jet.TimestampT(req.From.AsTime()),
		))
	}
	if req.To != nil {
		condition = condition.AND(auditLog.CreatedAt.LT_EQ(
			jet.TimestampT(req.To.AsTime()),
		))
	}

	countStmt := auditLog.
		SELECT(
			jet.COUNT(auditLog.ID).AS("datacount.totalcount"),
		).
		FROM(
			auditLog,
		).
		WHERE(condition)

	var count database.DataCount
	if err := countStmt.QueryContext(ctx, s.db, &count); err != nil {
		return nil, err
	}

	resp := &ViewAuditLogResponse{
		Pagination: database.EmptyPaginationResponseWithPageSize(req.Pagination.Offset, AuditLogPageSize),
	}
	if count.TotalCount <= 0 {
		return resp, nil
	}

	if count.TotalCount <= 0 {
		return resp, nil
	}

	stmt := auditLog.
		SELECT(
			auditLog.AllColumns,
			user.ID,
			user.Identifier,
			user.Job,
			user.JobGrade,
			user.Firstname,
			user.Lastname,
		).
		FROM(
			auditLog.
				LEFT_JOIN(user,
					user.ID.EQ(auditLog.UserID),
				),
		).
		WHERE(condition).
		ORDER_BY(
			auditLog.CreatedAt.DESC(),
		).
		OFFSET(req.Pagination.Offset).
		LIMIT(AuditLogPageSize)

	if err := stmt.QueryContext(ctx, s.db, &resp.Logs); err != nil {
		return nil, err
	}

	database.PaginationHelperWithPageSize(resp.Pagination,
		count.TotalCount,
		req.Pagination.Offset,
		len(resp.Logs),
		AuditLogPageSize)

	return resp, nil
}
