
class ServerPush {
  String push;
  String token;

  ServerPush(
      {this.push = '',
        this.token = ''});

  ServerPush.fromJson(Map<String, dynamic> json)
      : push = json['Push'] ?? '',
        token = json['Token'] ?? '';




  static List<ServerPush> fromJsonList(dynamic json) {
    return List<ServerPush>.from(json.map((model)=> ServerPush.fromJson(model)));
  }

}
