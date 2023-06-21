export function encodeError(message) {
  let bb = popByteBuffer();
  _encodeError(message, bb);
  return toUint8Array(bb);
}

function _encodeError(message, bb) {
  // optional int32 code = 1;
  let $code = message.code;
  if ($code !== undefined) {
    writeVarint32(bb, 8);
    writeVarint64(bb, intToLong($code));
  }

  // optional string description = 2;
  let $description = message.description;
  if ($description !== undefined) {
    writeVarint32(bb, 18);
    writeString(bb, $description);
  }
}

export function decodeError(binary) {
  return _decodeError(wrapByteBuffer(binary));
}

function _decodeError(bb) {
  let message = {};

  end_of_message: while (!isAtEnd(bb)) {
    let tag = readVarint32(bb);

    switch (tag >>> 3) {
      case 0:
        break end_of_message;

      // optional int32 code = 1;
      case 1: {
        message.code = readVarint32(bb);
        break;
      }

      // optional string description = 2;
      case 2: {
        message.description = readString(bb, readVarint32(bb));
        break;
      }

      default:
        skipUnknownField(bb, tag & 7);
    }
  }

  return message;
}

export function encodeCallbackStatus(message) {
  let bb = popByteBuffer();
  _encodeCallbackStatus(message, bb);
  return toUint8Array(bb);
}

function _encodeCallbackStatus(message, bb) {
  // optional int32 code = 1;
  let $code = message.code;
  if ($code !== undefined) {
    writeVarint32(bb, 8);
    writeVarint64(bb, intToLong($code));
  }

  // optional string description = 2;
  let $description = message.description;
  if ($description !== undefined) {
    writeVarint32(bb, 18);
    writeString(bb, $description);
  }

  // optional string updated_on = 3;
  let $updated_on = message.updated_on;
  if ($updated_on !== undefined) {
    writeVarint32(bb, 26);
    writeString(bb, $updated_on);
  }

  // optional string reference_id = 4;
  let $reference_id = message.reference_id;
  if ($reference_id !== undefined) {
    writeVarint32(bb, 34);
    writeString(bb, $reference_id);
  }

  // optional string last_channel = 5;
  let $last_channel = message.last_channel;
  if ($last_channel !== undefined) {
    writeVarint32(bb, 42);
    writeString(bb, $last_channel);
  }
}

export function decodeCallbackStatus(binary) {
  return _decodeCallbackStatus(wrapByteBuffer(binary));
}

function _decodeCallbackStatus(bb) {
  let message = {};

  end_of_message: while (!isAtEnd(bb)) {
    let tag = readVarint32(bb);

    switch (tag >>> 3) {
      case 0:
        break end_of_message;

      // optional int32 code = 1;
      case 1: {
        message.code = readVarint32(bb);
        break;
      }

      // optional string description = 2;
      case 2: {
        message.description = readString(bb, readVarint32(bb));
        break;
      }

      // optional string updated_on = 3;
      case 3: {
        message.updated_on = readString(bb, readVarint32(bb));
        break;
      }

      // optional string reference_id = 4;
      case 4: {
        message.reference_id = readString(bb, readVarint32(bb));
        break;
      }

      // optional string last_channel = 5;
      case 5: {
        message.last_channel = readString(bb, readVarint32(bb));
        break;
      }

      default:
        skipUnknownField(bb, tag & 7);
    }
  }

  return message;
}

export function encodeChannelStatus(message) {
  let bb = popByteBuffer();
  _encodeChannelStatus(message, bb);
  return toUint8Array(bb);
}

function _encodeChannelStatus(message, bb) {
  // optional map<string, CallbackStatus> channel_statuses = 1;
  let map$channel_statuses = message.channel_statuses;
  if (map$channel_statuses !== undefined) {
    for (let key in map$channel_statuses) {
      let nested = popByteBuffer();
      let value = map$channel_statuses[key];
      writeVarint32(nested, 10);
      writeString(nested, key);
      writeVarint32(nested, 18);
      let nestedValue = popByteBuffer();
      _encodeCallbackStatus(value, nestedValue);
      writeVarint32(nested, nestedValue.limit);
      writeByteBuffer(nested, nestedValue);
      pushByteBuffer(nestedValue);
      writeVarint32(bb, 10);
      writeVarint32(bb, nested.offset);
      writeByteBuffer(bb, nested);
      pushByteBuffer(nested);
    }
  }
}

export function decodeChannelStatus(binary) {
  return _decodeChannelStatus(wrapByteBuffer(binary));
}

function _decodeChannelStatus(bb) {
  let message = {};

  end_of_message: while (!isAtEnd(bb)) {
    let tag = readVarint32(bb);

    switch (tag >>> 3) {
      case 0:
        break end_of_message;

      // optional map<string, CallbackStatus> channel_statuses = 1;
      case 1: {
        let values = message.channel_statuses || (message.channel_statuses = {});
        let outerLimit = pushTemporaryLength(bb);
        let key;
        let value;
        end_of_entry: while (!isAtEnd(bb)) {
          let tag = readVarint32(bb);
          switch (tag >>> 3) {
            case 0:
              break end_of_entry;
            case 1: {
              key = readString(bb, readVarint32(bb));
              break;
            }
            case 2: {
              let valueLimit = pushTemporaryLength(bb);
              value = _decodeCallbackStatus(bb);
              bb.limit = valueLimit;
              break;
            }
            default:
              skipUnknownField(bb, tag & 7);
          }
        }
        if (key === undefined || value === undefined)
          throw new Error("Invalid data for map: channel_statuses");
        values[key] = value;
        bb.limit = outerLimit;
        break;
      }

      default:
        skipUnknownField(bb, tag & 7);
    }
  }

  return message;
}

export function encodeCallbackAdditionalInfo(message) {
  let bb = popByteBuffer();
  _encodeCallbackAdditionalInfo(message, bb);
  return toUint8Array(bb);
}

function _encodeCallbackAdditionalInfo(message, bb) {
  // optional string mcc = 1;
  let $mcc = message.mcc;
  if ($mcc !== undefined) {
    writeVarint32(bb, 10);
    writeString(bb, $mcc);
  }

  // optional string mnc = 2;
  let $mnc = message.mnc;
  if ($mnc !== undefined) {
    writeVarint32(bb, 18);
    writeString(bb, $mnc);
  }

  // optional string price = 3;
  let $price = message.price;
  if ($price !== undefined) {
    writeVarint32(bb, 26);
    writeString(bb, $price);
  }

  // optional string currency = 4;
  let $currency = message.currency;
  if ($currency !== undefined) {
    writeVarint32(bb, 34);
    writeString(bb, $currency);
  }

  // optional int32 message_parts_count = 5;
  let $message_parts_count = message.message_parts_count;
  if ($message_parts_count !== undefined) {
    writeVarint32(bb, 40);
    writeVarint64(bb, intToLong($message_parts_count));
  }

  // optional string message_parts_reference_ids = 6;
  let $message_parts_reference_ids = message.message_parts_reference_ids;
  if ($message_parts_reference_ids !== undefined) {
    writeVarint32(bb, 50);
    writeString(bb, $message_parts_reference_ids);
  }
}

