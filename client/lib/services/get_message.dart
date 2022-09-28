import 'package:cardio_messenger/models/chat.dart';
import 'package:cardio_messenger/models/error.dart';
import 'package:cardio_messenger/models/message.dart';
import 'package:cardio_messenger/services/api.dart';

class GetMessage {
  MMessage message;
  Errorr? error;

  GetMessage.fromJson(Map<String, dynamic> json)
      : message =MMessage.fromJson(json['Message']),
        error = json['Error'] == null ? null : Errorr.fromJson(json['Error']);

  static Future<MMessage> send(int chatID, String? text, List<int> forwardMessages, List<int> files) async {
    Map<String, dynamic> response = await Api.shared.sendWithToken({'ChatID': chatID, 'Text': text, 'ForwardMessages': forwardMessages, 'Attachments': files}, '/message/send');
    return GetMessage.fromJson(response).message;
  }
}