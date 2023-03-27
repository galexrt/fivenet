// source: services/dmv/vehicles.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global =
    (typeof globalThis !== 'undefined' && globalThis) ||
    (typeof window !== 'undefined' && window) ||
    (typeof global !== 'undefined' && global) ||
    (typeof self !== 'undefined' && self) ||
    (function () { return this; }).call(null) ||
    Function('return this')();

var resources_common_database_database_pb = require('../../resources/common/database/database_pb.js');
goog.object.extend(proto, resources_common_database_database_pb);
var resources_vehicles_vehicles_pb = require('../../resources/vehicles/vehicles_pb.js');
goog.object.extend(proto, resources_vehicles_vehicles_pb);
goog.exportSymbol('proto.services.dmv.FindVehiclesRequest', null, global);
goog.exportSymbol('proto.services.dmv.FindVehiclesResponse', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.services.dmv.FindVehiclesRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.services.dmv.FindVehiclesRequest.repeatedFields_, null);
};
goog.inherits(proto.services.dmv.FindVehiclesRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.services.dmv.FindVehiclesRequest.displayName = 'proto.services.dmv.FindVehiclesRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.services.dmv.FindVehiclesResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.services.dmv.FindVehiclesResponse.repeatedFields_, null);
};
goog.inherits(proto.services.dmv.FindVehiclesResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.services.dmv.FindVehiclesResponse.displayName = 'proto.services.dmv.FindVehiclesResponse';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.services.dmv.FindVehiclesRequest.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.services.dmv.FindVehiclesRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.services.dmv.FindVehiclesRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.services.dmv.FindVehiclesRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.services.dmv.FindVehiclesRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    pagination: (f = msg.getPagination()) && resources_common_database_database_pb.PaginationRequest.toObject(includeInstance, f),
    orderbyList: jspb.Message.toObjectList(msg.getOrderbyList(),
    resources_common_database_database_pb.OrderBy.toObject, includeInstance),
    search: jspb.Message.getFieldWithDefault(msg, 3, ""),
    model: jspb.Message.getFieldWithDefault(msg, 4, ""),
    userId: jspb.Message.getFieldWithDefault(msg, 5, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.services.dmv.FindVehiclesRequest}
 */
proto.services.dmv.FindVehiclesRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.services.dmv.FindVehiclesRequest;
  return proto.services.dmv.FindVehiclesRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.services.dmv.FindVehiclesRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.services.dmv.FindVehiclesRequest}
 */
proto.services.dmv.FindVehiclesRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new resources_common_database_database_pb.PaginationRequest;
      reader.readMessage(value,resources_common_database_database_pb.PaginationRequest.deserializeBinaryFromReader);
      msg.setPagination(value);
      break;
    case 2:
      var value = new resources_common_database_database_pb.OrderBy;
      reader.readMessage(value,resources_common_database_database_pb.OrderBy.deserializeBinaryFromReader);
      msg.addOrderby(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setSearch(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setModel(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setUserId(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.services.dmv.FindVehiclesRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.services.dmv.FindVehiclesRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.services.dmv.FindVehiclesRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.services.dmv.FindVehiclesRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPagination();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      resources_common_database_database_pb.PaginationRequest.serializeBinaryToWriter
    );
  }
  f = message.getOrderbyList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      resources_common_database_database_pb.OrderBy.serializeBinaryToWriter
    );
  }
  f = message.getSearch();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getModel();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getUserId();
  if (f !== 0) {
    writer.writeInt32(
      5,
      f
    );
  }
};


/**
 * optional resources.common.database.PaginationRequest pagination = 1;
 * @return {?proto.resources.common.database.PaginationRequest}
 */
proto.services.dmv.FindVehiclesRequest.prototype.getPagination = function() {
  return /** @type{?proto.resources.common.database.PaginationRequest} */ (
    jspb.Message.getWrapperField(this, resources_common_database_database_pb.PaginationRequest, 1));
};


/**
 * @param {?proto.resources.common.database.PaginationRequest|undefined} value
 * @return {!proto.services.dmv.FindVehiclesRequest} returns this
*/
proto.services.dmv.FindVehiclesRequest.prototype.setPagination = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.services.dmv.FindVehiclesRequest} returns this
 */
proto.services.dmv.FindVehiclesRequest.prototype.clearPagination = function() {
  return this.setPagination(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.services.dmv.FindVehiclesRequest.prototype.hasPagination = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * repeated resources.common.database.OrderBy orderBy = 2;
 * @return {!Array<!proto.resources.common.database.OrderBy>}
 */
proto.services.dmv.FindVehiclesRequest.prototype.getOrderbyList = function() {
  return /** @type{!Array<!proto.resources.common.database.OrderBy>} */ (
    jspb.Message.getRepeatedWrapperField(this, resources_common_database_database_pb.OrderBy, 2));
};


/**
 * @param {!Array<!proto.resources.common.database.OrderBy>} value
 * @return {!proto.services.dmv.FindVehiclesRequest} returns this
*/
proto.services.dmv.FindVehiclesRequest.prototype.setOrderbyList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.resources.common.database.OrderBy=} opt_value
 * @param {number=} opt_index
 * @return {!proto.resources.common.database.OrderBy}
 */
proto.services.dmv.FindVehiclesRequest.prototype.addOrderby = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.resources.common.database.OrderBy, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.services.dmv.FindVehiclesRequest} returns this
 */