export function decodeCallbackAdditionalInfo(binary) {
  return _decodeCallbackAdditionalInfo(wrapByteBuffer(binary));
}

function _decodeCallbackAdditionalInfo(bb) {
  let message = {};

  end_of_message: while (!isAtEnd(bb)) {
    let tag = readVarint32(bb);

    switch (tag >>> 3) {
      case 0:
        break end_of_message;

      // optional string mcc = 1;
      case 1: {
        message.mcc = readString(bb, readVarint32(bb));
        break;
      }

      // optional string mnc = 2;
      case 2: {
        message.mnc = readString(bb, readVarint32(bb));
        break;
      }

      // optional string price = 3;
      case 3: {
        message.price = readString(bb, readVarint32(bb));
        break;
      }

      // optional string currency = 4;
      case 4: {
        message.currency = readString(bb, readVarint32(bb));
        break;
      }

      // optional int32 message_parts_count = 5;
      case 5: {
        message.message_parts_count = readVarint32(bb);
        break;
      }

      // optional string message_parts_reference_ids = 6;
      case 6: {
        message.message_parts_reference_ids = readString(bb, readVarint32(bb));
        break;
      }

      default:
        skipUnknownField(bb, tag & 7);
    }
  }

  return message;
}

export function encodeUserResponse(message) {
  let bb = popByteBuffer();
  _encodeUserResponse(message, bb);
  return toUint8Array(bb);
}

function _encodeUserResponse(message, bb) {
  // optional string phone_number = 1;
  let $phone_number = message.phone_number;
  if ($phone_number !== undefined) {
    writeVarint32(bb, 10);
    writeString(bb, $phone_number);
  }

  // optional string iso_country_code = 2;
  let $iso_country_code = message.iso_country_code;
  if ($iso_country_code !== undefined) {
    writeVarint32(bb, 18);
    writeString(bb, $iso_country_code);
  }

  // optional string sender_id = 3;
  let $sender_id = message.sender_id;
  if ($sender_id !== undefined) {
    writeVarint32(bb, 26);
    writeString(bb, $sender_id);
  }

  // optional string mo_message = 4;
  let $mo_message = message.mo_message;
  if ($mo_message !== undefined) {
    writeVarint32(bb, 34);
    writeString(bb, $mo_message);
  }
}

export function decodeUserResponse(binary) {
  return _decodeUserResponse(wrapByteBuffer(binary));
}

function _decodeUserResponse(bb) {
  let message = {};

  end_of_message: while (!isAtEnd(bb)) {
    let tag = readVarint32(bb);

    switch (tag >>> 3) {
      case 0:
        break end_of_message;

      // optional string phone_number = 1;
      case 1: {
        message.phone_number = readString(bb, readVarint32(bb));
        break;
      }

      // optional string iso_country_code = 2;
      case 2: {
        message.iso_country_code = readString(bb, readVarint32(bb));
        break;
      }

      // optional string sender_id = 3;
      case 3: {
        message.sender_id = readString(bb, readVarint32(bb));
        break;
      }

      // optional string mo_message = 4;
      case 4: {
        message.mo_message = readString(bb, readVarint32(bb));
        break;
      }

      default:
        skipUnknownField(bb, tag & 7);
    }
  }

  return message;
}

export function encodeOutgoingCallbackRequest(message) {
  let bb = popByteBuffer();
  _encodeOutgoingCallbackRequest(message, bb);
  return toUint8Array(bb);
}

function _encodeOutgoingCallbackRequest(message, bb) {
  // optional string reference_id = 1;
  let $reference_id = message.reference_id;
  if ($reference_id !== undefined) {
    writeVarint32(bb, 10);
    writeString(bb, $reference_id);
  }

  // optional CallbackStatus status = 2;
  let $status = message.status;
  if ($status !== undefined) {
    writeVarint32(bb, 18);
    let nested = popByteBuffer();
    _encodeCallbackStatus($status, nested);
    writeVarint32(bb, nested.limit);
    writeByteBuffer(bb, nested);
    pushByteBuffer(nested);
  }

  // optional string submit_timestamp = 3;
  let $submit_timestamp = message.submit_timestamp;
  if ($submit_timestamp !== undefined) {
    writeVarint32(bb, 26);
    writeString(bb, $submit_timestamp);
  }

  // repeated Error errors = 4;
  let array$errors = message.errors;
  if (array$errors !== undefined) {
    for (let value of array$errors) {
      writeVarint32(bb, 34);
      let nested = popByteBuffer();
      _encodeError(value, nested);
      writeVarint32(bb, nested.limit);
      writeByteBuffer(bb, nested);
      pushByteBuffer(nested);
    }
  }

  // optional string sub_resource = 5;
  let $sub_resource = message.sub_resource;
  if ($sub_resource !== undefined) {
    writeVarint32(bb, 42);
    writeString(bb, $sub_resource);
  }

  // optional CallbackAdditionalInfo additional_info = 6;
  let $additional_info = message.additional_info;
  if ($additional_info !== undefined) {
    writeVarint32(bb, 50);
    let nested = popByteBuffer();
    _encodeCallbackAdditionalInfo($additional_info, nested);
    writeVarint32(bb, nested.limit);
    writeByteBuffer(bb, nested);
    pushByteBuffer(nested);
  }

  // optional string external_id = 7;
  let $external_id = message.external_id;
  if ($external_id !== undefined) {
    writeVarint32(bb, 58);
    writeString(bb, $external_id);
  }
}

export function decodeOutgoingCallbackRequest(binary) {
  return _decodeOutgoingCallbackRequest(wrapByteBuffer(binary));
}

function _decodeOutgoingCallbackRequest(bb) {
  let message = {};

  end_of_message: while (!isAtEnd(bb)) {
    let tag = readVarint32(bb);

    switch (tag >>> 3) {
      case 0:
        break end_of_message;

      // optional string reference_id = 1;
      case 1: {
        message.reference_id = readString(bb, readVarint32(bb));
        break;
      }

      // optional CallbackStatus status = 2;
      case 2: {
        let limit = pushTemporaryLength(bb);
        message.status = _decodeCallbackStatus(bb);
        bb.limit = limit;
        break;
      }

      // optional string submit_timestamp = 3;
      case 3: {
        message.submit_timestamp = readString(bb, readVarint32(bb));
        break;
      }

      // repeated Error errors = 4;
      case 4: {
        let limit = pushTemporaryLength(bb);
        let values = message.errors || (message.errors = []);
        values.push(_decodeError(bb));
        bb.limit = limit;
        break;
      }

      // optional string sub_resource = 5;
      case 5: {
        message.sub_resource = readString(bb, readVarint32(bb));
        break;
      }

      // optional CallbackAdditionalInfo additional_info = 6;
      case 6: {
        let limit = pushTemporaryLength(bb);
        message.additional_info = _decodeCallbackAdditionalInfo(bb);
        bb.limit = limit;
        break;
      }

      // optional string external_id = 7;
      case 7: {
        message.external_id = readString(bb, readVarint32(bb));
        break;
      }

      default:
        skipUnknownField(bb, tag & 7);
    }
  }

  return message;
}

