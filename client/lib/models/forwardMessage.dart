
import 'package:cardio_messenger/models/message.dart';


class ForwardMessage {
  int id;
  int messageId;
  int forwardMessageId;
  MMessage forwardMessage;

  ForwardMessage(
      {this.id = 0,
      this.messageId = 0,
      this.forwardMessageId = 0,
        MMessage? forwardMessage})
      : forwardMessage = forwardMessage ?? MMessage();

  static ForwardMessage init(MMessage forwardMessage){
    return ForwardMessage(forwardMessage: forwardMessage);
  }

  ForwardMessage.fromJson(Map<String, dynamic> json)
      : id = json['ID'] ?? 0,
        messageId = json['MessageId'] ?? 0,
        forwardMessage = MMessage.fromJson(json['ForwardMessage']),
        forwardMessageId = json['ForwardMessageId'] ?? 0;

  Map<String, dynamic> toJson() => {
        'ID': id,
        'MessageId': messageId,
        'ForwardMessageId': forwardMessageId,
        'ForwardMessage': forwardMessage
      };

  int getId() {
    return id;
  }

  static List<ForwardMessage> fromJsonList(dynamic json) {
    return List<ForwardMessage>.from(json.map((model) => ForwardMessage.fromJson(model)));
  }

  static List<int> getIds(List<ForwardMessage> messages) {
    List<int> ids = [];
    messages.forEach((message) {
      ids.add(message.getId());
    });
    return ids;
  }
}
