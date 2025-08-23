import 'package:flutter/material.dart';
import 'package:frontend/model/article.dart';

class ArticleScreen extends StatelessWidget {
  final Article article;
  const ArticleScreen({required this.article, super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Text(article.content),
      ),
      appBar: AppBar(
        title: Text(
          article.title,
          softWrap: true,
        ),
        centerTitle: true,
      ),
    );
  }
}