export function encodeIncomingSMSCallbackRequest(message) {
  let bb = popByteBuffer();
  _encodeIncomingSMSCallbackRequest(message, bb);
  return toUint8Array(bb);
}

function _encodeIncomingSMSCallbackRequest(message, bb) {
  // optional string reference_id = 1;
  let $reference_id = message.reference_id;
  if ($reference_id !== undefined) {
    writeVarint32(bb, 10);
    writeString(bb, $reference_id);
  }

  // optional string sub_resource = 2;
  let $sub_resource = message.sub_resource;
  if ($sub_resource !== undefined) {
    writeVarint32(bb, 18);
    writeString(bb, $sub_resource);
  }

  // repeated Error errors = 3;
  let array$errors = message.errors;
  if (array$errors !== undefined) {
    for (let value of array$errors) {
      writeVarint32(bb, 26);
      let nested = popByteBuffer();
      _encodeError(value, nested);
      writeVarint32(bb, nested.limit);
      writeByteBuffer(bb, nested);
      pushByteBuffer(nested);
    }
  }

  // optional CallbackStatus status = 4;
  let $status = message.status;
  if ($status !== undefined) {
    writeVarint32(bb, 34);
    let nested = popByteBuffer();
    _encodeCallbackStatus($status, nested);
    writeVarint32(bb, nested.limit);
    writeByteBuffer(bb, nested);
    pushByteBuffer(nested);
  }

  // optional string submit_timestamp = 5;
  let $submit_timestamp = message.submit_timestamp;
  if ($submit_timestamp !== undefined) {
    writeVarint32(bb, 42);
    writeString(bb, $submit_timestamp);
  }

  // optional UserResponse user_response = 6;
  let $user_response = message.user_response;
  if ($user_response !== undefined) {
    writeVarint32(bb, 50);
    let nested = popByteBuffer();
    _encodeUserResponse($user_response, nested);
    writeVarint32(bb, nested.limit);
    writeByteBuffer(bb, nested);
    pushByteBuffer(nested);
  }

  // optional string mo_price = 7;
  let $mo_price = message.mo_price;
  if ($mo_price !== undefined) {
    writeVarint32(bb, 58);
    writeString(bb, $mo_price);
  }
}

export function decodeIncomingSMSCallbackRequest(binary) {
  return _decodeIncomingSMSCallbackRequest(wrapByteBuffer(binary));
}

function _decodeIncomingSMSCallbackRequest(bb) {
  let message = {};

  end_of_message: while (!isAtEnd(bb)) {
    let tag = readVarint32(bb);

    switch (tag >>> 3) {
      case 0:
        break end_of_message;

      // optional string reference_id = 1;
      case 1: {
        message.reference_id = readString(bb, readVarint32(bb));
        break;
      }

      // optional string sub_resource = 2;
      case 2: {
        message.sub_resource = readString(bb, readVarint32(bb));
        break;
      }

      // repeated Error errors = 3;
      case 3: {
        let limit = pushTemporaryLength(bb);
        let values = message.errors || (message.errors = []);
        values.push(_decodeError(bb));
        bb.limit = limit;
        break;
      }

      // optional CallbackStatus status = 4;
      case 4: {
        let limit = pushTemporaryLength(bb);
        message.status = _decodeCallbackStatus(bb);
        bb.limit = limit;
        break;
      }

      // optional string submit_timestamp = 5;
      case 5: {
        message.submit_timestamp = readString(bb, readVarint32(bb));
        break;
      }

      // optional UserResponse user_response = 6;
      case 6: {
        let limit = pushTemporaryLength(bb);
        message.user_response = _decodeUserResponse(bb);
        bb.limit = limit;
        break;
      }

      // optional string mo_price = 7;
      case 7: {
        message.mo_price = readString(bb, readVarint32(bb));
        break;
      }

      default:
        skipUnknownField(bb, tag & 7);
    }
  }

  return message;
}

export function encodeMaskedNumberCallbackRequest(message) {
  let bb = popByteBuffer();
  _encodeMaskedNumberCallbackRequest(message, bb);
  return toUint8Array(bb);
}

function _encodeMaskedNumberCallbackRequest(message, bb) {
  // optional string submit_timestamp = 1;
  let $submit_timestamp = message.submit_timestamp;
  if ($submit_timestamp !== undefined) {
    writeVarint32(bb, 10);
    writeString(bb, $submit_timestamp);
  }

  // optional string from_phone_number = 2;
  let $from_phone_number = message.from_phone_number;
  if ($from_phone_number !== undefined) {
    writeVarint32(bb, 18);
    writeString(bb, $from_phone_number);
  }

  // optional string mo_message = 3;
  let $mo_message = message.mo_message;
  if ($mo_message !== undefined) {
    writeVarint32(bb, 26);
    writeString(bb, $mo_message);
  }

  // optional string mo_reference_id = 4;
  let $mo_reference_id = message.mo_reference_id;
  if ($mo_reference_id !== undefined) {
    writeVarint32(bb, 34);
    writeString(bb, $mo_reference_id);
  }

  // optional string session_reference_id = 5;
  let $session_reference_id = message.session_reference_id;
  if ($session_reference_id !== undefined) {
    writeVarint32(bb, 42);
    writeString(bb, $session_reference_id);
  }

  // optional string to_phone_number = 6;
  let $to_phone_number = message.to_phone_number;
  if ($to_phone_number !== undefined) {
    writeVarint32(bb, 50);
    writeString(bb, $to_phone_number);
  }
}

export function decodeMaskedNumberCallbackRequest(binary) {
  return _decodeMaskedNumberCallbackRequest(wrapByteBuffer(binary));
}

function _decodeMaskedNumberCallbackRequest(bb) {
  let message = {};

  end_of_message: while (!isAtEnd(bb)) {
    let tag = readVarint32(bb);

    switch (tag >>> 3) {
      case 0:
        break end_of_message;

      // optional string submit_timestamp = 1;
      case 1: {
        message.submit_timestamp = readString(bb, readVarint32(bb));
        break;
      }

      // optional string from_phone_number = 2;
      case 2: {
        message.from_phone_number = readString(bb, readVarint32(bb));
        break;
      }

      // optional string mo_message = 3;
      case 3: {
        message.mo_message = readString(bb, readVarint32(bb));
        break;
      }

      // optional string mo_reference_id = 4;
      case 4: {
        message.mo_reference_id = readString(bb, readVarint32(bb));
        break;
      }

      // optional string session_reference_id = 5;
      case 5: {
        message.session_reference_id = readString(bb, readVarint32(bb));
        break;
      }

      // optional string to_phone_number = 6;
      case 6: {
        message.to_phone_number = readString(bb, readVarint32(bb));
        break;
      }

      default:
        skipUnknownField(bb, tag & 7);
    }
  }

  return message;
}

