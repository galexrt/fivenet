// source: livemap/livemap.proto
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

var common_timestamp_timestamp_pb = require('../common/timestamp/timestamp_pb.js');
goog.object.extend(proto, common_timestamp_timestamp_pb);
goog.exportSymbol('proto.gen.livemap.Marker', null, global);
goog.exportSymbol('proto.gen.livemap.ServerStreamResponse', null, global);
goog.exportSymbol('proto.gen.livemap.StreamRequest', null, global);
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
proto.gen.livemap.StreamRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gen.livemap.StreamRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.gen.livemap.StreamRequest.displayName = 'proto.gen.livemap.StreamRequest';
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
proto.gen.livemap.ServerStreamResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.gen.livemap.ServerStreamResponse.repeatedFields_, null);
};
goog.inherits(proto.gen.livemap.ServerStreamResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.gen.livemap.ServerStreamResponse.displayName = 'proto.gen.livemap.ServerStreamResponse';
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
proto.gen.livemap.Marker = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gen.livemap.Marker, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.gen.livemap.Marker.displayName = 'proto.gen.livemap.Marker';
}



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
proto.gen.livemap.StreamRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.gen.livemap.StreamRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gen.livemap.StreamRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gen.livemap.StreamRequest.toObject = function(includeInstance, msg) {
  var f, obj = {

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
 * @return {!proto.gen.livemap.StreamRequest}
 */
proto.gen.livemap.StreamRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gen.livemap.StreamRequest;
  return proto.gen.livemap.StreamRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gen.livemap.StreamRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gen.livemap.StreamRequest}
 */
proto.gen.livemap.StreamRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
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
proto.gen.livemap.StreamRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gen.livemap.StreamRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gen.livemap.StreamRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gen.livemap.StreamRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.gen.livemap.ServerStreamResponse.repeatedFields_ = [1,2];



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
proto.gen.livemap.ServerStreamResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.gen.livemap.ServerStreamResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gen.livemap.ServerStreamResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gen.livemap.ServerStreamResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    usersList: jspb.Message.toObjectList(msg.getUsersList(),
    proto.gen.livemap.Marker.toObject, includeInstance),
    dispatchesList: jspb.Message.toObjectList(msg.getDispatchesList(),
    proto.gen.livemap.Marker.toObject, includeInstance)
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
 * @return {!proto.gen.livemap.ServerStreamResponse}
 */
proto.gen.livemap.ServerStreamResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gen.livemap.ServerStreamResponse;
  return proto.gen.livemap.ServerStreamResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gen.livemap.ServerStreamResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gen.livemap.ServerStreamResponse}
 */
proto.gen.livemap.ServerStreamResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.gen.livemap.Marker;
      reader.readMessage(value,proto.gen.livemap.Marker.deserializeBinaryFromReader);
      msg.addUsers(value);
      break;
    case 2:
      var value = new proto.gen.livemap.Marker;
      reader.readMessage(value,proto.gen.livemap.Marker.deserializeBinaryFromReader);
      msg.addDispatches(value);
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
proto.gen.livemap.ServerStreamResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gen.livemap.ServerStreamResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gen.livemap.ServerStreamResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gen.livemap.ServerStreamResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUsersList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.gen.livemap.Marker.serializeBinaryToWriter
    );
  }
  f = message.getDispatchesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.gen.livemap.Marker.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Marker users = 1;
 * @return {!Array<!proto.gen.livemap.Marker>}
 */
proto.gen.livemap.ServerStreamResponse.prototype.getUsersList = function() {
  return /** @type{!Array<!proto.gen.livemap.Marker>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.gen.livemap.Marker, 1));
};


/**
 * @param {!Array<!proto.gen.livemap.Marker>} value
 * @return {!proto.gen.livemap.ServerStreamResponse} returns this
*/
proto.gen.livemap.ServerStreamResponse.prototype.setUsersList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.gen.livemap.Marker=} opt_value
 * @param {number=} opt_index
 * @return {!proto.gen.livemap.Marker}
 */
proto.gen.livemap.ServerStreamResponse.prototype.addUsers = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.gen.livemap.Marker, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.gen.livemap.ServerStreamResponse} returns this
 */
proto.gen.livemap.ServerStreamResponse.prototype.clearUsersList = function() {
  return this.setUsersList([]);
};


/**
 * repeated Marker dispatches = 2;
 * @return {!Array<!proto.gen.livemap.Marker>}
 */
proto.gen.livemap.ServerStreamResponse.prototype.getDispatchesList = function() {
  return /** @type{!Array<!proto.gen.livemap.Marker>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.gen.livemap.Marker, 2));
};


/**
 * @param {!Array<!proto.gen.livemap.Marker>} value
 * @return {!proto.gen.livemap.ServerStreamResponse} returns this
*/
proto.gen.livemap.ServerStreamResponse.prototype.setDispatchesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.gen.livemap.Marker=} opt_value
 * @param {number=} opt_index
 * @return {!proto.gen.livemap.Marker}
 */
proto.gen.livemap.ServerStreamResponse.prototype.addDispatches = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.gen.livemap.Marker, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.gen.livemap.ServerStreamResponse} returns this
 */
