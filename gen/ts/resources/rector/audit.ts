// @generated by protobuf-ts 2.9.1 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "resources/rector/audit.proto" (package "resources.rector", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { UserShort } from "../users/users";
import { Timestamp } from "../timestamp/timestamp";
/**
 * @generated from protobuf message resources.rector.AuditEntry
 */
export interface AuditEntry {
    /**
     * @generated from protobuf field: uint64 id = 1 [jstype = JS_STRING];
     */
    id: string; // @gotags: alias:"id"
    /**
     * @generated from protobuf field: resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: uint64 user_id = 3 [jstype = JS_STRING];
     */
    userId: string; // @gotags: alias:"user_id"
    /**
     * @generated from protobuf field: optional resources.users.UserShort user = 4;
     */
    user?: UserShort;
    /**
     * @generated from protobuf field: string user_job = 5;
     */
    userJob: string; // @gotags: alias:"user_job"
    /**
     * @generated from protobuf field: optional int32 target_user_id = 6;
     */
    targetUserId?: number; // @gotags: alias:"target_user_id"
    /**
     * @generated from protobuf field: optional resources.users.UserShort target_user = 7;
     */
    targetUser?: UserShort;
    /**
     * @generated from protobuf field: string target_user_job = 8;
     */
    targetUserJob: string; // @gotags: alias:"target_user_job"
    /**
     * @generated from protobuf field: string service = 9;
     */
    service: string; // @gotags: alias:"service"
    /**
     * @generated from protobuf field: string method = 10;
     */
    method: string; // @gotags: alias:"method"
    /**
     * @generated from protobuf field: resources.rector.EventType state = 11;
     */
    state: EventType; // @gotags: alias:"state"
    /**
     * @generated from protobuf field: optional string data = 12;
     */
    data?: string; // @gotags: alias:"data"
}
/**
 * @generated from protobuf enum resources.rector.EventType
 */
export enum EventType {
    /**
     * @generated from protobuf enum value: EVENT_TYPE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: EVENT_TYPE_ERRORED = 1;
     */
    ERRORED = 1,
    /**
     * @generated from protobuf enum value: EVENT_TYPE_VIEWED = 2;
     */
    VIEWED = 2,
    /**
     * @generated from protobuf enum value: EVENT_TYPE_CREATED = 3;
     */
    CREATED = 3,
    /**
     * @generated from protobuf enum value: EVENT_TYPE_UPDATED = 4;
     */
    UPDATED = 4,
    /**
     * @generated from protobuf enum value: EVENT_TYPE_DELETED = 5;
     */
    DELETED = 5
}
// @generated message type with reflection information, may provide speed optimized methods
class AuditEntry$Type extends MessageType<AuditEntry> {
    constructor() {
        super("resources.rector.AuditEntry", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "user_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/ },
            { no: 4, name: "user", kind: "message", T: () => UserShort },
            { no: 5, name: "user_job", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 6, name: "target_user_id", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/ },
            { no: 7, name: "target_user", kind: "message", T: () => UserShort },
            { no: 8, name: "target_user_job", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 9, name: "service", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 10, name: "method", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 11, name: "state", kind: "enum", T: () => ["resources.rector.EventType", EventType, "EVENT_TYPE_"], options: { "validate.rules": { enum: { definedOnly: true } } } },
            { no: 12, name: "data", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.rector.AuditEntry
 */
export const AuditEntry = new AuditEntry$Type();
