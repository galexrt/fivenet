// @generated by protobuf-ts 2.9.3 with parameter optimize_speed,long_type_bigint
// @generated from protobuf file "services/filestore/filestore.proto" (package "services.filestore", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { FileStoreService } from "./filestore";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { DeleteResponse } from "./filestore";
import type { DeleteRequest } from "./filestore";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service services.filestore.FileStoreService
 */
export interface IFileStoreServiceClient {
    /**
     * @perm: Name=SuperUser
     *
     * @generated from protobuf rpc: Delete(services.filestore.DeleteRequest) returns (services.filestore.DeleteResponse);
     */
    delete(input: DeleteRequest, options?: RpcOptions): UnaryCall<DeleteRequest, DeleteResponse>;
}
/**
 * @generated from protobuf service services.filestore.FileStoreService
 */
export class FileStoreServiceClient implements IFileStoreServiceClient, ServiceInfo {
    typeName = FileStoreService.typeName;
    methods = FileStoreService.methods;
    options = FileStoreService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @perm: Name=SuperUser
     *
     * @generated from protobuf rpc: Delete(services.filestore.DeleteRequest) returns (services.filestore.DeleteResponse);
     */
    delete(input: DeleteRequest, options?: RpcOptions): UnaryCall<DeleteRequest, DeleteResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<DeleteRequest, DeleteResponse>("unary", this._transport, method, opt, input);
    }
}
