class Gearbox {
  final String code;
  final String serialNo;
  final num gears;
  final String type;

  Gearbox(this.code, this.serialNo, this.gears, this.type);

  Map<String, dynamic> toJson() {
    return {
      "Code": this.code,
      "SerialNo": this.serialNo,
      "Gears": this.gears,
      "Type": this.type,
    };
  }
}
