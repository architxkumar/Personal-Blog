// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'article.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

Article _$ArticleFromJson(Map<String, dynamic> json) => Article(
  id: json['id'] as String,
  title: json['title'] as String,
  content: json['content'] as String,
  date: json['date'] as String,
);

Map<String, dynamic> _$ArticleToJson(Article instance) => <String, dynamic>{
  'id': instance.id,
  'title': instance.title,
  'content': instance.content,
  'date': instance.date,
};
