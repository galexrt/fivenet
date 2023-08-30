// @generated by protobuf-ts 2.9.1 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "resources/documents/templates.proto" (package "resources.documents", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { ACCESS_LEVEL } from "./access.js";
import { Vehicle } from "../vehicles/vehicles.js";
import { DocumentShort } from "./documents.js";
import { User } from "../users/users.js";
import { DocumentAccess } from "./documents.js";
import { UserShort } from "../users/users.js";
import { Category } from "./category.js";
import { Timestamp } from "../timestamp/timestamp.js";
/**
 * @generated from protobuf message resources.documents.Template
 */
export interface Template {
    /**
     * @generated from protobuf field: uint64 id = 1;
     */
    id: bigint; // @gotags: alias:"id"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp updated_at = 3;
     */
    updatedAt?: Timestamp;
    /**
     * @generated from protobuf field: resources.documents.Category category = 4;
     */
    category?: Category; // @gotags: alias:"category"
    /**
     * @generated from protobuf field: uint32 weight = 5;
     */
    weight: number; // @gotags: alias:"weight"
    /**
     * @generated from protobuf field: string title = 6;
     */
    title: string; // @gotags: alias:"title"
    /**
     * @generated from protobuf field: string description = 7;
     */
    description: string; // @gotags: alias:"description"
    /**
     * @generated from protobuf field: string content_title = 8;
     */
    contentTitle: string; // @gotags: alias:"content_title"
    /**
     * @generated from protobuf field: string content = 9;
     */
    content: string; // @gotags: alias:"content"
    /**
     * @generated from protobuf field: string state = 10;
     */
    state: string; // @gotags: alias:"state"
    /**
     * @generated from protobuf field: resources.documents.TemplateSchema schema = 11;
     */
    schema?: TemplateSchema; // @gotags: alias:"schema"
    /**
     * @generated from protobuf field: optional int32 creator_id = 12;
     */
    creatorId?: number; // @gotags: alias:"creator_id"
    /**
     * @generated from protobuf field: optional resources.users.UserShort creator = 13;
     */
    creator?: UserShort;
    /**
     * @generated from protobuf field: optional string job = 14;
     */
    job?: string; // @gotags: alias:"job"
    /**
     * @generated from protobuf field: repeated resources.documents.TemplateJobAccess job_access = 15;
     */
    jobAccess: TemplateJobAccess[];
    /**
     * @generated from protobuf field: resources.documents.DocumentAccess content_access = 16;
     */
    contentAccess?: DocumentAccess; // @gotags: alias:"access"
}
/**
 * @generated from protobuf message resources.documents.TemplateShort
 */
export interface TemplateShort {
    /**
     * @generated from protobuf field: uint64 id = 1;
     */
    id: bigint; // @gotags: alias:"id"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp updated_at = 3;
     */
    updatedAt?: Timestamp;
    /**
     * @generated from protobuf field: uint32 weight = 4;
     */
    weight: number; // @gotags: alias:"weight"
    /**
     * @generated from protobuf field: resources.documents.Category category = 5;
     */
    category?: Category; // @gotags: alias:"category"
    /**
     * @generated from protobuf field: string title = 6;
     */
    title: string; // @gotags: alias:"title"
    /**
     * @generated from protobuf field: string description = 7;
     */
    description: string; // @gotags: alias:"description"
    /**
     * @generated from protobuf field: resources.documents.TemplateSchema schema = 8;
     */
    schema?: TemplateSchema; // @gotags: alias:"schema"
    /**
     * @generated from protobuf field: optional int32 creator_id = 9;
     */
    creatorId?: number; // @gotags: alias:"creator_id"
    /**
     * @generated from protobuf field: optional resources.users.UserShort creator = 10;
     */
    creator?: UserShort;
    /**
     * @generated from protobuf field: string job = 11;
     */
    job: string; // @gotags: alias:"job"
}
/**
 * @generated from protobuf message resources.documents.TemplateSchema
 */
export interface TemplateSchema {
    /**
     * @generated from protobuf field: resources.documents.TemplateRequirements requirements = 1;
     */
    requirements?: TemplateRequirements;
}
/**
 * @generated from protobuf message resources.documents.TemplateRequirements
 */
export interface TemplateRequirements {
    /**
     * @generated from protobuf field: optional resources.documents.ObjectSpecs documents = 1;
     */
    documents?: ObjectSpecs;
    /**
     * @generated from protobuf field: optional resources.documents.ObjectSpecs users = 2;
     */
    users?: ObjectSpecs;
    /**
     * @generated from protobuf field: optional resources.documents.ObjectSpecs vehicles = 3;
     */
    vehicles?: ObjectSpecs;
}
/**
 * @generated from protobuf message resources.documents.ObjectSpecs
 */
export interface ObjectSpecs {
    /**
     * @generated from protobuf field: optional bool required = 1;
     */
    required?: boolean;
    /**
     * @generated from protobuf field: optional int64 min = 2;
     */
    min?: bigint;
    /**
     * @generated from protobuf field: optional int64 max = 3;
     */
    max?: bigint;
}
/**
 * @generated from protobuf message resources.documents.TemplateData
 */
export interface TemplateData {
    /**
     * @generated from protobuf field: resources.users.User activeChar = 1;
     */
    activeChar?: User;
    /**
     * @generated from protobuf field: repeated resources.documents.DocumentShort documents = 2;
     */
    documents: DocumentShort[];
    /**
     * @generated from protobuf field: repeated resources.users.User users = 3;
     */
    users: User[];
    /**
     * @generated from protobuf field: repeated resources.vehicles.Vehicle vehicles = 4;
     */
    vehicles: Vehicle[];
}
/**
 * @generated from protobuf message resources.documents.TemplateJobAccess
 */