export function encodeMultiChannelCallbackRequest(message) {
  let bb = popByteBuffer();
  _encodeMultiChannelCallbackRequest(message, bb);
  return toUint8Array(bb);
}

function _encodeMultiChannelCallbackRequest(message, bb) {
  // optional string receiver_id = 1;
  let $receiver_id = message.receiver_id;
  if ($receiver_id !== undefined) {
    writeVarint32(bb, 10);
    writeString(bb, $receiver_id);
  }

  // optional string external_id = 2;
  let $external_id = message.external_id;
  if ($external_id !== undefined) {
    writeVarint32(bb, 18);
    writeString(bb, $external_id);
  }

  // optional CallbackStatus status = 3;
  let $status = message.status;
  if ($status !== undefined) {
    writeVarint32(bb, 26);
    let nested = popByteBuffer();
    _encodeCallbackStatus($status, nested);
    writeVarint32(bb, nested.limit);
    writeByteBuffer(bb, nested);
    pushByteBuffer(nested);
  }

  // repeated ChannelStatus channel_status = 4;
  let array$channel_status = message.channel_status;
  if (array$channel_status !== undefined) {
    for (let value of array$channel_status) {
      writeVarint32(bb, 34);
      let nested = popByteBuffer();
      _encodeChannelStatus(value, nested);
      writeVarint32(bb, nested.limit);
      writeByteBuffer(bb, nested);
      pushByteBuffer(nested);
    }
  }
}

export function decodeMultiChannelCallbackRequest(binary) {
  return _decodeMultiChannelCallbackRequest(wrapByteBuffer(binary));
}

function _decodeMultiChannelCallbackRequest(bb) {
  let message = {};

  end_of_message: while (!isAtEnd(bb)) {
    let tag = readVarint32(bb);

    switch (tag >>> 3) {
      case 0:
        break end_of_message;

      // optional string receiver_id = 1;
      case 1: {
        message.receiver_id = readString(bb, readVarint32(bb));
        break;
      }

      // optional string external_id = 2;
      case 2: {
        message.external_id = readString(bb, readVarint32(bb));
        break;
      }

      // optional CallbackStatus status = 3;
      case 3: {
        let limit = pushTemporaryLength(bb);
        message.status = _decodeCallbackStatus(bb);
        bb.limit = limit;
        break;
      }

      // repeated ChannelStatus channel_status = 4;
      case 4: {
        let limit = pushTemporaryLength(bb);
        let values = message.channel_status || (message.channel_status = []);
        values.push(_decodeChannelStatus(bb));
        bb.limit = limit;
        break;
      }

      default:
        skipUnknownField(bb, tag & 7);
    }
  }

  return message;
}

export function encodeCommsAppointmentInput(message) {
  let bb = popByteBuffer();
  _encodeCommsAppointmentInput(message, bb);
  return toUint8Array(bb);
}

function _encodeCommsAppointmentInput(message, bb) {
  // optional int64 provider_id = 1;
  let $provider_id = message.provider_id;
  if ($provider_id !== undefined) {
    writeVarint32(bb, 8);
    writeVarint64(bb, $provider_id);
  }

  // optional int64 customer_id = 2;
  let $customer_id = message.customer_id;
  if ($customer_id !== undefined) {
    writeVarint32(bb, 16);
    writeVarint64(bb, $customer_id);
  }

  // repeated int64 recipient_ids = 3;
  let array$recipient_ids = message.recipient_ids;
  if (array$recipient_ids !== undefined) {
    let packed = popByteBuffer();
    for (let value of array$recipient_ids) {
      writeVarint64(packed, value);
    }
    writeVarint32(bb, 26);
    writeVarint32(bb, packed.offset);
    writeByteBuffer(bb, packed);
    pushByteBuffer(packed);
  }

  // optional int64 appointment_type_id = 4;
  let $appointment_type_id = message.appointment_type_id;
  if ($appointment_type_id !== undefined) {
    writeVarint32(bb, 32);
    writeVarint64(bb, $appointment_type_id);
  }

  // optional int64 startDate = 5;
  let $startDate = message.startDate;
  if ($startDate !== undefined) {
    writeVarint32(bb, 40);
    writeVarint64(bb, $startDate);
  }

  // optional int64 endDate = 6;
  let $endDate = message.endDate;
  if ($endDate !== undefined) {
    writeVarint32(bb, 48);
    writeVarint64(bb, $endDate);
  }
}

export function decodeCommsAppointmentInput(binary) {
  return _decodeCommsAppointmentInput(wrapByteBuffer(binary));
}

function _decodeCommsAppointmentInput(bb) {
  let message = {};

  end_of_message: while (!isAtEnd(bb)) {
    let tag = readVarint32(bb);

    switch (tag >>> 3) {
      case 0:
        break end_of_message;

      // optional int64 provider_id = 1;
      case 1: {
        message.provider_id = readVarint64(bb, /* unsigned */ false);
        break;
      }

      // optional int64 customer_id = 2;
      case 2: {
        message.customer_id = readVarint64(bb, /* unsigned */ false);
        break;
      }

      // repeated int64 recipient_ids = 3;
      case 3: {
        let values = message.recipient_ids || (message.recipient_ids = []);
        if ((tag & 7) === 2) {
          let outerLimit = pushTemporaryLength(bb);
          while (!isAtEnd(bb)) {
            values.push(readVarint64(bb, /* unsigned */ false));
          }
          bb.limit = outerLimit;
        } else {
          values.push(readVarint64(bb, /* unsigned */ false));
        }
        break;
      }

      // optional int64 appointment_type_id = 4;
      case 4: {
        message.appointment_type_id = readVarint64(bb, /* unsigned */ false);
        break;
      }

      // optional int64 startDate = 5;
      case 5: {
        message.startDate = readVarint64(bb, /* unsigned */ false);
        break;
      }

      // optional int64 endDate = 6;
      case 6: {
        message.endDate = readVarint64(bb, /* unsigned */ false);
        break;
      }

      default:
        skipUnknownField(bb, tag & 7);
    }
  }

  return message;
}

export function encodeCommsPendingAppointment(message) {
  let bb = popByteBuffer();
  _encodeCommsPendingAppointment(message, bb);
  return toUint8Array(bb);
}

function _encodeCommsPendingAppointment(message, bb) {
  // optional string id = 1;
  let $id = message.id;
  if ($id !== undefined) {
    writeVarint32(bb, 10);
    writeString(bb, $id);
  }

  // optional string status = 2;
  let $status = message.status;
  if ($status !== undefined) {
    writeVarint32(bb, 18);
    writeString(bb, $status);
  }
}

