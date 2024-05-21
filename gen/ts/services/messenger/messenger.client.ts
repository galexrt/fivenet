// @generated by protobuf-ts 2.9.4 with parameter optimize_speed,long_type_number,force_server_none
// @generated from protobuf file "services/messenger/messenger.proto" (package "services.messenger", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { MessengerService } from "./messenger";
import type { DeleteMessageResponse } from "./messenger";
import type { DeleteMessageRequest } from "./messenger";
import type { PostMessageResponse } from "./messenger";
import type { PostMessageRequest } from "./messenger";
import type { GetThreadMessagesResponse } from "./messenger";
import type { GetThreadMessagesRequest } from "./messenger";
import type { SetThreadUserStateResponse } from "./messenger";
import type { SetThreadUserStateRequest } from "./messenger";
import type { DeleteThreadResponse } from "./messenger";
import type { DeleteThreadRequest } from "./messenger";
import type { CreateOrUpdateThreadResponse } from "./messenger";
import type { CreateOrUpdateThreadRequest } from "./messenger";
import type { GetThreadResponse } from "./messenger";
import type { GetThreadRequest } from "./messenger";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { ListThreadsResponse } from "./messenger";
import type { ListThreadsRequest } from "./messenger";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service services.messenger.MessengerService
 */
export interface IMessengerServiceClient {
    /**
     * @perm
     *
     * @generated from protobuf rpc: ListThreads(services.messenger.ListThreadsRequest) returns (services.messenger.ListThreadsResponse);
     */
    listThreads(input: ListThreadsRequest, options?: RpcOptions): UnaryCall<ListThreadsRequest, ListThreadsResponse>;
    /**
     * @perm: Name=ListThreads
     *
     * @generated from protobuf rpc: GetThread(services.messenger.GetThreadRequest) returns (services.messenger.GetThreadResponse);
     */
    getThread(input: GetThreadRequest, options?: RpcOptions): UnaryCall<GetThreadRequest, GetThreadResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: CreateOrUpdateThread(services.messenger.CreateOrUpdateThreadRequest) returns (services.messenger.CreateOrUpdateThreadResponse);
     */
    createOrUpdateThread(input: CreateOrUpdateThreadRequest, options?: RpcOptions): UnaryCall<CreateOrUpdateThreadRequest, CreateOrUpdateThreadResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: DeleteThread(services.messenger.DeleteThreadRequest) returns (services.messenger.DeleteThreadResponse);
     */
    deleteThread(input: DeleteThreadRequest, options?: RpcOptions): UnaryCall<DeleteThreadRequest, DeleteThreadResponse>;
    /**
     * @perm: Name=ListThreads
     *
     * @generated from protobuf rpc: SetThreadUserState(services.messenger.SetThreadUserStateRequest) returns (services.messenger.SetThreadUserStateResponse);
     */
    setThreadUserState(input: SetThreadUserStateRequest, options?: RpcOptions): UnaryCall<SetThreadUserStateRequest, SetThreadUserStateResponse>;
    /**
     * @perm: Name=ListThreads
     *
     * @generated from protobuf rpc: GetThreadMessages(services.messenger.GetThreadMessagesRequest) returns (services.messenger.GetThreadMessagesResponse);
     */
    getThreadMessages(input: GetThreadMessagesRequest, options?: RpcOptions): UnaryCall<GetThreadMessagesRequest, GetThreadMessagesResponse>;
    /**
     * @perm
     *
     * @generated from protobuf rpc: PostMessage(services.messenger.PostMessageRequest) returns (services.messenger.PostMessageResponse);
     */
    postMessage(input: PostMessageRequest, options?: RpcOptions): UnaryCall<PostMessageRequest, PostMessageResponse>;
    /**
     * @perm: Name=SuperUser
     *
     * @generated from protobuf rpc: DeleteMessage(services.messenger.DeleteMessageRequest) returns (services.messenger.DeleteMessageResponse);
     */
    deleteMessage(input: DeleteMessageRequest, options?: RpcOptions): UnaryCall<DeleteMessageRequest, DeleteMessageResponse>;
}
/**
 * @generated from protobuf service services.messenger.MessengerService
 */
export class MessengerServiceClient implements IMessengerServiceClient, ServiceInfo {
    typeName = MessengerService.typeName;
    methods = MessengerService.methods;
    options = MessengerService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: ListThreads(services.messenger.ListThreadsRequest) returns (services.messenger.ListThreadsResponse);
     */
    listThreads(input: ListThreadsRequest, options?: RpcOptions): UnaryCall<ListThreadsRequest, ListThreadsResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<ListThreadsRequest, ListThreadsResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=ListThreads
     *
     * @generated from protobuf rpc: GetThread(services.messenger.GetThreadRequest) returns (services.messenger.GetThreadResponse);
     */
    getThread(input: GetThreadRequest, options?: RpcOptions): UnaryCall<GetThreadRequest, GetThreadResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetThreadRequest, GetThreadResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: CreateOrUpdateThread(services.messenger.CreateOrUpdateThreadRequest) returns (services.messenger.CreateOrUpdateThreadResponse);
     */
    createOrUpdateThread(input: CreateOrUpdateThreadRequest, options?: RpcOptions): UnaryCall<CreateOrUpdateThreadRequest, CreateOrUpdateThreadResponse> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateOrUpdateThreadRequest, CreateOrUpdateThreadResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: DeleteThread(services.messenger.DeleteThreadRequest) returns (services.messenger.DeleteThreadResponse);
     */
    deleteThread(input: DeleteThreadRequest, options?: RpcOptions): UnaryCall<DeleteThreadRequest, DeleteThreadResponse> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<DeleteThreadRequest, DeleteThreadResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=ListThreads
     *
     * @generated from protobuf rpc: SetThreadUserState(services.messenger.SetThreadUserStateRequest) returns (services.messenger.SetThreadUserStateResponse);
     */
    setThreadUserState(input: SetThreadUserStateRequest, options?: RpcOptions): UnaryCall<SetThreadUserStateRequest, SetThreadUserStateResponse> {
        const method = this.methods[4], opt = this._transport.mergeOptions(options);
        return stackIntercept<SetThreadUserStateRequest, SetThreadUserStateResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=ListThreads
     *
     * @generated from protobuf rpc: GetThreadMessages(services.messenger.GetThreadMessagesRequest) returns (services.messenger.GetThreadMessagesResponse);
     */
    getThreadMessages(input: GetThreadMessagesRequest, options?: RpcOptions): UnaryCall<GetThreadMessagesRequest, GetThreadMessagesResponse> {
        const method = this.methods[5], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetThreadMessagesRequest, GetThreadMessagesResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm
     *
     * @generated from protobuf rpc: PostMessage(services.messenger.PostMessageRequest) returns (services.messenger.PostMessageResponse);
     */
    postMessage(input: PostMessageRequest, options?: RpcOptions): UnaryCall<PostMessageRequest, PostMessageResponse> {
        const method = this.methods[6], opt = this._transport.mergeOptions(options);
        return stackIntercept<PostMessageRequest, PostMessageResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @perm: Name=SuperUser
     *
     * @generated from protobuf rpc: DeleteMessage(services.messenger.DeleteMessageRequest) returns (services.messenger.DeleteMessageResponse);
     */
    deleteMessage(input: DeleteMessageRequest, options?: RpcOptions): UnaryCall<DeleteMessageRequest, DeleteMessageResponse> {
        const method = this.methods[7], opt = this._transport.mergeOptions(options);
        return stackIntercept<DeleteMessageRequest, DeleteMessageResponse>("unary", this._transport, method, opt, input);
    }
}
