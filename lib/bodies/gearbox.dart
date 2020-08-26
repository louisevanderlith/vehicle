class Gearbox {
  final String Code;
  final String SerialNo;
  final num Gears;
  final String Type;

  Gearbox(this.Code, this.SerialNo, this.Gears, this.Type);

  Map<String, dynamic> toJson() {
    return {
      'Code': Code,
      'SerialNo': SerialNo,
      'Gears': Gears,
      'Type': Type,
    };
  }
}