export function decodeCommsPendingAppointment(binary) {
  return _decodeCommsPendingAppointment(wrapByteBuffer(binary));
}

function _decodeCommsPendingAppointment(bb) {
  let message = {};

  end_of_message: while (!isAtEnd(bb)) {
    let tag = readVarint32(bb);

    switch (tag >>> 3) {
      case 0:
        break end_of_message;

      // optional string id = 1;
      case 1: {
        message.id = readString(bb, readVarint32(bb));
        break;
      }

      // optional string status = 2;
      case 2: {
        message.status = readString(bb, readVarint32(bb));
        break;
      }

      default:
        skipUnknownField(bb, tag & 7);
    }
  }

  return message;
}

export function encodeAcceptedAppointmentInput(message) {
  let bb = popByteBuffer();
  _encodeAcceptedAppointmentInput(message, bb);
  return toUint8Array(bb);
}

function _encodeAcceptedAppointmentInput(message, bb) {
  // optional int64 appointment_id = 1;
  let $appointment_id = message.appointment_id;
  if ($appointment_id !== undefined) {
    writeVarint32(bb, 8);
    writeVarint64(bb, $appointment_id);
  }

  // optional int64 status = 2;
  let $status = message.status;
  if ($status !== undefined) {
    writeVarint32(bb, 16);
    writeVarint64(bb, $status);
  }

  // optional CommsAppointmentInput appointment = 3;
  let $appointment = message.appointment;
  if ($appointment !== undefined) {
    writeVarint32(bb, 26);
    let nested = popByteBuffer();
    _encodeCommsAppointmentInput($appointment, nested);
    writeVarint32(bb, nested.limit);
    writeByteBuffer(bb, nested);
    pushByteBuffer(nested);
  }
}

export function decodeAcceptedAppointmentInput(binary) {
  return _decodeAcceptedAppointmentInput(wrapByteBuffer(binary));
}

function _decodeAcceptedAppointmentInput(bb) {
  let message = {};

  end_of_message: while (!isAtEnd(bb)) {
    let tag = readVarint32(bb);

    switch (tag >>> 3) {
      case 0:
        break end_of_message;

      // optional int64 appointment_id = 1;
      case 1: {
        message.appointment_id = readVarint64(bb, /* unsigned */ false);
        break;
      }

      // optional int64 status = 2;
      case 2: {
        message.status = readVarint64(bb, /* unsigned */ false);
        break;
      }

      // optional CommsAppointmentInput appointment = 3;
      case 3: {
        let limit = pushTemporaryLength(bb);
        message.appointment = _decodeCommsAppointmentInput(bb);
        bb.limit = limit;
        break;
      }

      default:
        skipUnknownField(bb, tag & 7);
    }
  }

  return message;
}

export function encodeRejectedAppointmentInput(message) {
  let bb = popByteBuffer();
  _encodeRejectedAppointmentInput(message, bb);
  return toUint8Array(bb);
}

function _encodeRejectedAppointmentInput(message, bb) {
  // optional int64 pending_appointment_id = 1;
  let $pending_appointment_id = message.pending_appointment_id;
  if ($pending_appointment_id !== undefined) {
    writeVarint32(bb, 8);
    writeVarint64(bb, $pending_appointment_id);
  }
}

export function decodeRejectedAppointmentInput(binary) {
  return _decodeRejectedAppointmentInput(wrapByteBuffer(binary));
}

function _decodeRejectedAppointmentInput(bb) {
  let message = {};

  end_of_message: while (!isAtEnd(bb)) {
    let tag = readVarint32(bb);

    switch (tag >>> 3) {
      case 0:
        break end_of_message;

      // optional int64 pending_appointment_id = 1;
      case 1: {
        message.pending_appointment_id = readVarint64(bb, /* unsigned */ false);
        break;
      }

      default:
        skipUnknownField(bb, tag & 7);
    }
  }

  return message;
}

export function encodeCommsAppointment(message) {
  let bb = popByteBuffer();
  _encodeCommsAppointment(message, bb);
  return toUint8Array(bb);
}

function _encodeCommsAppointment(message, bb) {
  // optional int64 id = 1;
  let $id = message.id;
  if ($id !== undefined) {
    writeVarint32(bb, 8);
    writeVarint64(bb, $id);
  }

  // optional int64 provider_id = 2;
  let $provider_id = message.provider_id;
  if ($provider_id !== undefined) {
    writeVarint32(bb, 16);
    writeVarint64(bb, $provider_id);
  }

  // optional int64 customer_id = 3;
  let $customer_id = message.customer_id;
  if ($customer_id !== undefined) {
    writeVarint32(bb, 24);
    writeVarint64(bb, $customer_id);
  }

  // repeated int64 recipient_ids = 4;
  let array$recipient_ids = message.recipient_ids;
  if (array$recipient_ids !== undefined) {
    let packed = popByteBuffer();
    for (let value of array$recipient_ids) {
      writeVarint64(packed, value);
    }
    writeVarint32(bb, 34);
    writeVarint32(bb, packed.offset);
    writeByteBuffer(bb, packed);
    pushByteBuffer(packed);
  }

  // optional int64 appointment_type_id = 5;
  let $appointment_type_id = message.appointment_type_id;
  if ($appointment_type_id !== undefined) {
    writeVarint32(bb, 40);
    writeVarint64(bb, $appointment_type_id);
  }

  // optional int64 start_date = 6;
  let $start_date = message.start_date;
  if ($start_date !== undefined) {
    writeVarint32(bb, 48);
    writeVarint64(bb, $start_date);
  }

  // optional int64 end_date = 7;
  let $end_date = message.end_date;
  if ($end_date !== undefined) {
    writeVarint32(bb, 56);
    writeVarint64(bb, $end_date);
  }
}

export function decodeCommsAppointment(binary) {
  return _decodeCommsAppointment(wrapByteBuffer(binary));
}

function _decodeCommsAppointment(bb) {
  let message = {};

  end_of_message: while (!isAtEnd(bb)) {
    let tag = readVarint32(bb);

    switch (tag >>> 3) {
      case 0:
        break end_of_message;

      // optional int64 id = 1;
      case 1: {
        message.id = readVarint64(bb, /* unsigned */ false);
        break;
      }

      // optional int64 provider_id = 2;
      case 2: {
        message.provider_id = readVarint64(bb, /* unsigned */ false);
        break;
      }

      // optional int64 customer_id = 3;
      case 3: {
        message.customer_id = readVarint64(bb, /* unsigned */ false);
        break;
      }

      // repeated int64 recipient_ids = 4;
      case 4: {
        let values = message.recipient_ids || (message.recipient_ids = []);
        if ((tag & 7) === 2) {
          let outerLimit = pushTemporaryLength(bb);
          while (!isAtEnd(bb)) {
            values.push(readVarint64(bb, /* unsigned */ false));
          }
          bb.limit = outerLimit;
        } else {
          values.push(readVarint64(bb, /* unsigned */ false));
        }
        break;
      }

      // optional int64 appointment_type_id = 5;
      case 5: {
        message.appointment_type_id = readVarint64(bb, /* unsigned */ false);
        break;
      }

      // optional int64 start_date = 6;
      case 6: {
        message.start_date = readVarint64(bb, /* unsigned */ false);
        break;
      }

      // optional int64 end_date = 7;
      case 7: {
        message.end_date = readVarint64(bb, /* unsigned */ false);
        break;
      }

      default:
        skipUnknownField(bb, tag & 7);
    }
  }

  return message;
}