proto.services.dmv.FindVehiclesRequest.prototype.clearOrderbyList = function() {
  return this.setOrderbyList([]);
};


/**
 * optional string search = 3;
 * @return {string}
 */
proto.services.dmv.FindVehiclesRequest.prototype.getSearch = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.services.dmv.FindVehiclesRequest} returns this
 */
proto.services.dmv.FindVehiclesRequest.prototype.setSearch = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string model = 4;
 * @return {string}
 */
proto.services.dmv.FindVehiclesRequest.prototype.getModel = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.services.dmv.FindVehiclesRequest} returns this
 */
proto.services.dmv.FindVehiclesRequest.prototype.setModel = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional int32 user_id = 5;
 * @return {number}
 */
proto.services.dmv.FindVehiclesRequest.prototype.getUserId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.services.dmv.FindVehiclesRequest} returns this
 */
proto.services.dmv.FindVehiclesRequest.prototype.setUserId = function(value) {
  return jspb.Message.setProto3IntField(this, 5, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.services.dmv.FindVehiclesResponse.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.services.dmv.FindVehiclesResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.services.dmv.FindVehiclesResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.services.dmv.FindVehiclesResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.services.dmv.FindVehiclesResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    pagination: (f = msg.getPagination()) && resources_common_database_database_pb.PaginationResponse.toObject(includeInstance, f),
    vehiclesList: jspb.Message.toObjectList(msg.getVehiclesList(),
    resources_vehicles_vehicles_pb.Vehicle.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.services.dmv.FindVehiclesResponse}
 */
proto.services.dmv.FindVehiclesResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.services.dmv.FindVehiclesResponse;
  return proto.services.dmv.FindVehiclesResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.services.dmv.FindVehiclesResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.services.dmv.FindVehiclesResponse}
 */
proto.services.dmv.FindVehiclesResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new resources_common_database_database_pb.PaginationResponse;
      reader.readMessage(value,resources_common_database_database_pb.PaginationResponse.deserializeBinaryFromReader);
      msg.setPagination(value);
      break;
    case 2:
      var value = new resources_vehicles_vehicles_pb.Vehicle;
      reader.readMessage(value,resources_vehicles_vehicles_pb.Vehicle.deserializeBinaryFromReader);
      msg.addVehicles(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.services.dmv.FindVehiclesResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.services.dmv.FindVehiclesResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.services.dmv.FindVehiclesResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.services.dmv.FindVehiclesResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPagination();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      resources_common_database_database_pb.PaginationResponse.serializeBinaryToWriter
    );
  }
  f = message.getVehiclesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      resources_vehicles_vehicles_pb.Vehicle.serializeBinaryToWriter
    );
  }
};


/**
 * optional resources.common.database.PaginationResponse pagination = 1;
 * @return {?proto.resources.common.database.PaginationResponse}
 */
proto.services.dmv.FindVehiclesResponse.prototype.getPagination = function() {
  return /** @type{?proto.resources.common.database.PaginationResponse} */ (
    jspb.Message.getWrapperField(this, resources_common_database_database_pb.PaginationResponse, 1));
};


/**
 * @param {?proto.resources.common.database.PaginationResponse|undefined} value
 * @return {!proto.services.dmv.FindVehiclesResponse} returns this
*/
proto.services.dmv.FindVehiclesResponse.prototype.setPagination = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.services.dmv.FindVehiclesResponse} returns this
 */
proto.services.dmv.FindVehiclesResponse.prototype.clearPagination = function() {
  return this.setPagination(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.services.dmv.FindVehiclesResponse.prototype.hasPagination = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * repeated resources.vehicles.Vehicle vehicles = 2;
 * @return {!Array<!proto.resources.vehicles.Vehicle>}
 */
proto.services.dmv.FindVehiclesResponse.prototype.getVehiclesList = function() {
  return /** @type{!Array<!proto.resources.vehicles.Vehicle>} */ (
    jspb.Message.getRepeatedWrapperField(this, resources_vehicles_vehicles_pb.Vehicle, 2));
};


/**
 * @param {!Array<!proto.resources.vehicles.Vehicle>} value
 * @return {!proto.services.dmv.FindVehiclesResponse} returns this
*/
proto.services.dmv.FindVehiclesResponse.prototype.setVehiclesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.resources.vehicles.Vehicle=} opt_value
 * @param {number=} opt_index
 * @return {!proto.resources.vehicles.Vehicle}
 */
proto.services.dmv.FindVehiclesResponse.prototype.addVehicles = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.resources.vehicles.Vehicle, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.services.dmv.FindVehiclesResponse} returns this
 */
proto.services.dmv.FindVehiclesResponse.prototype.clearVehiclesList = function() {
  return this.setVehiclesList([]);
};


goog.object.extend(exports, proto.services.dmv);
