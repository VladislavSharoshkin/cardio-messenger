import 'package:cardio_messenger/models/chat.dart';
import 'package:cardio_messenger/models/error.dart';
import 'package:cardio_messenger/models/user.dart';
import 'package:cardio_messenger/services/api.dart';

class GetUsers {
  List<MUser> users;
  Errorr? error;

  GetUsers.fromJson(Map<String, dynamic> json)
      : users = MUser.fromJsonList(json['Users'] ?? []),
        error = json['Error'] == null ? null : Errorr.fromJson(json['Error']);

  static Future<GetUsers> search(String searchString) async {
    Map<String, dynamic> response = await Api.shared.sendWithToken({'SearchString': searchString}, '/user/search');
    return GetUsers.fromJson(response);
  }
}