export function encodeConsoleMessageInput(message) {
  let bb = popByteBuffer();
  _encodeConsoleMessageInput(message, bb);
  return toUint8Array(bb);
}

function _encodeConsoleMessageInput(message, bb) {
  // optional int64 sender_id = 1;
  let $sender_id = message.sender_id;
  if ($sender_id !== undefined) {
    writeVarint32(bb, 8);
    writeVarint64(bb, $sender_id);
  }

  // optional int64 receiver_id = 2;
  let $receiver_id = message.receiver_id;
  if ($receiver_id !== undefined) {
    writeVarint32(bb, 16);
    writeVarint64(bb, $receiver_id);
  }

  // optional string app_message = 3;
  let $app_message = message.app_message;
  if ($app_message !== undefined) {
    writeVarint32(bb, 26);
    writeString(bb, $app_message);
  }

  // repeated string cdn_urls = 4;
  let array$cdn_urls = message.cdn_urls;
  if (array$cdn_urls !== undefined) {
    for (let value of array$cdn_urls) {
      writeVarint32(bb, 34);
      writeString(bb, value);
    }
  }
}

export function decodeConsoleMessageInput(binary) {
  return _decodeConsoleMessageInput(wrapByteBuffer(binary));
}

function _decodeConsoleMessageInput(bb) {
  let message = {};

  end_of_message: while (!isAtEnd(bb)) {
    let tag = readVarint32(bb);

    switch (tag >>> 3) {
      case 0:
        break end_of_message;

      // optional int64 sender_id = 1;
      case 1: {
        message.sender_id = readVarint64(bb, /* unsigned */ false);
        break;
      }

      // optional int64 receiver_id = 2;
      case 2: {
        message.receiver_id = readVarint64(bb, /* unsigned */ false);
        break;
      }

      // optional string app_message = 3;
      case 3: {
        message.app_message = readString(bb, readVarint32(bb));
        break;
      }

      // repeated string cdn_urls = 4;
      case 4: {
        let values = message.cdn_urls || (message.cdn_urls = []);
        values.push(readString(bb, readVarint32(bb)));
        break;
      }

      default:
        skipUnknownField(bb, tag & 7);
    }
  }

  return message;
}

export function encodeConsoleMessage(message) {
  let bb = popByteBuffer();
  _encodeConsoleMessage(message, bb);
  return toUint8Array(bb);
}

function _encodeConsoleMessage(message, bb) {
  // optional int64 id = 1;
  let $id = message.id;
  if ($id !== undefined) {
    writeVarint32(bb, 8);
    writeVarint64(bb, $id);
  }

  // optional int64 sender_id = 2;
  let $sender_id = message.sender_id;
  if ($sender_id !== undefined) {
    writeVarint32(bb, 16);
    writeVarint64(bb, $sender_id);
  }

  // optional int64 receiver_id = 3;
  let $receiver_id = message.receiver_id;
  if ($receiver_id !== undefined) {
    writeVarint32(bb, 24);
    writeVarint64(bb, $receiver_id);
  }

  // optional string app_message = 4;
  let $app_message = message.app_message;
  if ($app_message !== undefined) {
    writeVarint32(bb, 34);
    writeString(bb, $app_message);
  }

  // repeated string cdn_urls = 5;
  let array$cdn_urls = message.cdn_urls;
  if (array$cdn_urls !== undefined) {
    for (let value of array$cdn_urls) {
      writeVarint32(bb, 42);
      writeString(bb, value);
    }
  }
}

export function decodeConsoleMessage(binary) {
  return _decodeConsoleMessage(wrapByteBuffer(binary));
}

function _decodeConsoleMessage(bb) {
  let message = {};

  end_of_message: while (!isAtEnd(bb)) {
    let tag = readVarint32(bb);

    switch (tag >>> 3) {
      case 0:
        break end_of_message;

      // optional int64 id = 1;
      case 1: {
        message.id = readVarint64(bb, /* unsigned */ false);
        break;
      }

      // optional int64 sender_id = 2;
      case 2: {
        message.sender_id = readVarint64(bb, /* unsigned */ false);
        break;
      }

      // optional int64 receiver_id = 3;
      case 3: {
        message.receiver_id = readVarint64(bb, /* unsigned */ false);
        break;
      }

      // optional string app_message = 4;
      case 4: {
        message.app_message = readString(bb, readVarint32(bb));
        break;
      }

      // repeated string cdn_urls = 5;
      case 5: {
        let values = message.cdn_urls || (message.cdn_urls = []);
        values.push(readString(bb, readVarint32(bb)));
        break;
      }

      default:
        skipUnknownField(bb, tag & 7);
    }
  }

  return message;
}

function pushTemporaryLength(bb) {
  let length = readVarint32(bb);
  let limit = bb.limit;
  bb.limit = bb.offset + length;
  return limit;
}

function skipUnknownField(bb, type) {
  switch (type) {
    case 0: while (readByte(bb) & 0x80) { } break;
    case 2: skip(bb, readVarint32(bb)); break;
    case 5: skip(bb, 4); break;
    case 1: skip(bb, 8); break;
    default: throw new Error("Unimplemented type: " + type);
  }
}

function stringToLong(value) {
  return {
    low: value.charCodeAt(0) | (value.charCodeAt(1) << 16),
    high: value.charCodeAt(2) | (value.charCodeAt(3) << 16),
    unsigned: false,
  };
}

function longToString(value) {
  let low = value.low;
  let high = value.high;
  return String.fromCharCode(
    low & 0xFFFF,
    low >>> 16,
    high & 0xFFFF,
    high >>> 16);
}

// The code below was modified from https://github.com/protobufjs/bytebuffer.js
// which is under the Apache License 2.0.

let f32 = new Float32Array(1);
let f32_u8 = new Uint8Array(f32.buffer);

let f64 = new Float64Array(1);
let f64_u8 = new Uint8Array(f64.buffer);

function intToLong(value) {
  value |= 0;
  return {
    low: value,
    high: value >> 31,
    unsigned: value >= 0,
  };
}

let bbStack = [];

function popByteBuffer() {
  const bb = bbStack.pop();
  if (!bb) return { bytes: new Uint8Array(64), offset: 0, limit: 0 };
  bb.offset = bb.limit = 0;
  return bb;
}

