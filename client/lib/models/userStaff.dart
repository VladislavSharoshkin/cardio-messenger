
import 'package:cardio_messenger/models/branch.dart';
import 'package:cardio_messenger/models/staff.dart';

enum DeliveryStatus { Unknown, Unsent, DeliverToServer }

class UserStaff {
  int id;
  Branch branch;
  Staff staff;

  UserStaff(
      {this.id = 0,
        Branch? branch,
        Staff? staff})
      : branch = branch ?? Branch(), staff = staff ?? Staff();

  UserStaff.fromJson(Map<String, dynamic> json)
      : id = json['ID'] ?? 0,
        branch = Branch.fromJson(json['Branch']),
        staff = Staff.fromJson(json['Staff']);

  Map<String, dynamic> toJson() => {
    'ID': id,
    'Branch': branch,
    'Staff': staff
  };

  int getId(){
    return id;
  }

  static List<UserStaff> fromJsonList(dynamic json) {
    return List<UserStaff>.from(json.map((model)=> UserStaff.fromJson(model)));
  }

  static List<int> getIds(List<UserStaff> messages){
    List<int> ids = [];
    messages.forEach((message) {
      ids.add(message.getId());
    });
    return ids;
  }
}
