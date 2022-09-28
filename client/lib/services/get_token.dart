
import 'package:cardio_messenger/models/error.dart';
import 'package:cardio_messenger/models/token.dart';
import 'package:cardio_messenger/services/api.dart';

class GetToken {
  Token token;
  Errorr? error;

  GetToken.fromJson(Map<String, dynamic> json)
      : token = Token.fromJson(json['Token'] ?? {}),
        error = json['Error'] == null ? null : Errorr.fromJson(json['Error']);

  static Future<GetToken> login(String login, String pass, String? push) async {
    Map<String, dynamic> response = await Api.shared.send(
        {'Login': login, 'Pass': pass, 'Push': push},
        '/user/login');
    return GetToken.fromJson(response);
  }

  static Future<GetToken> reg(String login, String pass, String? push) async {
    Map<String, dynamic> response = await Api.shared.send(
        {'Login': login, 'Pass': pass, 'Push': push},
        '/user/reg');
    return GetToken.fromJson(response);
  }
}