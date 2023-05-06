import * as jspb from 'google-protobuf'

import * as resources_common_database_database_pb from '../../resources/common/database/database_pb';
import * as resources_documents_category_pb from '../../resources/documents/category_pb';
import * as resources_documents_documents_pb from '../../resources/documents/documents_pb';
import * as resources_documents_templates_pb from '../../resources/documents/templates_pb';


export class ListTemplatesRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListTemplatesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListTemplatesRequest): ListTemplatesRequest.AsObject;
  static serializeBinaryToWriter(message: ListTemplatesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListTemplatesRequest;
  static deserializeBinaryFromReader(message: ListTemplatesRequest, reader: jspb.BinaryReader): ListTemplatesRequest;
}

export namespace ListTemplatesRequest {
  export type AsObject = {
  }
}

export class ListTemplatesResponse extends jspb.Message {
  getTemplatesList(): Array<resources_documents_templates_pb.TemplateShort>;
  setTemplatesList(value: Array<resources_documents_templates_pb.TemplateShort>): ListTemplatesResponse;
  clearTemplatesList(): ListTemplatesResponse;
  addTemplates(value?: resources_documents_templates_pb.TemplateShort, index?: number): resources_documents_templates_pb.TemplateShort;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListTemplatesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListTemplatesResponse): ListTemplatesResponse.AsObject;
  static serializeBinaryToWriter(message: ListTemplatesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListTemplatesResponse;
  static deserializeBinaryFromReader(message: ListTemplatesResponse, reader: jspb.BinaryReader): ListTemplatesResponse;
}

export namespace ListTemplatesResponse {
  export type AsObject = {
    templatesList: Array<resources_documents_templates_pb.TemplateShort.AsObject>,
  }
}

export class GetTemplateRequest extends jspb.Message {
  getTemplateId(): number;
  setTemplateId(value: number): GetTemplateRequest;

  getData(): string;
  setData(value: string): GetTemplateRequest;

  getRender(): boolean;
  setRender(value: boolean): GetTemplateRequest;
  hasRender(): boolean;
  clearRender(): GetTemplateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTemplateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetTemplateRequest): GetTemplateRequest.AsObject;
  static serializeBinaryToWriter(message: GetTemplateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTemplateRequest;
  static deserializeBinaryFromReader(message: GetTemplateRequest, reader: jspb.BinaryReader): GetTemplateRequest;
}

export namespace GetTemplateRequest {
  export type AsObject = {
    templateId: number,
    data: string,
    render?: boolean,
  }

  export enum RenderCase { 
    _RENDER_NOT_SET = 0,
    RENDER = 3,
  }
}

export class GetTemplateResponse extends jspb.Message {
  getTemplate(): resources_documents_templates_pb.Template | undefined;
  setTemplate(value?: resources_documents_templates_pb.Template): GetTemplateResponse;
  hasTemplate(): boolean;
  clearTemplate(): GetTemplateResponse;

  getRendered(): boolean;
  setRendered(value: boolean): GetTemplateResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTemplateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetTemplateResponse): GetTemplateResponse.AsObject;
  static serializeBinaryToWriter(message: GetTemplateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTemplateResponse;
  static deserializeBinaryFromReader(message: GetTemplateResponse, reader: jspb.BinaryReader): GetTemplateResponse;
}

export namespace GetTemplateResponse {
  export type AsObject = {
    template?: resources_documents_templates_pb.Template.AsObject,
    rendered: boolean,
  }
}

export class CreateTemplateRequest extends jspb.Message {
  getTemplate(): resources_documents_templates_pb.Template | undefined;
  setTemplate(value?: resources_documents_templates_pb.Template): CreateTemplateRequest;
  hasTemplate(): boolean;
  clearTemplate(): CreateTemplateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateTemplateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateTemplateRequest): CreateTemplateRequest.AsObject;
  static serializeBinaryToWriter(message: CreateTemplateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateTemplateRequest;
  static deserializeBinaryFromReader(message: CreateTemplateRequest, reader: jspb.BinaryReader): CreateTemplateRequest;
}

export namespace CreateTemplateRequest {
  export type AsObject = {
    template?: resources_documents_templates_pb.Template.AsObject,
  }
}

export class CreateTemplateResponse extends jspb.Message {
  getId(): number;
  setId(value: number): CreateTemplateResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateTemplateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateTemplateResponse): CreateTemplateResponse.AsObject;
  static serializeBinaryToWriter(message: CreateTemplateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateTemplateResponse;
  static deserializeBinaryFromReader(message: CreateTemplateResponse, reader: jspb.BinaryReader): CreateTemplateResponse;
}

export namespace CreateTemplateResponse {
  export type AsObject = {
    id: number,
  }
}

export class UpdateTemplateRequest extends jspb.Message {
  getTemplate(): resources_documents_templates_pb.Template | undefined;
  setTemplate(value?: resources_documents_templates_pb.Template): UpdateTemplateRequest;
  hasTemplate(): boolean;
  clearTemplate(): UpdateTemplateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateTemplateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateTemplateRequest): UpdateTemplateRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateTemplateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateTemplateRequest;
  static deserializeBinaryFromReader(message: UpdateTemplateRequest, reader: jspb.BinaryReader): UpdateTemplateRequest;
}

export namespace UpdateTemplateRequest {
  export type AsObject = {
    template?: resources_documents_templates_pb.Template.AsObject,
  }
}

export class UpdateTemplateResponse extends jspb.Message {
  getId(): number;
  setId(value: number): UpdateTemplateResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateTemplateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateTemplateResponse): UpdateTemplateResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateTemplateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateTemplateResponse;
  static deserializeBinaryFromReader(message: UpdateTemplateResponse, reader: jspb.BinaryReader): UpdateTemplateResponse;
}