function pushByteBuffer(bb) {
  bbStack.push(bb);
}

function wrapByteBuffer(bytes) {
  return { bytes, offset: 0, limit: bytes.length };
}

function toUint8Array(bb) {
  let bytes = bb.bytes;
  let limit = bb.limit;
  return bytes.length === limit ? bytes : bytes.subarray(0, limit);
}

function skip(bb, offset) {
  if (bb.offset + offset > bb.limit) {
    throw new Error('Skip past limit');
  }
  bb.offset += offset;
}

function isAtEnd(bb) {
  return bb.offset >= bb.limit;
}

function grow(bb, count) {
  let bytes = bb.bytes;
  let offset = bb.offset;
  let limit = bb.limit;
  let finalOffset = offset + count;
  if (finalOffset > bytes.length) {
    let newBytes = new Uint8Array(finalOffset * 2);
    newBytes.set(bytes);
    bb.bytes = newBytes;
  }
  bb.offset = finalOffset;
  if (finalOffset > limit) {
    bb.limit = finalOffset;
  }
  return offset;
}

function advance(bb, count) {
  let offset = bb.offset;
  if (offset + count > bb.limit) {
    throw new Error('Read past limit');
  }
  bb.offset += count;
  return offset;
}

function readBytes(bb, count) {
  let offset = advance(bb, count);
  return bb.bytes.subarray(offset, offset + count);
}

function writeBytes(bb, buffer) {
  let offset = grow(bb, buffer.length);
  bb.bytes.set(buffer, offset);
}

function readString(bb, count) {
  // Sadly a hand-coded UTF8 decoder is much faster than subarray+TextDecoder in V8
  let offset = advance(bb, count);
  let fromCharCode = String.fromCharCode;
  let bytes = bb.bytes;
  let invalid = '\uFFFD';
  let text = '';

  for (let i = 0; i < count; i++) {
    let c1 = bytes[i + offset], c2, c3, c4, c;

    // 1 byte
    if ((c1 & 0x80) === 0) {
      text += fromCharCode(c1);
    }

    // 2 bytes
    else if ((c1 & 0xE0) === 0xC0) {
      if (i + 1 >= count) text += invalid;
      else {
        c2 = bytes[i + offset + 1];
        if ((c2 & 0xC0) !== 0x80) text += invalid;
        else {
          c = ((c1 & 0x1F) << 6) | (c2 & 0x3F);
          if (c < 0x80) text += invalid;
          else {
            text += fromCharCode(c);
            i++;
          }
        }
      }
    }

    // 3 bytes
    else if ((c1 & 0xF0) == 0xE0) {
      if (i + 2 >= count) text += invalid;
      else {
        c2 = bytes[i + offset + 1];
        c3 = bytes[i + offset + 2];
        if (((c2 | (c3 << 8)) & 0xC0C0) !== 0x8080) text += invalid;
        else {
          c = ((c1 & 0x0F) << 12) | ((c2 & 0x3F) << 6) | (c3 & 0x3F);
          if (c < 0x0800 || (c >= 0xD800 && c <= 0xDFFF)) text += invalid;
          else {
            text += fromCharCode(c);
            i += 2;
          }
        }
      }
    }

    // 4 bytes
    else if ((c1 & 0xF8) == 0xF0) {
      if (i + 3 >= count) text += invalid;
      else {
        c2 = bytes[i + offset + 1];
        c3 = bytes[i + offset + 2];
        c4 = bytes[i + offset + 3];
        if (((c2 | (c3 << 8) | (c4 << 16)) & 0xC0C0C0) !== 0x808080) text += invalid;
        else {
          c = ((c1 & 0x07) << 0x12) | ((c2 & 0x3F) << 0x0C) | ((c3 & 0x3F) << 0x06) | (c4 & 0x3F);
          if (c < 0x10000 || c > 0x10FFFF) text += invalid;
          else {
            c -= 0x10000;
            text += fromCharCode((c >> 10) + 0xD800, (c & 0x3FF) + 0xDC00);
            i += 3;
          }
        }
      }
    }

    else text += invalid;
  }

  return text;
}

function writeString(bb, text) {
  // Sadly a hand-coded UTF8 encoder is much faster than TextEncoder+set in V8
  let n = text.length;
  let byteCount = 0;

  // Write the byte count first
  for (let i = 0; i < n; i++) {
    let c = text.charCodeAt(i);
    if (c >= 0xD800 && c <= 0xDBFF && i + 1 < n) {
      c = (c << 10) + text.charCodeAt(++i) - 0x35FDC00;
    }
    byteCount += c < 0x80 ? 1 : c < 0x800 ? 2 : c < 0x10000 ? 3 : 4;
  }
  writeVarint32(bb, byteCount);

  let offset = grow(bb, byteCount);
  let bytes = bb.bytes;

  // Then write the bytes
  for (let i = 0; i < n; i++) {
    let c = text.charCodeAt(i);
    if (c >= 0xD800 && c <= 0xDBFF && i + 1 < n) {
      c = (c << 10) + text.charCodeAt(++i) - 0x35FDC00;
    }
    if (c < 0x80) {
      bytes[offset++] = c;
    } else {
      if (c < 0x800) {
        bytes[offset++] = ((c >> 6) & 0x1F) | 0xC0;
      } else {
        if (c < 0x10000) {
          bytes[offset++] = ((c >> 12) & 0x0F) | 0xE0;
        } else {
          bytes[offset++] = ((c >> 18) & 0x07) | 0xF0;
          bytes[offset++] = ((c >> 12) & 0x3F) | 0x80;
        }
        bytes[offset++] = ((c >> 6) & 0x3F) | 0x80;
      }
      bytes[offset++] = (c & 0x3F) | 0x80;
    }
  }
}

function writeByteBuffer(bb, buffer) {
  let offset = grow(bb, buffer.limit);
  let from = bb.bytes;
  let to = buffer.bytes;

  // This for loop is much faster than subarray+set on V8
  for (let i = 0, n = buffer.limit; i < n; i++) {
    from[i + offset] = to[i];
  }
}

function readByte(bb) {
  return bb.bytes[advance(bb, 1)];
}

function writeByte(bb, value) {
  let offset = grow(bb, 1);
  bb.bytes[offset] = value;
}

function readFloat(bb) {
  let offset = advance(bb, 4);
  let bytes = bb.bytes;

  // Manual copying is much faster than subarray+set in V8
  f32_u8[0] = bytes[offset++];
  f32_u8[1] = bytes[offset++];
  f32_u8[2] = bytes[offset++];
  f32_u8[3] = bytes[offset++];
  return f32[0];
}

function writeFloat(bb, value) {
  let offset = grow(bb, 4);
  let bytes = bb.bytes;
  f32[0] = value;

  // Manual copying is much faster than subarray+set in V8
  bytes[offset++] = f32_u8[0];
  bytes[offset++] = f32_u8[1];
  bytes[offset++] = f32_u8[2];
  bytes[offset++] = f32_u8[3];
}

