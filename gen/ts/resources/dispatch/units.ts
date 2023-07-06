// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "resources/dispatch/units.proto" (package "resources.dispatch", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { UserShort } from "../users/users.js";
import { Timestamp } from "../timestamp/timestamp.js";
/**
 * @generated from protobuf message resources.dispatch.Unit
 */
export interface Unit {
    /**
     * @generated from protobuf field: uint64 id = 1;
     */
    id: bigint; // @gotags: sql:"primary_key" alias:"id"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp updated_at = 3;
     */
    updatedAt?: Timestamp;
    /**
     * @generated from protobuf field: string name = 4;
     */
    name: string;
    /**
     * @generated from protobuf field: string initials = 5;
     */
    initials: string;
    /**
     * @generated from protobuf field: optional string color = 6;
     */
    color?: string;
    /**
     * @generated from protobuf field: optional string description = 7;
     */
    description?: string;
    /**
     * @generated from protobuf field: optional resources.dispatch.UnitStatus status = 8;
     */
    status?: UnitStatus;
    /**
     * @generated from protobuf field: repeated resources.users.UserShort users = 9;
     */
    users: UserShort[];
}
/**
 * @generated from protobuf message resources.dispatch.UnitStatus
 */
export interface UnitStatus {
    /**
     * @generated from protobuf field: uint64 id = 1;
     */
    id: bigint;
    /**
     * @generated from protobuf field: uint64 unit_id = 2;
     */
    unitId: bigint;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 3;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: optional resources.dispatch.UNIT_STATUS status = 4;
     */
    status?: UNIT_STATUS;
    /**
     * @generated from protobuf field: optional string reason = 5;
     */
    reason?: string;
    /**
     * @generated from protobuf field: optional string code = 6;
     */
    code?: string;
    /**
     * @generated from protobuf field: int32 user_id = 7;
     */
    userId: number;
    /**
     * @generated from protobuf field: resources.users.UserShort user = 8;
     */
    user?: UserShort;
    /**
     * @generated from protobuf field: optional bool in_squad = 9;
     */
    inSquad?: boolean;
    /**
     * @generated from protobuf field: optional float x = 10;
     */
    x?: number;
    /**
     * @generated from protobuf field: optional float y = 11;
     */
    y?: number;
}
/**
 * @generated from protobuf enum resources.dispatch.UNIT_STATUS
 */
export enum UNIT_STATUS {
    /**
     * @generated from protobuf enum value: UNAVAILABLE = 0;
     */
    UNAVAILABLE = 0,
    /**
     * @generated from protobuf enum value: AVAILABLE = 1;
     */
    AVAILABLE = 1,
    /**
     * @generated from protobuf enum value: USER_ADDED = 2;
     */
    USER_ADDED = 2,
    /**
     * @generated from protobuf enum value: ON_BREAK = 3;
     */
    ON_BREAK = 3,
    /**
     * @generated from protobuf enum value: BUSY = 4;
     */
    BUSY = 4
}
// @generated message type with reflection information, may provide speed optimized methods
class Unit$Type extends MessageType<Unit> {
    constructor() {
        super("resources.dispatch.Unit", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "3", maxLen: "24" } } } },
            { no: 5, name: "initials", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "2", maxLen: "4" } } } },
            { no: 6, name: "color", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "6" } } } },
            { no: 7, name: "description", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "255" } } } },
            { no: 8, name: "status", kind: "message", T: () => UnitStatus },
            { no: 9, name: "users", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => UserShort }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.dispatch.Unit
 */
export const Unit = new Unit$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UnitStatus$Type extends MessageType<UnitStatus> {
    constructor() {
        super("resources.dispatch.UnitStatus", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "unit_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 3, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "status", kind: "enum", opt: true, T: () => ["resources.dispatch.UNIT_STATUS", UNIT_STATUS], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 5, name: "reason", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "255" } } } },
            { no: 6, name: "code", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 7, name: "user_id", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 8, name: "user", kind: "message", T: () => UserShort },
            { no: 9, name: "in_squad", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ },
            { no: 10, name: "x", kind: "scalar", opt: true, T: 2 /*ScalarType.FLOAT*/ },
            { no: 11, name: "y", kind: "scalar", opt: true, T: 2 /*ScalarType.FLOAT*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.dispatch.UnitStatus
 */
export const UnitStatus = new UnitStatus$Type();
