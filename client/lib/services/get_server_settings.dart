import 'dart:convert';

import 'package:cardio_messenger/models/error.dart';
import 'package:cardio_messenger/models/server_settings.dart';
import 'package:http/http.dart' as http;

class GetServerSettings {
  ServerSettings serverSettings;
  Errorr? error;

  GetServerSettings.fromJson(Map<String, dynamic> json)
      : serverSettings = ServerSettings.fromJson(json['Settings']),
        error = json['Error'] == null ? null : Errorr.fromJson(json['Error']);

  static Future<GetServerSettings> send(String address) async {
    String url = 'https://$address/misc/info';
    var response = await http.post(Uri.parse(url));
    Map<String, dynamic> body = jsonDecode(response.body);
    return GetServerSettings.fromJson(body);
  }
}