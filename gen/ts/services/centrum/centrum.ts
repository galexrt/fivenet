// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "services/centrum/centrum.proto" (package "services.centrum", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import { MessageType } from "@protobuf-ts/runtime";
// Squad Management

/**
 * @generated from protobuf message services.centrum.CreateSquadRequest
 */
export interface CreateSquadRequest {
    /**
     * @generated from protobuf field: string name = 1;
     */
    name: string;
    /**
     * @generated from protobuf field: optional uint64 limit = 2;
     */
    limit?: bigint;
}
/**
 * @generated from protobuf message services.centrum.CreateSquadResponse
 */
export interface CreateSquadResponse {
    /**
     * @generated from protobuf field: uint64 id = 1;
     */
    id: bigint;
}
/**
 * @generated from protobuf message services.centrum.UpdateSquadRequest
 */
export interface UpdateSquadRequest {
    /**
     * @generated from protobuf field: uint64 id = 1;
     */
    id: bigint;
}
/**
 * @generated from protobuf message services.centrum.UpdateSquadResponse
 */
export interface UpdateSquadResponse {
    /**
     * @generated from protobuf field: uint64 id = 1;
     */
    id: bigint;
}
/**
 * @generated from protobuf message services.centrum.DeleteSquadRequest
 */
export interface DeleteSquadRequest {
    /**
     * @generated from protobuf field: uint64 id = 1;
     */
    id: bigint;
}
/**
 * @generated from protobuf message services.centrum.DeleteSquadResponse
 */
export interface DeleteSquadResponse {
}
/**
 * @generated from protobuf message services.centrum.AssignSquadRequest
 */
export interface AssignSquadRequest {
    /**
     * @generated from protobuf field: uint64 squad_id = 1;
     */
    squadId: bigint;
    /**
     * @generated from protobuf field: repeated int32 to_add = 2;
     */
    toAdd: number[];
    /**
     * @generated from protobuf field: repeated int32 to_remove = 3;
     */
    toRemove: number[];
}
/**
 * @generated from protobuf message services.centrum.AssignSquadResponse
 */
export interface AssignSquadResponse {
}
/**
 * @generated from protobuf message services.centrum.SquadStreamRequest
 */
export interface SquadStreamRequest {
}
/**
 * @generated from protobuf message services.centrum.SquadStreamResponse
 */
export interface SquadStreamResponse {
    /**
     * @generated from protobuf field: repeated services.centrum.SquadChanges changes = 1;
     */
    changes: SquadChanges[];
}
/**
 * @generated from protobuf message services.centrum.SquadChanges
 */
export interface SquadChanges {
    /**
     * @generated from protobuf field: uint64 id = 1;
     */
    id: bigint;
    /**
     * @generated from protobuf field: repeated int32 added = 2;
     */
    added: number[];
    /**
     * @generated from protobuf field: repeated int32 removed = 3;
     */
    removed: number[];
}
// Action + Dispatch Management

/**
 * @generated from protobuf message services.centrum.CreateActionRequest
 */
export interface CreateActionRequest {
}
/**
 * @generated from protobuf message services.centrum.CreateActionResponse
 */
export interface CreateActionResponse {
}
// TODO

/**
 * @generated from protobuf message services.centrum.CentrumStreamRequest
 */
export interface CentrumStreamRequest {
}
/**
 * @generated from protobuf message services.centrum.CentrumStreamResponse
 */
