import * as jspb from 'google-protobuf'

import * as resources_permissions_permissions_pb from '../../resources/permissions/permissions_pb';
import * as resources_rector_audit_pb from '../../resources/rector/audit_pb';


export class GetRolesRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetRolesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetRolesRequest): GetRolesRequest.AsObject;
  static serializeBinaryToWriter(message: GetRolesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetRolesRequest;
  static deserializeBinaryFromReader(message: GetRolesRequest, reader: jspb.BinaryReader): GetRolesRequest;
}

export namespace GetRolesRequest {
  export type AsObject = {
  }
}

export class GetRolesResponse extends jspb.Message {
  getRolesList(): Array<resources_permissions_permissions_pb.Role>;
  setRolesList(value: Array<resources_permissions_permissions_pb.Role>): GetRolesResponse;
  clearRolesList(): GetRolesResponse;
  addRoles(value?: resources_permissions_permissions_pb.Role, index?: number): resources_permissions_permissions_pb.Role;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetRolesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetRolesResponse): GetRolesResponse.AsObject;
  static serializeBinaryToWriter(message: GetRolesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetRolesResponse;
  static deserializeBinaryFromReader(message: GetRolesResponse, reader: jspb.BinaryReader): GetRolesResponse;
}

export namespace GetRolesResponse {
  export type AsObject = {
    rolesList: Array<resources_permissions_permissions_pb.Role.AsObject>,
  }
}

export class GetRoleRequest extends jspb.Message {
  getId(): number;
  setId(value: number): GetRoleRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetRoleRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetRoleRequest): GetRoleRequest.AsObject;
  static serializeBinaryToWriter(message: GetRoleRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetRoleRequest;
  static deserializeBinaryFromReader(message: GetRoleRequest, reader: jspb.BinaryReader): GetRoleRequest;
}

export namespace GetRoleRequest {
  export type AsObject = {
    id: number,
  }
}

export class GetRoleResponse extends jspb.Message {
  getRole(): resources_permissions_permissions_pb.Role | undefined;
  setRole(value?: resources_permissions_permissions_pb.Role): GetRoleResponse;
  hasRole(): boolean;
  clearRole(): GetRoleResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetRoleResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetRoleResponse): GetRoleResponse.AsObject;
  static serializeBinaryToWriter(message: GetRoleResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetRoleResponse;
  static deserializeBinaryFromReader(message: GetRoleResponse, reader: jspb.BinaryReader): GetRoleResponse;
}

export namespace GetRoleResponse {
  export type AsObject = {
    role?: resources_permissions_permissions_pb.Role.AsObject,
  }
}

export class CreateRoleRequest extends jspb.Message {
  getJob(): string;
  setJob(value: string): CreateRoleRequest;

  getGrade(): number;
  setGrade(value: number): CreateRoleRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateRoleRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateRoleRequest): CreateRoleRequest.AsObject;
  static serializeBinaryToWriter(message: CreateRoleRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateRoleRequest;
  static deserializeBinaryFromReader(message: CreateRoleRequest, reader: jspb.BinaryReader): CreateRoleRequest;
}

export namespace CreateRoleRequest {
  export type AsObject = {
    job: string,
    grade: number,
  }
}

export class CreateRoleResponse extends jspb.Message {
  getRole(): resources_permissions_permissions_pb.Role | undefined;
  setRole(value?: resources_permissions_permissions_pb.Role): CreateRoleResponse;
  hasRole(): boolean;
  clearRole(): CreateRoleResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateRoleResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateRoleResponse): CreateRoleResponse.AsObject;
  static serializeBinaryToWriter(message: CreateRoleResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateRoleResponse;
  static deserializeBinaryFromReader(message: CreateRoleResponse, reader: jspb.BinaryReader): CreateRoleResponse;
}

export namespace CreateRoleResponse {
  export type AsObject = {
    role?: resources_permissions_permissions_pb.Role.AsObject,
  }
}

export class DeleteRoleRequest extends jspb.Message {
  getId(): number;
  setId(value: number): DeleteRoleRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteRoleRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteRoleRequest): DeleteRoleRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteRoleRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteRoleRequest;
  static deserializeBinaryFromReader(message: DeleteRoleRequest, reader: jspb.BinaryReader): DeleteRoleRequest;
}

export namespace DeleteRoleRequest {
  export type AsObject = {
    id: number,
  }
}

export class DeleteRoleResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteRoleResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteRoleResponse): DeleteRoleResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteRoleResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteRoleResponse;
  static deserializeBinaryFromReader(message: DeleteRoleResponse, reader: jspb.BinaryReader): DeleteRoleResponse;
}

export namespace DeleteRoleResponse {
  export type AsObject = {
  }
}

export class AddPermToRoleRequest extends jspb.Message {
  getId(): number;
  setId(value: number): AddPermToRoleRequest;

