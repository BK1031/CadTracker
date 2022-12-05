import 'package:cad_tracker/models/user.dart';
import 'package:cad_tracker/models/version.dart';
import 'package:fluro/fluro.dart';

final router = FluroRouter();

Version appVersion = Version("0.0.1+1");

// ignore: non_constant_identifier_names
// String API_HOST = "https://api.storkecentr.al";
String API_HOST = "http://localhost:4001";
// String API_HOST = "https://ce01-169-231-9-80.ngrok.io";
// ignore: non_constant_identifier_names
String TRACKER_API_KEY = "tracker-api-key";
// ignore: non_constant_identifier_names
String TRACKER_AUTH_TOKEN = "tracker-auth-token";

User currentUser = User();