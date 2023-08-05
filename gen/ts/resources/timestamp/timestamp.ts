// @generated by protobuf-ts 2.9.1 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "resources/timestamp/timestamp.proto" (package "resources.timestamp", syntax proto3)
// tslint:disable
//
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
import { MessageType } from "@protobuf-ts/runtime";
import { Timestamp as Timestamp$ } from "../../google/protobuf/timestamp.js";
/**
 * Timestamp for storage messages.  We've defined a new local type wrapper
 * of google.protobuf.Timestamp so we can implement sql.Scanner and sql.Valuer
 * interfaces.  See:
 * https://golang.org/pkg/database/sql/#Scanner
 * https://golang.org/pkg/database/sql/driver/#Valuer
 *
 * @generated from protobuf message resources.timestamp.Timestamp
 */
export interface Timestamp {
    /**
     * @generated from protobuf field: google.protobuf.Timestamp timestamp = 1;
     */
    timestamp?: Timestamp$;
}
// @generated message type with reflection information, may provide speed optimized methods
class Timestamp$Type extends MessageType<Timestamp> {
    constructor() {
        super("resources.timestamp.Timestamp", [
            { no: 1, name: "timestamp", kind: "message", T: () => Timestamp$ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.timestamp.Timestamp
 */
export const Timestamp = new Timestamp$Type();
