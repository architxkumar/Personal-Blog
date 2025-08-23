import 'package:json_annotation/json_annotation.dart';

part 'article.g.dart';

@JsonSerializable()
class Article {
  final String id;
  final String title;
  final String content;
  final String date;

  Article({
    required this.id,
    required this.title,
    required this.content,
    required this.date,
  });

  factory Article.fromJson(Map<String, dynamic> json) =>
      _$ArticleFromJson(json);

  Map<String, dynamic> toJson() => _$ArticleToJson(this);
}

List<Article> articleList = [];
