# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: demo.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='demo.proto',
  package='proto',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\ndemo.proto\x12\x05proto\".\n\rDemonstration\x12\x1d\n\x07\x62\x61tches\x18\x01 \x03(\x0b\x32\x0c.proto.Batch\"<\n\x05\x42\x61tch\x12 \n\titerators\x18\x01 \x03(\x0b\x32\r.proto.Events\x12\x11\n\ttimestamp\x18\x02 \x01(\t\"\x19\n\tStartTime\x12\x0c\n\x04time\x18\x01 \x01(\r\"4\n\x06\x45vents\x12\x0c\n\x04type\x18\x02 \x01(\t\x12\x1c\n\x06\x65vents\x18\x01 \x03(\x0b\x32\x0c.proto.Event\"\x16\n\x05\x45vent\x12\r\n\x05\x65vent\x18\x01 \x01(\x0c\"H\n\x07Headers\x12\x11\n\tMessageId\x18\x02 \x01(\r\x12\x17\n\x0fParentMessageId\x18\x03 \x01(\r\x12\x11\n\tEpisodeId\x18\x04 \x01(\r\"\xd0\x01\n\x04\x42ody\x12\r\n\x05\x45nvId\x18\x01 \x01(\t\x12\x11\n\tEnvStatus\x18\x02 \x01(\t\x12\x0b\n\x03\x46ps\x18\x03 \x01(\x02\x12\x0e\n\x06Reward\x18\x04 \x01(\x02\x12\x0c\n\x04\x44one\x18\x05 \x01(\x08\x12\x0e\n\x06Record\x18\x06 \x01(\x08\x12\x0b\n\x03Obs\x18\x07 \x01(\t\x12\x0f\n\x07ObsType\x18\x08 \x01(\t\x12\x0c\n\x04Info\x18\t \x01(\t\x12\x10\n\x08InfoType\x18\n \x01(\t\x12\x0f\n\x07Message\x18\x0b \x01(\t\x12\x0c\n\x04Seed\x18\x0c \x01(\r\x12\x0e\n\x06\x41\x63tion\x18\r \x01(\x0c\"h\n\x07Message\x12\x0e\n\x06Method\x18\x01 \x01(\t\x12\x1f\n\x07Headers\x18\x02 \x01(\x0b\x32\x0e.proto.Headers\x12\x19\n\x04\x42ody\x18\x03 \x01(\x0b\x32\x0b.proto.Body\x12\x11\n\tTimestamp\x18\x04 \x01(\rb\x06proto3')
)




_DEMONSTRATION = _descriptor.Descriptor(
  name='Demonstration',
  full_name='proto.Demonstration',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='batches', full_name='proto.Demonstration.batches', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=21,
  serialized_end=67,
)


_BATCH = _descriptor.Descriptor(
  name='Batch',
  full_name='proto.Batch',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='iterators', full_name='proto.Batch.iterators', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='proto.Batch.timestamp', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=69,
  serialized_end=129,
)


_STARTTIME = _descriptor.Descriptor(
  name='StartTime',
  full_name='proto.StartTime',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='time', full_name='proto.StartTime.time', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=131,
  serialized_end=156,
)


_EVENTS = _descriptor.Descriptor(
  name='Events',
  full_name='proto.Events',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='proto.Events.type', index=0,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='events', full_name='proto.Events.events', index=1,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=158,
  serialized_end=210,
)


_EVENT = _descriptor.Descriptor(
  name='Event',
  full_name='proto.Event',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='event', full_name='proto.Event.event', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=212,
  serialized_end=234,
)


_HEADERS = _descriptor.Descriptor(
  name='Headers',
  full_name='proto.Headers',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='MessageId', full_name='proto.Headers.MessageId', index=0,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ParentMessageId', full_name='proto.Headers.ParentMessageId', index=1,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='EpisodeId', full_name='proto.Headers.EpisodeId', index=2,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=236,
  serialized_end=308,
)


