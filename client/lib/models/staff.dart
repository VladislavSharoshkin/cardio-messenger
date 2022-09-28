enum DeliveryStatus { Unknown, Unsent, DeliverToServer }

class Staff {
  int id;
  String name;

  Staff(
      {this.id = 0,
        this.name = ''});

  Staff.fromJson(Map<String, dynamic> json)
      : id = json['ID'] ?? 0,
        name = json['Name'] ?? '';

  Map<String, dynamic> toJson() => {
    'ID': id,
    'UserID': name,
  };

  int getId(){
    return id;
  }

  static List<Staff> fromJsonList(dynamic json) {
    return List<Staff>.from(json.map((model)=> Staff.fromJson(model)));
  }

  static List<int> getIds(List<Staff> messages){
    List<int> ids = [];
    messages.forEach((message) {
      ids.add(message.getId());
    });
    return ids;
  }
}
