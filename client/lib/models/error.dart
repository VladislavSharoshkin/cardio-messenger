enum ErrorCode{
  Unknown,
  BadToken,
  Database,
  DecodeRequest,
  NotFound,
  ErrorCodeNotValidation,
  ErrorCodeBadWord,
  ErrorCodeHZ
}

class Errorr {
  ErrorCode code;
  String text;
  int key;

  Errorr({this.code = ErrorCode.Unknown, this.text = '', this.key = 0});

  Errorr.fromJson(Map<String, dynamic> json)
      : code = ErrorCode.values[json['Code'] ?? 0],
        text = json['Text'] ?? '',
        key = json['Key'] ?? 0;

  bool isLoad(){
    return key != 0;
  }
}
