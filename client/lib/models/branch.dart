enum DeliveryStatus { Unknown, Unsent, DeliverToServer }

class Branch {
  int id;
  String name;

  Branch(
      {this.id = 0,
        this.name = ''});

  Branch.fromJson(Map<String, dynamic> json)
      : id = json['ID'] ?? 0,
        name = json['Name'] ?? '';

  Map<String, dynamic> toJson() => {
    'ID': id,
    'UserID': name,
  };

  int getId(){
    return id;
  }

  static List<Branch> fromJsonList(dynamic json) {
    return List<Branch>.from(json.map((model)=> Branch.fromJson(model)));
  }

  static List<int> getIds(List<Branch> messages){
    List<int> ids = [];
    messages.forEach((message) {
      ids.add(message.getId());
    });
    return ids;
  }
}
