import 'package:cad_tracker/utils/auth_service.dart';
import 'package:cad_tracker/utils/config.dart';
import 'package:cad_tracker/utils/theme.dart';
import 'package:firebase_auth/firebase_auth.dart';
import 'package:fluro/fluro.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class HomeNavbar extends StatefulWidget {
  @override
  _HomeNavbarState createState() => _HomeNavbarState();
}

class _HomeNavbarState extends State<HomeNavbar> {

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 75.0,
      color: Theme.of(context).cardColor,
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceAround,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: <Widget>[
          SizedBox(
              width: 300,
              child: Text(
                "CadTracker",
                style: TextStyle(fontSize: 30, fontWeight: FontWeight.bold, color: MAIN_COLOR),
              )
          ),
          Row(
            crossAxisAlignment: CrossAxisAlignment.center,
            children: <Widget>[
              Visibility(
                visible: currentUser.id == "",
                child: SizedBox(
                  height: 35,
                  width: 100,
                  child: CupertinoButton(
                    padding: EdgeInsets.zero,
                    color: MAIN_COLOR,
                    onPressed: () {
                      router.navigateTo(context, '/login', transition: TransitionType.fadeIn);
                    },
                    child: const Text("LOGIN", style: TextStyle(color: Colors.white, fontSize: 20, fontFamily: "Product Sans", fontWeight: FontWeight.bold)),
                  ),
                ),
              ),
              Visibility(
                visible: currentUser.id != "",
                child: SizedBox(
                  height: 35,
                  width: 120,
                  child: CupertinoButton(
                    padding: EdgeInsets.zero,
                    color: Colors.red,
                    onPressed: () async {
                      await AuthService.signOut();
                      router.navigateTo(context, "/", replace: true, transition: TransitionType.fadeIn);
                    },
                    child: const Text("SIGN OUT", style: TextStyle(color: Colors.white, fontSize: 20, fontFamily: "Product Sans", fontWeight: FontWeight.bold)),
                  ),
                ),
              ),
            ],
          ),
        ],
      ),
    );
  }
}