proto.gen.livemap.ServerStreamResponse.prototype.clearDispatchesList = function() {
  return this.setDispatchesList([]);
};





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
proto.gen.livemap.Marker.prototype.toObject = function(opt_includeInstance) {
  return proto.gen.livemap.Marker.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gen.livemap.Marker} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gen.livemap.Marker.toObject = function(includeInstance, msg) {
  var f, obj = {
    userid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    job: jspb.Message.getFieldWithDefault(msg, 2, ""),
    x: jspb.Message.getFloatingPointFieldWithDefault(msg, 3, 0.0),
    y: jspb.Message.getFloatingPointFieldWithDefault(msg, 4, 0.0),
    createdat: (f = msg.getCreatedat()) && common_timestamp_timestamp_pb.Timestamp.toObject(includeInstance, f),
    name: jspb.Message.getFieldWithDefault(msg, 6, ""),
    icon: jspb.Message.getFieldWithDefault(msg, 7, ""),
    popup: jspb.Message.getFieldWithDefault(msg, 8, ""),
    link: jspb.Message.getFieldWithDefault(msg, 9, "")
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
 * @return {!proto.gen.livemap.Marker}
 */
proto.gen.livemap.Marker.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gen.livemap.Marker;
  return proto.gen.livemap.Marker.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gen.livemap.Marker} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gen.livemap.Marker}
 */
proto.gen.livemap.Marker.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setUserid(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setJob(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setX(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setY(value);
      break;
    case 5:
      var value = new common_timestamp_timestamp_pb.Timestamp;
      reader.readMessage(value,common_timestamp_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setCreatedat(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setIcon(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setPopup(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.setLink(value);
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
proto.gen.livemap.Marker.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gen.livemap.Marker.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gen.livemap.Marker} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gen.livemap.Marker.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserid();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getJob();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getX();
  if (f !== 0.0) {
    writer.writeFloat(
      3,
      f
    );
  }
  f = message.getY();
  if (f !== 0.0) {
    writer.writeFloat(
      4,
      f
    );
  }
  f = message.getCreatedat();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      common_timestamp_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getIcon();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getPopup();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
  f = message.getLink();
  if (f.length > 0) {
    writer.writeString(
      9,
      f
    );
  }
};


/**
 * optional int32 userID = 1;
 * @return {number}
 */
proto.gen.livemap.Marker.prototype.getUserid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.gen.livemap.Marker} returns this
 */
proto.gen.livemap.Marker.prototype.setUserid = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string job = 2;
 * @return {string}
 */
proto.gen.livemap.Marker.prototype.getJob = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.gen.livemap.Marker} returns this
 */
proto.gen.livemap.Marker.prototype.setJob = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional float x = 3;
 * @return {number}
 */
proto.gen.livemap.Marker.prototype.getX = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 3, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.gen.livemap.Marker} returns this
 */
proto.gen.livemap.Marker.prototype.setX = function(value) {
  return jspb.Message.setProto3FloatField(this, 3, value);
};


/**
 * optional float y = 4;
 * @return {number}
 */
proto.gen.livemap.Marker.prototype.getY = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 4, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.gen.livemap.Marker} returns this
 */
proto.gen.livemap.Marker.prototype.setY = function(value) {
  return jspb.Message.setProto3FloatField(this, 4, value);
};


/**
 * optional gen.common.timestamp.Timestamp createdAt = 5;
 * @return {?proto.gen.common.timestamp.Timestamp}
 */
proto.gen.livemap.Marker.prototype.getCreatedat = function() {
  return /** @type{?proto.gen.common.timestamp.Timestamp} */ (
    jspb.Message.getWrapperField(this, common_timestamp_timestamp_pb.Timestamp, 5));
};


/**
 * @param {?proto.gen.common.timestamp.Timestamp|undefined} value
 * @return {!proto.gen.livemap.Marker} returns this
*/
proto.gen.livemap.Marker.prototype.setCreatedat = function(value) {
  return jspb.Message.setWrapperField(this, 5, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.gen.livemap.Marker} returns this
 */
proto.gen.livemap.Marker.prototype.clearCreatedat = function() {
  return this.setCreatedat(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.gen.livemap.Marker.prototype.hasCreatedat = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional string name = 6;
 * @return {string}
 */
proto.gen.livemap.Marker.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.gen.livemap.Marker} returns this
 */
proto.gen.livemap.Marker.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional string icon = 7;
 * @return {string}
 */
proto.gen.livemap.Marker.prototype.getIcon = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * @param {string} value
 * @return {!proto.gen.livemap.Marker} returns this
 */
proto.gen.livemap.Marker.prototype.setIcon = function(value) {
  return jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional string popup = 8;
 * @return {string}
 */
proto.gen.livemap.Marker.prototype.getPopup = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/**
 * @param {string} value
 * @return {!proto.gen.livemap.Marker} returns this
 */
proto.gen.livemap.Marker.prototype.setPopup = function(value) {
  return jspb.Message.setProto3StringField(this, 8, value);
};


/**
 * optional string link = 9;
 * @return {string}
 */
proto.gen.livemap.Marker.prototype.getLink = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 9, ""));
};


/**
 * @param {string} value
 * @return {!proto.gen.livemap.Marker} returns this
 */
proto.gen.livemap.Marker.prototype.setLink = function(value) {
  return jspb.Message.setProto3StringField(this, 9, value);
};


goog.object.extend(exports, proto.gen.livemap);
