// @generated by protobuf-ts 2.9.1 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "resources/livemap/livemap.proto" (package "resources.livemap", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { Unit } from "../dispatch/units.js";
import { UserShort } from "../users/users.js";
import { Timestamp } from "../timestamp/timestamp.js";
/**
 * @generated from protobuf message resources.livemap.MarkerInfo
 */
export interface MarkerInfo {
    /**
     * @generated from protobuf field: uint64 id = 1;
     */
    id: bigint;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp updated_at = 3;
     */
    updatedAt?: Timestamp;
    /**
     * @generated from protobuf field: string job = 4;
     */
    job: string;
    /**
     * @sanitize
     *
     * @generated from protobuf field: string name = 5;
     */
    name: string;
    /**
     * @sanitize
     *
     * @generated from protobuf field: optional string description = 6;
     */
    description?: string;
    /**
     * @generated from protobuf field: double x = 7;
     */
    x: number;
    /**
     * @generated from protobuf field: double y = 8;
     */
    y: number;
    /**
     * @sanitize
     *
     * @generated from protobuf field: optional string postal = 9;
     */
    postal?: string;
    /**
     * @generated from protobuf field: optional string color = 10;
     */
    color?: string;
    /**
     * @generated from protobuf field: optional string icon = 11;
     */
    icon?: string;
}
/**
 * @generated from protobuf message resources.livemap.UserMarker
 */
export interface UserMarker {
    /**
     * @generated from protobuf field: resources.livemap.MarkerInfo info = 1;
     */
    info?: MarkerInfo;
    /**
     * @generated from protobuf field: int32 user_id = 2;
     */
    userId: number;
    /**
     * @generated from protobuf field: resources.users.UserShort user = 3;
     */
    user?: UserShort; // @gotags: alias:"user"
    /**
     * @generated from protobuf field: optional uint64 unit_id = 4;
     */
    unitId?: bigint;
    /**
     * @generated from protobuf field: optional resources.dispatch.Unit unit = 5;
     */
    unit?: Unit;
}
/**
 * @generated from protobuf message resources.livemap.Marker
 */
export interface Marker {
    /**
     * @generated from protobuf field: resources.livemap.MarkerInfo info = 1;
     */
    info?: MarkerInfo;
    /**
     * @generated from protobuf field: resources.livemap.MarkerType type = 2;
     */
    type: MarkerType; // @gotags: alias:"markerType"
    /**
     * @generated from protobuf field: optional resources.livemap.MarkerData data = 3;
     */
    data?: MarkerData; // @gotags: alias:"markerData"
    /**
     * @generated from protobuf field: optional int32 creator_id = 4;
     */
    creatorId?: number;
    /**
     * @generated from protobuf field: optional resources.users.UserShort creator = 5;
     */
    creator?: UserShort;
}
/**
 * @generated from protobuf message resources.livemap.MarkerData
 */
export interface MarkerData {
    /**
     * @generated from protobuf oneof: data
     */
    data: {
        oneofKind: "circle";
        /**
         * @generated from protobuf field: resources.livemap.CircleMarker circle = 3;
         */
        circle: CircleMarker;
    } | {
        oneofKind: "icon";
        /**
         * @generated from protobuf field: resources.livemap.IconMarker icon = 4;
         */
        icon: IconMarker;
    } | {
        oneofKind: undefined;
    };
}
/**
 * @generated from protobuf message resources.livemap.CircleMarker
 */
export interface CircleMarker {
    /**
     * @generated from protobuf field: int32 radius = 1;
     */
    radius: number;
    /**
     * @generated from protobuf field: optional float oapcity = 2;
     */
    oapcity?: number;
}
/**
 * @generated from protobuf message resources.livemap.Coords
 */
export interface Coords {
    /**
     * @generated from protobuf field: double x = 1;
     */
    x: number;
    /**
     * @generated from protobuf field: double y = 2;
     */
    y: number;
}
/**
 * @generated from protobuf message resources.livemap.IconMarker
 */
export interface IconMarker {
    /**
     * @generated from protobuf field: string icon = 1;
     */
    icon: string;
}
/**
 * @generated from protobuf enum resources.livemap.MarkerType
 */
export enum MarkerType {
    /**
     * @generated from protobuf enum value: MARKER_TYPE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: MARKER_TYPE_DOT = 1;
     */
    DOT = 1,
    /**
     * @generated from protobuf enum value: MARKER_TYPE_CIRCLE = 2;
     */
    CIRCLE = 2,
    /**
     * @generated from protobuf enum value: MARKER_TYPE_ICON = 3;
     */
    ICON = 3
}
// @generated message type with reflection information, may provide speed optimized methods
class MarkerInfo$Type extends MessageType<MarkerInfo> {
    constructor() {
        super("resources.livemap.MarkerInfo", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 5, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 6, name: "description", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 7, name: "x", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 8, name: "y", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 9, name: "postal", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "48" } } } },
            { no: 10, name: "color", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { len: "6", pattern: "^[A-Fa-f0-9]{6}$" } } } },
            { no: 11, name: "icon", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.livemap.MarkerInfo
 */
export const MarkerInfo = new MarkerInfo$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UserMarker$Type extends MessageType<UserMarker> {
    constructor() {
        super("resources.livemap.UserMarker", [
            { no: 1, name: "info", kind: "message", T: () => MarkerInfo },
            { no: 2, name: "user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 3, name: "user", kind: "message", T: () => UserShort },
            { no: 4, name: "unit_id", kind: "scalar", opt: true, T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 5, name: "unit", kind: "message", T: () => Unit }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.livemap.UserMarker
 */
export const UserMarker = new UserMarker$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Marker$Type extends MessageType<Marker> {
    constructor() {
        super("resources.livemap.Marker", [
            { no: 1, name: "info", kind: "message", T: () => MarkerInfo },
            { no: 2, name: "type", kind: "enum", T: () => ["resources.livemap.MarkerType", MarkerType, "MARKER_TYPE_"] },
            { no: 3, name: "data", kind: "message", T: () => MarkerData },
            { no: 4, name: "creator_id", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "creator", kind: "message", T: () => UserShort }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.livemap.Marker
 */
export const Marker = new Marker$Type();
// @generated message type with reflection information, may provide speed optimized methods
class MarkerData$Type extends MessageType<MarkerData> {
    constructor() {
        super("resources.livemap.MarkerData", [
            { no: 3, name: "circle", kind: "message", oneof: "data", T: () => CircleMarker },
            { no: 4, name: "icon", kind: "message", oneof: "data", T: () => IconMarker }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.livemap.MarkerData
 */
export const MarkerData = new MarkerData$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CircleMarker$Type extends MessageType<CircleMarker> {
    constructor() {
        super("resources.livemap.CircleMarker", [
            { no: 1, name: "radius", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "oapcity", kind: "scalar", opt: true, T: 2 /*ScalarType.FLOAT*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.livemap.CircleMarker
 */
export const CircleMarker = new CircleMarker$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Coords$Type extends MessageType<Coords> {
    constructor() {
        super("resources.livemap.Coords", [
            { no: 1, name: "x", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 2, name: "y", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.livemap.Coords
 */
export const Coords = new Coords$Type();
// @generated message type with reflection information, may provide speed optimized methods
class IconMarker$Type extends MessageType<IconMarker> {
    constructor() {
        super("resources.livemap.IconMarker", [
            { no: 1, name: "icon", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.livemap.IconMarker
 */
export const IconMarker = new IconMarker$Type();
