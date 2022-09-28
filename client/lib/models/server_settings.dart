class ServerSettings {
  String? ip;
  int? port;
  String? fingerprint;
  String? name;
  bool? registration;
  int? version;

  ServerSettings(
      {this.ip,
        this.fingerprint,
        this.name});

  ServerSettings.fromJson(Map<String, dynamic> json)
      : ip = json['Ip'],
        fingerprint = json['Fingerprint'],
        name = json['Name'],
        registration = json['Registration'],
        version = json['Version'];

  Map<String, dynamic> toJson() => {
    'Ip': ip,
    'Fingerprint': fingerprint,
    'Name': name,
    'Registration': registration
  };

  String getAddress() {
    return '$ip:$port';
  }

  static List<ServerSettings> fromJsonList(dynamic json) {
    return List<ServerSettings>.from(json.map((model)=> ServerSettings.fromJson(model)));
  }
}
