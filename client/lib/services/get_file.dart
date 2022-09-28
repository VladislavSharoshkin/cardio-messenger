import 'package:cardio_messenger/models/chat.dart';
import 'package:cardio_messenger/models/error.dart';
import 'package:cardio_messenger/models/filee.dart';
import 'package:cardio_messenger/models/user.dart';
import 'package:cardio_messenger/services/api.dart';
import 'package:open_file/open_file.dart';
import 'package:path_provider/path_provider.dart';
import 'dart:io';
import 'package:http/http.dart' as http;

class GetFile {
  MFile file;
  Errorr? error;

  GetFile.fromJson(Map<String, dynamic> json)
      : file = MFile.fromJson(json['File']),
        error = json['Error'] == null ? null : Errorr.fromJson(json['Error']);

  static Future<MFile> send(String path) async {
    Map<String, dynamic> response = await Api.shared.sendFile(path);
    return GetFile.fromJson(response).file;
  }

  static Future<String> download(MFile file) async {
    String dir = (await getTemporaryDirectory()).path;
    String fileName = file.hash + file.name;
    File fileStore = File('$dir/$fileName');
    if (await fileStore.exists()) return fileStore.path;
    await fileStore.create(recursive: true);
    var response = await http.get(file.getURI()).timeout(const Duration(seconds: 60));

    if (response.statusCode == 200) {
      await fileStore.writeAsBytes(response.bodyBytes);
      return fileStore.path;
    }
    throw 'Download $fileName failed';
  }

  static Future<void> open(MFile file) async {
    OpenFile.open(await download(file));
  }
}