export interface TemplateJobAccess {
    /**
     * @generated from protobuf field: uint64 id = 1;
     */
    id: bigint; // @gotags: alias:"id"
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp created_at = 2;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: optional resources.timestamp.Timestamp updated_at = 3;
     */
    updatedAt?: Timestamp;
    /**
     * @generated from protobuf field: uint64 template_id = 4;
     */
    templateId: bigint; // @gotags: alias:"template_id"
    /**
     * @generated from protobuf field: string job = 5;
     */
    job: string; // @gotags: alias:"job"
    /**
     * @generated from protobuf field: optional string job_label = 6;
     */
    jobLabel?: string; // @gotags: alias:"job_label"
    /**
     * @generated from protobuf field: int32 minimum_grade = 7;
     */
    minimumGrade: number; // @gotags: alias:"minimum_grade"
    /**
     * @generated from protobuf field: optional string job_grade_label = 8;
     */
    jobGradeLabel?: string; // @gotags: alias:"job_grade_label"
    /**
     * @generated from protobuf field: resources.documents.ACCESS_LEVEL access = 9;
     */
    access: ACCESS_LEVEL; // @gotags: alias:"access"
}
// @generated message type with reflection information, may provide speed optimized methods
class Template$Type extends MessageType<Template> {
    constructor() {
        super("resources.documents.Template", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "category", kind: "message", T: () => Category },
            { no: 5, name: "weight", kind: "scalar", T: 13 /*ScalarType.UINT32*/ },
            { no: 6, name: "title", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "3" } } } },
            { no: 7, name: "description", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "255" } } } },
            { no: 8, name: "content_title", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "3", maxBytes: "21845" } } } },
            { no: 9, name: "content", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "0", maxBytes: "1500000" } } } },
            { no: 10, name: "state", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "24" } } } },
            { no: 11, name: "schema", kind: "message", T: () => TemplateSchema },
            { no: 12, name: "creator_id", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/ },
            { no: 13, name: "creator", kind: "message", T: () => UserShort },
            { no: 14, name: "job", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 15, name: "job_access", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => TemplateJobAccess },
            { no: 16, name: "content_access", kind: "message", T: () => DocumentAccess }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.documents.Template
 */
export const Template = new Template$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TemplateShort$Type extends MessageType<TemplateShort> {
    constructor() {
        super("resources.documents.TemplateShort", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "weight", kind: "scalar", T: 13 /*ScalarType.UINT32*/ },
            { no: 5, name: "category", kind: "message", T: () => Category },
            { no: 6, name: "title", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { minLen: "3" } } } },
            { no: 7, name: "description", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "255" } } } },
            { no: 8, name: "schema", kind: "message", T: () => TemplateSchema },
            { no: 9, name: "creator_id", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/ },
            { no: 10, name: "creator", kind: "message", T: () => UserShort },
            { no: 11, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.documents.TemplateShort
 */
export const TemplateShort = new TemplateShort$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TemplateSchema$Type extends MessageType<TemplateSchema> {
    constructor() {
        super("resources.documents.TemplateSchema", [
            { no: 1, name: "requirements", kind: "message", T: () => TemplateRequirements }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.documents.TemplateSchema
 */
export const TemplateSchema = new TemplateSchema$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TemplateRequirements$Type extends MessageType<TemplateRequirements> {
    constructor() {
        super("resources.documents.TemplateRequirements", [
            { no: 1, name: "documents", kind: "message", T: () => ObjectSpecs },
            { no: 2, name: "users", kind: "message", T: () => ObjectSpecs },
            { no: 3, name: "vehicles", kind: "message", T: () => ObjectSpecs }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.documents.TemplateRequirements
 */
export const TemplateRequirements = new TemplateRequirements$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ObjectSpecs$Type extends MessageType<ObjectSpecs> {
    constructor() {
        super("resources.documents.ObjectSpecs", [
            { no: 1, name: "required", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "min", kind: "scalar", opt: true, T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 3, name: "max", kind: "scalar", opt: true, T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.documents.ObjectSpecs
 */
export const ObjectSpecs = new ObjectSpecs$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TemplateData$Type extends MessageType<TemplateData> {
    constructor() {
        super("resources.documents.TemplateData", [
            { no: 1, name: "activeChar", kind: "message", T: () => User },
            { no: 2, name: "documents", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => DocumentShort, options: { "validate.rules": { repeated: { maxItems: "5" } } } },
            { no: 3, name: "users", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => User, options: { "validate.rules": { repeated: { maxItems: "5" } } } },
            { no: 4, name: "vehicles", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Vehicle, options: { "validate.rules": { repeated: { maxItems: "5" } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.documents.TemplateData
 */
export const TemplateData = new TemplateData$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TemplateJobAccess$Type extends MessageType<TemplateJobAccess> {
    constructor() {
        super("resources.documents.TemplateJobAccess", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 3, name: "updated_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "template_id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 5, name: "job", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "20" } } } },
            { no: 6, name: "job_label", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 7, name: "minimum_grade", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gt: 0 } } } },
            { no: 8, name: "job_grade_label", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxLen: "50" } } } },
            { no: 9, name: "access", kind: "enum", T: () => ["resources.documents.ACCESS_LEVEL", ACCESS_LEVEL], options: { "validate.rules": { enum: { definedOnly: true } } } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.documents.TemplateJobAccess
 */
export const TemplateJobAccess = new TemplateJobAccess$Type();
