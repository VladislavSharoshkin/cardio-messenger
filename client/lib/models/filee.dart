

import 'package:cardio_messenger/store_helper.dart';

enum FileType {unknown, file, image}

class MFile {
  int id;
  String name;
  int size;
  String hash;
  String token;
  FileType type;
  String path;
  String? url;

  MFile(
      {this.id = 0,
        this.name = '',
        this.size = 0,
        this.hash = '',
        this.token = '',
        this.type = FileType.unknown,
        this.path = '',
        this.url = ''});

  MFile.fromJson(Map<String, dynamic> json)
      : id = json['ID'] ?? 0,
        name = json['Name'] ?? '',
        size = json['Size'] ?? 0,
        hash = json['Hash'] ?? '',
        token = json['Token'] ?? '',
        type = FileType.values[json['Type'] ?? 0],
        path = json['Path'] ?? '';
        //url = 'https://${StoreHelper.shared.server.ip}:${StoreHelper.shared.server.port}/file/download/${json['ID'].toString()}';

  Map<String, dynamic> toJson() => {
    'ID': id,
    'Name': name,
    'Size': size,
    'Hash': hash,
    'Type': type.index,
    'Path': path,
    'Url': url
  };

  Uri getURI() {
    return Uri.parse('https://${StoreHelper.shared.domain}/file/download/$token/$name');
  }

  int getId(){
    return this.id;
  }

  bool isLoad(){
    return id != 0;
  }

  static List<int> getIds(List<MFile> files){
    List<int> ids = [];
    files.forEach((file) {
      ids.add(file.getId());
    });
    return ids;
  }

  static MFile init(String patch, int size, String name){
    return MFile(path: patch, size: size, name: name);
  }

  static List<MFile> fromJsonList(dynamic json) {
    return List<MFile>.from(json.map((model)=> MFile.fromJson(model)));
  }

  static List<Map<String, dynamic>> toJsonList(List<MFile> files) {
    return files.map((i) => i.toJson()).toList();
  }
}