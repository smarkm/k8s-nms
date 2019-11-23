var chart = am4core.create(
    "chartdiv",
    am4charts.PieChart
);
chart.data = [{
    "status": "ready",
    "count": 3,
    "color": "red",
}, {
    "status": "not ready",
    "count": 0
}];

var series = chart.series.push(new am4charts.PieSeries());
series.dataFields.value = "count";
series.dataFields.category = "status";
series.dataFields.color = "color";
series.labels.template.disabled = true;

// this creates initial animation
series.hiddenState.properties.opacity = 1;
series.hiddenState.properties.endAngle = -90;
series.hiddenState.properties.startAngle = -90;
series.legendSettings.valueText = "{count}";

chart.legend = new am4charts.Legend();
chart.legend.position = "right";