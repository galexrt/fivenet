// @generated by protobuf-ts 2.9.3 with parameter optimize_speed,long_type_bigint
// @generated from protobuf file "services/jobs/qualifications.proto" (package "services.jobs", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { JobsQualificationsService } from "./qualifications";
/**
 * @generated from protobuf service services.jobs.JobsQualificationsService
 */
export interface IJobsQualificationsServiceClient {
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
}
