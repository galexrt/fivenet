/**
 * @fileoverview gRPC-Web generated client stub for services.documents
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v3.21.12
// source: services/documents/documents.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as services_documents_documents_pb from '../../services/documents/documents_pb';


export class DocumentsServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname.replace(/\/+$/, '');
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorFindDocuments = new grpcWeb.MethodDescriptor(
    '/services.documents.DocumentsService/FindDocuments',
    grpcWeb.MethodType.UNARY,
    services_documents_documents_pb.FindDocumentsRequest,
    services_documents_documents_pb.FindDocumentsResponse,
    (request: services_documents_documents_pb.FindDocumentsRequest) => {
      return request.serializeBinary();
    },
    services_documents_documents_pb.FindDocumentsResponse.deserializeBinary
  );

  findDocuments(
    request: services_documents_documents_pb.FindDocumentsRequest,
    metadata: grpcWeb.Metadata | null): Promise<services_documents_documents_pb.FindDocumentsResponse>;

  findDocuments(
    request: services_documents_documents_pb.FindDocumentsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: services_documents_documents_pb.FindDocumentsResponse) => void): grpcWeb.ClientReadableStream<services_documents_documents_pb.FindDocumentsResponse>;

  findDocuments(
    request: services_documents_documents_pb.FindDocumentsRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: services_documents_documents_pb.FindDocumentsResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/services.documents.DocumentsService/FindDocuments',
        request,
        metadata || {},
        this.methodDescriptorFindDocuments,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/services.documents.DocumentsService/FindDocuments',
    request,
    metadata || {},
    this.methodDescriptorFindDocuments);
  }

  methodDescriptorGetDocument = new grpcWeb.MethodDescriptor(
    '/services.documents.DocumentsService/GetDocument',
    grpcWeb.MethodType.UNARY,
    services_documents_documents_pb.GetDocumentRequest,
    services_documents_documents_pb.GetDocumentResponse,
    (request: services_documents_documents_pb.GetDocumentRequest) => {
      return request.serializeBinary();
    },
    services_documents_documents_pb.GetDocumentResponse.deserializeBinary
  );

  getDocument(
    request: services_documents_documents_pb.GetDocumentRequest,
    metadata: grpcWeb.Metadata | null): Promise<services_documents_documents_pb.GetDocumentResponse>;

  getDocument(
    request: services_documents_documents_pb.GetDocumentRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: services_documents_documents_pb.GetDocumentResponse) => void): grpcWeb.ClientReadableStream<services_documents_documents_pb.GetDocumentResponse>;

  getDocument(
    request: services_documents_documents_pb.GetDocumentRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: services_documents_documents_pb.GetDocumentResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/services.documents.DocumentsService/GetDocument',
        request,
        metadata || {},
        this.methodDescriptorGetDocument,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/services.documents.DocumentsService/GetDocument',
    request,
    metadata || {},
    this.methodDescriptorGetDocument);
  }

  methodDescriptorCreateDocument = new grpcWeb.MethodDescriptor(
    '/services.documents.DocumentsService/CreateDocument',
    grpcWeb.MethodType.UNARY,
    services_documents_documents_pb.CreateDocumentRequest,
    services_documents_documents_pb.CreateDocumentResponse,
    (request: services_documents_documents_pb.CreateDocumentRequest) => {
      return request.serializeBinary();
    },
    services_documents_documents_pb.CreateDocumentResponse.deserializeBinary
  );

  createDocument(
    request: services_documents_documents_pb.CreateDocumentRequest,
    metadata: grpcWeb.Metadata | null): Promise<services_documents_documents_pb.CreateDocumentResponse>;

  createDocument(
    request: services_documents_documents_pb.CreateDocumentRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: services_documents_documents_pb.CreateDocumentResponse) => void): grpcWeb.ClientReadableStream<services_documents_documents_pb.CreateDocumentResponse>;

  createDocument(
    request: services_documents_documents_pb.CreateDocumentRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: services_documents_documents_pb.CreateDocumentResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/services.documents.DocumentsService/CreateDocument',
        request,
        metadata || {},
        this.methodDescriptorCreateDocument,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/services.documents.DocumentsService/CreateDocument',
    request,
    metadata || {},
    this.methodDescriptorCreateDocument);
  }

  methodDescriptorUpdateDocument = new grpcWeb.MethodDescriptor(
    '/services.documents.DocumentsService/UpdateDocument',
    grpcWeb.MethodType.UNARY,
    services_documents_documents_pb.UpdateDocumentRequest,
    services_documents_documents_pb.UpdateDocumentResponse,
    (request: services_documents_documents_pb.UpdateDocumentRequest) => {
      return request.serializeBinary();
    },
    services_documents_documents_pb.UpdateDocumentResponse.deserializeBinary
  );

  updateDocument(
    request: services_documents_documents_pb.UpdateDocumentRequest,
    metadata: grpcWeb.Metadata | null): Promise<services_documents_documents_pb.UpdateDocumentResponse>;

  updateDocument(
    request: services_documents_documents_pb.UpdateDocumentRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: services_documents_documents_pb.UpdateDocumentResponse) => void): grpcWeb.ClientReadableStream<services_documents_documents_pb.UpdateDocumentResponse>;

  updateDocument(
    request: services_documents_documents_pb.UpdateDocumentRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: services_documents_documents_pb.UpdateDocumentResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/services.documents.DocumentsService/UpdateDocument',
        request,
        metadata || {},
        this.methodDescriptorUpdateDocument,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/services.documents.DocumentsService/UpdateDocument',
    request,
    metadata || {},
    this.methodDescriptorUpdateDocument);
  }

  methodDescriptorGetDocumentResponses = new grpcWeb.MethodDescriptor(
    '/services.documents.DocumentsService/GetDocumentResponses',
    grpcWeb.MethodType.UNARY,
    services_documents_documents_pb.GetDocumentResponsesRequest,
    services_documents_documents_pb.GetDocumentResponsesResponse,
    (request: services_documents_documents_pb.GetDocumentResponsesRequest) => {
      return request.serializeBinary();
    },
    services_documents_documents_pb.GetDocumentResponsesResponse.deserializeBinary
  );

  getDocumentResponses(
    request: services_documents_documents_pb.GetDocumentResponsesRequest,
    metadata: grpcWeb.Metadata | null): Promise<services_documents_documents_pb.GetDocumentResponsesResponse>;

  getDocumentResponses(
    request: services_documents_documents_pb.GetDocumentResponsesRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: services_documents_documents_pb.GetDocumentResponsesResponse) => void): grpcWeb.ClientReadableStream<services_documents_documents_pb.GetDocumentResponsesResponse>;

  getDocumentResponses(
    request: services_documents_documents_pb.GetDocumentResponsesRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: services_documents_documents_pb.GetDocumentResponsesResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/services.documents.DocumentsService/GetDocumentResponses',
        request,
        metadata || {},
        this.methodDescriptorGetDocumentResponses,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/services.documents.DocumentsService/GetDocumentResponses',
    request,
    metadata || {},
    this.methodDescriptorGetDocumentResponses);
  }

  methodDescriptorListTemplates = new grpcWeb.MethodDescriptor(
    '/services.documents.DocumentsService/ListTemplates',
    grpcWeb.MethodType.UNARY,
    services_documents_documents_pb.ListTemplatesRequest,
    services_documents_documents_pb.ListTemplatesResponse,
    (request: services_documents_documents_pb.ListTemplatesRequest) => {
      return request.serializeBinary();
    },
    services_documents_documents_pb.ListTemplatesResponse.deserializeBinary
  );

  listTemplates(
    request: services_documents_documents_pb.ListTemplatesRequest,
    metadata: grpcWeb.Metadata | null): Promise<services_documents_documents_pb.ListTemplatesResponse>;

  listTemplates(
    request: services_documents_documents_pb.ListTemplatesRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: services_documents_documents_pb.ListTemplatesResponse) => void): grpcWeb.ClientReadableStream<services_documents_documents_pb.ListTemplatesResponse>;

  listTemplates(
    request: services_documents_documents_pb.ListTemplatesRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: services_documents_documents_pb.ListTemplatesResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/services.documents.DocumentsService/ListTemplates',
        request,
        metadata || {},
        this.methodDescriptorListTemplates,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/services.documents.DocumentsService/ListTemplates',
    request,
    metadata || {},
    this.methodDescriptorListTemplates);
  }

  methodDescriptorGetTemplate = new grpcWeb.MethodDescriptor(
    '/services.documents.DocumentsService/GetTemplate',
    grpcWeb.MethodType.UNARY,
    services_documents_documents_pb.GetTemplateRequest,
    services_documents_documents_pb.GetTemplateResponse,
    (request: services_documents_documents_pb.GetTemplateRequest) => {
      return request.serializeBinary();
    },
    services_documents_documents_pb.GetTemplateResponse.deserializeBinary
  );

  getTemplate(
    request: services_documents_documents_pb.GetTemplateRequest,
    metadata: grpcWeb.Metadata | null): Promise<services_documents_documents_pb.GetTemplateResponse>;

  getTemplate(
    request: services_documents_documents_pb.GetTemplateRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: services_documents_documents_pb.GetTemplateResponse) => void): grpcWeb.ClientReadableStream<services_documents_documents_pb.GetTemplateResponse>;

  getTemplate(
    request: services_documents_documents_pb.GetTemplateRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: services_documents_documents_pb.GetTemplateResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/services.documents.DocumentsService/GetTemplate',
        request,
        metadata || {},
        this.methodDescriptorGetTemplate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/services.documents.DocumentsService/GetTemplate',
    request,
    metadata || {},
    this.methodDescriptorGetTemplate);
  }

  methodDescriptorGetDocumentAccess = new grpcWeb.MethodDescriptor(
    '/services.documents.DocumentsService/GetDocumentAccess',
    grpcWeb.MethodType.UNARY,
    services_documents_documents_pb.GetDocumentAccessRequest,
    services_documents_documents_pb.GetDocumentAccessResponse,
    (request: services_documents_documents_pb.GetDocumentAccessRequest) => {
      return request.serializeBinary();
    },
    services_documents_documents_pb.GetDocumentAccessResponse.deserializeBinary
  );

  getDocumentAccess(
    request: services_documents_documents_pb.GetDocumentAccessRequest,
    metadata: grpcWeb.Metadata | null): Promise<services_documents_documents_pb.GetDocumentAccessResponse>;

  getDocumentAccess(
    request: services_documents_documents_pb.GetDocumentAccessRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: services_documents_documents_pb.GetDocumentAccessResponse) => void): grpcWeb.ClientReadableStream<services_documents_documents_pb.GetDocumentAccessResponse>;

  getDocumentAccess(
    request: services_documents_documents_pb.GetDocumentAccessRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: services_documents_documents_pb.GetDocumentAccessResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/services.documents.DocumentsService/GetDocumentAccess',
        request,
        metadata || {},
        this.methodDescriptorGetDocumentAccess,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/services.documents.DocumentsService/GetDocumentAccess',
    request,
    metadata || {},
    this.methodDescriptorGetDocumentAccess);
  }

  methodDescriptorSetDocumentAccess = new grpcWeb.MethodDescriptor(
    '/services.documents.DocumentsService/SetDocumentAccess',
    grpcWeb.MethodType.UNARY,
    services_documents_documents_pb.SetDocumentAccessRequest,
    services_documents_documents_pb.SetDocumentAccessResponse,
    (request: services_documents_documents_pb.SetDocumentAccessRequest) => {
      return request.serializeBinary();
    },
    services_documents_documents_pb.SetDocumentAccessResponse.deserializeBinary
  );

  setDocumentAccess(
    request: services_documents_documents_pb.SetDocumentAccessRequest,
    metadata: grpcWeb.Metadata | null): Promise<services_documents_documents_pb.SetDocumentAccessResponse>;

  setDocumentAccess(
    request: services_documents_documents_pb.SetDocumentAccessRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: services_documents_documents_pb.SetDocumentAccessResponse) => void): grpcWeb.ClientReadableStream<services_documents_documents_pb.SetDocumentAccessResponse>;

  setDocumentAccess(
    request: services_documents_documents_pb.SetDocumentAccessRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: services_documents_documents_pb.SetDocumentAccessResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/services.documents.DocumentsService/SetDocumentAccess',
        request,
        metadata || {},
        this.methodDescriptorSetDocumentAccess,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/services.documents.DocumentsService/SetDocumentAccess',
    request,
    metadata || {},
    this.methodDescriptorSetDocumentAccess);
  }

}

