import 'dart:convert';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

class StoreHelper {
  static final shared = StoreHelper();
  final storage = const FlutterSecureStorage();
  late int? userID;
  late String? domain;
  String? token;

  void setLogin(String token, int userID) {
    this.token = token;
    this.userID = userID;
    storage.write(key: 'token', value: token);
    storage.write(key: 'domain', value: domain);
    storage.write(key: 'userID', value: userID.toString());
  }

  Future<void> init() async {
    token = await shared.storage.read(key: 'token');
    domain = await storage.read(key: 'domain');
    String? userIDString =  await storage.read(key: 'userID');
    userID = userIDString == null ? null : int.parse(userIDString);
  }

  bool isLogged() {
    return token != null && domain != null && userID != null;
  }

  // Token token = Token();
  // late ServerInfo server;
  // Map<String, String> authorizationHeader = {};
  //
  // Future<void> loadToken() async {
  //   token = Token.fromJson(jsonDecode(await StoreHelper.shared.storage.read(key: "token") ?? '{}'));
  //   server = ServerInfo.fromJson(jsonDecode(await StoreHelper.shared.storage.read(key: "server") ?? '{}'));
  //   authorizationHeader = {'Authorization': token.token};
  // }
  //
  // Future<void> saveToken(Token token, ServerInfo server) async {
  //   this.token = token;
  //   this.server = server;
  //   storage.write(key: 'token', value: json.encode(token));
  //   storage.write(key: 'server', value: json.encode(server));
  //   authorizationHeader = {'Authorization': token.token};
  // }
}