export namespace UpdateTemplateResponse {
  export type AsObject = {
    id: number,
  }
}

export class DeleteTemplateRequest extends jspb.Message {
  getId(): number;
  setId(value: number): DeleteTemplateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteTemplateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteTemplateRequest): DeleteTemplateRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteTemplateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteTemplateRequest;
  static deserializeBinaryFromReader(message: DeleteTemplateRequest, reader: jspb.BinaryReader): DeleteTemplateRequest;
}

export namespace DeleteTemplateRequest {
  export type AsObject = {
    id: number,
  }
}

export class DeleteTemplateResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteTemplateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteTemplateResponse): DeleteTemplateResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteTemplateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteTemplateResponse;
  static deserializeBinaryFromReader(message: DeleteTemplateResponse, reader: jspb.BinaryReader): DeleteTemplateResponse;
}

export namespace DeleteTemplateResponse {
  export type AsObject = {
  }
}

export class ListDocumentsRequest extends jspb.Message {
  getPagination(): resources_common_database_database_pb.PaginationRequest | undefined;
  setPagination(value?: resources_common_database_database_pb.PaginationRequest): ListDocumentsRequest;
  hasPagination(): boolean;
  clearPagination(): ListDocumentsRequest;

  getOrderbyList(): Array<resources_common_database_database_pb.OrderBy>;
  setOrderbyList(value: Array<resources_common_database_database_pb.OrderBy>): ListDocumentsRequest;
  clearOrderbyList(): ListDocumentsRequest;
  addOrderby(value?: resources_common_database_database_pb.OrderBy, index?: number): resources_common_database_database_pb.OrderBy;

  getSearch(): string;
  setSearch(value: string): ListDocumentsRequest;

  getCategoryIdsList(): Array<number>;
  setCategoryIdsList(value: Array<number>): ListDocumentsRequest;
  clearCategoryIdsList(): ListDocumentsRequest;
  addCategoryIds(value: number, index?: number): ListDocumentsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListDocumentsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListDocumentsRequest): ListDocumentsRequest.AsObject;
  static serializeBinaryToWriter(message: ListDocumentsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListDocumentsRequest;
  static deserializeBinaryFromReader(message: ListDocumentsRequest, reader: jspb.BinaryReader): ListDocumentsRequest;
}

export namespace ListDocumentsRequest {
  export type AsObject = {
    pagination?: resources_common_database_database_pb.PaginationRequest.AsObject,
    orderbyList: Array<resources_common_database_database_pb.OrderBy.AsObject>,
    search: string,
    categoryIdsList: Array<number>,
  }
}

export class ListDocumentsResponse extends jspb.Message {
  getPagination(): resources_common_database_database_pb.PaginationResponse | undefined;
  setPagination(value?: resources_common_database_database_pb.PaginationResponse): ListDocumentsResponse;
  hasPagination(): boolean;
  clearPagination(): ListDocumentsResponse;

  getDocumentsList(): Array<resources_documents_documents_pb.Document>;
  setDocumentsList(value: Array<resources_documents_documents_pb.Document>): ListDocumentsResponse;
  clearDocumentsList(): ListDocumentsResponse;
  addDocuments(value?: resources_documents_documents_pb.Document, index?: number): resources_documents_documents_pb.Document;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListDocumentsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListDocumentsResponse): ListDocumentsResponse.AsObject;
  static serializeBinaryToWriter(message: ListDocumentsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListDocumentsResponse;
  static deserializeBinaryFromReader(message: ListDocumentsResponse, reader: jspb.BinaryReader): ListDocumentsResponse;
}

export namespace ListDocumentsResponse {
  export type AsObject = {
    pagination?: resources_common_database_database_pb.PaginationResponse.AsObject,
    documentsList: Array<resources_documents_documents_pb.Document.AsObject>,
  }
}

export class GetDocumentRequest extends jspb.Message {
  getDocumentId(): number;
  setDocumentId(value: number): GetDocumentRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDocumentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetDocumentRequest): GetDocumentRequest.AsObject;
  static serializeBinaryToWriter(message: GetDocumentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDocumentRequest;
  static deserializeBinaryFromReader(message: GetDocumentRequest, reader: jspb.BinaryReader): GetDocumentRequest;
}

export namespace GetDocumentRequest {
  export type AsObject = {
    documentId: number,
  }
}

export class GetDocumentResponse extends jspb.Message {
  getDocument(): resources_documents_documents_pb.Document | undefined;
  setDocument(value?: resources_documents_documents_pb.Document): GetDocumentResponse;
  hasDocument(): boolean;
  clearDocument(): GetDocumentResponse;

  getAccess(): resources_documents_documents_pb.DocumentAccess | undefined;
  setAccess(value?: resources_documents_documents_pb.DocumentAccess): GetDocumentResponse;
  hasAccess(): boolean;
  clearAccess(): GetDocumentResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDocumentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetDocumentResponse): GetDocumentResponse.AsObject;
  static serializeBinaryToWriter(message: GetDocumentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDocumentResponse;
  static deserializeBinaryFromReader(message: GetDocumentResponse, reader: jspb.BinaryReader): GetDocumentResponse;
}

export namespace GetDocumentResponse {
  export type AsObject = {
    document?: resources_documents_documents_pb.Document.AsObject,
    access?: resources_documents_documents_pb.DocumentAccess.AsObject,
  }
}

