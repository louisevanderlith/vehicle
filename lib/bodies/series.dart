class Series {
  final String model;
  final String manufacturer;
  final String assemblyPlant;
  final num month;
  final num year;
  final String trim;

  Series(this.year, this.month, this.model, this.manufacturer, this.trim,
      this.assemblyPlant);

  Map<String, dynamic> toJson() {
    return {
      "Model": this.model,
      "Manufacturer": this.manufacturer,
      "AssemblyPlant": this.assemblyPlant,
      "Month": this.month,
      "Year": this.year,
      "Trim": this.trim,
    };
  }
}
