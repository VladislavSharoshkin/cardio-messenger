
class Token {
  int userId;
  String token;

  Token(
      {this.userId = 0,
        this.token = ''});

  Token.fromJson(Map<String, dynamic> json)
      : userId = json['UserID'] ?? 0,
        token = json['Token'] ?? '';

  Map<String, dynamic> toJson() => {
    'UserID': userId,
    'Token': token,
  };

  static List<Token> fromJsonList(dynamic json) {
    return List<Token>.from(json.map((model)=> Token.fromJson(model)));
  }
}
