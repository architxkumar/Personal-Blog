import 'package:flutter/material.dart';
import 'package:frontend/model/article.dart';
import 'package:frontend/service/article.dart';
import 'package:frontend/ui/article.dart';
import 'package:result_dart/result_dart.dart';

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  late Future<Result<List<Article>>> _result;

  @override
  void initState() {
    super.initState();
    _result = getArticleList();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Blogify'),
        centerTitle: true,
      ),
      body: Center(
        child: FutureBuilder(
          future: _result,
          builder: (context, snapshot) {
            if (snapshot.connectionState == ConnectionState.waiting) {
              return CircularProgressIndicator();
            } else {
              if (snapshot.data!.isSuccess()) {
                final articles = snapshot.data!.getOrNull()!;
                return ListView.builder(
                  itemCount: articles.length,
                  itemBuilder: (context, index) => ListTile(
                    onTap: () {
                      final article = articles[index];
                      Navigator.push(
                        context,
                        MaterialPageRoute(
                          builder: (context) => ArticleScreen(
                            article: article,
                          ),
                        ),
                      );
                    },
                    title: Text(articles[index].title),
                  ),
                );
              } else {
                return Text(
                  'Error fetching Results',
                  style: TextStyle(color: Colors.red),
                );
              }
            }
          },
        ),
      ),
    );
  }
}