export class GetDocumentReferencesRequest extends jspb.Message {
  getDocumentId(): number;
  setDocumentId(value: number): GetDocumentReferencesRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDocumentReferencesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetDocumentReferencesRequest): GetDocumentReferencesRequest.AsObject;
  static serializeBinaryToWriter(message: GetDocumentReferencesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDocumentReferencesRequest;
  static deserializeBinaryFromReader(message: GetDocumentReferencesRequest, reader: jspb.BinaryReader): GetDocumentReferencesRequest;
}

export namespace GetDocumentReferencesRequest {
  export type AsObject = {
    documentId: number,
  }
}

export class GetDocumentReferencesResponse extends jspb.Message {
  getReferencesList(): Array<resources_documents_documents_pb.DocumentReference>;
  setReferencesList(value: Array<resources_documents_documents_pb.DocumentReference>): GetDocumentReferencesResponse;
  clearReferencesList(): GetDocumentReferencesResponse;
  addReferences(value?: resources_documents_documents_pb.DocumentReference, index?: number): resources_documents_documents_pb.DocumentReference;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDocumentReferencesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetDocumentReferencesResponse): GetDocumentReferencesResponse.AsObject;
  static serializeBinaryToWriter(message: GetDocumentReferencesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDocumentReferencesResponse;
  static deserializeBinaryFromReader(message: GetDocumentReferencesResponse, reader: jspb.BinaryReader): GetDocumentReferencesResponse;
}

export namespace GetDocumentReferencesResponse {
  export type AsObject = {
    referencesList: Array<resources_documents_documents_pb.DocumentReference.AsObject>,
  }
}

export class GetDocumentRelationsRequest extends jspb.Message {
  getDocumentId(): number;
  setDocumentId(value: number): GetDocumentRelationsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDocumentRelationsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetDocumentRelationsRequest): GetDocumentRelationsRequest.AsObject;
  static serializeBinaryToWriter(message: GetDocumentRelationsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDocumentRelationsRequest;
  static deserializeBinaryFromReader(message: GetDocumentRelationsRequest, reader: jspb.BinaryReader): GetDocumentRelationsRequest;
}

export namespace GetDocumentRelationsRequest {
  export type AsObject = {
    documentId: number,
  }
}

export class GetDocumentRelationsResponse extends jspb.Message {
  getRelationsList(): Array<resources_documents_documents_pb.DocumentRelation>;
  setRelationsList(value: Array<resources_documents_documents_pb.DocumentRelation>): GetDocumentRelationsResponse;
  clearRelationsList(): GetDocumentRelationsResponse;
  addRelations(value?: resources_documents_documents_pb.DocumentRelation, index?: number): resources_documents_documents_pb.DocumentRelation;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDocumentRelationsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetDocumentRelationsResponse): GetDocumentRelationsResponse.AsObject;
  static serializeBinaryToWriter(message: GetDocumentRelationsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDocumentRelationsResponse;
  static deserializeBinaryFromReader(message: GetDocumentRelationsResponse, reader: jspb.BinaryReader): GetDocumentRelationsResponse;
}

export namespace GetDocumentRelationsResponse {
  export type AsObject = {
    relationsList: Array<resources_documents_documents_pb.DocumentRelation.AsObject>,
  }
}

export class AddDocumentReferenceRequest extends jspb.Message {
  getReference(): resources_documents_documents_pb.DocumentReference | undefined;
  setReference(value?: resources_documents_documents_pb.DocumentReference): AddDocumentReferenceRequest;
  hasReference(): boolean;
  clearReference(): AddDocumentReferenceRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddDocumentReferenceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddDocumentReferenceRequest): AddDocumentReferenceRequest.AsObject;
  static serializeBinaryToWriter(message: AddDocumentReferenceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddDocumentReferenceRequest;
  static deserializeBinaryFromReader(message: AddDocumentReferenceRequest, reader: jspb.BinaryReader): AddDocumentReferenceRequest;
}

export namespace AddDocumentReferenceRequest {
  export type AsObject = {
    reference?: resources_documents_documents_pb.DocumentReference.AsObject,
  }
}

export class AddDocumentReferenceResponse extends jspb.Message {
  getId(): number;
  setId(value: number): AddDocumentReferenceResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddDocumentReferenceResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddDocumentReferenceResponse): AddDocumentReferenceResponse.AsObject;
  static serializeBinaryToWriter(message: AddDocumentReferenceResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddDocumentReferenceResponse;
  static deserializeBinaryFromReader(message: AddDocumentReferenceResponse, reader: jspb.BinaryReader): AddDocumentReferenceResponse;
}

export namespace AddDocumentReferenceResponse {
  export type AsObject = {
    id: number,
  }
}

export class RemoveDocumentReferenceRequest extends jspb.Message {
  getId(): number;
  setId(value: number): RemoveDocumentReferenceRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemoveDocumentReferenceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RemoveDocumentReferenceRequest): RemoveDocumentReferenceRequest.AsObject;
  static serializeBinaryToWriter(message: RemoveDocumentReferenceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemoveDocumentReferenceRequest;
  static deserializeBinaryFromReader(message: RemoveDocumentReferenceRequest, reader: jspb.BinaryReader): RemoveDocumentReferenceRequest;
}

export namespace RemoveDocumentReferenceRequest {
  export type AsObject = {
    id: number,
  }
}

export class RemoveDocumentReferenceResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemoveDocumentReferenceResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RemoveDocumentReferenceResponse): RemoveDocumentReferenceResponse.AsObject;
  static serializeBinaryToWriter(message: RemoveDocumentReferenceResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemoveDocumentReferenceResponse;
  static deserializeBinaryFromReader(message: RemoveDocumentReferenceResponse, reader: jspb.BinaryReader): RemoveDocumentReferenceResponse;
}

