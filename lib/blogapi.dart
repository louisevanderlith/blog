import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/keys.dart';
import 'package:mango_ui/requester.dart';

import 'bodies/article.dart';

Future<HttpRequest> createArticle(Article obj) async {
  var apiroute = getEndpoint("blog");
  var url = "${apiroute}/articles";

  return invokeService("POST", url, jsonEncode(obj.toJson()));
}

Future<HttpRequest> updateArticle(Key key, Article obj) async {
  var route = getEndpoint("blog");
  var url = "${route}/articles/${key.toJson()}";

  final data = jsonEncode(obj.toJson());

  return invokeService("PUT", url, data);
}

Future<HttpRequest> deleteArticle(Key key) async {
  var route = getEndpoint("blog");
  var url = "${route}/articles/${key.toJson()}";

  return invokeService("DELETE", url, "");
}
