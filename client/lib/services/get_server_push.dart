import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:cardio_messenger/config/app_config.dart';
import 'package:cardio_messenger/models/chat.dart';
import 'package:cardio_messenger/models/error.dart';
import 'package:cardio_messenger/models/server_push.dart';
import 'package:cardio_messenger/services/api.dart';

class GetServerPush {
  ServerPush serverPush;
  Errorr? error;

  GetServerPush.fromJson(Map<String, dynamic> json)
      : serverPush = ServerPush.fromJson(json['ServerPush']),
        error = json['Error'] == null ? null : Errorr.fromJson(json['Error']);

  static Future<ServerPush> get(String push) async {
    var body = json.encode({'Push': push});

    String url = 'https://${AppConfig.pushServer}/push/add';
    var response = await http.post(Uri.parse(url), body: body);
    Map<String, dynamic> responseBody = jsonDecode(response.body);

    return GetServerPush.fromJson(responseBody).serverPush;
  }


}