export namespace RemoveDocumentReferenceResponse {
  export type AsObject = {
  }
}

export class AddDocumentRelationRequest extends jspb.Message {
  getRelation(): resources_documents_documents_pb.DocumentRelation | undefined;
  setRelation(value?: resources_documents_documents_pb.DocumentRelation): AddDocumentRelationRequest;
  hasRelation(): boolean;
  clearRelation(): AddDocumentRelationRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddDocumentRelationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddDocumentRelationRequest): AddDocumentRelationRequest.AsObject;
  static serializeBinaryToWriter(message: AddDocumentRelationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddDocumentRelationRequest;
  static deserializeBinaryFromReader(message: AddDocumentRelationRequest, reader: jspb.BinaryReader): AddDocumentRelationRequest;
}

export namespace AddDocumentRelationRequest {
  export type AsObject = {
    relation?: resources_documents_documents_pb.DocumentRelation.AsObject,
  }
}

export class AddDocumentRelationResponse extends jspb.Message {
  getId(): number;
  setId(value: number): AddDocumentRelationResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddDocumentRelationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddDocumentRelationResponse): AddDocumentRelationResponse.AsObject;
  static serializeBinaryToWriter(message: AddDocumentRelationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddDocumentRelationResponse;
  static deserializeBinaryFromReader(message: AddDocumentRelationResponse, reader: jspb.BinaryReader): AddDocumentRelationResponse;
}

export namespace AddDocumentRelationResponse {
  export type AsObject = {
    id: number,
  }
}

export class RemoveDocumentRelationRequest extends jspb.Message {
  getId(): number;
  setId(value: number): RemoveDocumentRelationRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemoveDocumentRelationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RemoveDocumentRelationRequest): RemoveDocumentRelationRequest.AsObject;
  static serializeBinaryToWriter(message: RemoveDocumentRelationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemoveDocumentRelationRequest;
  static deserializeBinaryFromReader(message: RemoveDocumentRelationRequest, reader: jspb.BinaryReader): RemoveDocumentRelationRequest;
}

export namespace RemoveDocumentRelationRequest {
  export type AsObject = {
    id: number,
  }
}

export class RemoveDocumentRelationResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemoveDocumentRelationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RemoveDocumentRelationResponse): RemoveDocumentRelationResponse.AsObject;
  static serializeBinaryToWriter(message: RemoveDocumentRelationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemoveDocumentRelationResponse;
  static deserializeBinaryFromReader(message: RemoveDocumentRelationResponse, reader: jspb.BinaryReader): RemoveDocumentRelationResponse;
}

export namespace RemoveDocumentRelationResponse {
  export type AsObject = {
  }
}

export class GetDocumentCommentsRequest extends jspb.Message {
  getPagination(): resources_common_database_database_pb.PaginationRequest | undefined;
  setPagination(value?: resources_common_database_database_pb.PaginationRequest): GetDocumentCommentsRequest;
  hasPagination(): boolean;
  clearPagination(): GetDocumentCommentsRequest;

  getDocumentId(): number;
  setDocumentId(value: number): GetDocumentCommentsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDocumentCommentsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetDocumentCommentsRequest): GetDocumentCommentsRequest.AsObject;
  static serializeBinaryToWriter(message: GetDocumentCommentsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDocumentCommentsRequest;
  static deserializeBinaryFromReader(message: GetDocumentCommentsRequest, reader: jspb.BinaryReader): GetDocumentCommentsRequest;
}

export namespace GetDocumentCommentsRequest {
  export type AsObject = {
    pagination?: resources_common_database_database_pb.PaginationRequest.AsObject,
    documentId: number,
  }
}

export class GetDocumentCommentsResponse extends jspb.Message {
  getPagination(): resources_common_database_database_pb.PaginationResponse | undefined;
  setPagination(value?: resources_common_database_database_pb.PaginationResponse): GetDocumentCommentsResponse;
  hasPagination(): boolean;
  clearPagination(): GetDocumentCommentsResponse;

  getCommentsList(): Array<resources_documents_documents_pb.DocumentComment>;
  setCommentsList(value: Array<resources_documents_documents_pb.DocumentComment>): GetDocumentCommentsResponse;
  clearCommentsList(): GetDocumentCommentsResponse;
  addComments(value?: resources_documents_documents_pb.DocumentComment, index?: number): resources_documents_documents_pb.DocumentComment;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDocumentCommentsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetDocumentCommentsResponse): GetDocumentCommentsResponse.AsObject;
  static serializeBinaryToWriter(message: GetDocumentCommentsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDocumentCommentsResponse;
  static deserializeBinaryFromReader(message: GetDocumentCommentsResponse, reader: jspb.BinaryReader): GetDocumentCommentsResponse;
}

export namespace GetDocumentCommentsResponse {
  export type AsObject = {
    pagination?: resources_common_database_database_pb.PaginationResponse.AsObject,
    commentsList: Array<resources_documents_documents_pb.DocumentComment.AsObject>,
  }
}

export class PostDocumentCommentRequest extends jspb.Message {
  getComment(): resources_documents_documents_pb.DocumentComment | undefined;
  setComment(value?: resources_documents_documents_pb.DocumentComment): PostDocumentCommentRequest;
  hasComment(): boolean;
  clearComment(): PostDocumentCommentRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PostDocumentCommentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PostDocumentCommentRequest): PostDocumentCommentRequest.AsObject;
  static serializeBinaryToWriter(message: PostDocumentCommentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PostDocumentCommentRequest;
  static deserializeBinaryFromReader(message: PostDocumentCommentRequest, reader: jspb.BinaryReader): PostDocumentCommentRequest;
}

