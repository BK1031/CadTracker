class UserPrivacy {
  String userID = "";
  String status = "PUBLIC";
  String statsBasic = "PUBLIC";
  String statsDetailed = "PUBLIC";
  DateTime updatedAt = DateTime.now().toUtc();
  DateTime createdAt = DateTime.now().toUtc();

  UserPrivacy();

  UserPrivacy.fromJson(Map<String, dynamic> json) {
    userID = json["user_id"] ?? "";
    status = json["status"] ?? "";
    statsBasic = json["stats_basic"] ?? "";
    statsDetailed = json["stats_detailed"] ?? "";
    updatedAt = DateTime.tryParse(json["updated_at"]) ?? DateTime.now().toUtc();
    createdAt = DateTime.tryParse(json["created_at"]) ?? DateTime.now().toUtc();
  }

  Map<String, dynamic> toJson() {
    return {
      "user_id": userID,
      "status": status,
      "stats_basic": statsBasic,
      "stats_detailed": statsDetailed,
      "updated_at": updatedAt.toIso8601String(),
      "created_at": createdAt.toIso8601String()
    };
  }
}