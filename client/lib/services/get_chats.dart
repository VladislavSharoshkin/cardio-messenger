import 'package:cardio_messenger/models/chat.dart';
import 'package:cardio_messenger/models/error.dart';
import 'package:cardio_messenger/services/api.dart';

class GetChats {
  List<MChat> chats;
  Errorr? error;

  GetChats.fromJson(Map<String, dynamic> json)
      : chats = MChat.fromJsonList(json['Chats'] ?? []),
        error = json['Error'] == null ? null : Errorr.fromJson(json['Error']);

  static Future<GetChats> my() async {
    Map<String, dynamic> response = await Api.shared.sendWithToken({}, '/chat/my');
    return GetChats.fromJson(response);
  }

  static Future<List<MChat>> del(List<int> ids) async {
    Map<String, dynamic> response = await Api.shared.sendWithToken({'Ids': ids}, '/chat/del');
    return GetChats.fromJson(response).chats;
  }
}