export namespace PostDocumentCommentRequest {
  export type AsObject = {
    comment?: resources_documents_documents_pb.DocumentComment.AsObject,
  }
}

export class PostDocumentCommentResponse extends jspb.Message {
  getId(): number;
  setId(value: number): PostDocumentCommentResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PostDocumentCommentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PostDocumentCommentResponse): PostDocumentCommentResponse.AsObject;
  static serializeBinaryToWriter(message: PostDocumentCommentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PostDocumentCommentResponse;
  static deserializeBinaryFromReader(message: PostDocumentCommentResponse, reader: jspb.BinaryReader): PostDocumentCommentResponse;
}

export namespace PostDocumentCommentResponse {
  export type AsObject = {
    id: number,
  }
}

export class EditDocumentCommentRequest extends jspb.Message {
  getComment(): resources_documents_documents_pb.DocumentComment | undefined;
  setComment(value?: resources_documents_documents_pb.DocumentComment): EditDocumentCommentRequest;
  hasComment(): boolean;
  clearComment(): EditDocumentCommentRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EditDocumentCommentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: EditDocumentCommentRequest): EditDocumentCommentRequest.AsObject;
  static serializeBinaryToWriter(message: EditDocumentCommentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EditDocumentCommentRequest;
  static deserializeBinaryFromReader(message: EditDocumentCommentRequest, reader: jspb.BinaryReader): EditDocumentCommentRequest;
}

export namespace EditDocumentCommentRequest {
  export type AsObject = {
    comment?: resources_documents_documents_pb.DocumentComment.AsObject,
  }
}

export class EditDocumentCommentResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EditDocumentCommentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: EditDocumentCommentResponse): EditDocumentCommentResponse.AsObject;
  static serializeBinaryToWriter(message: EditDocumentCommentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EditDocumentCommentResponse;
  static deserializeBinaryFromReader(message: EditDocumentCommentResponse, reader: jspb.BinaryReader): EditDocumentCommentResponse;
}

export namespace EditDocumentCommentResponse {
  export type AsObject = {
  }
}

export class DeleteDocumentCommentRequest extends jspb.Message {
  getCommentId(): number;
  setCommentId(value: number): DeleteDocumentCommentRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteDocumentCommentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteDocumentCommentRequest): DeleteDocumentCommentRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteDocumentCommentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteDocumentCommentRequest;
  static deserializeBinaryFromReader(message: DeleteDocumentCommentRequest, reader: jspb.BinaryReader): DeleteDocumentCommentRequest;
}

export namespace DeleteDocumentCommentRequest {
  export type AsObject = {
    commentId: number,
  }
}

export class DeleteDocumentCommentResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteDocumentCommentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteDocumentCommentResponse): DeleteDocumentCommentResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteDocumentCommentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteDocumentCommentResponse;
  static deserializeBinaryFromReader(message: DeleteDocumentCommentResponse, reader: jspb.BinaryReader): DeleteDocumentCommentResponse;
}

export namespace DeleteDocumentCommentResponse {
  export type AsObject = {
  }
}

export class CreateDocumentRequest extends jspb.Message {
  getCategoryId(): number;
  setCategoryId(value: number): CreateDocumentRequest;
  hasCategoryId(): boolean;
  clearCategoryId(): CreateDocumentRequest;

  getTitle(): string;
  setTitle(value: string): CreateDocumentRequest;

  getContent(): string;
  setContent(value: string): CreateDocumentRequest;

  getContentType(): resources_documents_documents_pb.DOC_CONTENT_TYPE;
  setContentType(value: resources_documents_documents_pb.DOC_CONTENT_TYPE): CreateDocumentRequest;

  getData(): string;
  setData(value: string): CreateDocumentRequest;
  hasData(): boolean;
  clearData(): CreateDocumentRequest;

  getState(): string;
  setState(value: string): CreateDocumentRequest;

  getClosed(): boolean;
  setClosed(value: boolean): CreateDocumentRequest;

  getPublic(): boolean;
  setPublic(value: boolean): CreateDocumentRequest;

  getAccess(): resources_documents_documents_pb.DocumentAccess | undefined;
  setAccess(value?: resources_documents_documents_pb.DocumentAccess): CreateDocumentRequest;
  hasAccess(): boolean;
  clearAccess(): CreateDocumentRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateDocumentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateDocumentRequest): CreateDocumentRequest.AsObject;
  static serializeBinaryToWriter(message: CreateDocumentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateDocumentRequest;
  static deserializeBinaryFromReader(message: CreateDocumentRequest, reader: jspb.BinaryReader): CreateDocumentRequest;
}

export namespace CreateDocumentRequest {
  export type AsObject = {
    categoryId?: number,
    title: string,
    content: string,
    contentType: resources_documents_documents_pb.DOC_CONTENT_TYPE,
    data?: string,
    state: string,
    closed: boolean,
    pb_public: boolean,
    access?: resources_documents_documents_pb.DocumentAccess.AsObject,
  }

  export enum CategoryIdCase { 
    _CATEGORY_ID_NOT_SET = 0,
    CATEGORY_ID = 1,
  }

  export enum DataCase { 
    _DATA_NOT_SET = 0,
    DATA = 5,
  }

  export enum AccessCase { 
    _ACCESS_NOT_SET = 0,
    ACCESS = 9,
  }
}

export class CreateDocumentResponse extends jspb.Message {
  getDocumentId(): number;
  setDocumentId(value: number): CreateDocumentResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateDocumentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateDocumentResponse): CreateDocumentResponse.AsObject;
  static serializeBinaryToWriter(message: CreateDocumentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateDocumentResponse;
  static deserializeBinaryFromReader(message: CreateDocumentResponse, reader: jspb.BinaryReader): CreateDocumentResponse;
}

