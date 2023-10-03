import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';

Color MAIN_COLOR = const Color(0xFF2eb2ff);

// Dark theme
const darkTextColor = Color(0xFFFFFFFF);
const darkBackgroundColor = Color(0xFF1F1F1F);
const darkCanvasColor = Color(0xFF242424);
const darkCardColor = Color(0xFF272727);
const darkDividerColor = Color(0xFF545454);

/// Dark style
final ThemeData theme = ThemeData(
  brightness: Brightness.dark,
  fontFamily: "Product Sans",
  primaryColor: MAIN_COLOR,
  canvasColor: darkCanvasColor,
  scaffoldBackgroundColor: darkBackgroundColor,
  cardColor: darkCardColor,
  cardTheme: CardTheme(
    color: darkCardColor,
    elevation: 0,
    shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(8)),
  ),
  listTileTheme: ListTileThemeData(
    shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(8)),
  ),
  buttonTheme: ButtonThemeData(
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(8))
  ),
  dividerColor: darkDividerColor,
  dialogBackgroundColor: darkCardColor,
  // textTheme: GoogleFonts.openSansTextTheme(ThemeData.dark().textTheme),
  popupMenuTheme: PopupMenuThemeData(
    shape: RoundedRectangleBorder(
      borderRadius: BorderRadius.circular(6),
    ),
  ),
  colorScheme: const ColorScheme.dark().copyWith(
    primary: MAIN_COLOR,
    secondary: MAIN_COLOR,
  ).copyWith(background: darkBackgroundColor),
);