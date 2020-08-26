class Engine {
  final String Code;
  final String SerialNo;
  final num Output;

  Engine(this.Code, this.SerialNo, this.Output);

  Map<String, dynamic> toJson() {
    return {'Code': Code, 'SerialNo': SerialNo, 'Output': Output};
  }
}
