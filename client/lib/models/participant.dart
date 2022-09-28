
import 'package:cardio_messenger/models/user.dart';

class MParticipant {
  int id;
  int chatId;
  int userId;
  DateTime createdAt;
  MUser user;

  MParticipant(
      {this.id = 0,
        this.chatId = 0,
        MUser? user,
        this.userId = 0,
        DateTime? createdAt})
  : this.createdAt = createdAt?? DateTime.now(), user = user ?? MUser();

  MParticipant.fromJson(Map<String, dynamic> json)
      : id = json['ID'] ?? 0,
        chatId = json['ChatId'] ?? 0,
        userId = json['UserID'] ?? 0,
        user = MUser.fromJson(json['User']),
        createdAt = DateTime.parse(json['CreatedAt'] ?? '');

  Map<String, dynamic> toJson() => {
    'ID': id,
    'chatId': chatId,
    'userId': userId,
    'User': user,

  };

  bool isI(int myId) {
    return userId == myId;
  }

  static MParticipant getParticipantSingleChat(List<MParticipant> participantList, int myId) {
    for (var participant in participantList) {
      if (!participant.isI(myId)) {
        return participant;
      }
    }
    return participantList.first;
  }

  int getId(){
    return id;
  }

  static MParticipant init(int userId, int chatId) {
    return MParticipant(userId: userId, chatId: chatId);
  }

  static List<int> getIds(List<MParticipant> list) {
    return [...list.map((el) => el.id)];
  }

  static List<MParticipant> fromJsonList(dynamic json) {
    return List<MParticipant>.from(json.map((model)=> MParticipant.fromJson(model)));
  }


}
