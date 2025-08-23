import 'package:flutter/material.dart';
import 'package:frontend/ui/home.dart';
import 'package:logger/logger.dart';

const String baseUrl =
    'https://f09ebe61-3e06-428c-a313-e88fb4987f66.mock.pstmn.io';

var logger = Logger();

void main() => runApp(Blogify());

class Blogify extends StatelessWidget {
  const Blogify({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(title: 'Blogify', home: HomeScreen());
  }
}
