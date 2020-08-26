class Series {
  final String Model;
  final String Manufacturer;
  final String AssemblyPlant;

  Series(this.Model, this.Manufacturer, this.AssemblyPlant);

  Map<String, dynamic> toJson() {
    return {
      'Model': Model,
      'Manufacturer': Manufacturer,
      'AssemblyPlant': AssemblyPlant,
    };
  }
}
