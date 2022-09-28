import 'package:cardio_messenger/models/chat.dart';
import 'package:cardio_messenger/models/error.dart';
import 'package:cardio_messenger/models/participant.dart';
import 'package:cardio_messenger/models/user.dart';
import 'package:cardio_messenger/services/api.dart';

class GetParticipants {
  List<MParticipant> participants;
  Errorr? error;

  GetParticipants.fromJson(Map<String, dynamic> json)
      : participants = MParticipant.fromJsonList(json['Participants'] ?? []),
        error = json['Error'] == null ? null : Errorr.fromJson(json['Error']);

  static Future<List<MParticipant>> getParticipants(int id) async {
    Map<String, dynamic> response = await Api.shared.sendWithToken({'id': id}, '/chat/getParticipants');
    return GetParticipants.fromJson(response).participants;
  }

  static Future<List<MParticipant>> invite(int chatID, int id) async {
    Map<String, dynamic> response = await Api.shared.sendWithToken({'ChatID': chatID, 'Ids': [id]}, '/chat/invite');
    return GetParticipants.fromJson(response).participants;
  }

  static Future<List<MParticipant>> kick(int chatID, int id) async {
    Map<String, dynamic> response = await Api.shared.sendWithToken({'ChatID': chatID, 'Ids': [id]}, '/chat/invite');
    return GetParticipants.fromJson(response).participants;
  }
}