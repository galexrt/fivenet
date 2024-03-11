// @generated by protobuf-ts 2.9.3 with parameter optimize_speed,long_type_bigint
// @generated from protobuf file "services/jobs/qualifications.proto" (package "services.jobs", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { JobsQualificationsService } from "./qualifications";
import type { DeleteQualificationResponse } from "./qualifications";
import type { DeleteQualificationRequest } from "./qualifications";
import type { UpdateQualificationResponse } from "./qualifications";
import type { UpdateQualificationRequest } from "./qualifications";
import type { CreateQualificationResponse } from "./qualifications";
import type { CreateQualificationRequest } from "./qualifications";
import type { GetQualificationResponse } from "./qualifications";
import type { GetQualificationRequest } from "./qualifications";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { ListQualificationsResponse } from "./qualifications";
import type { ListQualificationsRequest } from "./qualifications";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service services.jobs.JobsQualificationsService
 */
export interface IJobsQualificationsServiceClient {
    /**
     * @perm
     *
     * @generated from protobuf rpc: ListQualifications(services.jobs.ListQualificationsRequest) returns (services.jobs.ListQualificationsResponse);
     */
    listQualifications(input: ListQualificationsRequest, options?: RpcOptions): UnaryCall<ListQualificationsRequest, ListQualificationsResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: GetQualification(services.jobs.GetQualificationRequest) returns (services.jobs.GetQualificationResponse);
     */
    getQualification(input: GetQualificationRequest, options?: RpcOptions): UnaryCall<GetQualificationRequest, GetQualificationResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: CreateQualification(services.jobs.CreateQualificationRequest) returns (services.jobs.CreateQualificationResponse);
     */
    createQualification(input: CreateQualificationRequest, options?: RpcOptions): UnaryCall<CreateQualificationRequest, CreateQualificationResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: UpdateQualification(services.jobs.UpdateQualificationRequest) returns (services.jobs.UpdateQualificationResponse);
     */
    updateQualification(input: UpdateQualificationRequest, options?: RpcOptions): UnaryCall<UpdateQualificationRequest, UpdateQualificationResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: DeleteQualification(services.jobs.DeleteQualificationRequest) returns (services.jobs.DeleteQualificationResponse);
     */
    deleteQualification(input: DeleteQualificationRequest, options?: RpcOptions): UnaryCall<DeleteQualificationRequest, DeleteQualificationResponse>;
}
/**
 * @generated from protobuf service services.jobs.JobsQualificationsService
 */
export class JobsQualificationsServiceClient implements IJobsQualificationsServiceClient, ServiceInfo {
    typeName = JobsQualificationsService.typeName;
    methods = JobsQualificationsService.methods;
    options = JobsQualificationsService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: ListQualifications(services.jobs.ListQualificationsRequest) returns (services.jobs.ListQualificationsResponse);
     */
    listQualifications(input: ListQualificationsRequest, options?: RpcOptions): UnaryCall<ListQualificationsRequest, ListQualificationsResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<ListQualificationsRequest, ListQualificationsResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: GetQualification(services.jobs.GetQualificationRequest) returns (services.jobs.GetQualificationResponse);
     */
    getQualification(input: GetQualificationRequest, options?: RpcOptions): UnaryCall<GetQualificationRequest, GetQualificationResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetQualificationRequest, GetQualificationResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: CreateQualification(services.jobs.CreateQualificationRequest) returns (services.jobs.CreateQualificationResponse);
     */
    createQualification(input: CreateQualificationRequest, options?: RpcOptions): UnaryCall<CreateQualificationRequest, CreateQualificationResponse> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateQualificationRequest, CreateQualificationResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: UpdateQualification(services.jobs.UpdateQualificationRequest) returns (services.jobs.UpdateQualificationResponse);
     */
    updateQualification(input: UpdateQualificationRequest, options?: RpcOptions): UnaryCall<UpdateQualificationRequest, UpdateQualificationResponse> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<UpdateQualificationRequest, UpdateQualificationResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: DeleteQualification(services.jobs.DeleteQualificationRequest) returns (services.jobs.DeleteQualificationResponse);
     */
    deleteQualification(input: DeleteQualificationRequest, options?: RpcOptions): UnaryCall<DeleteQualificationRequest, DeleteQualificationResponse> {
        const method = this.methods[4], opt = this._transport.mergeOptions(options);
        return stackIntercept<DeleteQualificationRequest, DeleteQualificationResponse>("unary", this._transport, method, opt, input);
    }
}
