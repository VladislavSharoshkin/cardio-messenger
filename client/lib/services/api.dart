import 'package:cardio_messenger/multipart_request.dart';
import 'package:cardio_messenger/store_helper.dart';
import 'package:hive_flutter/hive_flutter.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

class Api {
  static final shared = Api();
  final header = {'Content-Type': 'application/json'};

  Future<Map<String, dynamic>> send(request, path) async {
    print("path: $path request: $request");
    //header['Authorization'] = token!;

    var body = json.encode(request);

    String url = 'https://${StoreHelper.shared.domain}$path';
    var response = await http.post(Uri.parse(url), // https://127.0.0.1:27991/token/add
        headers: header, body: body);
    Map<String, dynamic> responseBody = jsonDecode(response.body);
    print("response: $responseBody");

    return responseBody;
  }

  Future<Map<String, dynamic>> sendWithToken(request, path) async {
    print("path: $path request: $request");
    header['Authorization'] = StoreHelper.shared.token!;

    var body = json.encode(request);

    String url = 'https://${StoreHelper.shared.domain}$path';
    var response = await http.post(Uri.parse(url), // https://127.0.0.1:27991/token/add
        headers: header, body: body);
    Map<String, dynamic> responseBody = jsonDecode(response.body);
    print("response: $responseBody");

    return responseBody;
  }

  Future<Map<String, dynamic>> sendFile(String filePath) async {

    String url = 'https://${StoreHelper.shared.domain}/file/upload';
    final request = MultipartRequest(
      'POST',
      Uri.parse(url),
      onProgress: (int bytes, int total) {
        final progress = bytes / total;
        print('progress: $progress ($bytes/$total)');
      },
    );

    request.headers['HeaderKey'] = 'header_value';
    request.headers['Authorization'] = StoreHelper.shared.token!;
    request.fields['form_key'] = 'form_value';
    request.files.add(
      await http.MultipartFile.fromPath(
        'multipleFiles',
        filePath,
        //contentType: MediaType('image', 'jpeg'),
      ),
    );

    final streamedResponse = await request.send();
    final respStr = await streamedResponse.stream.bytesToString();
    var responseBody = jsonDecode(respStr);
    print("response: $responseBody");
    return responseBody;
  }
}