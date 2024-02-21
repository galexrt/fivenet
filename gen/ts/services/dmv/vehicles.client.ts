// @generated by protobuf-ts 2.9.3 with parameter optimize_speed,long_type_bigint
// @generated from protobuf file "services/dmv/vehicles.proto" (package "services.dmv", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { DMVService } from "./vehicles";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { ListVehiclesResponse } from "./vehicles";
import type { ListVehiclesRequest } from "./vehicles";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service services.dmv.DMVService
 */
export interface IDMVServiceClient {
    /**
     * @perm
     *
     * @generated from protobuf rpc: ListVehicles(services.dmv.ListVehiclesRequest) returns (services.dmv.ListVehiclesResponse);
     */
    listVehicles(input: ListVehiclesRequest, options?: RpcOptions): UnaryCall<ListVehiclesRequest, ListVehiclesResponse>;
}
/**
 * @generated from protobuf service services.dmv.DMVService
 */
export class DMVServiceClient implements IDMVServiceClient, ServiceInfo {
    typeName = DMVService.typeName;
    methods = DMVService.methods;
    options = DMVService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: ListVehicles(services.dmv.ListVehiclesRequest) returns (services.dmv.ListVehiclesResponse);
     */
    listVehicles(input: ListVehiclesRequest, options?: RpcOptions): UnaryCall<ListVehiclesRequest, ListVehiclesResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<ListVehiclesRequest, ListVehiclesResponse>("unary", this._transport, method, opt, input);
    }
}
