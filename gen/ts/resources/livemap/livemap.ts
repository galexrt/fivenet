// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "resources/livemap/livemap.proto" (package "resources.livemap", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { UserShort } from "../users/users.js";
import { Timestamp } from "../timestamp/timestamp.js";
/**
 * @generated from protobuf message resources.livemap.GenericMarker
 */
export interface GenericMarker {
    /**
     * @generated from protobuf field: int32 id = 1;
     */
    id: number;
    /**
     * @generated from protobuf field: resources.timestamp.Timestamp updated_at = 2;
     */
    updatedAt?: Timestamp; // @gotags: alias:"updated_at"
    /**
     * @generated from protobuf field: float x = 3;
     */
    x: number;
    /**
     * @generated from protobuf field: float y = 4;
     */
    y: number;
    /**
     * @generated from protobuf field: string name = 5;
     */
    name: string;
    /**
     * @generated from protobuf field: string icon = 6;
     */
    icon: string;
    /**
     * @generated from protobuf field: string icon_color = 7;
     */
    iconColor: string;
    /**
     * @generated from protobuf field: string popup = 8;
     */
    popup: string;
}
/**
 * @generated from protobuf message resources.livemap.DispatchMarker
 */
export interface DispatchMarker {
    /**
     * @generated from protobuf field: resources.livemap.GenericMarker marker = 1;
     */
    marker?: GenericMarker;
    /**
     * @generated from protobuf field: string job = 2;
     */
    job: string; // @gotags: alias:"job"
    /**
     * @generated from protobuf field: optional string job_label = 3;
     */
    jobLabel?: string; // @gotags: alias:"job_label"
    /**
     * @generated from protobuf field: bool active = 4;
     */
    active: boolean;
}
/**
 * @generated from protobuf message resources.livemap.UserMarker
 */
export interface UserMarker {
    /**
     * @generated from protobuf field: resources.livemap.GenericMarker marker = 1;
     */
    marker?: GenericMarker;
    /**
     * @generated from protobuf field: resources.users.UserShort user = 2;
     */
    user?: UserShort; // @gotags: alias:"user"
}
/**
 * @generated from protobuf enum resources.livemap.MARKER_TYPE
 */
export enum MARKER_TYPE {
    /**
     * @generated from protobuf enum value: GENERIC = 0;
     */
    GENERIC = 0,
    /**
     * @generated from protobuf enum value: USER = 1;
     */
    USER = 1,
    /**
     * @generated from protobuf enum value: DISPATCH = 2;
     */
    DISPATCH = 2,
    /**
     * @generated from protobuf enum value: CIRCLE = 3;
     */
    CIRCLE = 3
}
// @generated message type with reflection information, may provide speed optimized methods
class GenericMarker$Type extends MessageType<GenericMarker> {
    constructor() {
        super("resources.livemap.GenericMarker", [
            { no: 1, name: "id", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 2, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "x", kind: "scalar", T: 2 /*ScalarType.FLOAT*/ },
            { no: 4, name: "y", kind: "scalar", T: 2 /*ScalarType.FLOAT*/ },
            { no: 5, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 6, name: "icon", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 7, name: "icon_color", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 8, name: "popup", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.livemap.GenericMarker
 */
export const GenericMarker = new GenericMarker$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DispatchMarker$Type extends MessageType<DispatchMarker> {
    constructor() {
        super("resources.livemap.DispatchMarker", [
            { no: 1, name: "marker", kind: "message", T: () => GenericMarker },
            { no: 2, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 3, name: "job_label", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 4, name: "active", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.livemap.DispatchMarker
 */
export const DispatchMarker = new DispatchMarker$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UserMarker$Type extends MessageType<UserMarker> {
    constructor() {
        super("resources.livemap.UserMarker", [
            { no: 1, name: "marker", kind: "message", T: () => GenericMarker },
            { no: 2, name: "user", kind: "message", T: () => UserShort }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.livemap.UserMarker
 */
export const UserMarker = new UserMarker$Type();
