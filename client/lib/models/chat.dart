

import 'package:cardio_messenger/models/message.dart';
import 'package:cardio_messenger/models/participant.dart';
import 'package:cardio_messenger/models/user.dart';

enum ChatType { Unknown, Single, Group }

class MChat {
  int id;
  int creatorId;
  ChatType type;
  DateTime createdAt;
  List<MParticipant> participants;
  int? avatarID;
  String name;
  int lastMessageId;
  String about;
  MUser? companion;

  int unreadCount;
  MMessage? lastMessage;
  int participantCount;


  MChat(
      {this.id = 0,
        this.creatorId = 0,
        this.type = ChatType.Unknown,
        DateTime? createdAt,
        List<MParticipant>? participants,
        List<MMessage>? messages,


        this.avatarID = 0,
        this.name = '',
        this.lastMessageId = 0,
        this.about = '',
        this.participantCount = 0,
        this.unreadCount = 0})
  : this.createdAt = createdAt?? DateTime.now(),
        participants = participants ?? [];

  void setName(String name){
    this.name = name;
  }

  bool isSingle(){
    return type == ChatType.Single;
  }

  bool isLoad(){
    return id != 0;
  }

  bool isDirectChat() {
    return type == ChatType.Single;
  }

  static MChat init(ChatType type, List<MParticipant> participants) {
    return MChat(type: type, participants: participants);
  }

  MChat.fromJson(Map<String, dynamic> json)
      : id = json['ID'] ?? 0,
        creatorId = json['CreatorID'] ?? 0,
        type = ChatType.values[json['Type'] ?? 0],
        createdAt = DateTime.parse(json['CreatedAt'] ?? '0001-01-01T00:00:00Z').toLocal(),
        participants = MParticipant.fromJsonList(json['ParticipantList'] ?? []),
        avatarID = json['AvatarID'],
        companion = json['Companion'] == null ? null: MUser.fromJson(json['Companion']),
        lastMessage = json['LastMessage'] == null ? null: MMessage.fromJson(json['LastMessage']),
        name = json['Name'] ?? '',
        participantCount = json['ParticipantCount'] ?? 0,
        lastMessageId = json['LastMessageId'] ?? 0,
        unreadCount = json['UnreadCount'] ?? 0,
        about = json['About'] ?? '';

  Map<String, dynamic> toJson() => {
    'ID': id,
    'CreatorID': creatorId,
    'Type': type.index,

    'ParticipantList': participants,
    'AvatarID': avatarID,
    'Name': name,


    'LastMessageId': lastMessageId,
    'About': about,
    'UnreadCount': unreadCount
  };

  int getId(){
    return id;
  }

  int getParticipantSingleChat(int myId) {
    if (participants.isEmpty){
      return 0;
    }

    if (participants.length == 1) {
      return participants.first.userId;
    }
    for (var i = 0; i < participants.length; i++) {
      MParticipant participant = participants[i];
      if (participant.userId != myId){
        return participant.userId;
      }
    }
    return 0;
  }

  static List<int> getIds(List<MChat> chats){
    return [...chats.map((el) => el.id)];
  }

  void setLastMessageId(int id){
    lastMessageId = id;
  }

  static List<MChat> fromJsonList(dynamic json) {
    return List<MChat>.from(json.map((model)=> MChat.fromJson(model)));
  }


}