function readDouble(bb) {
  let offset = advance(bb, 8);
  let bytes = bb.bytes;

  // Manual copying is much faster than subarray+set in V8
  f64_u8[0] = bytes[offset++];
  f64_u8[1] = bytes[offset++];
  f64_u8[2] = bytes[offset++];
  f64_u8[3] = bytes[offset++];
  f64_u8[4] = bytes[offset++];
  f64_u8[5] = bytes[offset++];
  f64_u8[6] = bytes[offset++];
  f64_u8[7] = bytes[offset++];
  return f64[0];
}

function writeDouble(bb, value) {
  let offset = grow(bb, 8);
  let bytes = bb.bytes;
  f64[0] = value;

  // Manual copying is much faster than subarray+set in V8
  bytes[offset++] = f64_u8[0];
  bytes[offset++] = f64_u8[1];
  bytes[offset++] = f64_u8[2];
  bytes[offset++] = f64_u8[3];
  bytes[offset++] = f64_u8[4];
  bytes[offset++] = f64_u8[5];
  bytes[offset++] = f64_u8[6];
  bytes[offset++] = f64_u8[7];
}

function readInt32(bb) {
  let offset = advance(bb, 4);
  let bytes = bb.bytes;
  return (
    bytes[offset] |
    (bytes[offset + 1] << 8) |
    (bytes[offset + 2] << 16) |
    (bytes[offset + 3] << 24)
  );
}

function writeInt32(bb, value) {
  let offset = grow(bb, 4);
  let bytes = bb.bytes;
  bytes[offset] = value;
  bytes[offset + 1] = value >> 8;
  bytes[offset + 2] = value >> 16;
  bytes[offset + 3] = value >> 24;
}

function readInt64(bb, unsigned) {
  return {
    low: readInt32(bb),
    high: readInt32(bb),
    unsigned,
  };
}

function writeInt64(bb, value) {
  writeInt32(bb, value.low);
  writeInt32(bb, value.high);
}

function readVarint32(bb) {
  let c = 0;
  let value = 0;
  let b;
  do {
    b = readByte(bb);
    if (c < 32) value |= (b & 0x7F) << c;
    c += 7;
  } while (b & 0x80);
  return value;
}

function writeVarint32(bb, value) {
  value >>>= 0;
  while (value >= 0x80) {
    writeByte(bb, (value & 0x7f) | 0x80);
    value >>>= 7;
  }
  writeByte(bb, value);
}

function readVarint64(bb, unsigned) {
  let part0 = 0;
  let part1 = 0;
  let part2 = 0;
  let b;

  b = readByte(bb); part0 = (b & 0x7F); if (b & 0x80) {
    b = readByte(bb); part0 |= (b & 0x7F) << 7; if (b & 0x80) {
      b = readByte(bb); part0 |= (b & 0x7F) << 14; if (b & 0x80) {
        b = readByte(bb); part0 |= (b & 0x7F) << 21; if (b & 0x80) {

          b = readByte(bb); part1 = (b & 0x7F); if (b & 0x80) {
            b = readByte(bb); part1 |= (b & 0x7F) << 7; if (b & 0x80) {
              b = readByte(bb); part1 |= (b & 0x7F) << 14; if (b & 0x80) {
                b = readByte(bb); part1 |= (b & 0x7F) << 21; if (b & 0x80) {

                  b = readByte(bb); part2 = (b & 0x7F); if (b & 0x80) {
                    b = readByte(bb); part2 |= (b & 0x7F) << 7;
                  }
                }
              }
            }
          }
        }
      }
    }
  }

  return {
    low: part0 | (part1 << 28),
    high: (part1 >>> 4) | (part2 << 24),
    unsigned,
  };
}

function writeVarint64(bb, value) {
  let part0 = value.low >>> 0;
  let part1 = ((value.low >>> 28) | (value.high << 4)) >>> 0;
  let part2 = value.high >>> 24;

  // ref: src/google/protobuf/io/coded_stream.cc
  let size =
    part2 === 0 ?
      part1 === 0 ?
        part0 < 1 << 14 ?
          part0 < 1 << 7 ? 1 : 2 :
          part0 < 1 << 21 ? 3 : 4 :
        part1 < 1 << 14 ?
          part1 < 1 << 7 ? 5 : 6 :
          part1 < 1 << 21 ? 7 : 8 :
      part2 < 1 << 7 ? 9 : 10;

  let offset = grow(bb, size);
  let bytes = bb.bytes;

  switch (size) {
    case 10: bytes[offset + 9] = (part2 >>> 7) & 0x01;
    case 9: bytes[offset + 8] = size !== 9 ? part2 | 0x80 : part2 & 0x7F;
    case 8: bytes[offset + 7] = size !== 8 ? (part1 >>> 21) | 0x80 : (part1 >>> 21) & 0x7F;
    case 7: bytes[offset + 6] = size !== 7 ? (part1 >>> 14) | 0x80 : (part1 >>> 14) & 0x7F;
    case 6: bytes[offset + 5] = size !== 6 ? (part1 >>> 7) | 0x80 : (part1 >>> 7) & 0x7F;
    case 5: bytes[offset + 4] = size !== 5 ? part1 | 0x80 : part1 & 0x7F;
    case 4: bytes[offset + 3] = size !== 4 ? (part0 >>> 21) | 0x80 : (part0 >>> 21) & 0x7F;
    case 3: bytes[offset + 2] = size !== 3 ? (part0 >>> 14) | 0x80 : (part0 >>> 14) & 0x7F;
    case 2: bytes[offset + 1] = size !== 2 ? (part0 >>> 7) | 0x80 : (part0 >>> 7) & 0x7F;
    case 1: bytes[offset] = size !== 1 ? part0 | 0x80 : part0 & 0x7F;
  }
}

function readVarint32ZigZag(bb) {
  let value = readVarint32(bb);

  // ref: src/google/protobuf/wire_format_lite.h
  return (value >>> 1) ^ -(value & 1);
}

function writeVarint32ZigZag(bb, value) {
  // ref: src/google/protobuf/wire_format_lite.h
  writeVarint32(bb, (value << 1) ^ (value >> 31));
}

function readVarint64ZigZag(bb) {
  let value = readVarint64(bb, /* unsigned */ false);
  let low = value.low;
  let high = value.high;
  let flip = -(low & 1);

  // ref: src/google/protobuf/wire_format_lite.h
  return {
    low: ((low >>> 1) | (high << 31)) ^ flip,
    high: (high >>> 1) ^ flip,
    unsigned: false,
  };
}

function writeVarint64ZigZag(bb, value) {
  let low = value.low;
  let high = value.high;
  let flip = high >> 31;

  // ref: src/google/protobuf/wire_format_lite.h
  writeVarint64(bb, {
    low: (low << 1) ^ flip,
    high: ((high << 1) | (low >>> 31)) ^ flip,
    unsigned: false,
  });
}