export namespace CreateDocumentResponse {
  export type AsObject = {
    documentId: number,
  }
}

export class UpdateDocumentRequest extends jspb.Message {
  getDocumentId(): number;
  setDocumentId(value: number): UpdateDocumentRequest;

  getCategoryId(): number;
  setCategoryId(value: number): UpdateDocumentRequest;
  hasCategoryId(): boolean;
  clearCategoryId(): UpdateDocumentRequest;

  getTitle(): string;
  setTitle(value: string): UpdateDocumentRequest;
  hasTitle(): boolean;
  clearTitle(): UpdateDocumentRequest;

  getContent(): string;
  setContent(value: string): UpdateDocumentRequest;
  hasContent(): boolean;
  clearContent(): UpdateDocumentRequest;

  getContentType(): resources_documents_documents_pb.DOC_CONTENT_TYPE;
  setContentType(value: resources_documents_documents_pb.DOC_CONTENT_TYPE): UpdateDocumentRequest;
  hasContentType(): boolean;
  clearContentType(): UpdateDocumentRequest;

  getData(): string;
  setData(value: string): UpdateDocumentRequest;
  hasData(): boolean;
  clearData(): UpdateDocumentRequest;

  getState(): string;
  setState(value: string): UpdateDocumentRequest;
  hasState(): boolean;
  clearState(): UpdateDocumentRequest;

  getClosed(): boolean;
  setClosed(value: boolean): UpdateDocumentRequest;
  hasClosed(): boolean;
  clearClosed(): UpdateDocumentRequest;

  getPublic(): boolean;
  setPublic(value: boolean): UpdateDocumentRequest;
  hasPublic(): boolean;
  clearPublic(): UpdateDocumentRequest;

  getAccess(): resources_documents_documents_pb.DocumentAccess | undefined;
  setAccess(value?: resources_documents_documents_pb.DocumentAccess): UpdateDocumentRequest;
  hasAccess(): boolean;
  clearAccess(): UpdateDocumentRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateDocumentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateDocumentRequest): UpdateDocumentRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateDocumentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateDocumentRequest;
  static deserializeBinaryFromReader(message: UpdateDocumentRequest, reader: jspb.BinaryReader): UpdateDocumentRequest;
}

export namespace UpdateDocumentRequest {
  export type AsObject = {
    documentId: number,
    categoryId?: number,
    title?: string,
    content?: string,
    contentType?: resources_documents_documents_pb.DOC_CONTENT_TYPE,
    data?: string,
    state?: string,
    closed?: boolean,
    pb_public?: boolean,
    access?: resources_documents_documents_pb.DocumentAccess.AsObject,
  }

  export enum CategoryIdCase { 
    _CATEGORY_ID_NOT_SET = 0,
    CATEGORY_ID = 2,
  }

  export enum TitleCase { 
    _TITLE_NOT_SET = 0,
    TITLE = 3,
  }

  export enum ContentCase { 
    _CONTENT_NOT_SET = 0,
    CONTENT = 4,
  }

  export enum ContentTypeCase { 
    _CONTENT_TYPE_NOT_SET = 0,
    CONTENT_TYPE = 5,
  }

  export enum DataCase { 
    _DATA_NOT_SET = 0,
    DATA = 6,
  }

  export enum StateCase { 
    _STATE_NOT_SET = 0,
    STATE = 7,
  }

  export enum ClosedCase { 
    _CLOSED_NOT_SET = 0,
    CLOSED = 8,
  }

  export enum PublicCase { 
    _PUBLIC_NOT_SET = 0,
    PUBLIC = 9,
  }

  export enum AccessCase { 
    _ACCESS_NOT_SET = 0,
    ACCESS = 10,
  }
}

export class UpdateDocumentResponse extends jspb.Message {
  getDocumentId(): number;
  setDocumentId(value: number): UpdateDocumentResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateDocumentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateDocumentResponse): UpdateDocumentResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateDocumentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateDocumentResponse;
  static deserializeBinaryFromReader(message: UpdateDocumentResponse, reader: jspb.BinaryReader): UpdateDocumentResponse;
}

export namespace UpdateDocumentResponse {
  export type AsObject = {
    documentId: number,
  }
}

export class DeleteDocumentRequest extends jspb.Message {
  getDocumentId(): number;
  setDocumentId(value: number): DeleteDocumentRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteDocumentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteDocumentRequest): DeleteDocumentRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteDocumentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteDocumentRequest;
  static deserializeBinaryFromReader(message: DeleteDocumentRequest, reader: jspb.BinaryReader): DeleteDocumentRequest;
}

export namespace DeleteDocumentRequest {
  export type AsObject = {
    documentId: number,
  }
}

export class DeleteDocumentResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteDocumentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteDocumentResponse): DeleteDocumentResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteDocumentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteDocumentResponse;
  static deserializeBinaryFromReader(message: DeleteDocumentResponse, reader: jspb.BinaryReader): DeleteDocumentResponse;
}

export namespace DeleteDocumentResponse {
  export type AsObject = {
  }
}

export class GetDocumentAccessRequest extends jspb.Message {
  getDocumentId(): number;
  setDocumentId(value: number): GetDocumentAccessRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDocumentAccessRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetDocumentAccessRequest): GetDocumentAccessRequest.AsObject;
  static serializeBinaryToWriter(message: GetDocumentAccessRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDocumentAccessRequest;
  static deserializeBinaryFromReader(message: GetDocumentAccessRequest, reader: jspb.BinaryReader): GetDocumentAccessRequest;
}

