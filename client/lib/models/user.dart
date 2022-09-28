import 'package:cardio_messenger/models/filee.dart';
import 'package:cardio_messenger/models/online.dart';
import 'package:cardio_messenger/models/userStaff.dart';

class MUser {
  int id;
  String username;
  String firstName;
  String middleName;
  String lastName;
  int sex;
  int? avatarID;
  String about;
  Online? lastOnline;

  List<String> roleIds;
  List<UserStaff> userStaffList;

  MUser(
      {this.id = 0,
        this.username = '',
        this.firstName = '',
        this.middleName = '',
        this.lastName = '',
        this.sex = 0,


        List<String>? roleIds,
        List<UserStaff>? userStaffList,
        this.avatarID,
        this.about = ''})
      : roleIds = roleIds ?? [],
        userStaffList = userStaffList ?? [];

  MUser.fromJson(Map<String, dynamic> json)
      : id = json['ID'] ?? 0,
        username = json['Login'] ?? '',
        firstName = json['FirstName'] ?? '',
        middleName = json['MiddleName'] ?? '',
        lastName = json['LastName'] ?? '',
        sex = json['Sex'] ?? 0,
        lastOnline = json['LastOnline'] != null ? Online.fromJson(json['LastOnline']) : null,
        avatarID = json['AvatarID'],
        about = json['About'] ?? '',
        userStaffList = UserStaff.fromJsonList(json['UserStaffs'] ?? []),
        roleIds = List<String>.from(json['RoleIds'] ?? []);

  Map<String, dynamic> toJson() => {
    'ID': id,
    'Login': username,
    'firstName': firstName,
    'middleName': middleName,
    'lastName': lastName,
    'sex': sex,
    'avatarID': avatarID,
    'About': about,
    'RoleIds': roleIds,
    'UserStaffList': userStaffList,
  };

  String getFullName() {
    return this.firstName + ' ' + this.lastName;
  }

  String getFirstName() {
    return this.firstName;
  }

  bool isLoad(){
    return id != 0;
  }

  int getId() {
    return id;
  }

  void setFirstName(String newFirstName) {
    this.firstName = newFirstName;
  }

  void setLastName(String newLastName) {
    this.lastName = newLastName;
  }

  bool isI(String id){
    return this.id == id;
  }

  static List<int> getIds(List<MUser> users) {
    return [...users.map((el) => el.id)];
  }

  static List<MUser> fromJsonList(dynamic json) {
    return List<MUser>.from(json.map((model) => MUser.fromJson(model)));
  }
}