import 'dart:convert';
import 'package:cad_tracker/utils/logger.dart';
import 'package:http/http.dart' as http;
import 'package:cad_tracker/models/user.dart';
import 'package:cad_tracker/utils/auth_service.dart';
import 'package:cad_tracker/utils/config.dart';
import 'package:cad_tracker/utils/theme.dart';
import 'package:cool_alert/cool_alert.dart';
import 'package:fluro/fluro.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:firebase_auth/firebase_auth.dart' as fb;

class RegisterPage extends StatefulWidget {
  const RegisterPage({Key? key}) : super(key: key);

  @override
  State<RegisterPage> createState() => _RegisterPageState();
}

class _RegisterPageState extends State<RegisterPage> {

  User registerUser = User();
  String password = "";
  String confirmPassword = "";

  Future<void> register() async {
    fb.FirebaseAuth.instance.setPersistence(fb.Persistence.LOCAL);
    if (registerUser.firstName == "" || registerUser.lastName == "" || registerUser.email == "" || password == "") {}
    else if (password != confirmPassword) {
      CoolAlert.show(
        context: context,
        type: CoolAlertType.error,
        title: "Error",
        text: "Passwords do not match",
      );
    }
    else {
      try {
        await fb.FirebaseAuth.instance.createUserWithEmailAndPassword(email: registerUser.email, password: password).then((value) async {
          if (value.user?.uid != null) {
            registerUser.id = value.user!.uid;
            registerUser.privacy.userID = value.user!.uid;
            await AuthService.getAuthToken();
            await http.post(Uri.parse("$API_HOST/users/${registerUser.id}"), headers: {"Authorization": "Bearer $TRACKER_AUTH_TOKEN"}, body: jsonEncode(registerUser));
            await AuthService.getUser(value.user!.uid);
            Future.delayed(Duration.zero, () => router.navigateTo(context, "/", transition: TransitionType.fadeIn, replace: true, clearStack: true));
          }
          else {
            CoolAlert.show(
              context: context,
              type: CoolAlertType.error,
              title: "Error",
              text: "Account creation error",
            );
          }
        });
      } catch (error) {
        log(error, LogLevel.error);
        fb.FirebaseAuth.instance.currentUser?.delete();
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
                Text("Register", style: TextStyle(fontSize: 35, fontWeight: FontWeight.bold, color: MAIN_COLOR), textAlign: TextAlign.center,),
                const Padding(padding: EdgeInsets.all(12.0),),
                const Text("Create to your CadTracker account below!", textAlign: TextAlign.center, style: TextStyle(fontSize: 16),),
                Row(
                  children: [
                    Expanded(
                      child: TextField(
                        decoration: const InputDecoration(
                            icon: Icon(Icons.person),
                            labelText: "First Name",
                            hintText: "Enter your first name"
                        ),
                        autocorrect: false,
                        keyboardType: TextInputType.emailAddress,
                        textCapitalization: TextCapitalization.none,
                        onChanged: (value) {
                          setState(() {
                            registerUser.firstName = value;
                          });
                        },
                      ),
                    ),
                    const Padding(padding: EdgeInsets.all(16)),
                    Expanded(
                      child: TextField(
                        decoration: const InputDecoration(
                            labelText: "Last Name",
                            hintText: "Enter your last name"
                        ),
                        autocorrect: false,
                        keyboardType: TextInputType.emailAddress,
                        textCapitalization: TextCapitalization.none,
                        onChanged: (value) {
                          setState(() {
                            registerUser.lastName = value;
                          });
                        },
                      ),
                    ),
                  ],
                ),
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
                      registerUser.email = value;
                    });
                  },
                ),
                const Padding(padding: EdgeInsets.all(4)),
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    const Text("Gender", style: TextStyle(fontSize: 18),),
                    const Padding(padding: EdgeInsets.all(2)),
                    Card(
                      child: Padding(
                        padding: const EdgeInsets.only(left: 8.0),
                        child: DropdownButton<String>(
                          value: registerUser.gender,
                          alignment: Alignment.centerRight,
                          underline: Container(),
                          style: TextStyle(fontSize: 18, color: Theme.of(context).textTheme.bodyLarge!.color),
                          items: const [
                            DropdownMenuItem(
                              value: "Male",
                              child: Text("Male"),
                            ),
                            DropdownMenuItem(
                              value: "Female",
                              child: Text("Female"),
                            ),
                            DropdownMenuItem(
                              value: "Other",
                              child: Text("Other"),
                            ),
                            DropdownMenuItem(
                              value: "Prefer not to say",
                              child: Text("Prefer not to say"),
                            ),
                          ],
                          borderRadius: BorderRadius.circular(8),
                          onChanged: (item) {
                            if (mounted) {
                              setState(() {
                                registerUser.gender = item!;
                              });
                            }
                          },
                        ),
                      ),
                    )
                  ],
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
                TextField(
                  decoration: const InputDecoration(
                      icon: Icon(Icons.lock),
                      labelText: "Confirm Password",
                      hintText: "Re-enter your password"
                  ),
                  autocorrect: false,
                  keyboardType: TextInputType.number,
                  obscureText: true,
                  onChanged: (value) {
                    setState(() {
                      confirmPassword = value;
                    });
                  },
                ),
                const Padding(padding: EdgeInsets.all(16.0)),
                SizedBox(
                  width: double.infinity,
                  child: CupertinoButton(
                    color: MAIN_COLOR,
                    onPressed: () {
                      register();
                    },
                    child: const Text("REGISTER", style: TextStyle(color: Colors.white, fontSize: 20, fontFamily: "Product Sans", fontWeight: FontWeight.bold)),
                  ),
                ),
                const Padding(padding: EdgeInsets.all(8.0)),
                CupertinoButton(
                  child: const Text("Already have an account?", style: TextStyle(fontSize: 17, fontFamily: "Product Sans"),),
                  onPressed: () {
                    router.navigateTo(context, "/login", transition: TransitionType.fadeIn);
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