export namespace GetDocumentAccessRequest {
  export type AsObject = {
    documentId: number,
  }
}

export class GetDocumentAccessResponse extends jspb.Message {
  getAccess(): resources_documents_documents_pb.DocumentAccess | undefined;
  setAccess(value?: resources_documents_documents_pb.DocumentAccess): GetDocumentAccessResponse;
  hasAccess(): boolean;
  clearAccess(): GetDocumentAccessResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDocumentAccessResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetDocumentAccessResponse): GetDocumentAccessResponse.AsObject;
  static serializeBinaryToWriter(message: GetDocumentAccessResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDocumentAccessResponse;
  static deserializeBinaryFromReader(message: GetDocumentAccessResponse, reader: jspb.BinaryReader): GetDocumentAccessResponse;
}

export namespace GetDocumentAccessResponse {
  export type AsObject = {
    access?: resources_documents_documents_pb.DocumentAccess.AsObject,
  }
}

export class SetDocumentAccessRequest extends jspb.Message {
  getDocumentId(): number;
  setDocumentId(value: number): SetDocumentAccessRequest;

  getMode(): ACCESS_LEVEL_UPDATE_MODE;
  setMode(value: ACCESS_LEVEL_UPDATE_MODE): SetDocumentAccessRequest;

  getAccess(): resources_documents_documents_pb.DocumentAccess | undefined;
  setAccess(value?: resources_documents_documents_pb.DocumentAccess): SetDocumentAccessRequest;
  hasAccess(): boolean;
  clearAccess(): SetDocumentAccessRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SetDocumentAccessRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SetDocumentAccessRequest): SetDocumentAccessRequest.AsObject;
  static serializeBinaryToWriter(message: SetDocumentAccessRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SetDocumentAccessRequest;
  static deserializeBinaryFromReader(message: SetDocumentAccessRequest, reader: jspb.BinaryReader): SetDocumentAccessRequest;
}

export namespace SetDocumentAccessRequest {
  export type AsObject = {
    documentId: number,
    mode: ACCESS_LEVEL_UPDATE_MODE,
    access?: resources_documents_documents_pb.DocumentAccess.AsObject,
  }
}

export class SetDocumentAccessResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SetDocumentAccessResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SetDocumentAccessResponse): SetDocumentAccessResponse.AsObject;
  static serializeBinaryToWriter(message: SetDocumentAccessResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SetDocumentAccessResponse;
  static deserializeBinaryFromReader(message: SetDocumentAccessResponse, reader: jspb.BinaryReader): SetDocumentAccessResponse;
}

export namespace SetDocumentAccessResponse {
  export type AsObject = {
  }
}

export class ListUserDocumentsRequest extends jspb.Message {
  getPagination(): resources_common_database_database_pb.PaginationRequest | undefined;
  setPagination(value?: resources_common_database_database_pb.PaginationRequest): ListUserDocumentsRequest;
  hasPagination(): boolean;
  clearPagination(): ListUserDocumentsRequest;

  getUserId(): number;
  setUserId(value: number): ListUserDocumentsRequest;

  getRelationsList(): Array<resources_documents_documents_pb.DOC_RELATION>;
  setRelationsList(value: Array<resources_documents_documents_pb.DOC_RELATION>): ListUserDocumentsRequest;
  clearRelationsList(): ListUserDocumentsRequest;
  addRelations(value: resources_documents_documents_pb.DOC_RELATION, index?: number): ListUserDocumentsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListUserDocumentsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListUserDocumentsRequest): ListUserDocumentsRequest.AsObject;
  static serializeBinaryToWriter(message: ListUserDocumentsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListUserDocumentsRequest;
  static deserializeBinaryFromReader(message: ListUserDocumentsRequest, reader: jspb.BinaryReader): ListUserDocumentsRequest;
}

export namespace ListUserDocumentsRequest {
  export type AsObject = {
    pagination?: resources_common_database_database_pb.PaginationRequest.AsObject,
    userId: number,
    relationsList: Array<resources_documents_documents_pb.DOC_RELATION>,
  }
}

export class ListUserDocumentsResponse extends jspb.Message {
  getPagination(): resources_common_database_database_pb.PaginationResponse | undefined;
  setPagination(value?: resources_common_database_database_pb.PaginationResponse): ListUserDocumentsResponse;
  hasPagination(): boolean;
  clearPagination(): ListUserDocumentsResponse;

  getRelationsList(): Array<resources_documents_documents_pb.DocumentRelation>;
  setRelationsList(value: Array<resources_documents_documents_pb.DocumentRelation>): ListUserDocumentsResponse;
  clearRelationsList(): ListUserDocumentsResponse;
  addRelations(value?: resources_documents_documents_pb.DocumentRelation, index?: number): resources_documents_documents_pb.DocumentRelation;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListUserDocumentsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListUserDocumentsResponse): ListUserDocumentsResponse.AsObject;
  static serializeBinaryToWriter(message: ListUserDocumentsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListUserDocumentsResponse;
  static deserializeBinaryFromReader(message: ListUserDocumentsResponse, reader: jspb.BinaryReader): ListUserDocumentsResponse;
}

export namespace ListUserDocumentsResponse {
  export type AsObject = {
    pagination?: resources_common_database_database_pb.PaginationResponse.AsObject,
    relationsList: Array<resources_documents_documents_pb.DocumentRelation.AsObject>,
  }
}