export interface CentrumStreamResponse {
}
// @generated message type with reflection information, may provide speed optimized methods
class CreateSquadRequest$Type extends MessageType<CreateSquadRequest> {
    constructor() {
        super("services.centrum.CreateSquadRequest", [
            { no: 1, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "limit", kind: "scalar", opt: true, T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.CreateSquadRequest
 */
export const CreateSquadRequest = new CreateSquadRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CreateSquadResponse$Type extends MessageType<CreateSquadResponse> {
    constructor() {
        super("services.centrum.CreateSquadResponse", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.CreateSquadResponse
 */
export const CreateSquadResponse = new CreateSquadResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpdateSquadRequest$Type extends MessageType<UpdateSquadRequest> {
    constructor() {
        super("services.centrum.UpdateSquadRequest", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.UpdateSquadRequest
 */
export const UpdateSquadRequest = new UpdateSquadRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpdateSquadResponse$Type extends MessageType<UpdateSquadResponse> {
    constructor() {
        super("services.centrum.UpdateSquadResponse", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.UpdateSquadResponse
 */
export const UpdateSquadResponse = new UpdateSquadResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteSquadRequest$Type extends MessageType<DeleteSquadRequest> {
    constructor() {
        super("services.centrum.DeleteSquadRequest", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.DeleteSquadRequest
 */
export const DeleteSquadRequest = new DeleteSquadRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteSquadResponse$Type extends MessageType<DeleteSquadResponse> {
    constructor() {
        super("services.centrum.DeleteSquadResponse", []);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.DeleteSquadResponse
 */
export const DeleteSquadResponse = new DeleteSquadResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class AssignSquadRequest$Type extends MessageType<AssignSquadRequest> {
    constructor() {
        super("services.centrum.AssignSquadRequest", [
            { no: 1, name: "squad_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "to_add", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "to_remove", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 5 /*ScalarType.INT32*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.AssignSquadRequest
 */
export const AssignSquadRequest = new AssignSquadRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class AssignSquadResponse$Type extends MessageType<AssignSquadResponse> {
    constructor() {
        super("services.centrum.AssignSquadResponse", []);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.AssignSquadResponse
 */
export const AssignSquadResponse = new AssignSquadResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SquadStreamRequest$Type extends MessageType<SquadStreamRequest> {
    constructor() {
        super("services.centrum.SquadStreamRequest", []);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.SquadStreamRequest
 */
export const SquadStreamRequest = new SquadStreamRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SquadStreamResponse$Type extends MessageType<SquadStreamResponse> {
    constructor() {
        super("services.centrum.SquadStreamResponse", [
            { no: 1, name: "changes", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => SquadChanges }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.SquadStreamResponse
 */
export const SquadStreamResponse = new SquadStreamResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SquadChanges$Type extends MessageType<SquadChanges> {
    constructor() {
        super("services.centrum.SquadChanges", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "added", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "removed", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 5 /*ScalarType.INT32*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.SquadChanges
 */
export const SquadChanges = new SquadChanges$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CreateActionRequest$Type extends MessageType<CreateActionRequest> {
    constructor() {
        super("services.centrum.CreateActionRequest", []);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.CreateActionRequest
 */
export const CreateActionRequest = new CreateActionRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CreateActionResponse$Type extends MessageType<CreateActionResponse> {
    constructor() {
        super("services.centrum.CreateActionResponse", []);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.CreateActionResponse
 */
export const CreateActionResponse = new CreateActionResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CentrumStreamRequest$Type extends MessageType<CentrumStreamRequest> {
    constructor() {
        super("services.centrum.CentrumStreamRequest", []);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.CentrumStreamRequest
 */
export const CentrumStreamRequest = new CentrumStreamRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CentrumStreamResponse$Type extends MessageType<CentrumStreamResponse> {
    constructor() {
        super("services.centrum.CentrumStreamResponse", []);
    }
}
/**
 * @generated MessageType for protobuf message services.centrum.CentrumStreamResponse
 */
export const CentrumStreamResponse = new CentrumStreamResponse$Type();
/**
 * @generated ServiceType for protobuf service services.centrum.SquadService
 */
export const SquadService = new ServiceType("services.centrum.SquadService", [
    { name: "CreateSquad", options: {}, I: CreateSquadRequest, O: CreateSquadResponse },
    { name: "UpdateSquad", options: {}, I: UpdateSquadRequest, O: UpdateSquadResponse },
    { name: "DeleteSquad", options: {}, I: DeleteSquadRequest, O: DeleteSquadResponse },
    { name: "AssignSquad", options: {}, I: AssignSquadRequest, O: AssignSquadResponse },
    { name: "StreamSquads", serverStreaming: true, options: {}, I: SquadStreamRequest, O: SquadStreamResponse }
]);
/**
 * @generated ServiceType for protobuf service services.centrum.CentrumService
 */
export const CentrumService = new ServiceType("services.centrum.CentrumService", [
    { name: "CreateAction", options: {}, I: CreateActionRequest, O: CreateActionResponse },
    { name: "Stream", serverStreaming: true, options: {}, I: CentrumStreamRequest, O: CentrumStreamResponse }
]);
