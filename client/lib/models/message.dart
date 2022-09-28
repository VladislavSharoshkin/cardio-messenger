
import 'package:cardio_messenger/models/forwardMessage.dart';
import 'package:cardio_messenger/models/messageFile.dart';
import 'package:cardio_messenger/models/user.dart';



enum MessageDeliveryStatus {unknown, unsent, deliverToServer, sending}

class MMessage {
  int id;
  int senderId;
  int chatId;
  String? text;
  int type;
  DateTime createdAt;
  bool isDeleted;
  List<String> readIds;
  List<MessageFile> messageFiles;
  List<ForwardMessage> forwardMessages;
  MUser sender;
  MessageDeliveryStatus deliveryStatus;
  bool isMy;
  int readState;

  MMessage(
      {this.id = 0,
        this.senderId = 0,
        this.chatId = 0,
        this.text = '',
        this.isMy = false,
        this.type = 0,
        this.isDeleted = false,
        this.readState = 0,
        this.deliveryStatus = MessageDeliveryStatus.unknown,
        MUser? sender,
        List<MessageFile>? messageFiles,
        List<ForwardMessage>? forwardMessages,
        List<String>? readIds,
        DateTime? createdAt})
  : messageFiles = messageFiles ?? [], readIds = readIds ?? [],
        createdAt = createdAt?? DateTime.utc(1), sender = sender ?? MUser(),
        forwardMessages = forwardMessages ?? [];

  // setFiles(List<String> fileIds) {
  //   this.fileIds = fileIds;
  // }

  // static Message copy(Message message) {
  //   return Message(id: message.id, senderId: message.senderId,
  //       chatId: message.chatId, text: message.text, type: message.type,
  //       isDeleted: message.isDeleted,attachList: message.attachList,readIds: message.readIds,
  //       createdAt: message.createdAt);
  // }


  static MMessage init(int chatId, String text, MessageDeliveryStatus deliveryStatus, List<MessageFile> messageFiles) {
    return MMessage(chatId: chatId, text: text, deliveryStatus: deliveryStatus, messageFiles: messageFiles);
  }

  MMessage.fromJson(Map<String, dynamic> json)
      : id = json['ID'] ?? 0,
        senderId = json['SenderID'] ?? 0,
        chatId = json['ChatId'] ?? 0,
        text = json['Text'] ?? '',
        type = json['Type'] ?? 0,
        readState = json['ReadState'] ?? 0,
        deliveryStatus = MessageDeliveryStatus.values[json['DeliveryStatus'] ?? 0],
        sender = MUser.fromJson(json['Sender'] ?? {}),
        isDeleted = json['IsDeleted'] ?? false,
        isMy = json['IsMy'] ?? false,
        messageFiles = MessageFile.fromJsonList(json['Attachments'] ?? []),
        forwardMessages = ForwardMessage.fromJsonList(json['ForwardMessages'] ?? []),
        createdAt = DateTime.parse(json['CreatedAt'] ?? '0001-01-01T00:00:00Z').toLocal(),
        readIds = List<String>.from(json['ReadIds'] ?? []);

  Map<String, dynamic> toJson() => {
    'ID': id,
    'SenderID': senderId,
    'ChatId': chatId,
    'Text': text,
    'Type': type,
    'ReadIds': readIds,
    'Sender': sender.toJson(),
    'ForwardMessages': forwardMessages,
    'DeliveryStatus': deliveryStatus.index,

    'IsDeleted': isDeleted,
    'Attachments': messageFiles,
  };

  int getId(){
    return id;
  }

  bool isLoad(){
    return id != 0;
  }

  bool isSent() {
    return true;
  }

  bool isRead() {
    return this.readIds.isNotEmpty;
  }

  // void encrypt() {
  //   this.text = base64.encode(utf8.encode(this.text));
  // }
  //
  // void decrypt() {
  //   this.text = utf8.decode(base64.decode(this.text));
  // }

  static List<MMessage> fromJsonList(dynamic json) {
    return List<MMessage>.from(json.map((model)=> MMessage.fromJson(model)));
  }

  static List<int> getIds(List<MMessage> messages){
    return [...messages.map((el) => el.id)];
  }
}
