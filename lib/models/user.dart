import 'package:cad_tracker/models/user_privacy.dart';

class User {
  String id = "";
  String userName = "";
  String firstName = "";
  String lastName = "";
  String email = "";
  String profilePictureURL = "";
  String gender = "Male";
  String discordID = "";
  UserPrivacy privacy = UserPrivacy();
  DateTime updatedAt = DateTime.now().toUtc();
  DateTime createdAt = DateTime.now().toUtc();

  User();

  User.fromJson(Map<String, dynamic> json) {
    id = json["id"] ?? "";
    userName = json["user_name"] ?? "";
    firstName = json["first_name"] ?? "";
    lastName = json["last_name"] ?? "";
    email = json["email"] ?? "";
    profilePictureURL = json["profile_picture_url"] ?? "";
    gender = json["gender"] ?? "";
    discordID = json["discord_id"] ?? "";
    privacy = UserPrivacy.fromJson(json["privacy"]);
    updatedAt = DateTime.tryParse(json["updated_at"]) ?? DateTime.now().toUtc();
    createdAt = DateTime.tryParse(json["created_at"]) ?? DateTime.now().toUtc();
  }

  Map<String, dynamic> toJson() {
    return {
      "id": id,
      "user_name": userName,
      "first_name": firstName,
      "last_name": lastName,
      "email": email,
      "profile_picture_url": profilePictureURL,
      "gender": gender,
      "discord_id": discordID,
      "privacy": privacy,
      "updated_at": updatedAt.toIso8601String(),
      "created_at": createdAt.toIso8601String()
    };
  }

  @override
  String toString() {
    return "[$id] $firstName $lastName (@$userName)";
  }
}
