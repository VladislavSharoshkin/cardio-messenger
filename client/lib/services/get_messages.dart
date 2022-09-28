import 'package:cardio_messenger/models/chat.dart';
import 'package:cardio_messenger/models/error.dart';
import 'package:cardio_messenger/models/message.dart';
import 'package:cardio_messenger/services/api.dart';

class GetMessages {
  List<MMessage> messages;
  Errorr? error;

  GetMessages.fromJson(Map<String, dynamic> json)
      : messages = MMessage.fromJsonList(json['Messages'] ?? []),
        error = json['Error'] == null ? null : Errorr.fromJson(json['Error']);

  static Future<List<MMessage>> getMessages(int chatId, int? startMessageId, int? count) async {
    Map<String, dynamic> response = await Api.shared.sendWithToken({'ChatId': chatId, 'StartMessageId': startMessageId, 'Count': count}, '/chat/getMessages');
    return GetMessages.fromJson(response).messages;
  }

  static Future<void> readMessage(int chatId, int messageId) async {
    Map<String, dynamic> response = await Api.shared.sendWithToken({'MessageId': messageId, 'ChatId': chatId}, '/message/read');
    //return GetMessages.fromJson(response).messages;
  }

  static Future<void> del(int chatId, List<int> ids) async {
    Map<String, dynamic> response = await Api.shared.sendWithToken({'ChatId': chatId, 'Ids': ids}, '/message/del');
    //return GetMessages.fromJson(response).messages;
  }
}