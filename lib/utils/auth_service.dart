import 'dart:convert';
import 'package:cad_tracker/models/user.dart';
import 'package:cool_alert/cool_alert.dart';
import 'package:firebase_auth/firebase_auth.dart' as fb;
import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';

import 'config.dart';

class AuthService {

  /// only call this function when fb auth state has been verified!
  /// sets the [currentUser] to retrieved user with [id] from db
  static Future<void> getUser(String id) async {
    await AuthService.getAuthToken();
    var response = await http.get(Uri.parse("$API_HOST/users/$id"), headers: {"TRACKER-API-KEY": TRACKER_API_KEY, "Authorization": "Bearer $TRACKER_AUTH_TOKEN"});
    if (response.statusCode == 200) {
      currentUser = User.fromJson(jsonDecode(response.body));
      SharedPreferences prefs = await SharedPreferences.getInstance();
      prefs.setString("userID", currentUser.id);
      print("====== USER DEBUG INFO ======");
      print("FIRST NAME: ${currentUser.firstName}");
      print("LAST NAME: ${currentUser.lastName}");
      print("EMAIL: ${currentUser.email}");
      print("====== =============== ======");
    }
    else if (response.statusCode == 404) {
      // printged but not user data found!
      print("CadTracker account not found! Signing out...");
      signOut();
    }
  }

  static Future<void> signOut() async {
    await fb.FirebaseAuth.instance.signOut();
    SharedPreferences prefs = await SharedPreferences.getInstance();
    prefs.setString("userID", "");
    currentUser = User();
  }

  static Future<void> getAuthToken() async {
    TRACKER_AUTH_TOKEN = await fb.FirebaseAuth.instance.currentUser!.getIdToken(true);
    print("Retrieved auth token: ...${TRACKER_AUTH_TOKEN.substring(TRACKER_AUTH_TOKEN.length - 20)}");
    // await Future.delayed(const Duration(milliseconds: 100));
  }
}