package docstore

import (
	context "context"
	"errors"

	"github.com/galexrt/arpanet/pkg/auth"
	"github.com/galexrt/arpanet/proto/resources/documents"
	"github.com/galexrt/arpanet/query/arpanet/table"
	jet "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/qrm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	docRef = table.ArpanetDocumentsReferences.AS("documentreference")
	docRel = table.ArpanetDocumentsRelations.AS("documentrelation")
)

func (s *Server) GetDocumentReferences(ctx context.Context, req *GetDocumentReferencesRequest) (*GetDocumentReferencesResponse, error) {
	userId, job, jobGrade := auth.GetUserInfoFromContext(ctx)
	check, err := s.checkIfUserHasAccessToDoc(ctx, req.DocumentId, userId, job, jobGrade, true, documents.DOC_ACCESS_VIEW)
	if err != nil {
		return nil, err
	}
	if !check {
		return nil, status.Error(codes.PermissionDenied, "You don't have permission to view this document!")
	}

	references, err := s.getDocumentReferences(ctx, req.DocumentId)
	if err != nil {
		return nil, err
	}

	return &GetDocumentReferencesResponse{
		References: references,
	}, nil
}

func (s *Server) GetDocumentRelations(ctx context.Context, req *GetDocumentRelationsRequest) (*GetDocumentRelationsResponse, error) {
	userId, job, jobGrade := auth.GetUserInfoFromContext(ctx)
	check, err := s.checkIfUserHasAccessToDoc(ctx, req.DocumentId, userId, job, jobGrade, true, documents.DOC_ACCESS_VIEW)
	if err != nil {
		return nil, err
	}
	if !check {
		return nil, status.Error(codes.PermissionDenied, "You don't have permission to view this document!")
	}

	relations, err := s.getDocumentRelations(ctx, req.DocumentId)
	if err != nil {
		return nil, err
	}

	return &GetDocumentRelationsResponse{
		Relations: relations,
	}, nil
}

func (s *Server) AddDocumentReference(ctx context.Context, req *AddDocumentReferenceRequest) (*AddDocumentReferenceResponse, error) {
	userId, job, jobGrade := auth.GetUserInfoFromContext(ctx)
	// Check if user has access to both documents
	check, err := s.checkIfUserHasAccessToDocs(ctx, userId, job, jobGrade, false, documents.DOC_ACCESS_EDIT,
		req.Reference.SourceDocumentId, req.Reference.TargetDocumentId)
	if err != nil {
		return nil, err
	}
	if !check {
		return nil, status.Error(codes.PermissionDenied, "You don't have permission to add references from/to this document!")
	}

	req.Reference.CreatorId = userId
	stmt := docRef.
		INSERT().
		MODEL(req.Reference)

	if _, err := stmt.ExecContext(ctx, s.db); err != nil {
		return nil, err
	}

	resp := &AddDocumentReferenceResponse{}

	return resp, nil
}
func (s *Server) RemoveDcoumentReference(ctx context.Context, req *RemoveDcoumentReferenceRequest) (*RemoveDcoumentReferenceResponse, error) {
	userId, job, jobGrade := auth.GetUserInfoFromContext(ctx)
	var docIDs struct {
		Source uint64
		Target uint64
	}

	// Get document IDs of reference entry
	docsStmt := docRef.
		SELECT(
			docRef.SourceDocumentID.AS("source"),
			docRef.TargetDocumentID.AS("target"),
		).
		FROM(docRef).
		WHERE(docRef.ID.EQ(jet.Uint64(req.Id))).
		LIMIT(1)

	if err := docsStmt.QueryContext(ctx, s.db, &docIDs); err != nil {
		return nil, err
	}

	check, err := s.checkIfUserHasAccessToDocs(ctx, userId, job, jobGrade, false, documents.DOC_ACCESS_EDIT, docIDs.Source, docIDs.Target)
	if err != nil {
		return nil, err
	}
	if !check {
		return nil, status.Error(codes.PermissionDenied, "You don't have permission to remove references from this document!")
	}

	stmt := docRef.
		UPDATE().
		SET(
			docRef.DeletedAt.SET(jet.CURRENT_TIMESTAMP()),
		).
		WHERE(
			docRef.ID.EQ(jet.Uint64(req.Id)),
		)

	if _, err := stmt.ExecContext(ctx, s.db); err != nil {
		return nil, err
	}

	resp := &RemoveDcoumentReferenceResponse{}
	return resp, nil
}
func (s *Server) AddDocumentRelation(ctx context.Context, req *AddDocumentRelationRequest) (*AddDocumentRelationResponse, error) {
	userId, job, jobGrade := auth.GetUserInfoFromContext(ctx)
	check, err := s.checkIfUserHasAccessToDoc(ctx, req.Relation.DocumentId, userId, job, jobGrade, false, documents.DOC_ACCESS_EDIT)
	if err != nil {
		return nil, err
	}
	if !check {
		return nil, status.Error(codes.PermissionDenied, "You don't have permission to add relation from/to this document!")
	}

	req.Relation.SourceUserId = userId
	stmt := docRef.
		INSERT().
		MODEL(req.Relation)

	if _, err := stmt.ExecContext(ctx, s.db); err != nil {
		return nil, err
	}

	resp := &AddDocumentRelationResponse{}

	return resp, nil
}
func (s *Server) RemoveDcoumentRelation(ctx context.Context, req *RemoveDcoumentRelationRequest) (*RemoveDcoumentRelationResponse, error) {
	userId, job, jobGrade := auth.GetUserInfoFromContext(ctx)
	var docID struct {
		ID uint64
	}

	// Get document IDs of reference entry
	docsStmt := docRel.
		SELECT(
			docRel.DocumentID.AS("id"),
		).
		FROM(docRel).
		WHERE(docRel.ID.EQ(jet.Uint64(req.Id))).
		LIMIT(1)

	if err := docsStmt.QueryContext(ctx, s.db, &docID); err != nil {
		return nil, err
	}

	check, err := s.checkIfUserHasAccessToDoc(ctx, docID.ID, userId, job, jobGrade, false, documents.DOC_ACCESS_EDIT)
	if err != nil {
		return nil, err
	}
	if !check {
		return nil, status.Error(codes.PermissionDenied, "You don't have permission to remove references from this document!")
	}

	stmt := docRel.
		UPDATE().
		SET(
			docRel.DeletedAt.SET(jet.CURRENT_TIMESTAMP()),
		).
		WHERE(
			docRel.ID.EQ(jet.Uint64(req.Id)),
		)

	if _, err := stmt.ExecContext(ctx, s.db); err != nil {
		return nil, err
	}

	resp := &RemoveDcoumentRelationResponse{}

	return resp, nil
}

