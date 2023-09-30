// @generated by protobuf-ts 2.9.1 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "services/rector/rector.proto" (package "services.rector", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { RectorService } from "./rector.js";
import type { DeleteLawResponse } from "./rector.js";
import type { DeleteLawRequest } from "./rector.js";
import type { CreateOrUpdateLawResponse } from "./rector.js";
import type { CreateOrUpdateLawRequest } from "./rector.js";
import type { DeleteLawBookResponse } from "./rector.js";
import type { DeleteLawBookRequest } from "./rector.js";
import type { CreateOrUpdateLawBookResponse } from "./rector.js";
import type { CreateOrUpdateLawBookRequest } from "./rector.js";
import type { UpdateRoleLimitsResponse } from "./rector.js";
import type { UpdateRoleLimitsRequest } from "./rector.js";
import type { ViewAuditLogResponse } from "./rector.js";
import type { ViewAuditLogRequest } from "./rector.js";
import type { GetPermissionsResponse } from "./rector.js";
import type { GetPermissionsRequest } from "./rector.js";
import type { UpdateRolePermsResponse } from "./rector.js";
import type { UpdateRolePermsRequest } from "./rector.js";
import type { DeleteRoleResponse } from "./rector.js";
import type { DeleteRoleRequest } from "./rector.js";
import type { CreateRoleResponse } from "./rector.js";
import type { CreateRoleRequest } from "./rector.js";
import type { GetRoleResponse } from "./rector.js";
import type { GetRoleRequest } from "./rector.js";
import type { GetRolesResponse } from "./rector.js";
import type { GetRolesRequest } from "./rector.js";
import type { SetJobPropsResponse } from "./rector.js";
import type { SetJobPropsRequest } from "./rector.js";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { GetJobPropsResponse } from "./rector.js";
import type { GetJobPropsRequest } from "./rector.js";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service services.rector.RectorService
 */
