import 'package:mango_ui/keys.dart';
import 'package:mango_vehicle/bodies/engine.dart';
import 'package:mango_vehicle/bodies/engine.dart';
import 'package:mango_vehicle/bodies/series.dart';

class Vehicle {
  final Key VinKey;
  final String FullVin;
  final Series Series;
  final String Colour;
  final String PaintNo;
  final num Month;
  final num Year;
  final Engine Engine;
  final Gearbox Gearbox;
  final String BodyStyle;
  final num Doors;
  final String Trim;
  final List<String> Extra;

  Vehicle(
      this.VinKey,
      this.FullVin,
      this.Series,
      this.Colour,
      this.PaintNo,
      this.Month,
      this.Year,
      this.Engine,
      this.Gearbox,
      this.BodyStyle,
      this.Doors,
      this.Trim,
      this.Extra);

  Map<String, dynamic> toJson() {
    return {
      'VINKey': VinKey,
      'FullVIN': FullVin,
      'Series': Series,
      'Colour': Colour,
      'PaintNo': PaintNo,
      'Month': Month,
      'Year': Year,
      'Engine': Engine,
      'Gearbox': Gearbox,
      'BodyStyle': BodyStyle,
      'Doors': Doors,
      'Trim': Trim,
      'Extra': Extra,
    };
  }
}
