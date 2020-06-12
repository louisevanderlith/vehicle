import 'dart:async';
import 'dart:convert';
import 'dart:html';

import '../pathlookup.dart';
import 'requester.dart';

Future<HttpRequest> submitVehicle(String vinKey, String vin) async {
  final url = await buildPath("Vehicle.API", "vehicle", []);

  var data = jsonEncode({
    'VINKey': vinKey,
    'FullVIN': vin,
    'Series': {
      'Manufacturer': '',
    },
    'Colour': '',
    'PaintNo': '',
    'Month': '',
    'Year': '',
    'Engine': '',
    'Gearbox': '',
    'BodyStyle': '',
    'Doors': '',
    'Trim': '',
    'Extra': [],
  });

  return invokeService("POST", url, true, data);
}
