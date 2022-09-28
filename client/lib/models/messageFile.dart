import 'package:cardio_messenger/models/filee.dart';


class MessageFile {
  int id;
  int messageId;
  int fileId;
  MFile file;

  MessageFile(
      {this.id = 0,
      this.messageId = 0,
      this.fileId = 0,
      MFile? file})
      : file = file ?? MFile();

  static MessageFile init(MFile file){
    return MessageFile(file: file);
  }

  MessageFile.fromJson(Map<String, dynamic> json)
      : id = json['ID'] ?? 0,
        messageId = json['MessageId'] ?? 0,
        file = MFile.fromJson(json['File']),
        fileId = json['FileId'] ?? 0;

  Map<String, dynamic> toJson() => {
        'ID': id,
        'MessageId': messageId,
        'FileId': fileId,
        'File': file
      };

  int getId() {
    return id;
  }

  static List<MessageFile> fromJsonList(dynamic json) {
    return List<MessageFile>.from(json.map((model) => MessageFile.fromJson(model)));
  }

  static List<int> getIds(List<MessageFile> messages) {
    List<int> ids = [];
    messages.forEach((message) {
      ids.add(message.getId());
    });
    return ids;
  }
}
