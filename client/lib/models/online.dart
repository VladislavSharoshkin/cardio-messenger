enum OnlineType {unknown, offline, online}

class Online {
  int userId;
  OnlineType type;
  DateTime createdAt;

  Online({this.userId = 0, this.type = OnlineType.unknown, DateTime? createdAt})
  : this.createdAt = createdAt ?? DateTime.utc(1);

  bool isOnline(){
    return type == OnlineType.online && DateTime.now().difference(createdAt).inMinutes < 5;
  }

  Online.fromJson(Map<String, dynamic> json)
      : userId = json['UserID'] ?? 0,
        type = OnlineType.values[json['Type'] ?? 0],
        createdAt = DateTime.parse(json['CreatedAt'] ?? '0001-01-01T00:00:00Z').toLocal();

  static List<Online> fromJsonList(dynamic json) {
    return List<Online>.from(json.map((model)=> Online.fromJson(model)));
  }
}