func (s *Server) getDocumentReferences(ctx context.Context, documentID uint64) ([]*documents.DocumentReference, error) {
	sourceDoc := docs.AS("source_document")
	uCreator := u.AS("ref_creator")
	stmt := docRef.
		SELECT(
			docRef.ID,
			docRef.CreatedAt,
			docRef.SourceDocumentID,
			docRef.Reference,
			docRef.TargetDocumentID,
			docRef.CreatorID,
			sourceDoc.ID,
			sourceDoc.CreatedAt,
			sourceDoc.UpdatedAt,
			sourceDoc.CategoryID,
			sourceDoc.Title,
			sourceDoc.CreatorID,
			sourceDoc.State,
			sourceDoc.Closed,
			dCategory.ID,
			dCategory.Name,
			dCategory.Description,
			uCreator.ID,
			uCreator.Identifier,
			uCreator.Job,
			uCreator.JobGrade,
			uCreator.Firstname,
			uCreator.Lastname,
		).
		FROM(
			docRef.
				LEFT_JOIN(sourceDoc,
					docRef.SourceDocumentID.EQ(sourceDoc.ID),
				).
				LEFT_JOIN(dCategory,
					dCategory.ID.EQ(sourceDoc.CategoryID),
				).
				LEFT_JOIN(uCreator,
					docRef.CreatorID.EQ(uCreator.ID),
				),
		).
		WHERE(
			jet.AND(
				docRef.TargetDocumentID.EQ(jet.Uint64(documentID)),
				docRef.DeletedAt.IS_NULL(),
			),
		).
		LIMIT(25)

	var dest []*documents.DocumentReference
	if err := stmt.QueryContext(ctx, s.db, &dest); err != nil {
		if !errors.Is(qrm.ErrNoRows, err) {
			return nil, err
		}
	}

	return dest, nil
}

func (s *Server) getDocumentRelations(ctx context.Context, documentID uint64) ([]*documents.DocumentRelation, error) {
	uSource := u.AS("source_user")
	uTarget := u.AS("target_user")
	stmt := docRel.
		SELECT(
			docRel.ID,
			docRel.CreatedAt,
			docRel.DocumentID,
			docRel.SourceUserID,
			docRel.Relation,
			docRel.TargetUserID,
			docs.ID,
			docs.CreatedAt,
			docs.UpdatedAt,
			docs.CategoryID,
			docs.Title,
			docs.CreatorID,
			docs.State,
			docs.Closed,
			dCategory.ID,
			dCategory.Name,
			dCategory.Description,
			uSource.ID,
			uSource.Identifier,
			uSource.Job,
			uSource.JobGrade,
			uSource.Firstname,
			uSource.Lastname,
			uTarget.ID,
			uTarget.Identifier,
			uTarget.Job,
			uTarget.JobGrade,
			uTarget.Firstname,
			uTarget.Lastname,
		).
		FROM(
			docRel.
				LEFT_JOIN(docs,
					docs.ID.EQ(docRel.DocumentID),
				).
				LEFT_JOIN(dCategory,
					docs.CategoryID.EQ(dCategory.ID),
				).
				LEFT_JOIN(uSource,
					uSource.ID.EQ(docRel.SourceUserID),
				).
				LEFT_JOIN(uTarget,
					uTarget.ID.EQ(docRel.TargetUserID),
				),
		).
		WHERE(
			jet.AND(
				docRel.DocumentID.EQ(jet.Uint64(documentID)),
				docRel.DeletedAt.IS_NULL(),
			),
		).
		LIMIT(25)

	var dest []*documents.DocumentRelation
	if err := stmt.QueryContext(ctx, s.db, &dest); err != nil {
		if !errors.Is(qrm.ErrNoRows, err) {
			return nil, err
		}
	}

	return dest, nil
}
