import 'package:mango_ui/keys.dart';

import 'engine.dart';
import 'gearbox.dart';
import 'series.dart';

class Vehicle {
  final Key vinKey;
  final String fullVin;
  final Series series;
  final String colour;
  final String paintNo;
  final Engine engine;
  final Gearbox gearbox;
  final String bodyStyle;
  final num doors;
  final List<String> extra;
  final bool spare;
  final bool service;
  final String condition;
  final String issues;
  final num mileage;

  Vehicle(
      this.vinKey,
      this.fullVin,
      this.series,
      this.colour,
      this.paintNo,
      this.engine,
      this.gearbox,
      this.bodyStyle,
      this.doors,
      this.extra,
      this.spare,
      this.service,
      this.condition,
      this.issues,
      this.mileage);

  Map<String, dynamic> toJson() {
    return {
      "VINKey": this.vinKey,
      "FullVIN": this.fullVin,
      "Series": this.series,
      "Colour": this.colour,
      "PaintNo": this.paintNo,
      "Engine": this.engine,
      "Gearbox": this.gearbox,
      "BodyStyle": this.bodyStyle,
      "Doors": this.doors,
      "Extra": this.extra,
      "Spare": this.spare,
      "Service": this.service,
      "Condition": this.condition,
      "Issues": this.issues,
      "Mileage": this.mileage
    };
  }
}