_BODY = _descriptor.Descriptor(
  name='Body',
  full_name='proto.Body',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='EnvId', full_name='proto.Body.EnvId', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='EnvStatus', full_name='proto.Body.EnvStatus', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='Fps', full_name='proto.Body.Fps', index=2,
      number=3, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='Reward', full_name='proto.Body.Reward', index=3,
      number=4, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='Done', full_name='proto.Body.Done', index=4,
      number=5, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='Record', full_name='proto.Body.Record', index=5,
      number=6, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='Obs', full_name='proto.Body.Obs', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ObsType', full_name='proto.Body.ObsType', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='Info', full_name='proto.Body.Info', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='InfoType', full_name='proto.Body.InfoType', index=9,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='Message', full_name='proto.Body.Message', index=10,
      number=11, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='Seed', full_name='proto.Body.Seed', index=11,
      number=12, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='Action', full_name='proto.Body.Action', index=12,
      number=13, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=311,
  serialized_end=519,
)


_MESSAGE = _descriptor.Descriptor(
  name='Message',
  full_name='proto.Message',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Method', full_name='proto.Message.Method', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='Headers', full_name='proto.Message.Headers', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='Body', full_name='proto.Message.Body', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='Timestamp', full_name='proto.Message.Timestamp', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=521,
  serialized_end=625,
)

_DEMONSTRATION.fields_by_name['batches'].message_type = _BATCH
_BATCH.fields_by_name['iterators'].message_type = _EVENTS
_EVENTS.fields_by_name['events'].message_type = _EVENT
_MESSAGE.fields_by_name['Headers'].message_type = _HEADERS
_MESSAGE.fields_by_name['Body'].message_type = _BODY
DESCRIPTOR.message_types_by_name['Demonstration'] = _DEMONSTRATION
DESCRIPTOR.message_types_by_name['Batch'] = _BATCH
DESCRIPTOR.message_types_by_name['StartTime'] = _STARTTIME
DESCRIPTOR.message_types_by_name['Events'] = _EVENTS
DESCRIPTOR.message_types_by_name['Event'] = _EVENT
DESCRIPTOR.message_types_by_name['Headers'] = _HEADERS
DESCRIPTOR.message_types_by_name['Body'] = _BODY
DESCRIPTOR.message_types_by_name['Message'] = _MESSAGE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Demonstration = _reflection.GeneratedProtocolMessageType('Demonstration', (_message.Message,), dict(
  DESCRIPTOR = _DEMONSTRATION,
  __module__ = 'demo_pb2'
  # @@protoc_insertion_point(class_scope:proto.Demonstration)
  ))
_sym_db.RegisterMessage(Demonstration)

Batch = _reflection.GeneratedProtocolMessageType('Batch', (_message.Message,), dict(
  DESCRIPTOR = _BATCH,
  __module__ = 'demo_pb2'
  # @@protoc_insertion_point(class_scope:proto.Batch)
  ))
_sym_db.RegisterMessage(Batch)

StartTime = _reflection.GeneratedProtocolMessageType('StartTime', (_message.Message,), dict(
  DESCRIPTOR = _STARTTIME,
  __module__ = 'demo_pb2'
  # @@protoc_insertion_point(class_scope:proto.StartTime)
  ))
_sym_db.RegisterMessage(StartTime)

Events = _reflection.GeneratedProtocolMessageType('Events', (_message.Message,), dict(
  DESCRIPTOR = _EVENTS,
  __module__ = 'demo_pb2'
  # @@protoc_insertion_point(class_scope:proto.Events)
  ))
_sym_db.RegisterMessage(Events)

Event = _reflection.GeneratedProtocolMessageType('Event', (_message.Message,), dict(
  DESCRIPTOR = _EVENT,
  __module__ = 'demo_pb2'
  # @@protoc_insertion_point(class_scope:proto.Event)
  ))
_sym_db.RegisterMessage(Event)

Headers = _reflection.GeneratedProtocolMessageType('Headers', (_message.Message,), dict(
  DESCRIPTOR = _HEADERS,
  __module__ = 'demo_pb2'
  # @@protoc_insertion_point(class_scope:proto.Headers)
  ))
_sym_db.RegisterMessage(Headers)

Body = _reflection.GeneratedProtocolMessageType('Body', (_message.Message,), dict(
  DESCRIPTOR = _BODY,
  __module__ = 'demo_pb2'
  # @@protoc_insertion_point(class_scope:proto.Body)
  ))
_sym_db.RegisterMessage(Body)

Message = _reflection.GeneratedProtocolMessageType('Message', (_message.Message,), dict(
  DESCRIPTOR = _MESSAGE,
  __module__ = 'demo_pb2'
  # @@protoc_insertion_point(class_scope:proto.Message)
  ))
_sym_db.RegisterMessage(Message)


# @@protoc_insertion_point(module_scope)