export interface IRectorServiceClient {
    /**
     * @perm
     *
     * @generated from protobuf rpc: GetJobProps(services.rector.GetJobPropsRequest) returns (services.rector.GetJobPropsResponse);
     */
    getJobProps(input: GetJobPropsRequest, options?: RpcOptions): UnaryCall<GetJobPropsRequest, GetJobPropsResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: SetJobProps(services.rector.SetJobPropsRequest) returns (services.rector.SetJobPropsResponse);
     */
    setJobProps(input: SetJobPropsRequest, options?: RpcOptions): UnaryCall<SetJobPropsRequest, SetJobPropsResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: GetRoles(services.rector.GetRolesRequest) returns (services.rector.GetRolesResponse);
     */
    getRoles(input: GetRolesRequest, options?: RpcOptions): UnaryCall<GetRolesRequest, GetRolesResponse>;
    /**
     * @perm: Name=GetRoles
     *
     * @generated from protobuf rpc: GetRole(services.rector.GetRoleRequest) returns (services.rector.GetRoleResponse);
     */
    getRole(input: GetRoleRequest, options?: RpcOptions): UnaryCall<GetRoleRequest, GetRoleResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: CreateRole(services.rector.CreateRoleRequest) returns (services.rector.CreateRoleResponse);
     */
    createRole(input: CreateRoleRequest, options?: RpcOptions): UnaryCall<CreateRoleRequest, CreateRoleResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: DeleteRole(services.rector.DeleteRoleRequest) returns (services.rector.DeleteRoleResponse);
     */
    deleteRole(input: DeleteRoleRequest, options?: RpcOptions): UnaryCall<DeleteRoleRequest, DeleteRoleResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: UpdateRolePerms(services.rector.UpdateRolePermsRequest) returns (services.rector.UpdateRolePermsResponse);
     */
    updateRolePerms(input: UpdateRolePermsRequest, options?: RpcOptions): UnaryCall<UpdateRolePermsRequest, UpdateRolePermsResponse>;
    /**
     * @perm: Name=GetRoles
     *
     * @generated from protobuf rpc: GetPermissions(services.rector.GetPermissionsRequest) returns (services.rector.GetPermissionsResponse);
     */
    getPermissions(input: GetPermissionsRequest, options?: RpcOptions): UnaryCall<GetPermissionsRequest, GetPermissionsResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: ViewAuditLog(services.rector.ViewAuditLogRequest) returns (services.rector.ViewAuditLogResponse);
     */
    viewAuditLog(input: ViewAuditLogRequest, options?: RpcOptions): UnaryCall<ViewAuditLogRequest, ViewAuditLogResponse>;
    /**
     * @perm: Name=SuperUser
     *
     * @generated from protobuf rpc: UpdateRoleLimits(services.rector.UpdateRoleLimitsRequest) returns (services.rector.UpdateRoleLimitsResponse);
     */
    updateRoleLimits(input: UpdateRoleLimitsRequest, options?: RpcOptions): UnaryCall<UpdateRoleLimitsRequest, UpdateRoleLimitsResponse>;
    /**
     * @perm: Name=SuperUser
     *
     * @generated from protobuf rpc: CreateOrUpdateLawBook(services.rector.CreateOrUpdateLawBookRequest) returns (services.rector.CreateOrUpdateLawBookResponse);
     */
    createOrUpdateLawBook(input: CreateOrUpdateLawBookRequest, options?: RpcOptions): UnaryCall<CreateOrUpdateLawBookRequest, CreateOrUpdateLawBookResponse>;
    /**
     * @perm: Name=SuperUser
     *
     * @generated from protobuf rpc: DeleteLawBook(services.rector.DeleteLawBookRequest) returns (services.rector.DeleteLawBookResponse);
     */
    deleteLawBook(input: DeleteLawBookRequest, options?: RpcOptions): UnaryCall<DeleteLawBookRequest, DeleteLawBookResponse>;
    /**
     * @perm: Name=SuperUser
     *
     * @generated from protobuf rpc: CreateOrUpdateLaw(services.rector.CreateOrUpdateLawRequest) returns (services.rector.CreateOrUpdateLawResponse);
     */
    createOrUpdateLaw(input: CreateOrUpdateLawRequest, options?: RpcOptions): UnaryCall<CreateOrUpdateLawRequest, CreateOrUpdateLawResponse>;
    /**
     * @perm: Name=SuperUser
     *
     * @generated from protobuf rpc: DeleteLaw(services.rector.DeleteLawRequest) returns (services.rector.DeleteLawResponse);
     */
    deleteLaw(input: DeleteLawRequest, options?: RpcOptions): UnaryCall<DeleteLawRequest, DeleteLawResponse>;
}
/**
 * @generated from protobuf service services.rector.RectorService
 */
