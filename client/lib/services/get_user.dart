import 'package:cardio_messenger/models/chat.dart';
import 'package:cardio_messenger/models/error.dart';
import 'package:cardio_messenger/models/user.dart';
import 'package:cardio_messenger/services/api.dart';

class GetUser {
  MUser user;
  Errorr? error;

  GetUser.fromJson(Map<String, dynamic> json)
      : user = MUser.fromJson(json['User']),
        error = json['Error'] == null ? null : Errorr.fromJson(json['Error']);

  static Future<MUser> get(int id) async {
    Map<String, dynamic> response = await Api.shared.sendWithToken({'Id': id}, '/user/get');
    return GetUser.fromJson(response).user;
  }

  static Future<MUser> edit(String? firstName, String? middleName, String? lastName, int? avatarID) async {
    Map<String, dynamic> response = await Api.shared.sendWithToken({'FirstName': firstName, 'LastName': lastName, 'MiddleName': middleName, 'AvatarID': avatarID}, '/user/edit');
    return GetUser.fromJson(response).user;
  }

  static Future<void> logout() async {
    Map<String, dynamic> response = await Api.shared.sendWithToken({}, '/user/logout');
  }
}