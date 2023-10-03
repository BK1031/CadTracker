import 'package:cad_tracker/utils/auth_service.dart';
import 'package:cad_tracker/utils/config.dart';
import 'package:cad_tracker/utils/logger.dart';
import 'package:cad_tracker/utils/theme.dart';
import 'package:cool_alert/cool_alert.dart';
import 'package:firebase_auth/firebase_auth.dart';
import 'package:fluro/fluro.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class LoginPage extends StatefulWidget {
  const LoginPage({Key? key}) : super(key: key);

  @override
  State<LoginPage> createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {

  String email = "";
  String password = "";

  Future<void> login() async {
    FirebaseAuth.instance.setPersistence(Persistence.LOCAL);
    if (email == "" || password == "") {}
    else {
      try {
        await FirebaseAuth.instance.signInWithEmailAndPassword(email: email, password: password).then((value) async {
          if (value.user?.uid != null) {
            await AuthService.getUser(value.user!.uid);
            Future.delayed(Duration.zero, () => router.navigateTo(context, "/", transition: TransitionType.fadeIn, replace: true, clearStack: true));
          }
          else {
            CoolAlert.show(
              context: context,
              type: CoolAlertType.error,
              title: "Error",
              text: "Invalid username/password",
            );
          }
        });
      } catch (error) {
        log(error, LogLevel.error);
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Card(
          shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(16.0)),
          child: Container(
            padding: const EdgeInsets.all(16.0),
            width: (MediaQuery.of(context).size.width > 500) ? 500.0 : MediaQuery.of(context).size.width - 25,
            child: Column(
              mainAxisSize: MainAxisSize.min,
              children: <Widget>[
                Text("Login", style: TextStyle(fontSize: 35, fontWeight: FontWeight.bold, color: MAIN_COLOR), textAlign: TextAlign.center,),
                const Padding(padding: EdgeInsets.all(12.0),),
                const Text("Login to your CadTracker account below!", textAlign: TextAlign.center, style: TextStyle(fontSize: 16),),
                TextField(
                  decoration: const InputDecoration(
                      icon: Icon(Icons.email),
                      labelText: "Email",
                      hintText: "Enter your email"
                  ),
                  autocorrect: false,
                  keyboardType: TextInputType.emailAddress,
                  textCapitalization: TextCapitalization.none,
                  onChanged: (value) {
                    setState(() {
                      email = value;
                    });
                  },
                ),
                TextField(
                  decoration: const InputDecoration(
                      icon: Icon(Icons.lock),
                      labelText: "Password",
                      hintText: "Enter your password"
                  ),
                  autocorrect: false,
                  keyboardType: TextInputType.number,
                  obscureText: true,
                  onChanged: (value) {
                    setState(() {
                      password = value;
                    });
                  },
                ),
                const Padding(padding: EdgeInsets.all(16.0)),
                SizedBox(
                  width: double.infinity,
                  child: CupertinoButton(
                    color: MAIN_COLOR,
                    onPressed: () {
                      login();
                    },
                    child: const Text("LOGIN", style: TextStyle(color: Colors.white, fontSize: 20, fontFamily: "Product Sans", fontWeight: FontWeight.bold)),
                  ),
                ),
                const Padding(padding: EdgeInsets.all(8.0)),
                CupertinoButton(
                  child: const Text("Don't have an account?", style: TextStyle(fontSize: 18, fontFamily: "Product Sans"),),
                  onPressed: () {
                    router.navigateTo(context, "/register", transition: TransitionType.fadeIn);
                  },
                )
              ],
            ),
          ),
        ),
      ),
    );
  }
}