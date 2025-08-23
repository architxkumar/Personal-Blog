import 'dart:convert';

import 'package:frontend/main.dart';
import 'package:frontend/model/article.dart';
import 'package:http/http.dart' as http;
import 'package:result_dart/result_dart.dart';

Future<Result<List<Article>>> getArticleList() async {
  try {
    final response = await http.get(
      Uri.parse(
        '$baseUrl/articles',
      ),
    );
    if (response.statusCode == 200) {
      final decoded = jsonDecode(response.body) as Map<String, dynamic>;
      final List<dynamic> rawArticles = decoded['blogs'] as List<dynamic>;
      final List<Article> articles = rawArticles
          .map((e) => Article.fromJson(e as Map<String, dynamic>))
          .toList();
      logger.i('Articles Retrieved Successfully');
      return Success(articles);
    } else {
      throw Exception('Incorrect Status code: ${response.statusCode}');
    }
  } catch (e, s) {
    logger.e('Error fetching Articles', error: e, stackTrace: s);
    return Failure(Exception(e));
  }
}