  getPermissionsList(): Array<number>;
  setPermissionsList(value: Array<number>): AddPermToRoleRequest;
  clearPermissionsList(): AddPermToRoleRequest;
  addPermissions(value: number, index?: number): AddPermToRoleRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddPermToRoleRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddPermToRoleRequest): AddPermToRoleRequest.AsObject;
  static serializeBinaryToWriter(message: AddPermToRoleRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddPermToRoleRequest;
  static deserializeBinaryFromReader(message: AddPermToRoleRequest, reader: jspb.BinaryReader): AddPermToRoleRequest;
}

export namespace AddPermToRoleRequest {
  export type AsObject = {
    id: number,
    permissionsList: Array<number>,
  }
}

export class AddPermToRoleResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddPermToRoleResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddPermToRoleResponse): AddPermToRoleResponse.AsObject;
  static serializeBinaryToWriter(message: AddPermToRoleResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddPermToRoleResponse;
  static deserializeBinaryFromReader(message: AddPermToRoleResponse, reader: jspb.BinaryReader): AddPermToRoleResponse;
}

export namespace AddPermToRoleResponse {
  export type AsObject = {
  }
}

export class RemovePermFromRoleRequest extends jspb.Message {
  getId(): number;
  setId(value: number): RemovePermFromRoleRequest;

  getPermissionsList(): Array<number>;
  setPermissionsList(value: Array<number>): RemovePermFromRoleRequest;
  clearPermissionsList(): RemovePermFromRoleRequest;
  addPermissions(value: number, index?: number): RemovePermFromRoleRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemovePermFromRoleRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RemovePermFromRoleRequest): RemovePermFromRoleRequest.AsObject;
  static serializeBinaryToWriter(message: RemovePermFromRoleRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemovePermFromRoleRequest;
  static deserializeBinaryFromReader(message: RemovePermFromRoleRequest, reader: jspb.BinaryReader): RemovePermFromRoleRequest;
}

export namespace RemovePermFromRoleRequest {
  export type AsObject = {
    id: number,
    permissionsList: Array<number>,
  }
}

export class RemovePermFromRoleResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemovePermFromRoleResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RemovePermFromRoleResponse): RemovePermFromRoleResponse.AsObject;
  static serializeBinaryToWriter(message: RemovePermFromRoleResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemovePermFromRoleResponse;
  static deserializeBinaryFromReader(message: RemovePermFromRoleResponse, reader: jspb.BinaryReader): RemovePermFromRoleResponse;
}

export namespace RemovePermFromRoleResponse {
  export type AsObject = {
  }
}

export class GetPermissionsRequest extends jspb.Message {
  getSearch(): string;
  setSearch(value: string): GetPermissionsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetPermissionsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetPermissionsRequest): GetPermissionsRequest.AsObject;
  static serializeBinaryToWriter(message: GetPermissionsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetPermissionsRequest;
  static deserializeBinaryFromReader(message: GetPermissionsRequest, reader: jspb.BinaryReader): GetPermissionsRequest;
}

export namespace GetPermissionsRequest {
  export type AsObject = {
    search: string,
  }
}

export class GetPermissionsResponse extends jspb.Message {
  getPermissionsList(): Array<resources_permissions_permissions_pb.Permission>;
  setPermissionsList(value: Array<resources_permissions_permissions_pb.Permission>): GetPermissionsResponse;
  clearPermissionsList(): GetPermissionsResponse;
  addPermissions(value?: resources_permissions_permissions_pb.Permission, index?: number): resources_permissions_permissions_pb.Permission;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetPermissionsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetPermissionsResponse): GetPermissionsResponse.AsObject;
  static serializeBinaryToWriter(message: GetPermissionsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetPermissionsResponse;
  static deserializeBinaryFromReader(message: GetPermissionsResponse, reader: jspb.BinaryReader): GetPermissionsResponse;
}

export namespace GetPermissionsResponse {
  export type AsObject = {
    permissionsList: Array<resources_permissions_permissions_pb.Permission.AsObject>,
  }
}

export class ViewAuditLogRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ViewAuditLogRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ViewAuditLogRequest): ViewAuditLogRequest.AsObject;
  static serializeBinaryToWriter(message: ViewAuditLogRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ViewAuditLogRequest;
  static deserializeBinaryFromReader(message: ViewAuditLogRequest, reader: jspb.BinaryReader): ViewAuditLogRequest;
}

export namespace ViewAuditLogRequest {
  export type AsObject = {
  }
}

export class ViewAuditLogResponse extends jspb.Message {
  getEntriesList(): Array<resources_rector_audit_pb.AuditEntry>;
  setEntriesList(value: Array<resources_rector_audit_pb.AuditEntry>): ViewAuditLogResponse;
  clearEntriesList(): ViewAuditLogResponse;
  addEntries(value?: resources_rector_audit_pb.AuditEntry, index?: number): resources_rector_audit_pb.AuditEntry;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ViewAuditLogResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ViewAuditLogResponse): ViewAuditLogResponse.AsObject;
  static serializeBinaryToWriter(message: ViewAuditLogResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ViewAuditLogResponse;
  static deserializeBinaryFromReader(message: ViewAuditLogResponse, reader: jspb.BinaryReader): ViewAuditLogResponse;
}

export namespace ViewAuditLogResponse {
  export type AsObject = {
    entriesList: Array<resources_rector_audit_pb.AuditEntry.AsObject>,
  }
}

