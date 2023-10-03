import 'package:cad_tracker/pages/auth/login_page.dart';
import 'package:cad_tracker/pages/auth/register_page.dart';
import 'package:cad_tracker/pages/home/home_page.dart';
import 'package:cad_tracker/utils/auth_service.dart';
import 'package:cad_tracker/utils/config.dart';
import 'package:cad_tracker/utils/firebase_options.dart';
import 'package:cad_tracker/utils/theme.dart';
import 'package:firebase_analytics/firebase_analytics.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:fluro/fluro.dart';
import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:shared_preferences/shared_preferences.dart';

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();

  await dotenv.load(fileName: ".env");

  print("BK CadTracker v${appVersion.toString()}");
  FirebaseApp app = await Firebase.initializeApp(options: DefaultFirebaseOptions.currentPlatform);
  print("Initialized default app $app");
  FirebaseAnalytics analytics = FirebaseAnalytics.instance;

  SharedPreferences prefs = await SharedPreferences.getInstance();
  if (!prefs.containsKey("userID")) prefs.setString("userID", "");
  currentUser.id = prefs.getString("userID")!;
  if (currentUser.id != "") await AuthService.getUser(currentUser.id);

  // ROUTE DEFINITIONS
  router.define("/", handler: Handler(handlerFunc: (BuildContext? context, Map<String, dynamic>? params) {
    return const HomePage();
  }));
  router.define("/login", handler: Handler(handlerFunc: (BuildContext? context, Map<String, dynamic>? params) {
    return const LoginPage();
  }));
  router.define("/register", handler: Handler(handlerFunc: (BuildContext? context, Map<String, dynamic>? params) {
    return const RegisterPage();
  }));

  runApp(MaterialApp(
    title: "CadTracker",
    initialRoute: "/",
    onGenerateRoute: router.generator,
    theme: theme,
    darkTheme: theme,
    themeMode: ThemeMode.dark,
    debugShowCheckedModeBanner: false,
    navigatorObservers: [
      FirebaseAnalyticsObserver(analytics: analytics),
    ],
  ),);
}