class Engine {
  final String code;
  final String serialNo;
  final num output;

  Engine(this.code, this.serialNo, this.output);

  Map<String, dynamic> toJson() {
    return {
      "Code": this.code,
      "SerialNo": this.serialNo,
      "Output": this.output
    };
  }
}
