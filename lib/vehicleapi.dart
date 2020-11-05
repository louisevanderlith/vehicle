import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/keys.dart';
import 'package:mango_ui/requester.dart';

import 'bodies/vehicle.dart';

Future<HttpRequest> submitVehicle(Vehicle obj) async {
  var apiroute = getEndpoint("vehicle");
  var url = "${apiroute}/info";

  return invokeService("POST", url, jsonEncode(obj.toJson()));
}

Future<HttpRequest> updateInfo(Key key, Vehicle obj) async {
  var route = getEndpoint("vehicle");
  var url = "${route}/info/${key.toJson()}";

  final data = jsonEncode(obj.toJson());

  return invokeService("PUT", url, data);
}

Future<HttpRequest> deleteVehicle(Key key) async {
  var route = getEndpoint("vehicle");
  var url = "${route}/info/${key.toJson()}";

  return invokeService("DELETE", url, "");
}
