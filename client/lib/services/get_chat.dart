import 'package:cardio_messenger/models/chat.dart';
import 'package:cardio_messenger/models/error.dart';
import 'package:cardio_messenger/services/api.dart';

class GetChat {
  MChat chat;
  Errorr? error;
  
  GetChat.fromJson(Map<String, dynamic> json)
      : chat = MChat.fromJson(json['Chat']),
        error = json['Error'] == null ? null : Errorr.fromJson(json['Error']);

  static Future<MChat> get(int id) async {
    Map<String, dynamic> response = await Api.shared.sendWithToken({'Id': id}, '/chat/get');
    return GetChat.fromJson(response).chat;
  }

  static Future<MChat> startDirectChat(int id) async {
    Map<String, dynamic> response = await Api.shared.sendWithToken({'Type': 1, 'Participants': [id]}, '/chat/create');
    return GetChat.fromJson(response).chat;
  }

  static Future<GetChat> createGroupChat(String? name, List<int> participants) async {
    Map<String, dynamic> response = await Api.shared.sendWithToken({'Type': 2, 'Name': name, 'Participants': participants}, '/chat/create');
    return GetChat.fromJson(response);
  }

  static Future<MChat> edit(int chatID, String? name, String? about, int? avatarID) async {
    Map<String, dynamic> response = await Api.shared.sendWithToken({'ChatID':chatID, 'Name': name, 'About': about, 'AvatarID': avatarID}, '/chat/edit');
    return GetChat.fromJson(response).chat;
  }
}