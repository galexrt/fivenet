// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size,long_type_bigint
// @generated from protobuf file "resources/vehicles/vehicles.proto" (package "resources.vehicles", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { UserShort } from "../users/users.js";
/**
 * @generated from protobuf message resources.vehicles.Vehicle
 */
export interface Vehicle {
    /**
     * @generated from protobuf field: string plate = 1;
     */
    plate: string;
    /**
     * @generated from protobuf field: string model = 2;
     */
    model: string;
    /**
     * @generated from protobuf field: string type = 3;
     */
    type: string;
    /**
     * @generated from protobuf field: resources.users.UserShort owner = 4;
     */
    owner?: UserShort;
}
// @generated message type with reflection information, may provide speed optimized methods
class Vehicle$Type extends MessageType<Vehicle> {
    constructor() {
        super("resources.vehicles.Vehicle", [
            { no: 1, name: "plate", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "model", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "type", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "owner", kind: "message", T: () => UserShort }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message resources.vehicles.Vehicle
 */
export const Vehicle = new Vehicle$Type();
