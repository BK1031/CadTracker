import 'package:cad_tracker/models/user.dart';
import 'package:cad_tracker/models/version.dart';
import 'package:fluro/fluro.dart';

final router = FluroRouter();

Version appVersion = Version("1.2.0+1");

// ignore: non_constant_identifier_names
String API_HOST = "https://cad.bk1031.dev/api";
// ignore: non_constant_identifier_names
String TRACKER_AUTH_TOKEN = "tracker-auth-token";

User currentUser = User();