export class RectorServiceClient implements IRectorServiceClient, ServiceInfo {
    typeName = RectorService.typeName;
    methods = RectorService.methods;
    options = RectorService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: GetJobProps(services.rector.GetJobPropsRequest) returns (services.rector.GetJobPropsResponse);
     */
    getJobProps(input: GetJobPropsRequest, options?: RpcOptions): UnaryCall<GetJobPropsRequest, GetJobPropsResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetJobPropsRequest, GetJobPropsResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: SetJobProps(services.rector.SetJobPropsRequest) returns (services.rector.SetJobPropsResponse);
     */
    setJobProps(input: SetJobPropsRequest, options?: RpcOptions): UnaryCall<SetJobPropsRequest, SetJobPropsResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<SetJobPropsRequest, SetJobPropsResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: GetRoles(services.rector.GetRolesRequest) returns (services.rector.GetRolesResponse);
     */
    getRoles(input: GetRolesRequest, options?: RpcOptions): UnaryCall<GetRolesRequest, GetRolesResponse> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetRolesRequest, GetRolesResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=GetRoles
     *
     * @generated from protobuf rpc: GetRole(services.rector.GetRoleRequest) returns (services.rector.GetRoleResponse);
     */
    getRole(input: GetRoleRequest, options?: RpcOptions): UnaryCall<GetRoleRequest, GetRoleResponse> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetRoleRequest, GetRoleResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: CreateRole(services.rector.CreateRoleRequest) returns (services.rector.CreateRoleResponse);
     */
    createRole(input: CreateRoleRequest, options?: RpcOptions): UnaryCall<CreateRoleRequest, CreateRoleResponse> {
        const method = this.methods[4], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateRoleRequest, CreateRoleResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: DeleteRole(services.rector.DeleteRoleRequest) returns (services.rector.DeleteRoleResponse);
     */
    deleteRole(input: DeleteRoleRequest, options?: RpcOptions): UnaryCall<DeleteRoleRequest, DeleteRoleResponse> {
        const method = this.methods[5], opt = this._transport.mergeOptions(options);
        return stackIntercept<DeleteRoleRequest, DeleteRoleResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: UpdateRolePerms(services.rector.UpdateRolePermsRequest) returns (services.rector.UpdateRolePermsResponse);
     */
    updateRolePerms(input: UpdateRolePermsRequest, options?: RpcOptions): UnaryCall<UpdateRolePermsRequest, UpdateRolePermsResponse> {
        const method = this.methods[6], opt = this._transport.mergeOptions(options);
        return stackIntercept<UpdateRolePermsRequest, UpdateRolePermsResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=GetRoles
     *
     * @generated from protobuf rpc: GetPermissions(services.rector.GetPermissionsRequest) returns (services.rector.GetPermissionsResponse);
     */
    getPermissions(input: GetPermissionsRequest, options?: RpcOptions): UnaryCall<GetPermissionsRequest, GetPermissionsResponse> {
        const method = this.methods[7], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetPermissionsRequest, GetPermissionsResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: ViewAuditLog(services.rector.ViewAuditLogRequest) returns (services.rector.ViewAuditLogResponse);
     */
    viewAuditLog(input: ViewAuditLogRequest, options?: RpcOptions): UnaryCall<ViewAuditLogRequest, ViewAuditLogResponse> {
        const method = this.methods[8], opt = this._transport.mergeOptions(options);
        return stackIntercept<ViewAuditLogRequest, ViewAuditLogResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=SuperUser
     *
     * @generated from protobuf rpc: UpdateRoleLimits(services.rector.UpdateRoleLimitsRequest) returns (services.rector.UpdateRoleLimitsResponse);
     */
    updateRoleLimits(input: UpdateRoleLimitsRequest, options?: RpcOptions): UnaryCall<UpdateRoleLimitsRequest, UpdateRoleLimitsResponse> {
        const method = this.methods[9], opt = this._transport.mergeOptions(options);
        return stackIntercept<UpdateRoleLimitsRequest, UpdateRoleLimitsResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=SuperUser
     *
     * @generated from protobuf rpc: CreateOrUpdateLawBook(services.rector.CreateOrUpdateLawBookRequest) returns (services.rector.CreateOrUpdateLawBookResponse);
     */
    createOrUpdateLawBook(input: CreateOrUpdateLawBookRequest, options?: RpcOptions): UnaryCall<CreateOrUpdateLawBookRequest, CreateOrUpdateLawBookResponse> {
        const method = this.methods[10], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateOrUpdateLawBookRequest, CreateOrUpdateLawBookResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=SuperUser
     *
     * @generated from protobuf rpc: DeleteLawBook(services.rector.DeleteLawBookRequest) returns (services.rector.DeleteLawBookResponse);
     */
    deleteLawBook(input: DeleteLawBookRequest, options?: RpcOptions): UnaryCall<DeleteLawBookRequest, DeleteLawBookResponse> {
        const method = this.methods[11], opt = this._transport.mergeOptions(options);
        return stackIntercept<DeleteLawBookRequest, DeleteLawBookResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=SuperUser
     *
     * @generated from protobuf rpc: CreateOrUpdateLaw(services.rector.CreateOrUpdateLawRequest) returns (services.rector.CreateOrUpdateLawResponse);
     */
    createOrUpdateLaw(input: CreateOrUpdateLawRequest, options?: RpcOptions): UnaryCall<CreateOrUpdateLawRequest, CreateOrUpdateLawResponse> {
        const method = this.methods[12], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateOrUpdateLawRequest, CreateOrUpdateLawResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=SuperUser
     *
     * @generated from protobuf rpc: DeleteLaw(services.rector.DeleteLawRequest) returns (services.rector.DeleteLawResponse);
     */
    deleteLaw(input: DeleteLawRequest, options?: RpcOptions): UnaryCall<DeleteLawRequest, DeleteLawResponse> {
        const method = this.methods[13], opt = this._transport.mergeOptions(options);
        return stackIntercept<DeleteLawRequest, DeleteLawResponse>("unary", this._transport, method, opt, input);
    }
}