export class ListDocumentCategoriesRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListDocumentCategoriesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListDocumentCategoriesRequest): ListDocumentCategoriesRequest.AsObject;
  static serializeBinaryToWriter(message: ListDocumentCategoriesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListDocumentCategoriesRequest;
  static deserializeBinaryFromReader(message: ListDocumentCategoriesRequest, reader: jspb.BinaryReader): ListDocumentCategoriesRequest;
}

export namespace ListDocumentCategoriesRequest {
  export type AsObject = {
  }
}

export class ListDocumentCategoriesResponse extends jspb.Message {
  getCategoryList(): Array<resources_documents_category_pb.DocumentCategory>;
  setCategoryList(value: Array<resources_documents_category_pb.DocumentCategory>): ListDocumentCategoriesResponse;
  clearCategoryList(): ListDocumentCategoriesResponse;
  addCategory(value?: resources_documents_category_pb.DocumentCategory, index?: number): resources_documents_category_pb.DocumentCategory;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListDocumentCategoriesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListDocumentCategoriesResponse): ListDocumentCategoriesResponse.AsObject;
  static serializeBinaryToWriter(message: ListDocumentCategoriesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListDocumentCategoriesResponse;
  static deserializeBinaryFromReader(message: ListDocumentCategoriesResponse, reader: jspb.BinaryReader): ListDocumentCategoriesResponse;
}

export namespace ListDocumentCategoriesResponse {
  export type AsObject = {
    categoryList: Array<resources_documents_category_pb.DocumentCategory.AsObject>,
  }
}

export class CreateDocumentCategoryRequest extends jspb.Message {
  getCategory(): resources_documents_category_pb.DocumentCategory | undefined;
  setCategory(value?: resources_documents_category_pb.DocumentCategory): CreateDocumentCategoryRequest;
  hasCategory(): boolean;
  clearCategory(): CreateDocumentCategoryRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateDocumentCategoryRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateDocumentCategoryRequest): CreateDocumentCategoryRequest.AsObject;
  static serializeBinaryToWriter(message: CreateDocumentCategoryRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateDocumentCategoryRequest;
  static deserializeBinaryFromReader(message: CreateDocumentCategoryRequest, reader: jspb.BinaryReader): CreateDocumentCategoryRequest;
}

export namespace CreateDocumentCategoryRequest {
  export type AsObject = {
    category?: resources_documents_category_pb.DocumentCategory.AsObject,
  }
}

export class CreateDocumentCategoryResponse extends jspb.Message {
  getId(): number;
  setId(value: number): CreateDocumentCategoryResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateDocumentCategoryResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateDocumentCategoryResponse): CreateDocumentCategoryResponse.AsObject;
  static serializeBinaryToWriter(message: CreateDocumentCategoryResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateDocumentCategoryResponse;
  static deserializeBinaryFromReader(message: CreateDocumentCategoryResponse, reader: jspb.BinaryReader): CreateDocumentCategoryResponse;
}

export namespace CreateDocumentCategoryResponse {
  export type AsObject = {
    id: number,
  }
}

export class UpdateDocumentCategoryRequest extends jspb.Message {
  getCategory(): resources_documents_category_pb.DocumentCategory | undefined;
  setCategory(value?: resources_documents_category_pb.DocumentCategory): UpdateDocumentCategoryRequest;
  hasCategory(): boolean;
  clearCategory(): UpdateDocumentCategoryRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateDocumentCategoryRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateDocumentCategoryRequest): UpdateDocumentCategoryRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateDocumentCategoryRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateDocumentCategoryRequest;
  static deserializeBinaryFromReader(message: UpdateDocumentCategoryRequest, reader: jspb.BinaryReader): UpdateDocumentCategoryRequest;
}

export namespace UpdateDocumentCategoryRequest {
  export type AsObject = {
    category?: resources_documents_category_pb.DocumentCategory.AsObject,
  }
}

export class UpdateDocumentCategoryResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateDocumentCategoryResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateDocumentCategoryResponse): UpdateDocumentCategoryResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateDocumentCategoryResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateDocumentCategoryResponse;
  static deserializeBinaryFromReader(message: UpdateDocumentCategoryResponse, reader: jspb.BinaryReader): UpdateDocumentCategoryResponse;
}

export namespace UpdateDocumentCategoryResponse {
  export type AsObject = {
  }
}

export class DeleteDocumentCategoryRequest extends jspb.Message {
  getIdsList(): Array<number>;
  setIdsList(value: Array<number>): DeleteDocumentCategoryRequest;
  clearIdsList(): DeleteDocumentCategoryRequest;
  addIds(value: number, index?: number): DeleteDocumentCategoryRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteDocumentCategoryRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteDocumentCategoryRequest): DeleteDocumentCategoryRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteDocumentCategoryRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteDocumentCategoryRequest;
  static deserializeBinaryFromReader(message: DeleteDocumentCategoryRequest, reader: jspb.BinaryReader): DeleteDocumentCategoryRequest;
}

export namespace DeleteDocumentCategoryRequest {
  export type AsObject = {
    idsList: Array<number>,
  }
}

export class DeleteDocumentCategoryResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteDocumentCategoryResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteDocumentCategoryResponse): DeleteDocumentCategoryResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteDocumentCategoryResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteDocumentCategoryResponse;
  static deserializeBinaryFromReader(message: DeleteDocumentCategoryResponse, reader: jspb.BinaryReader): DeleteDocumentCategoryResponse;
}

export namespace DeleteDocumentCategoryResponse {
  export type AsObject = {
  }
}

export enum ACCESS_LEVEL_UPDATE_MODE { 
  UPDATE = 0,
  DELETE = 1,
  CLEAR = 2,
}
