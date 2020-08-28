import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/requester.dart';
import 'bodies/vehicle.dart';

Future<HttpRequest> submitVehicle(Vehicle obj) async {
  var apiroute = getEndpoint("vehicle");
  var url = "${apiroute}/info";

  return invokeService("POST", url, jsonEncode(obj.toJson()));
}
